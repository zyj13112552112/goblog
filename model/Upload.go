package model

import (
	"context"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"goblog/utils"
	"goblog/utils/errmsg"
	"mime/multipart"
)

var AccessKey = utils.AccessKey
var ScretKey = utils.SecretKey
var Bucket = utils.Bucket
var ImgUrl = utils.QiNiuServer

// UpLoadFile 上传图片
func UpLoadFile(file multipart.File, fileSize int64) (string, int) {
	policy := storage.PutPolicy{
		Scope: Bucket,
	}
	mac := qbox.NewMac(AccessKey, ScretKey)
	token := policy.UploadToken(mac)

	config := storage.Config{
		Zone:          &storage.ZoneHuanan,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}

	putExtra := storage.PutExtra{}

	uploader := storage.NewFormUploader(&config)

	ret := storage.PutRet{}

	err2 := uploader.PutWithoutKey(context.Background(), &ret, token, file, fileSize, &putExtra)

	if err2 != nil {
		return "", errmsg.ERROR
	}

	url := ImgUrl + ret.Key

	return url, errmsg.SUCCSE
}
