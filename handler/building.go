package handler

import (
	"bytes"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/tobgu/qframe"
	"github.com/tobgu/qframe/config/groupby"
	qframeFunction "github.com/tobgu/qframe/function"
	gonumFloat "gonum.org/v1/gonum/floats"
	"net/http"
	"os"
	"time"
)

func jsonFileQFrame(filePath string) qframe.QFrame {
	jsonData, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Unable to read input file "+filePath, err)
	}
	defer jsonData.Close()
	df := qframe.ReadJSON(jsonData)
	return df
}

func getBuildingData(buildingId string) qframe.QFrame {
	electricity := jsonFileQFrame("data/" + buildingId + ".json")
	electricity = electricity.Apply(
		qframe.Instruction{Fn: tsToTime, DstCol: "ts", SrcCol1: "timestamp"},
		qframe.Instruction{Fn: tsToDay, DstCol: "day", SrcCol1: "timestamp"})
	return electricity
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

func filterDataFrame(df qframe.QFrame, start string, end string) qframe.QFrame {
	df = df.Filter(qframe.And(
		qframe.Filter{Column: "ts", Comparator: ">=", Arg: dayToTs(start)},
		qframe.Filter{Column: "ts", Comparator: "<=", Arg: dayToTs(end)},
	),
	)
	return df
}

func dayToTs(day string) int {
	tm, err := time.Parse("2006-01-02", day)
	if err != nil {
		panic(err)
	}
	return int(tm.Unix())
}

// BuildingsList godoc
// @Summary      ListBookmarks
// @Description  List Bookmarks from user
// @Accept       json
// @Produce      json
// @Success      200  {array}  model.Building
// @Router       /buildings [get]
func (h *Handler) BuildingsList(c echo.Context) (err error) {
	df := jsonFileQFrame("data/buildings.json")
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
// @Summary      ListBookmarks
// @Description  List Bookmarks from user
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
	df := getBuildingData(buildingId)
	df = filterDataFrame(df, start, end)

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
