import sys
sys.path.append('identify_proto')

from identify_proto import identify_pb2_grpc, identify_pb2
import grpc
from concurrent import futures
from PIL import Image
import io
import torch
import pathlib
from PIL import ImageDraw

# 临时替换PosixPath以适应Windows系统
temp = pathlib.PosixPath
# 加载模型
model = torch.hub.load("./", "custom", path="runs/train/expse/weights/best.pt", source="local")
# 恢复PosixPath
pathlib.PosixPath = temp

disease_dict = {
    0: "矮化叶病",
    1: "叶斑病",
    2: "叶斑病",
    3: "灰斑病",
    4: "灰斑病",
    5: "锈病",
    6: "锈病",

}

# 对模型进行预测的示例
def predict_image(img):
    # 使用模型对图像进行推理
    results = model(img)
    return results

# 定义服务类
class ImageService(identify_pb2_grpc.ImageServiceServicer):
    def ProcessImage(self, request, context):
        # 将接收到的字节数据转换为图像对象
        image = Image.open(io.BytesIO(request.image_data))
        print(1)
        print(image)
        # gRPC 传输代码
        # 对图像进行预测
        results = predict_image(image)
        retImg = Image.fromarray(results.render()[0])
        output_bytes = io.BytesIO()
        retImg.save(output_bytes, format='JPEG')
        output_bytes.seek(0)
        byte_data = output_bytes.read()
        # 显示结果
        # 构造并返回响应
        response = identify_pb2.ImageResponse()
        response.image_data=byte_data

        df = results.pandas().xyxy[0]
        # 创建一个空字典用于存储class对应的最大confidence
        class_max_confidence_dict = {}

        # 遍历DataFrame的每一行
        for index, row in df.iterrows():
            # 获取当前行的class和confidence
            class_id = row['class']
            confidence = float(row['confidence'])  # 将字符串类型转换为数值类型

            # 如果当前类别的置信度大于之前记录的最大值，则更新最大值
            if class_id not in class_max_confidence_dict or confidence > class_max_confidence_dict[class_id]:
                class_max_confidence_dict[class_id] = confidence

        # 打印存储结果
        for class_id, max_confidence in class_max_confidence_dict.items():
            print("Class ID:", class_id)
            print("Max Confidence:", max_confidence)
            image_objects = response.objects.add()
            name = disease_dict.get(class_id, "未知")
            image_objects.category = name
            image_objects.value = confidence


        # for prediction in results.pandas().xyxy:
        #     confidence = prediction["confidence"].item()  # 提取置信度值
        #     class_id = int(prediction["class"].item())  # 提取类别ID并转换为整数
        #     image_objects = response.objects.add()
        #     name = disease_dict.get(class_id, "未知")
        #     image_objects.category =name
        #     image_objects.value=confidence
        #     # 打印置信度和类别
        #     print("Confidence:", confidence)
        #     print("Class ID:", class_id)
        return response


# 启动服务器
def serve():
    model_path = "runs/train/expse/weights/best.pt"
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    identify_pb2_grpc.add_ImageServiceServicer_to_server(ImageService(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    server.wait_for_termination()

if __name__ == '__main__':
    serve()
