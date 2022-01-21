// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: api/services/v1/notification_service.proto

/*
Package v1 is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package v1

import (
	"context"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = metadata.Join

func request_UpdateNotificationService_CreateEventStreamingGroup_0(ctx context.Context, marshaler runtime.Marshaler, client UpdateNotificationServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq CreateEventStreamingGroupRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.CreateEventStreamingGroup(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_UpdateNotificationService_CreateEventStreamingGroup_0(ctx context.Context, marshaler runtime.Marshaler, server UpdateNotificationServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq CreateEventStreamingGroupRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.CreateEventStreamingGroup(ctx, &protoReq)
	return msg, metadata, err

}

func request_UpdateNotificationService_NotificationStreamGroup_0(ctx context.Context, marshaler runtime.Marshaler, client UpdateNotificationServiceClient, req *http.Request, pathParams map[string]string) (UpdateNotificationService_NotificationStreamGroupClient, runtime.ServerMetadata, error) {
	var metadata runtime.ServerMetadata
	stream, err := client.NotificationStreamGroup(ctx)
	if err != nil {
		grpclog.Infof("Failed to start streaming: %v", err)
		return nil, metadata, err
	}
	dec := marshaler.NewDecoder(req.Body)
	handleSend := func() error {
		var protoReq NotificationStreamGroupRequest
		err := dec.Decode(&protoReq)
		if err == io.EOF {
			return err
		}
		if err != nil {
			grpclog.Infof("Failed to decode request: %v", err)
			return err
		}
		if err := stream.Send(&protoReq); err != nil {
			grpclog.Infof("Failed to send request: %v", err)
			return err
		}
		return nil
	}
	go func() {
		for {
			if err := handleSend(); err != nil {
				break
			}
		}
		if err := stream.CloseSend(); err != nil {
			grpclog.Infof("Failed to terminate client stream: %v", err)
		}
	}()
	header, err := stream.Header()
	if err != nil {
		grpclog.Infof("Failed to get header from client: %v", err)
		return nil, metadata, err
	}
	metadata.HeaderMD = header
	return stream, metadata, nil
}

// RegisterUpdateNotificationServiceHandlerServer registers the http handlers for service UpdateNotificationService to "mux".
// UnaryRPC     :call UpdateNotificationServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterUpdateNotificationServiceHandlerFromEndpoint instead.
func RegisterUpdateNotificationServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server UpdateNotificationServiceServer) error {

	mux.Handle("POST", pattern_UpdateNotificationService_CreateEventStreamingGroup_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/api.services.v1.UpdateNotificationService/CreateEventStreamingGroup", runtime.WithHTTPPathPattern("/api/v1/eventstream/createstreaminggroup"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_UpdateNotificationService_CreateEventStreamingGroup_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_UpdateNotificationService_CreateEventStreamingGroup_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_UpdateNotificationService_NotificationStreamGroup_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		err := status.Error(codes.Unimplemented, "streaming calls are not yet supported in the in-process transport")
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
		return
	})

	return nil
}

// RegisterUpdateNotificationServiceHandlerFromEndpoint is same as RegisterUpdateNotificationServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterUpdateNotificationServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterUpdateNotificationServiceHandler(ctx, mux, conn)
}

// RegisterUpdateNotificationServiceHandler registers the http handlers for service UpdateNotificationService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterUpdateNotificationServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterUpdateNotificationServiceHandlerClient(ctx, mux, NewUpdateNotificationServiceClient(conn))
}

// RegisterUpdateNotificationServiceHandlerClient registers the http handlers for service UpdateNotificationService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "UpdateNotificationServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "UpdateNotificationServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "UpdateNotificationServiceClient" to call the correct interceptors.
func RegisterUpdateNotificationServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client UpdateNotificationServiceClient) error {

	mux.Handle("POST", pattern_UpdateNotificationService_CreateEventStreamingGroup_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/api.services.v1.UpdateNotificationService/CreateEventStreamingGroup", runtime.WithHTTPPathPattern("/api/v1/eventstream/createstreaminggroup"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_UpdateNotificationService_CreateEventStreamingGroup_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_UpdateNotificationService_CreateEventStreamingGroup_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_UpdateNotificationService_NotificationStreamGroup_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/api.services.v1.UpdateNotificationService/NotificationStreamGroup", runtime.WithHTTPPathPattern("/api/v1/eventstreamgroup/notifications"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_UpdateNotificationService_NotificationStreamGroup_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_UpdateNotificationService_NotificationStreamGroup_0(ctx, mux, outboundMarshaler, w, req, func() (proto.Message, error) { return resp.Recv() }, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_UpdateNotificationService_CreateEventStreamingGroup_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"api", "v1", "eventstream", "createstreaminggroup"}, ""))

	pattern_UpdateNotificationService_NotificationStreamGroup_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"api", "v1", "eventstreamgroup", "notifications"}, ""))
)

var (
	forward_UpdateNotificationService_CreateEventStreamingGroup_0 = runtime.ForwardResponseMessage

	forward_UpdateNotificationService_NotificationStreamGroup_0 = runtime.ForwardResponseStream
)
