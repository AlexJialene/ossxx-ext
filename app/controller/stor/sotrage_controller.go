package stor

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"ossxx-ext/app/model"
)

type SotrListReq struct {
	MaxKey     int
	NextMarker string
}

func List(r *ghttp.Request) {
	var (
		req  *SotrListReq
		list oss.ListObjectsResult
		err  error
	)
	if err = r.Parse(&req); err != nil {
		glog.Error(err)
	}
	if req.MaxKey == 0 {
		req.MaxKey = 100
	}
	get := r.Session.Get("auth")
	session := get.(*model.UserSession)
	marker := oss.Marker(req.NextMarker)
	if list, err = session.Bucket.ListObjects(oss.MaxKeys(req.MaxKey), marker); err != nil {
		glog.Error(err)
	}
	r.Response.WriteTpl("stor-list.html", g.Map{
		"maxKey":     req.MaxKey,
		"list":       list.Objects,
		"nextMarker": req.NextMarker,
	})
}
