package middleware

import (
	"context"
	"fmt"
	"github.com/linzijie1998/bytedance_camp_douyin/global"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
)

func AccessLog() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		start := time.Now()
		ctx.Next(c)
		end := time.Now()
		latency := end.Sub(start).Microseconds
		global.DOUYIN_LOGGER.Info(fmt.Sprintf(
			"status=%d cost=%d method=%s full_path=%s client_ip=%s host=%s",
			ctx.Response.StatusCode(),
			latency,
			ctx.Request.Header.Method(),
			ctx.Request.URI().PathOriginal(), ctx.ClientIP(), ctx.Request.Host()))
		//hlog.CtxTracef(c, "status=%d cost=%d method=%s full_path=%s client_ip=%s host=%s",
		//	ctx.Response.StatusCode(), latency,
		//	ctx.Request.Header.Method(), ctx.Request.URI().PathOriginal(), ctx.ClientIP(), ctx.Request.Host())
	}
}
