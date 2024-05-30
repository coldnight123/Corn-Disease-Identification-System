package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"project/cornleafdisease/logic"
)

func GetSearchHandler(c *gin.Context) {
	name := getDiseaseName(c)

	data, err := logic.GetSearch(name)
	if err != nil {
		zap.L().Error("logic.GetSearch() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
