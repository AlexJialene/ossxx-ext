package stor

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"io"
	"os"
	"ossxx-ext/app/model"
)

type SotrListReq struct {
	MaxKey     int
	NextMarker string
}

func Size(size int64) int64 {
	return size / 1024
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
	req.NextMarker = list.NextMarker
	_ = r.Response.WriteTpl("stor-list.html", g.Map{
		"maxKey":      req.MaxKey,
		"list":        list.Objects,
		"nextMarker":  req.NextMarker,
		"IsTruncated": list.IsTruncated,
	})
}

func Detail(r *ghttp.Request) {
	get := r.Session.Get("auth")
	session := get.(*model.UserSession)
	s := r.GetString("fileKey")
	props, err := session.Bucket.GetObjectDetailedMeta(s)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	fmt.Println("Object Meta:", props)
}

func DownLoad(r *ghttp.Request) {
	get := r.Session.Get("auth")
	session := get.(*model.UserSession)
	s := r.GetString("fileKey")
	if s == "" {
		r.Response.WriteExit("Error: fileKey is not null")
	}
	body, e := session.Bucket.GetObject(s)
	if e != nil {
		glog.Error(e)
		r.Response.WriteExit("Error:", e)
	}
	defer body.Close()

	//Download
	r.Response.Header().Add("Content-Type", "application/octet-stream")
	r.Response.Header().Add("content-disposition", "attachment; filename=\""+s+"\"")
	if _, e := io.Copy(r.Response.ResponseWriter, body); e != nil {
		glog.Error(e)
		return
	}

}

func init() {
	g.View().BindFunc("fileSize", Size)
}
