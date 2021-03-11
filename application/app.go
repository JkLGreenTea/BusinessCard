package application

import (
	"../config"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"strings"
)

type Application struct {
	Config *config.Config
}

func NewApplication() *Application {
	cfg := config.NewConfig()

	return &Application{
		Config: cfg,
	}
}

func (app *Application) Run() {
	router := gin.Default()

	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	router.LoadHTMLGlob(strings.Replace(path, "\\", "/", -1) + "/application/templates/*")
	router.Use(static.Serve("/static/", static.LocalFile(strings.Replace(path, "\\", "/", -1) + "/application/static/", true)))

	router.GET("/", HomePage)

	err = router.Run(app.Config.IP + ":" + app.Config.Port)
	if err != nil {
		log.Fatal(err)
	}
}

func HomePage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"main.html",
		gin.H{
		},
	)
}