#! /usr/bin/python
from __future__ import print_function

from PIL import Image, ImageDraw
import logging
import service_pb2
import service_pb2_grpc
import sys
import grpc
import base64



def run():
    channel = grpc.insecure_channel('localhost:50051')
    stub = service_pb2_grpc.TransformationStub(channel)
    response = stub.ToImage(service_pb2.ToImageRequest(text='Hello world!'))
    print("Greeter client received: " + response.image)


if __name__ == '__main__':
    logging.basicConfig()
    run()
