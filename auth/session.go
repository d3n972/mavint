package auth

import (
	"context"
	secRnd "crypto/rand"
	"encoding/base32"
	"encoding/json"
	"github.com/d3n972/mavint/scheduledTasks"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Session struct {
	User       *User             `json:"user"`
	ID         string            `json:"id"`
	Properties map[string]string `json:"props"`
}

func SessionMiddleware(ctx *gin.Context) {
	w, _ := ctx.Get("appctx")
	appCtx := w.(scheduledTasks.AppContext)
	sessid, err := ctx.Cookie("session_id")
	if err == http.ErrNoCookie || appCtx.Redis.Exists(context.TODO(), "session:"+sessid).Val() == 0 {
		o, _ := StartSession()
		if o.Properties == nil {
			o.Properties = map[string]string{}
		}
		o.Properties["remoteIP"] = ctx.RemoteIP()
		serialized, _ := json.Marshal(o)
		appCtx.Redis.Set(context.TODO(), "session:"+o.ID, serialized, 24*time.Hour)
		ctx.SetCookie(
			"session_id",
			o.ID,
			60*60*24,
			"/",
			"127.0.0.1:12700", false, true,
		)
	}
	sSession, _ := appCtx.Redis.Get(context.TODO(), "session:"+sessid).Bytes()
	o := &Session{}
	json.Unmarshal(sSession, o)
	if o.Properties["remote"] != ctx.RemoteIP() {
		appCtx.Redis.Del(context.TODO(), "session:"+sessid)
		return
	}
	ctx.Set("session", o)
	ctx.Next()
}
func StartSession() (*Session, error) {
	sess := &Session{}
	sess.ID = sess.GenerateSessionID()
	return sess, nil
}
func (s Session) GenerateSessionID() string {
	id := make([]byte, 15)
	secRnd.Read(id)
	return base32.HexEncoding.EncodeToString(id)
}
