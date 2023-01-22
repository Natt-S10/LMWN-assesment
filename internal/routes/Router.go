package routes

import (
	_ "fmt"
	_ "net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	CovidSummary_Registry(router)
	return router
}
