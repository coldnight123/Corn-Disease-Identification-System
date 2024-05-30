import os
from PIL import Image

classes = ['矮化叶病','叶斑病' ,'叶斑病（严重）','灰斑病',  '灰斑病（严重）', '锈病','锈病（严重）','健康',...]
class_to_index = {cls: idx for idx, cls in enumerate(classes)}

def save_as_yolo_format(folder_path, filename, object_data):
    # 读取图像文件
    image_path = os.path.join(folder_path, filename + '.jpg')
    img = Image.open(image_path)

    # 获取图像的宽度和高度
    image_width, image_height = img.size

    # 创建或打开YOLO格式的.txt文件
    yolo_file_path = os.path.join(folder_path, filename + '.txt')
    with open(yolo_file_path, 'w') as out_file:
        for obj in object_data:
            label, x_min, y_min, x_max, y_max = obj

            # 计算中心点的x和y坐标
            center_x = (x_min + x_max) / 2 / image_width
            center_y = (y_min + y_max) / 2 / image_height

            # 计算宽度和高度
            width = (x_max - x_min) / image_width
            height = (y_max - y_min) / image_height

            # 将数据写入文件，格式为：class x_center y_center width height
            out_file.write("{} {:.6f} {:.6f} {:.6f} {:.6f}\n".format(label, center_x, center_y, width, height))

        # 主程序部分

if __name__ == '__main__':
    path = r"E:\DemoCode\pythoncode\tools\yolov5-master\datasets\images\train"
    for root, dirs, files in os.walk(path):
        for file in files:
            filename, file_extension = os.path.splitext(file)
            if file.lower().startswith('xiubings') and file.lower().endswith(('.jpg', '.jpeg', '.png')):
                image_path = os.path.join(root, filename + '.jpg')
                img = Image.open(image_path)
                # 获取图像的宽度和高度
                image_width, image_height = img.size
                # 假设的标注数据，这里需要根据实际情况获取
                object_data = [['锈病（严重）', 0, 0, image_width-1, image_height-1]]  # [label, x_min, y_min, x_max, y_max]
                object_data_with_index = [[class_to_index[label], *coords] for label, *coords in object_data]
                # 保存为YOLO格式的.txt文件
                save_as_yolo_format(root, filename, object_data_with_index)