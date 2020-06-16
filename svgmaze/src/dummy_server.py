import time
import grpc

from gen.pb import maze_pb2,maze_pb2_grpc

import base64
from concurrent import futures

# シンプルサービスサーバーの定義
class SimpleServiceServicer(maze_pb2_grpc.MazeServicer):
    # 初期化
    def __init__(self):
      pass

    # 受信時の処理
    def Gen(self, request, context):
        print("req: {}".format(request))
        return maze_pb2.GenResponse(field='###\n# #\n###')

url = '[::]:50051'
# サーバーの開始
server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
maze_pb2_grpc.add_MazeServicer_to_server(SimpleServiceServicer(), server)
server.add_insecure_port(url)
server.start()
print('[START Server] URL:{}'.format(url))

# 待機
try:
   while True:
      time.sleep(3600)
except KeyboardInterrupt:
   # サーバーの停止
   server.stop(0)
