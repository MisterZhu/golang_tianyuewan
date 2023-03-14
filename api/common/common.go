package common

import (
	"fmt"
	"gindiary/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

/*
获取七牛云token
*/
func GetQiNiuUpToken(c *gin.Context) {
	accessKey := "txJmTY6iccAkgsmWUD8ax1yl2TLzNGPgVo2L2wmH"
	secretKey := "eeZwWtLPwDmVjzYGCcZkCAowCL64Zjwy_-rUqOTy"
	bucket := "zlxpicture"

	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	putPolicy.Expires = 7200 //示例2小时有效期

	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	fmt.Println(upToken)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": upToken,
		"msg":  "获取成功",
	})
}

/*
删除七牛云文件
*/
func DeleteQiNiuFile(c *gin.Context) {
	accessKey := "txJmTY6iccAkgsmWUD8ax1yl2TLzNGPgVo2L2wmH"
	secretKey := "eeZwWtLPwDmVjzYGCcZkCAowCL64Zjwy_-rUqOTy"
	bucket := "zlxpicture"
	// key := "github.png"
	key := c.PostForm("key")
	fmt.Println(key)

	mac := qbox.NewMac(accessKey, secretKey)
	// cfg := storage.Config{
	// 	// 是否使用https域名进行资源管理
	// 	UseHTTPS: true,
	// }
	cfg := storage.Config{}

	// 指定空间所在的区域，如果不指定将自动探测
	// 如果没有特殊需求，默认不需要指定
	//cfg.Region=&storage.ZoneHuabei
	bucketManager := storage.NewBucketManager(mac, &cfg)
	err := bucketManager.Delete(bucket, key)
	if err != nil {
		fmt.Println(err)
		response.Fail(c, "删除失败", nil)
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"data": nil,
			"msg":  "删除成功",
		})
	}

}
