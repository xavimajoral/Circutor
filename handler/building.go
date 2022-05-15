package handler

import (
	"bytes"
	"compress/gzip"
	"embed"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/tobgu/qframe"
	"github.com/tobgu/qframe/config/groupby"
	qframeFunction "github.com/tobgu/qframe/function"
	gonumFloat "gonum.org/v1/gonum/floats"
	"net/http"
	"time"
)

func jsonFileQFrame(dataFiles embed.FS, filePath string) (qframe.QFrame, error) {
	jsonData, err := dataFiles.Open(filePath)
	if err != nil {
		//fmt.Println("Unable to read input file "+filePath, err)
		return qframe.QFrame{}, err
	}
	defer jsonData.Close()
	df := qframe.ReadJSON(jsonData)
	return df, nil
}

func jsonFileQFrameZ(dataFiles embed.FS, filePath string) (qframe.QFrame, error) {
	f, err := dataFiles.Open(filePath)
	if err != nil {
		return qframe.QFrame{}, err
	}
	defer f.Close()
	gr, err := gzip.NewReader(f)
	if err != nil {
		return qframe.QFrame{}, err
	}
	defer gr.Close()

	df := qframe.ReadJSON(gr)
	return df, err
}

func getBuildingData(dataFiles embed.FS, buildingId string) (qframe.QFrame, error) {
	electricity, err := jsonFileQFrameZ(dataFiles, "dataz/"+buildingId+".json.gz")
	if err != nil {
		return qframe.QFrame{}, err
	}
	electricity = electricity.Apply(
		qframe.Instruction{Fn: tsToTime, DstCol: "ts", SrcCol1: "timestamp"},
		qframe.Instruction{Fn: tsToDay, DstCol: "day", SrcCol1: "timestamp"})
	return electricity, nil
}

func tsToTime(ts *string) int {
	tm, err := time.Parse("2006-01-02T15:04:05", *ts)
	if err != nil {
		panic(err)
	}
	return int(tm.Unix())
}

func tsToDay(ts *string) *string {
	tm, err := time.Parse("2006-01-02T15:04:05", *ts)
	if err != nil {
		panic(err)
	}
	s := tm.Format("2006-01-02")
	s = s + "T00:00:00"
	return &s
}

func filterDataFrame(df qframe.QFrame, start string, end string) (qframe.QFrame, error) {
	startTs, err := dayToTs(start)
	if err != nil {
		return qframe.QFrame{}, fmt.Errorf("Unable to parse start day: %s", err)
	}
	endTs, err := dayToTs(end)
	if err != nil {
		return qframe.QFrame{}, fmt.Errorf("Unable to parse end day: %s", err)
	}

	df = df.Filter(qframe.And(
		qframe.Filter{Column: "ts", Comparator: ">=", Arg: startTs},
		qframe.Filter{Column: "ts", Comparator: "<=", Arg: endTs},
	),
	)
	return df, nil
}

func dayToTs(day string) (int, error) {
	tm, err := time.Parse("2006-01-02", day)
	if err != nil {
		return 0, err
	}
	return int(tm.Unix()), nil
}

// BuildingsList godoc
// @Summary      ListBookmarks
// @Description  List Bookmarks from user
// @Accept       json
// @Produce      json
// @Success      200  {array}  model.Building
// @Router       /buildings [get]
func (h *Handler) BuildingsList(c echo.Context) (err error) {

	df, err := jsonFileQFrame(h.DataFiles, "dataz/buildings.json")
	if err != nil {
		return err
	}
	df = df.Copy("name", "locationNameEnglish")
	df = df.Select("id", "name")
	var buf bytes.Buffer
	if err := df.ToJSON(&buf); err != nil {
		fmt.Println("There has been an error")
	}

	//return c.JSON(http.StatusOK, buf.String())
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	return c.String(http.StatusOK, buf.String())
}

// BuildingsData godoc
// @Summary      Building Data
// @Description  Returns data from a specific building
// @Accept       json
// @Produce      json
// @Success      200  {array}  model.BuildingData
// @Param        id   path      string  true  "Building ID"
// @Param   	 start  query     string     false  "Start date YYYY-MM-DD"
// @Param   	 end  	query     string     false "End date YYYY-MM-DD"
// @Param        period   path      string  true  "Data aggeration period" Enums(hourly, daily)
// @Router       /buildings/{id}/{period} [get]
func (h *Handler) BuildingsData(c echo.Context) (err error) {
	start := c.QueryParam("start")
	end := c.QueryParam("end")
	period := c.Param("period")
	buildingId := c.Param("id")

	fmt.Println("Getting ", period, "data from building:", buildingId, "from:", start, "to:", end)
	df, err := getBuildingData(h.DataFiles, buildingId)
	if err != nil {
		return err
	}
	df, err = filterDataFrame(df, start, end)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unable to parse start or end string: %s", err))
	}

	if period == "daily" {
		df = df.GroupBy(groupby.Columns("day")).Aggregate(qframe.Aggregation{Fn: gonumFloat.Sum, Column: "value"})
		df = df.Apply(
			qframe.Instruction{Fn: qframeFunction.StrS, DstCol: "timestamp", SrcCol1: "day"},
		)
	}
	df = df.Select("timestamp", "value")
	var buf bytes.Buffer
	if err := df.ToJSON(&buf); err != nil {
		fmt.Println("There has been an error")
	}
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	return c.String(http.StatusOK, buf.String())
}
