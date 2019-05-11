#! /usr/bin/python
from PIL import Image, ImageDraw, ImageFont

import service_pb2
import service_pb2_grpc
import sys
import grpc
from concurrent import futures
import time
import io


_ONE_DAY_IN_SECONDS = 60 * 60 * 24

LETTER_SIZE = 20


class Transformer(service_pb2_grpc.TransformationServicer):
    def ToImage(self, request, context):
        text = request.text
        y = int(LETTER_SIZE * 3.5)
        x = len(text) * LETTER_SIZE
        image = Image.new('RGBA', (x, y), color=(255, 255, 255, 0))
        d = ImageDraw.Draw(image)
        fnt = ImageFont.truetype('fonts/DejaVuSans.ttf', size=LETTER_SIZE)

        d.text((0, 0), text, font=fnt, fill=(0, 0, 0, 255))
        image = image.rotate(10)
        image.save('test_data/test.png')
        bytesarray = io.BytesIO()
        image.save(bytesarray, format="PNG")
        bytesarray = bytesarray.getvalue()
        return service_pb2.ToImageResponse(image=bytesarray, x=x, y=y)


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
