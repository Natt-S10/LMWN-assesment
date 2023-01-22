package covid

import (
	"net/http"

	"github.com/Natt-S10/LMWN-assesment/internal/services/covid"
	"github.com/gin-gonic/gin"
)

func GenerateSummary(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, covid.GetCovidSummary())
}
