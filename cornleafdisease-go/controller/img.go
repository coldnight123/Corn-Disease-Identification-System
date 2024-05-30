package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"project/cornleafdisease/logic"
)

func ImgAnalysisHandler(c *gin.Context) {
	img, format, err := getImage(c)
	if err != nil {
		zap.L().Error("getImage() failed", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
		return
	}
	processedImg, result, diseaseName, err := logic.ProcessImage(img, format)
	if err != nil {
		zap.L().Error("logic.ProcessImage() failed", zap.Error(err))
		ResponseError(c, CodeIdentificationFailed)
		return
	}
	userID, err := getCurrentUserID(c)
	logic.SaveToRecord(result, userID)
	buffer, err := EncodeImageToBytes(processedImg)
	if err != nil {
		zap.L().Error(" EncodeImageToBytes()", zap.Error(err))
		ResponseError(c, CodeIdentificationFailed)
		return
	}
	fmt.Println(diseaseName)
	ResponseSuccess(c, gin.H{
		"result_image": buffer.Bytes(),
		"result":       diseaseName,
	})
}
