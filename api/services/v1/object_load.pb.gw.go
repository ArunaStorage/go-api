// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: api/services/v1/object_load.proto

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

func request_ObjectLoadService_CreateUploadLink_0(ctx context.Context, marshaler runtime.Marshaler, client ObjectLoadServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq CreateUploadLinkRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "id")
	}

	protoReq.Id, err = runtime.Uint64(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "id", err)
	}

	msg, err := client.CreateUploadLink(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_ObjectLoadService_CreateUploadLink_0(ctx context.Context, marshaler runtime.Marshaler, server ObjectLoadServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq CreateUploadLinkRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "id")
	}

	protoReq.Id, err = runtime.Uint64(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "id", err)
	}

	msg, err := server.CreateUploadLink(ctx, &protoReq)
	return msg, metadata, err

}

func request_ObjectLoadService_CreateDownloadLink_0(ctx context.Context, marshaler runtime.Marshaler, client ObjectLoadServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq CreateDownloadLinkRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "id")
	}

	protoReq.Id, err = runtime.Uint64(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "id", err)
	}

	msg, err := client.CreateDownloadLink(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_ObjectLoadService_CreateDownloadLink_0(ctx context.Context, marshaler runtime.Marshaler, server ObjectLoadServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq CreateDownloadLinkRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "id")
	}

	protoReq.Id, err = runtime.Uint64(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "id", err)
	}

	msg, err := server.CreateDownloadLink(ctx, &protoReq)
	return msg, metadata, err

}

func request_ObjectLoadService_CreateDownloadLinksStream_0(ctx context.Context, marshaler runtime.Marshaler, client ObjectLoadServiceClient, req *http.Request, pathParams map[string]string) (ObjectLoadService_CreateDownloadLinksStreamClient, runtime.ServerMetadata, error) {
	var metadata runtime.ServerMetadata
	stream, err := client.CreateDownloadLinksStream(ctx)
	if err != nil {
		grpclog.Infof("Failed to start streaming: %v", err)
		return nil, metadata, err
	}
	dec := marshaler.NewDecoder(req.Body)
	handleSend := func() error {
		var protoReq CreateDownloadLinksStreamRequest
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
	if err := handleSend(); err != nil {
		if cerr := stream.CloseSend(); cerr != nil {
			grpclog.Infof("Failed to terminate client stream: %v", cerr)
		}
		if err == io.EOF {
			return stream, metadata, nil
		}
		return nil, metadata, err
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

func request_ObjectLoadService_StartMultipartUpload_0(ctx context.Context, marshaler runtime.Marshaler, client ObjectLoadServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq StartMultipartUploadRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "id")
	}

	protoReq.Id, err = runtime.Uint64(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "id", err)
	}

	msg, err := client.StartMultipartUpload(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_ObjectLoadService_StartMultipartUpload_0(ctx context.Context, marshaler runtime.Marshaler, server ObjectLoadServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq StartMultipartUploadRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "id")
	}

	protoReq.Id, err = runtime.Uint64(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "id", err)
	}

	msg, err := server.StartMultipartUpload(ctx, &protoReq)
	return msg, metadata, err

}

func request_ObjectLoadService_GetMultipartUploadLink_0(ctx context.Context, marshaler runtime.Marshaler, client ObjectLoadServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetMultipartUploadLinkRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["object_id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "object_id")
	}

	protoReq.ObjectId, err = runtime.Uint64(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "object_id", err)
	}

	val, ok = pathParams["upload_part"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "upload_part")
	}

	protoReq.UploadPart, err = runtime.Int64(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "upload_part", err)
	}

	msg, err := client.GetMultipartUploadLink(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_ObjectLoadService_GetMultipartUploadLink_0(ctx context.Context, marshaler runtime.Marshaler, server ObjectLoadServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq GetMultipartUploadLinkRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["object_id"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "object_id")
	}

	protoReq.ObjectId, err = runtime.Uint64(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "object_id", err)
	}

	val, ok = pathParams["upload_part"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "upload_part")
	}

	protoReq.UploadPart, err = runtime.Int64(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "upload_part", err)
	}

	msg, err := server.GetMultipartUploadLink(ctx, &protoReq)
	return msg, metadata, err

}

func request_ObjectLoadService_CompleteMultipartUpload_0(ctx context.Context, marshaler runtime.Marshaler, client ObjectLoadServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq CompleteMultipartUploadRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.CompleteMultipartUpload(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_ObjectLoadService_CompleteMultipartUpload_0(ctx context.Context, marshaler runtime.Marshaler, server ObjectLoadServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq CompleteMultipartUploadRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.CompleteMultipartUpload(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterObjectLoadServiceHandlerServer registers the http handlers for service ObjectLoadService to "mux".
// UnaryRPC     :call ObjectLoadServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterObjectLoadServiceHandlerFromEndpoint instead.
func RegisterObjectLoadServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server ObjectLoadServiceServer) error {

	mux.Handle("GET", pattern_ObjectLoadService_CreateUploadLink_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/api.services.v1.ObjectLoadService/CreateUploadLink", runtime.WithHTTPPathPattern("/api/v1/objectload/upload/{id}"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_ObjectLoadService_CreateUploadLink_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ObjectLoadService_CreateUploadLink_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_ObjectLoadService_CreateDownloadLink_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/api.services.v1.ObjectLoadService/CreateDownloadLink", runtime.WithHTTPPathPattern("/api/v1/objectload/download/{id}"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_ObjectLoadService_CreateDownloadLink_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ObjectLoadService_CreateDownloadLink_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_ObjectLoadService_CreateDownloadLinksStream_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		err := status.Error(codes.Unimplemented, "streaming calls are not yet supported in the in-process transport")
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
		return
	})

	mux.Handle("POST", pattern_ObjectLoadService_StartMultipartUpload_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/api.services.v1.ObjectLoadService/StartMultipartUpload", runtime.WithHTTPPathPattern("/api/v1/objectload/init_multipart/{id}"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_ObjectLoadService_StartMultipartUpload_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ObjectLoadService_StartMultipartUpload_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_ObjectLoadService_GetMultipartUploadLink_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/api.services.v1.ObjectLoadService/GetMultipartUploadLink", runtime.WithHTTPPathPattern("/api/v1/objectload/upload_multipart_part/{object_id}/{upload_part}"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_ObjectLoadService_GetMultipartUploadLink_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ObjectLoadService_GetMultipartUploadLink_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_ObjectLoadService_CompleteMultipartUpload_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateIncomingContext(ctx, mux, req, "/api.services.v1.ObjectLoadService/CompleteMultipartUpload", runtime.WithHTTPPathPattern("/api/v1/objectload/complete_multipart"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_ObjectLoadService_CompleteMultipartUpload_0(rctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ObjectLoadService_CompleteMultipartUpload_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterObjectLoadServiceHandlerFromEndpoint is same as RegisterObjectLoadServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterObjectLoadServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
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

	return RegisterObjectLoadServiceHandler(ctx, mux, conn)
}

// RegisterObjectLoadServiceHandler registers the http handlers for service ObjectLoadService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterObjectLoadServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterObjectLoadServiceHandlerClient(ctx, mux, NewObjectLoadServiceClient(conn))
}

// RegisterObjectLoadServiceHandlerClient registers the http handlers for service ObjectLoadService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "ObjectLoadServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "ObjectLoadServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "ObjectLoadServiceClient" to call the correct interceptors.
func RegisterObjectLoadServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client ObjectLoadServiceClient) error {

	mux.Handle("GET", pattern_ObjectLoadService_CreateUploadLink_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/api.services.v1.ObjectLoadService/CreateUploadLink", runtime.WithHTTPPathPattern("/api/v1/objectload/upload/{id}"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_ObjectLoadService_CreateUploadLink_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ObjectLoadService_CreateUploadLink_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_ObjectLoadService_CreateDownloadLink_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/api.services.v1.ObjectLoadService/CreateDownloadLink", runtime.WithHTTPPathPattern("/api/v1/objectload/download/{id}"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_ObjectLoadService_CreateDownloadLink_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ObjectLoadService_CreateDownloadLink_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_ObjectLoadService_CreateDownloadLinksStream_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/api.services.v1.ObjectLoadService/CreateDownloadLinksStream", runtime.WithHTTPPathPattern("/api/v1/objectload/download_stream"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_ObjectLoadService_CreateDownloadLinksStream_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ObjectLoadService_CreateDownloadLinksStream_0(ctx, mux, outboundMarshaler, w, req, func() (proto.Message, error) { return resp.Recv() }, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_ObjectLoadService_StartMultipartUpload_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/api.services.v1.ObjectLoadService/StartMultipartUpload", runtime.WithHTTPPathPattern("/api/v1/objectload/init_multipart/{id}"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_ObjectLoadService_StartMultipartUpload_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ObjectLoadService_StartMultipartUpload_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_ObjectLoadService_GetMultipartUploadLink_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/api.services.v1.ObjectLoadService/GetMultipartUploadLink", runtime.WithHTTPPathPattern("/api/v1/objectload/upload_multipart_part/{object_id}/{upload_part}"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_ObjectLoadService_GetMultipartUploadLink_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ObjectLoadService_GetMultipartUploadLink_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_ObjectLoadService_CompleteMultipartUpload_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/api.services.v1.ObjectLoadService/CompleteMultipartUpload", runtime.WithHTTPPathPattern("/api/v1/objectload/complete_multipart"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_ObjectLoadService_CompleteMultipartUpload_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ObjectLoadService_CompleteMultipartUpload_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_ObjectLoadService_CreateUploadLink_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3, 1, 0, 4, 1, 5, 4}, []string{"api", "v1", "objectload", "upload", "id"}, ""))

	pattern_ObjectLoadService_CreateDownloadLink_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3, 1, 0, 4, 1, 5, 4}, []string{"api", "v1", "objectload", "download", "id"}, ""))

	pattern_ObjectLoadService_CreateDownloadLinksStream_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"api", "v1", "objectload", "download_stream"}, ""))

	pattern_ObjectLoadService_StartMultipartUpload_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3, 1, 0, 4, 1, 5, 4}, []string{"api", "v1", "objectload", "init_multipart", "id"}, ""))

	pattern_ObjectLoadService_GetMultipartUploadLink_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3, 1, 0, 4, 1, 5, 4, 1, 0, 4, 1, 5, 5}, []string{"api", "v1", "objectload", "upload_multipart_part", "object_id", "upload_part"}, ""))

	pattern_ObjectLoadService_CompleteMultipartUpload_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"api", "v1", "objectload", "complete_multipart"}, ""))
)

var (
	forward_ObjectLoadService_CreateUploadLink_0 = runtime.ForwardResponseMessage

	forward_ObjectLoadService_CreateDownloadLink_0 = runtime.ForwardResponseMessage

	forward_ObjectLoadService_CreateDownloadLinksStream_0 = runtime.ForwardResponseStream

	forward_ObjectLoadService_StartMultipartUpload_0 = runtime.ForwardResponseMessage

	forward_ObjectLoadService_GetMultipartUploadLink_0 = runtime.ForwardResponseMessage

	forward_ObjectLoadService_CompleteMultipartUpload_0 = runtime.ForwardResponseMessage
)
