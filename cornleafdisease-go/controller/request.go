package controller

import (
	"bytes"
	"errors"
	"github.com/gin-gonic/gin"
	"image"
	"image/jpeg"
	"strconv"
)

const CtxUserIDKey = "userID"
const CtxRecordPage = "page"
const CtxRecordSize = "size"
const CtxDiseaseName = "diseaseName"

var ErrorUserNotLogin = errors.New("用户未登录")
var ErrorImgNotUpload = errors.New("未上传图像")
var ErrorImgDecode = errors.New("无法解码上传图像")
var ErrorImgDecodeF = errors.New("无法对检测图像编码")
var ErrorImgConvert = errors.New("格式转换错误")

// getPageInfo 返回指定页码和每页大小的数据
func getPageInfo(c *gin.Context) (int64, int64) {
	pageStr := c.Query(CtxRecordPage)
	sizeStr := c.Query(CtxRecordSize)
	var (
		page int64
		size int64
		err  error
	)
	page, err = strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		page = 1
	}
	size, err = strconv.ParseInt(sizeStr, 10, 64)
	if err != nil {
		size = 10
	}
	return page, size
}

// getDiseaseName 获取病害名字
func getDiseaseName(c *gin.Context) string {
	DisName := c.Query(CtxDiseaseName)
	return DisName
}

// getCurrentUser 获取当前用户ID
func getCurrentUserID(c *gin.Context) (UserID int64, err error) {
	uid, ok := c.Get(CtxUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	UserID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}

// getImage 获取上传图片
func getImage(c *gin.Context) (img image.Image, format string, errw error) {
	file, _, err := c.Request.FormFile("image")
	if err != nil {
		errw = ErrorImgNotUpload
		return
	}
	defer file.Close()

	// 读取图片数据
	img, format, err = image.Decode(file)
	if err != nil {
		errw = ErrorImgDecode
		return
	}
	return
}

// EncodeImageToBytes 将处理后的图片转换为字节流
func EncodeImageToBytes(img image.Image) (resultBuffer bytes.Buffer, err error) {
	ok := jpeg.Encode(&resultBuffer, img, nil)
	if ok != nil {
		err = ErrorImgDecode
		return
	}
	return
}
