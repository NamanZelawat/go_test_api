import sys
import grpc

sys.path.append('../../../')

from proto.image import test_pb2, test_pb2_grpc


def run():
    # NOTE(gRPC Python Team): .close() is possible on a channel and should be
    # used in circumstances in which the with statement does not fit the needs
    # of the code.
    print("Will try to greet world ...")
    with grpc.insecure_channel("localhost:8080") as channel:
        stub = test_pb2_grpc.GreeterStub(channel)
        response = stub.SayHello(test_pb2.HelloRequest(name="you"))
    print("Greeter client received: " + response.message)
