package interceptor

import (
	"fmt"
	"github.com/gogf/gf/net/ghttp"
	"regexp"
)

const authUri = "/auth*"

func DefaultInterceptor(r *ghttp.Request) {

	if matchString, _ := regexp.MatchString(authUri, r.RequestURI); matchString {
		r.Middleware.Next()
		return
	}

	if get := r.Session.Get("auth"); get != nil {
		fmt.Println(get)
		r.Middleware.Next()
		return
	}

	r.Response.WriteTpl("not-auth.html")
}
