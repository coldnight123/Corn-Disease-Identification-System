package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"project/cornleafdisease/logic"
)

func GetRecordListHandler(c *gin.Context) {
	page, size := getPageInfo(c)
	userID, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("logic.getCurrentUserID() failed", zap.Error(err))
		ResponseError(c, CodeNeedLogin)
		return
	}
	data, err := logic.GetRecordList(page, size, userID)
	if err != nil {
		zap.L().Error("logic.GetRecordList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

func GetCountHandler(c *gin.Context) {
	userID, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("logic.getCurrentUserID() failed", zap.Error(err))
		ResponseError(c, CodeNeedLogin)
		return
	}
	data, err := logic.GetCount(userID)
	if err != nil {
		zap.L().Error("logic.GetCount() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
