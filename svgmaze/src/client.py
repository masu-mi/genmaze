# vim: fileencoding=utf-8

import sys
import grpc
from gen.pb import maze_pb2,maze_pb2_grpc
from maze.maze import Maze

url = 'localhost:50051'
if len(sys.argv) == 2:
    url = sys.argv[1]

def main():
    print("URL:{}".format(url))
    with grpc.insecure_channel(url) as channel:
        stub = maze_pb2_grpc.MazeStub(channel)
        response = stub.Gen(maze_pb2.GenRequest(x=97, y=45))
        print(response.field)
        m = Maze(response.field)
        m.renderMaze('maze.svg')

if __name__ == '__main__':
    main()
