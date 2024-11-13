# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc
import warnings

from chromadb.proto import logservice_pb2 as chromadb_dot_proto_dot_logservice__pb2

GRPC_GENERATED_VERSION = '1.66.1'
GRPC_VERSION = grpc.__version__
_version_not_supported = False

try:
    from grpc._utilities import first_version_is_lower
    _version_not_supported = first_version_is_lower(GRPC_VERSION, GRPC_GENERATED_VERSION)
except ImportError:
    _version_not_supported = True

if _version_not_supported:
    raise RuntimeError(
        f'The grpc package installed is at version {GRPC_VERSION},'
        + f' but the generated code in chromadb/proto/logservice_pb2_grpc.py depends on'
        + f' grpcio>={GRPC_GENERATED_VERSION}.'
        + f' Please upgrade your grpc module to grpcio>={GRPC_GENERATED_VERSION}'
        + f' or downgrade your generated code using grpcio-tools<={GRPC_VERSION}.'
    )


class LogServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.PushLogs = channel.unary_unary(
                '/chroma.LogService/PushLogs',
                request_serializer=chromadb_dot_proto_dot_logservice__pb2.PushLogsRequest.SerializeToString,
                response_deserializer=chromadb_dot_proto_dot_logservice__pb2.PushLogsResponse.FromString,
                _registered_method=True)
        self.PullLogs = channel.unary_unary(
                '/chroma.LogService/PullLogs',
                request_serializer=chromadb_dot_proto_dot_logservice__pb2.PullLogsRequest.SerializeToString,
                response_deserializer=chromadb_dot_proto_dot_logservice__pb2.PullLogsResponse.FromString,
                _registered_method=True)
        self.GetAllCollectionInfoToCompact = channel.unary_unary(
                '/chroma.LogService/GetAllCollectionInfoToCompact',
                request_serializer=chromadb_dot_proto_dot_logservice__pb2.GetAllCollectionInfoToCompactRequest.SerializeToString,
                response_deserializer=chromadb_dot_proto_dot_logservice__pb2.GetAllCollectionInfoToCompactResponse.FromString,
                _registered_method=True)
        self.UpdateCollectionLogOffset = channel.unary_unary(
                '/chroma.LogService/UpdateCollectionLogOffset',
                request_serializer=chromadb_dot_proto_dot_logservice__pb2.UpdateCollectionLogOffsetRequest.SerializeToString,
                response_deserializer=chromadb_dot_proto_dot_logservice__pb2.UpdateCollectionLogOffsetResponse.FromString,
                _registered_method=True)


class LogServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def PushLogs(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def PullLogs(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetAllCollectionInfoToCompact(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def UpdateCollectionLogOffset(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_LogServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'PushLogs': grpc.unary_unary_rpc_method_handler(
                    servicer.PushLogs,
                    request_deserializer=chromadb_dot_proto_dot_logservice__pb2.PushLogsRequest.FromString,
                    response_serializer=chromadb_dot_proto_dot_logservice__pb2.PushLogsResponse.SerializeToString,
            ),
            'PullLogs': grpc.unary_unary_rpc_method_handler(
                    servicer.PullLogs,
                    request_deserializer=chromadb_dot_proto_dot_logservice__pb2.PullLogsRequest.FromString,
                    response_serializer=chromadb_dot_proto_dot_logservice__pb2.PullLogsResponse.SerializeToString,
            ),
            'GetAllCollectionInfoToCompact': grpc.unary_unary_rpc_method_handler(
                    servicer.GetAllCollectionInfoToCompact,
                    request_deserializer=chromadb_dot_proto_dot_logservice__pb2.GetAllCollectionInfoToCompactRequest.FromString,
                    response_serializer=chromadb_dot_proto_dot_logservice__pb2.GetAllCollectionInfoToCompactResponse.SerializeToString,
            ),
            'UpdateCollectionLogOffset': grpc.unary_unary_rpc_method_handler(
                    servicer.UpdateCollectionLogOffset,
                    request_deserializer=chromadb_dot_proto_dot_logservice__pb2.UpdateCollectionLogOffsetRequest.FromString,
                    response_serializer=chromadb_dot_proto_dot_logservice__pb2.UpdateCollectionLogOffsetResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'chroma.LogService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))
    server.add_registered_method_handlers('chroma.LogService', rpc_method_handlers)


 # This class is part of an EXPERIMENTAL API.
class LogService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def PushLogs(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/chroma.LogService/PushLogs',
            chromadb_dot_proto_dot_logservice__pb2.PushLogsRequest.SerializeToString,
            chromadb_dot_proto_dot_logservice__pb2.PushLogsResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def PullLogs(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/chroma.LogService/PullLogs',
            chromadb_dot_proto_dot_logservice__pb2.PullLogsRequest.SerializeToString,
            chromadb_dot_proto_dot_logservice__pb2.PullLogsResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def GetAllCollectionInfoToCompact(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/chroma.LogService/GetAllCollectionInfoToCompact',
            chromadb_dot_proto_dot_logservice__pb2.GetAllCollectionInfoToCompactRequest.SerializeToString,
            chromadb_dot_proto_dot_logservice__pb2.GetAllCollectionInfoToCompactResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def UpdateCollectionLogOffset(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/chroma.LogService/UpdateCollectionLogOffset',
            chromadb_dot_proto_dot_logservice__pb2.UpdateCollectionLogOffsetRequest.SerializeToString,
            chromadb_dot_proto_dot_logservice__pb2.UpdateCollectionLogOffsetResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)
