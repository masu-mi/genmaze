from grpc.tools import protoc
protoc.main(
    (
        '',
        '-I../gen/grpc/maze/pb',
        '--python_out=./src/gen/pb',
        '--grpc_python_out=./src/gen/pb',
        '../gen/grpc/maze/pb/maze.proto',
    )
)
