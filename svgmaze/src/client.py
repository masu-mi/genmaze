# vim: fileencoding=utf-8

import grpc
from gen.pb import maze_pb2,maze_pb2_grpc
import base64


def main():
    with grpc.insecure_channel('localhost:50051') as channel:
        stub = maze_pb2_grpc.MazeStub(channel)
        response = stub.Gen(maze_pb2.GenRequest(x=25, y=25))
        print('[GENERATED]\n'+response.field)

if __name__ == '__main__':
    main()
