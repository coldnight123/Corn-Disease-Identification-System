# 玉米病害识别系统

**其他语言版本: [English](README.md), [中文](README_zh.md).**

### 项目介绍

本系统的开发旨在应对玉米叶片病害给农作物产量和质量带来的负面影响。本研究详细分析了玉米病害识别在农业科学领域内的重要意义及其广阔的应用前景。在研究与设计YOLOv5模型的过程中，本文对比了两种引入的注意力机制，从中选择了最适合的类型和方式。本系统基于深度学习技术，通过在Kaggle平台上训练模型，使用MySQL数据库管理数据，并采用了B/S架构。前端采用Vue框架，后端选择Gin框架提供接口服务。此外，系统通过gRPC技术实现了Go语言与Python的通信，结合YOLOv5模型进行高效的目标检测，致力于开发高精度的玉米病害识别系统。

### 数据来源

[玉米病害数据集_数据集-飞桨AI Studio星河社区 (baidu.com)](https://aistudio.baidu.com/datasetdetail/111048/0)