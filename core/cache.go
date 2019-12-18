package core

import (
	"bytes"
	"encoding/gob"
	"errors"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/logs"

	_ "github.com/astaxie/beego/cache/redis"
)

var cc cache.Cache

func init() {
	var err error

	defer func() {
		if r := recover(); r != nil {
			cc = nil
		}
	}()

	host := beego.AppConfig.String("redis_host")
	cc, err = cache.NewCache("redis", StringsJoin(`{"conn":"`, host, `"}`))
	if err != nil {
		logs.Error("【Init】 redis cache failure! ==>", err)
		return
	}

	logs.Info("【Init】 redis cache ok!")
}

// SetCache 设置缓存
func SetCache(key string, value interface{}, timeout int) error {
	data, err := Encode(value)
	if err != nil {
		return err
	}
	if cc == nil {
		return errors.New("cc is nil")
	}

	defer func() {
		if r := recover(); r != nil {
			logs.Error(r)
			cc = nil
		}
	}()

	timeouts := time.Duration(timeout) * time.Second
	err = cc.Put(key, data, timeouts)
	if err != nil {
		logs.Error(StringsJoin("Set Cache failure，key:", key, "==>"), err)
	}

	return err
}

// GetCache 获得缓存信息
func GetCache(key string, to interface{}) error {
	if cc == nil {
		return errors.New("cc is nil")
	}

	defer func() {
		if r := recover(); r != nil {
			logs.Error(r)
			cc = nil
		}
	}()

	data := cc.Get(key)
	if data == nil {
		return errors.New("Cache not found")
	}

	err := Decode(data.([]byte), to)
	if err != nil {
		logs.Error(StringsJoin("Get Cache failure，key:", key, "==>"), err)
	}

	return err
}

// DelCache 删除缓存信息
func DelCache(key string) error {
	if cc == nil {
		return errors.New("cc is nil")
	}

	defer func() {
		if r := recover(); r != nil {
			cc = nil
		}
	}()

	err := cc.Delete(key)
	if err != nil {
		logs.Error(StringsJoin("Delete Cache failure，key:", key, "==>"), err)
	}

	return err
}

// Encode 用gob进行数据编码
func Encode(data interface{}) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// Decode 用gob进行数据解码
func Decode(data []byte, to interface{}) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	return dec.Decode(to)
}
