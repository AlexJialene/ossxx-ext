package controller

import (
	"github.com/gogf/gf/net/ghttp"
)

func Index(r *ghttp.Request) {
	_ = r.Response.WriteTpl("index.html")
}

func Auth(r *ghttp.Request) {
	_ = r.Response.WriteTpl("index.html")
}
