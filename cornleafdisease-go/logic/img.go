package logic

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/disintegration/imaging"
	"google.golang.org/grpc"
	"image"
	"image/jpeg"
	"image/png"
	service "project/cornleafdisease/pkg/grpc/proto"
	"strings"
)

const host = "localhost"
const port = 50051

var ErrorGrpcConnectToServer = errors.New("连接 gRPC 服务器失败")
var ErrorGrpcCallToServer = errors.New("调用 gRPC 服务失败")
var ErrorImgDecode = errors.New("解码图像数据为图像对象失败")

// ProcessImageFunction 图片处理，检测
func ProcessImage(image image.Image, format string) (img image.Image, result map[float64]string, diseaseName []string, err error) {
	var buffer bytes.Buffer
	if format == "jpeg" {
		err = jpeg.Encode(&buffer, image, nil)
	} else if format == "png" {
		err = png.Encode(&buffer, image)
	}
	if err != nil {
		return
	}
	// 获取字节流
	imageBytes := buffer.Bytes()
	// 创建 gRPC 连接
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", host, port), grpc.WithInsecure())
	//, grpc.WithDefaultCallOptions(
	//	grpc.MaxCallRecvMsgSize(1024*1024*10), // 设置最大接收消息大小为10MB
	//	grpc.MaxCallSendMsgSize(1024*1024*10), // 设置最大发送消息大小为10MB
	//) )
	defer conn.Close()
	if err != nil {
		err = ErrorGrpcConnectToServer
		return
	}
	// 创建 gRPC 客户端
	client := service.NewImageServiceClient(conn)

	// 构造图像请求
	ctx := context.Background()
	req := &service.ImageRequest{
		ImageData: imageBytes,
	}
	//fmt.Println(imageBytes)
	// 发送图像请求
	res, err := client.ProcessImage(ctx, req)
	if err != nil {
		err = ErrorGrpcCallToServer
		return
	}
	// 解析响应
	imageData := res.GetImageData()
	objects := res.GetObjects()

	// 解码图像数据为图像对象
	img, err = imaging.Decode(bytes.NewReader(imageData))
	if err != nil {
		err = ErrorImgDecode
		return
	}

	// 存放检测到的对象信息到 map 中
	result = make(map[float64]string)
	for _, obj := range objects {
		confidence := obj.GetValue()
		category := obj.GetCategory()
		category = strings.Replace(category, "（严重）", "", -1)
		fmt.Println(confidence, category)
		result[float64(confidence)] = category
		diseaseName = append(diseaseName, category)
	}
	return
}
