// Code generated by hertz generator.

package Feed

import (
	"context"
	"offer_tiktok/biz/mw/jwt"

	"github.com/cloudwego/hertz/pkg/app"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _douyinMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _feedMw() []app.HandlerFunc {
	// your code...
	return nil
}

func feedMiddlewareFunc() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		token := c.Query("token")
		if len(token) == 0 {
			return
		} else {
			jwt.JwtMiddleware.MiddlewareFunc()(ctx, c)
			return
		}
	}
}

func _feed0Mw() []app.HandlerFunc {
	return []app.HandlerFunc{
		feedMiddlewareFunc(),
	}
}
