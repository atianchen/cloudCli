package server

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
)

/**
 *
 * @author jensen.chen
 * @date 2022/7/27
 */
type BaseController struct {
}

func (*BaseController) GetPayload(context *gin.Context, v any) error {
	content, exits := context.Get("payload")
	if !exits {
		return errors.New("Miss Param")
	}
	json.Unmarshal([]byte(content.(string)), v)
	return nil
}
