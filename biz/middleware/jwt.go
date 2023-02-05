package middleware

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/linzijie1998/bytedance_camp_douyin/util"
)

func JWTAuthMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// fmt.Println("进入JWT鉴权中间件")
		// fmt.Printf("status=%d method=%s full_path=%s client_ip=%s host=%s",
		// 	c.Response.StatusCode(),
		// 	c.Request.Header.Method(), c.Request.URI().PathOriginal(), c.ClientIP(), c.Request.Host())
		// pre-handle
		// ...
		token := c.Query("token")
		if token == "" {
			token = c.PostForm("token")
			// 未携带Token
			if token == "" {
				c.Abort()
			}
		}
		j := util.NewJWT()
		claim, err := j.ParseToken(token)
		if err != nil {
			// Token解析错误
			c.Abort()
		}
		c.Set("token_user_id", claim.UserInfo.ID)
		c.Next(ctx)
	}
}
