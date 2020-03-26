package context

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"sync"
)

var ossxx *OssXX
var isInit bool

func GetOssXX() *OssXX {
	return ossxx
}

type OssXX struct {
	m       sync.Mutex
	context map[string]*oss.Client
}

func NewOssXXClient() {
	if !isInit {
		m := make(map[string]*oss.Client)
		ossxx = &OssXX{context: m}
		isInit = true
	}
}

func (o *OssXX) NewClient(key, secret, bucket, endpoint string) error {
	o.m.Lock()
	defer o.m.Unlock()
	client, err := oss.New(endpoint, key, secret)
	if err != nil {
		return err
	}
	o.context[key] = client

	return nil
}

func (o *OssXX) Remove(key string) {
	o.m.Lock()
	defer o.m.Unlock()
	delete(o.context, key)
}

func (o *OssXX) Get(key string) *oss.Client {
	o.m.Lock()
	defer o.m.Unlock()
	return o.context[key]
}
