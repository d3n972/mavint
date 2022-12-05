package controllers

import (
	"github.com/d3n972/mavint/domain"
	"github.com/d3n972/mavint/domain/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type VpeController struct{}

func (e *VpeController) Render(ctx *gin.Context) {
	res := []models.VPEEntry{}

	if appCtx, exists := ctx.Get("appctx"); exists {
		appCtx.(domain.AppContext).
			Db.Find(&res)
	}

	ctx.HTML(http.StatusOK, "vpe/list", gin.H{
		"e": res,
	})
}
