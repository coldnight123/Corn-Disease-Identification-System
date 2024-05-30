# YOLOv5 🚀 by Ultralytics, AGPL-3.0 license
"""Run a Flask REST API exposing one or more YOLOv5s models."""
#迁移学习
import argparse
import io
import numpy as np
import cv2
import base64
import torch
from flask import Flask, request
from PIL import Image

app = Flask(__name__)
models = {}

DETECTION_URL = "/v1/object-detection/<model>"

@app.route(DETECTION_URL, methods=["POST"])
def predict(model):
    """Predict and return object detections in JSON format given an image and model name via a Flask REST API POST
    request.
    """
    if request.method != "POST":
        return

    if request.data:
        img =cv2.imdecode(np.frombuffer(request.data,dtype=np.uint8),cv2.IMREAD_COLOR)
        if model in models:
            results = models[model](img)
            results =results.render()[0]
            return cv2.imencode(".jpg",results)[1].tobytes()
    # if request.files.get("image"):
        # im_file = request.files["image"]
        # im_bytes = im_file.read()
        # im = Image.open(io.BytesIO(im_bytes))
        #
        # if model in models:
        #     results = models[model](im)  # reduce size=320 for faster inference
        #
        #     # Get image with detections
        #     image_with_detections = results.render()[0]
        #
        #     # Get xywgn text
        #     xywgn_text = results.pandas().xyxy[0].to_json(orient="records")
        #
        #     return {
        #         "image_with_detections": base64.b64encode(cv2.imencode(".jpg", image_with_detections)[1]).decode(),
        #         "xywgn_text": xywgn_text
        #     }

if __name__ == "__main__":
    parser = argparse.ArgumentParser(description="Flask API exposing YOLOv5 model")
    parser.add_argument("--port", default=5000, type=int, help="port number")
    opt = parser.parse_args()

    # 加载指定路径下的模型
    model_path = "runs/train/exp2/weights/best.pt"
    models["custom"] = torch.hub.load("./", "custom", path=model_path, source="local")
    app.run(host="0.0.0.0", port=opt.port)  # debug=True causes Restarting with stat