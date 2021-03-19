package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	MaxPageSize     = 50
	DefaultPageSize = 10
)

type Rsp struct {
	Ok   bool        `json:"ok"`
	Data interface{} `json:"data"`
}

func Ok(c *gin.Context, datas ...interface{}) {
	if len(datas) > 0 {
		c.JSON(http.StatusOK, gin.H{
			"ok":   true,
			"data": datas[0],
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok": true,
	})

}

func Err(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"ok":   false,
		"data": data,
	})
}

//错误情况增加一个uuid标识，前端反馈uuid来查找错误
func ErrU(c *gin.Context, data interface{}) string {
	uuid := Uuid()

	c.JSON(http.StatusOK, gin.H{
		"ok":   false,
		"data": data,
		"id":   uuid,
	})
	return uuid
}
