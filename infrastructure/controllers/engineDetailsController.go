package controllers

import (
	"fmt"
	"github.com/d3n972/mavint/domain"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type EngineDetailsController struct{}
type internal_UIC_COUNT struct {
	UIC   string
	Count int
}

func (e *EngineDetailsController) getCountsForDay(ctx *gin.Context) ([]internal_UIC_COUNT, error) {
	if appCtx, exists := ctx.Get("appctx"); exists {
		if r, e := appCtx.(domain.AppContext).
			Db.Model(&domain.EngineWorkday{}).
			Where("date = ?", time.Now().Format("2006-01-02")).
			Select("ui_c, count(ui_c) as count").
			Group("ui_c").
			Order("count desc").
			Rows(); e == nil {
			res := []internal_UIC_COUNT{}
			for r.Next() {
				var uic string
				var count int
				r.Scan(&uic, &count)
				res = append(res, internal_UIC_COUNT{
					UIC:   uic,
					Count: count,
				})
			}
			return res, nil
		}

	}
	return nil, nil
}

type resfmt struct {
	JobType        string
	TrainNumber    string
	NearestStation string
	LoggedAt       time.Time
}

func (e *EngineDetailsController) GetDetailsByDayAndUIC(ctx *gin.Context) ([]resfmt, error) {

	if appCtx, exists := ctx.Get("appctx"); exists {
		if r, e := appCtx.(domain.AppContext).
			Db.Model(&domain.EngineWorkday{}).
			Where("ui_c=? AND date=?", ctx.Param("uic"), ctx.Param("date")).
			Select("job_type, train_number,created_at,nearest_station").
			Rows(); e == nil {
			res := []resfmt{}
			for r.Next() {
				re := resfmt{}
				r.Scan(&re.JobType, &re.TrainNumber, &re.LoggedAt, &re.NearestStation)
				res = append(res, re)
			}
			return res, nil
		}

	}
	return nil, nil
}
func (e *EngineDetailsController) CountsForDay(ctx *gin.Context) {
	res, er := e.getCountsForDay(ctx)
	if er != nil {
		fmt.Errorf("%s", er.Error())
	}
	ctx.HTML(http.StatusOK, "engineworks/list", gin.H{
		"date": time.Now().Format("2006-01-02"),
		"r":    res,
	})
}

func (e *EngineDetailsController) Render(ctx *gin.Context) {
	res, er := e.GetDetailsByDayAndUIC(ctx)
	if er != nil {
		fmt.Errorf("%s", er.Error())
	}
	ctx.HTML(http.StatusOK, "engineworks/details", gin.H{
		"r": res,
	})
}
