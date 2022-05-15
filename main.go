package main

import (
	_ "cloud-front-test/docs"
	"cloud-front-test/handler"
	"cloud-front-test/model"
	"embed"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"
	"github.com/swaggo/echo-swagger"
	"strconv"
	"xorm.io/xorm"
)

type Config struct {
	Port     int
	Database string
}

//go:embed dataz/buildings.json dataz/*.json.gz
var dataFiles embed.FS

func loadConfig() *Config {
	cfg := &Config{
		Port:     1234,
		Database: "sqlite://myfile.db",
	}

	viper.SetConfigName("front-test")
	viper.SetEnvPrefix("ft")

	viper.BindEnv("PORT")
	viper.BindEnv("DATABASE")

	//Flags
	viper.AutomaticEnv()

	viper.ReadInConfig()

	if err := viper.Unmarshal(cfg); err != nil {
		fmt.Println("cannot unmarshal config: %s", err)
	}

	return cfg
}

// @title Circutor Frontend TEST API
// @version 1.0
// @description This is a sample server for the Frontend TEST API.

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /
func main() {
	config := loadConfig()
	e := echo.New()
	e.Logger.SetLevel(log.ERROR)
	e.Use(middleware.Logger())

	// Corswith default config allows *
	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))

	jwtMiddleware := middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(handler.Key),
		Skipper: func(c echo.Context) bool {
			// Skip authentication for and signup login requests
			if c.Path() == "/login" || c.Path() == "/signup" {
				return true
			}
			return false
		},
	})

	//e.Use(jwtMiddleware)

	// Database connection
	db, err := xorm.NewEngine("sqlite3", "frontend-test.db")
	//db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		e.Logger.Fatal(err)
	}
	// migration
	//db.AutoMigrate(&model.User{})
	//db.AutoMigrate(&model.Bookmark{})
	if err := db.Sync2(new(model.User)); err != nil {
		fmt.Println(err)
	}
	if err := db.Sync2(new(model.Bookmark)); err != nil {
		fmt.Println(err)
	}

	// Initialize handler
	h := &handler.Handler{DB: db, DataFiles: dataFiles}

	// Routes
	e.POST("/signup", h.Signup)
	e.POST("/login", h.Login)
	e.GET("/user/bookmarks", h.BookmarksList, jwtMiddleware)
	e.POST("/user/bookmarks", h.BookmarksAdd, jwtMiddleware)
	e.GET("/buildings", h.BuildingsList)
	e.GET("/buildings/:id/:period", h.BuildingsData)
	//e.POST("/follow/:id", h.Follow)
	//e.POST("/posts", h.CreatePost)
	//e.GET("/feed", h.FetchPost)
	e.GET("/docs/*", echoSwagger.WrapHandler)
	// Start server

	address := ":" + strconv.Itoa(config.Port)
	fmt.Println("Documentation is at localhost" + address + "/docs/")

	e.Logger.Fatal(e.Start(":" + strconv.Itoa(config.Port)))
}
