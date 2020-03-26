package interceptor

import (
	"github.com/gogf/gf/net/ghttp"
)

const authUri = "/auth"

func DefaultInterceptor(r *ghttp.Request) {
	if r.RequestURI == authUri {
		r.Middleware.Next()
	}

}
