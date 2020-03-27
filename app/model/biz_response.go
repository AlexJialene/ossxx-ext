package model

import "github.com/aliyun/aliyun-oss-go-sdk/oss"

type BizResponse struct {
	Ret  string  `json:"ret"`
	Data UserOss `json:"data"`
}

type UserOss struct {
	UserId     string `json:"userId"`
	OssKey     string `json:"ossKey"`
	OssSecret  string `json:"ossSecret"`
	BucketName string `json:"bucketName"`
	Endpoint   string `json:"endpoint"`
}

type UserSession struct {
	OssKey     string `json:"ossKey"`
	BucketName string `json:"bucketName"`
	Bucket     *oss.Bucket
}
