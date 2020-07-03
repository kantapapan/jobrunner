package main

import (
	"fmt"
	"net/http"

	"github.com/bamzi/jobrunner"
	"github.com/gin-gonic/gin"
)

func main() {
	jobrunner.Start()
	jobrunner.Schedule("@every 5s", MyJob{})

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/jobrunner/status", JobJSON)
	r.Run(":8080")
}

// JobJSON ...
func JobJSON(c *gin.Context) {
	c.JSON(http.StatusOK, jobrunner.StatusJson())
}

// MyJob ...
type MyJob struct {
}

// Run ...
func (e MyJob) Run() {
	fmt.Println("Run MyJob!")
}
