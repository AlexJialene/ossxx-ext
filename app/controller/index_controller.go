package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"io/ioutil"
	"net/http"
	"ossxx-ext/app/model"
	ossxx "ossxx-ext/app/service"
)

type AuthStr struct {
	KEY  string
	USER string
}

type AuthReq struct {
	SystemCode string `p:"systemCode" v:"required#The system code is not null"`
	Secret     string `p:"secret" v:"required#The system secret is not null"`
}

type UserKeyReq struct {
	UserKey string `p:"userKey" v:"required#The userKey is not null"`
}

func Index(r *ghttp.Request) {
	_ = r.Response.WriteTpl("index.html")
}

func Auth(r *ghttp.Request) {
	var req *UserKeyReq
	if err := r.Parse(&req); err != nil {
		r.Response.WriteExit("Error :", err.Error())
	}

	//Get oss config
	biz, err := post(req.UserKey)
	if err == nil && biz.Ret == "success" {
		//Init oss Client
		if err := ossxx.GetOssXX().NewClient(biz.Data.OssKey, biz.Data.OssSecret, "",
			biz.Data.Endpoint); err != nil {
			glog.Error(err)
			//todo
		}

		get := ossxx.GetOssXX().Get(biz.Data.OssKey)
		bucket, err := get.Bucket(biz.Data.BucketName)
		if err != nil {
			glog.Error(err)
			//todo
		}

		r.Session.Set("auth", &model.UserSession{
			OssKey:     biz.Data.OssKey,
			BucketName: biz.Data.BucketName,
			Bucket:     bucket,
		})

		r.Response.RedirectTo("/stor/list", 302)
		//_ = r.Response.WriteTpl("index.html")
	} else {
		_ = r.Response.WriteTpl("not-auth.html")
	}
}

//Logout
func Logout(r *ghttp.Request) {
	get := r.Session.Get("auth")
	session := get.(*model.UserSession)
	ossxx.GetOssXX().Remove(session.OssKey)
	_ = r.Session.Clear()
	_ = r.Response.WriteTpl("hello.html")
}

// Request 9004 start
const ossReqUrl = "http://127.0.0.1:9004/df/sys/systemUserOss/%s"
const contentType = "application/x-www-form-urlencoded"

func post(userKey string) (biz model.BizResponse, err error) {
	var resp *http.Response
	initUrl := fmt.Sprintf(ossReqUrl, userKey)
	resp, err = http.Post(initUrl, contentType, nil)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	if bytes, err := ioutil.ReadAll(resp.Body); err == nil {
		if err := json.Unmarshal(bytes, &biz); err != nil {
			fmt.Println(err)
		}
	}
	return

}
