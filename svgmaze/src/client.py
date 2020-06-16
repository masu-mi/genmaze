# vim: fileencoding=utf-8

import sys
import grpc
from gen.pb import maze_pb2,maze_pb2_grpc

url = 'localhost:50051'
if len(sys.argv) == 2:
    url = sys.argv[1]

def main():
    print("URL:{}".format(url))
    with grpc.insecure_channel(url) as channel:
        stub = maze_pb2_grpc.MazeStub(channel)
        response = stub.Gen(maze_pb2.GenRequest(x=25, y=25))
        print('[GENERATED]\n'+response.field)

if __name__ == '__main__':
    main()
