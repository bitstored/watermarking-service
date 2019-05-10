#! /usr/bin/python
from PIL import Image, ImageDraw

import service_pb2
import service_pb2_grpc
import sys
import grpc
from concurrent import futures
import time
import io


_ONE_DAY_IN_SECONDS = 60 * 60 * 24


class Transformer(service_pb2_grpc.TransformationServicer):
    def ToImage(self, request, context):
        text = request.text
        y = 20
        x = len(text) * 10
        image = Image.new('RGBA', (x, y), color=(255, 255, 255, 0))
        d = ImageDraw.Draw(image)
        d.text((10, 10), text, fill=(0, 0, 0, 255))
        # image.save('test.png')
        bar = io.BytesIO()
        image.save(bar, format=image.format)
        bar = bar.getvalue()
        return service_pb2.ToImageResponse("", 0, 0)


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    service_pb2_grpc.add_TransformationServicer_to_server(
        Transformer(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    try:
        while True:
            time.sleep(_ONE_DAY_IN_SECONDS)
    except KeyboardInterrupt:
        server.stop(0)


if __name__ == '__main__':
    serve()
