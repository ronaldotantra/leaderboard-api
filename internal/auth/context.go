package auth

import "github.com/gin-gonic/gin"

const (
	accessKey = "access_detail_key"
)

func GetAccessDetailFromContext(c *gin.Context) *AccessDetail {
	val, exists := c.Get(accessKey)
	if !exists {
		return nil
	}

	accessDetail, ok := val.(*AccessDetail)
	if !ok {
		return nil
	}

	return accessDetail
}
