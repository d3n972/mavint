package controllers

import (
	"github.com/d3n972/mavint/domain"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type TrainWatchController struct{}

func (t *TrainWatchController) getTimeFromParam(ctx *gin.Context) time.Time {
	if ctx.PostForm("until") != "" {
		//2022-09-21T18:18
		t, _ := time.Parse("2006-01-02T15:04", ctx.PostForm("until"))

		offset, _ := time.ParseDuration("-2h")
		return t.UTC().Add(offset)
	}
	return time.Time{}
}
func (t *TrainWatchController) Render(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "trainwatch/form", gin.H{})
}
func (t *TrainWatchController) Save(ctx *gin.Context) {
	c, _ := ctx.Get("appctx")
	q := (c.(domain.AppContext))

	q.Db.Create(&domain.WatchedTrain{
		TrainID:    ctx.PostForm("train"),
		WatchUntil: t.getTimeFromParam(ctx),
	})
	ctx.JSON(http.StatusOK, gin.H{
		"ok": 1,
	})
}
