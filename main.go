package main

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
	"ossxx-ext/context"
)

const (
	key      = ""
	secret   = ""
	bucket   = ""
	endpoint = ""
)

func main() {
	context.NewOssXXClient()
	err2 := context.GetOssXX().NewClient(key, secret, bucket, endpoint)
	if err2 != nil {
		fmt.Println(err2)
		return
	}

	client := context.GetOssXX().Get(key)

	// 获取存储空间。
	bucket, err := client.Bucket(bucket)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 分页列举所有文件。每页列举200个。
	marker := oss.Marker("")

	lsRes, err := bucket.ListObjects(oss.MaxKeys(5), marker)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	fmt.Println(len(lsRes.Objects))

	for _, v := range lsRes.Objects {
		fmt.Println(v)
	}

	fmt.Println(lsRes.NextMarker)

	if !lsRes.IsTruncated {
		fmt.Println("IsTruncated")
	}

}
