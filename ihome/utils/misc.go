package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	_ "github.com/garyburd/redigo/redis"
	_ "github.com/gomodule/redigo/redis"
)

/* 将url加上 http://IP:PROT/  前缀 */
//http:// + 127.0.0.1 + ：+ 8080 + 请求
func AddDomain2Url(url string) (domain_url string) {
	domain_url = "http://" + G_fastdfs_addr + ":" + G_fastdfs_port + "/" + url

	return domain_url
}

func RedisServer(key, addr, port, dbnum string) (r cache.Cache, err error) {

	/* Redis */
	// 1 info
	redis_info := map[string]string{
		"key":   key,
		"conn":  addr + ":" + port,
		"dbnum": dbnum,
	}

	// 2 marshal
	redis_config, _ := json.Marshal(redis_info)

	// 3 new
	r, err = cache.NewCache("redis", string(redis_config))
	if err != nil {
		return nil, err
	}

	/* return */
	return r, nil
}

func CrotoMd5(str string) (string, error) {

	hash := md5.New()

	_, err := hash.Write([]byte(str))
	if err != nil {
		return "", err
	}

	str = hex.EncodeToString(hash.Sum(nil))

	return str, nil
}
