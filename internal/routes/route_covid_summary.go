package routes

import (
	"github.com/Natt-S10/LMWN-assesment/internal/controllers/covid"
	"github.com/gin-gonic/gin"
)

func CovidSummary_Registry(route *gin.Engine) { // if not working try *gin.Engine
	route.GET("/covid/summary", covid.GenerateSummary)

}
