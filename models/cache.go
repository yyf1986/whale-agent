package models

import (
	"github.com/astaxie/beego/cache"
	"time"
)

//文件缓存
//var mem, _ = cache.NewCache("memory", `{"interval":10}`)
var fc, _ = cache.NewCache("file",
					`{"CachePath":"./logs/cache",
								"FileSuffix":".cache",
								"DirectoryLevel":"1",
								"EmbedExpiry":"120"}`)

func setCache2File(key string, value interface{}, timeout time.Duration) {
	fc.Put(key, value, timeout)
}

func GetCacheFromFile(key string) interface{} {
	return fc.Get(key)
}

func IsExistInFie(key string) bool {
	return fc.IsExist(key)
}

func DelCacheFromFile(key string) {
	fc.Delete(key)
}
