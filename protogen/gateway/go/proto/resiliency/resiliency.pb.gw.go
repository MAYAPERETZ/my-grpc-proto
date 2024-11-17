// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: proto/resiliency/resiliency.proto

/*
Package resiliency is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package resiliency

import (
	"context"
	"io"
	"net/http"

	extResiliency "github.com/MAYAPERETZ/my-grpc-proto/protogen/go/resiliency"
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

var (
	filter_ResiliencyService_UnaryResiliency_0 = &utilities.DoubleArray{Encoding: map[string]int{}, Base: []int(nil), Check: []int(nil)}
)

func request_ResiliencyService_UnaryResiliency_0(ctx context.Context, marshaler runtime.Marshaler, client extResiliency.ResiliencyServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq extResiliency.ResiliencyRequest
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_ResiliencyService_UnaryResiliency_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.UnaryResiliency(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_ResiliencyService_UnaryResiliency_0(ctx context.Context, marshaler runtime.Marshaler, server extResiliency.ResiliencyServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq extResiliency.ResiliencyRequest
	var metadata runtime.ServerMetadata

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_ResiliencyService_UnaryResiliency_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.UnaryResiliency(ctx, &protoReq)
	return msg, metadata, err

}

var (
	filter_ResiliencyService_ServerStreamingResiliency_0 = &utilities.DoubleArray{Encoding: map[string]int{"min_delay_second": 0, "max_delay_second": 1}, Base: []int{1, 1, 2, 0, 0}, Check: []int{0, 1, 1, 2, 3}}
)

func request_ResiliencyService_ServerStreamingResiliency_0(ctx context.Context, marshaler runtime.Marshaler, client extResiliency.ResiliencyServiceClient, req *http.Request, pathParams map[string]string) (extResiliency.ResiliencyService_ServerStreamingResiliencyClient, runtime.ServerMetadata, error) {
	var protoReq extResiliency.ResiliencyRequest
	var metadata runtime.ServerMetadata

	var (
		val string
		ok  bool
		err error
		_   = err
	)

	val, ok = pathParams["min_delay_second"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "min_delay_second")
	}

	protoReq.MinDelaySecond, err = runtime.Int32(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "min_delay_second", err)
	}

	val, ok = pathParams["max_delay_second"]
	if !ok {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "missing parameter %s", "max_delay_second")
	}

	protoReq.MaxDelaySecond, err = runtime.Int32(val)
	if err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "type mismatch, parameter: %s, error: %v", "max_delay_second", err)
	}

	if err := req.ParseForm(); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}
	if err := runtime.PopulateQueryParameters(&protoReq, req.Form, filter_ResiliencyService_ServerStreamingResiliency_0); err != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	stream, err := client.ServerStreamingResiliency(ctx, &protoReq)
	if err != nil {
		return nil, metadata, err
	}
	header, err := stream.Header()
	if err != nil {
		return nil, metadata, err
	}
	metadata.HeaderMD = header
	return stream, metadata, nil

}

func request_ResiliencyService_ClientStreamingResiliency_0(ctx context.Context, marshaler runtime.Marshaler, client extResiliency.ResiliencyServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var metadata runtime.ServerMetadata
	stream, err := client.ClientStreamingResiliency(ctx)
	if err != nil {
		grpclog.Errorf("Failed to start streaming: %v", err)
		return nil, metadata, err
	}
	dec := marshaler.NewDecoder(req.Body)
	for {
		var protoReq extResiliency.ResiliencyRequest
		err = dec.Decode(&protoReq)
		if err == io.EOF {
			break
		}
		if err != nil {
			grpclog.Errorf("Failed to decode request: %v", err)
			return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
		}
		if err = stream.Send(&protoReq); err != nil {
			if err == io.EOF {
				break
			}
			grpclog.Errorf("Failed to send request: %v", err)
			return nil, metadata, err
		}
	}

	if err := stream.CloseSend(); err != nil {
		grpclog.Errorf("Failed to terminate client stream: %v", err)
		return nil, metadata, err
	}
	header, err := stream.Header()
	if err != nil {
		grpclog.Errorf("Failed to get header from client: %v", err)
		return nil, metadata, err
	}
	metadata.HeaderMD = header

	msg, err := stream.CloseAndRecv()
	metadata.TrailerMD = stream.Trailer()
	return msg, metadata, err

}

func request_ResiliencyService_BiDirectionalResiliency_0(ctx context.Context, marshaler runtime.Marshaler, client extResiliency.ResiliencyServiceClient, req *http.Request, pathParams map[string]string) (extResiliency.ResiliencyService_BiDirectionalResiliencyClient, runtime.ServerMetadata, chan error, error) {
	var metadata runtime.ServerMetadata
	errChan := make(chan error, 1)
	stream, err := client.BiDirectionalResiliency(ctx)
	if err != nil {
		grpclog.Errorf("Failed to start streaming: %v", err)
		close(errChan)
		return nil, metadata, errChan, err
	}
	dec := marshaler.NewDecoder(req.Body)
	handleSend := func() error {
		var protoReq extResiliency.ResiliencyRequest
		err := dec.Decode(&protoReq)
		if err == io.EOF {
			return err
		}
		if err != nil {
			grpclog.Errorf("Failed to decode request: %v", err)
			return status.Errorf(codes.InvalidArgument, "Failed to decode request: %v", err)
		}
		if err := stream.Send(&protoReq); err != nil {
			grpclog.Errorf("Failed to send request: %v", err)
			return err
		}
		return nil
	}
	go func() {
		defer close(errChan)
		for {
			if err := handleSend(); err != nil {
				errChan <- err
				break
			}
		}
		if err := stream.CloseSend(); err != nil {
			grpclog.Errorf("Failed to terminate client stream: %v", err)
		}
	}()
	header, err := stream.Header()
	if err != nil {
		grpclog.Errorf("Failed to get header from client: %v", err)
		return nil, metadata, errChan, err
	}
	metadata.HeaderMD = header
	return stream, metadata, errChan, nil
}

func request_ResiliencyWithMetadataService_UnaryResiliencyWithMetadata_0(ctx context.Context, marshaler runtime.Marshaler, client extResiliency.ResiliencyWithMetadataServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq extResiliency.ResiliencyRequest
	var metadata runtime.ServerMetadata

	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.UnaryResiliencyWithMetadata(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_ResiliencyWithMetadataService_UnaryResiliencyWithMetadata_0(ctx context.Context, marshaler runtime.Marshaler, server extResiliency.ResiliencyWithMetadataServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq extResiliency.ResiliencyRequest
	var metadata runtime.ServerMetadata

	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.UnaryResiliencyWithMetadata(ctx, &protoReq)
	return msg, metadata, err

}

func request_ResiliencyWithMetadataService_ServerStreamingResiliencyWithMetadata_0(ctx context.Context, marshaler runtime.Marshaler, client extResiliency.ResiliencyWithMetadataServiceClient, req *http.Request, pathParams map[string]string) (extResiliency.ResiliencyWithMetadataService_ServerStreamingResiliencyWithMetadataClient, runtime.ServerMetadata, error) {
	var protoReq extResiliency.ResiliencyRequest
	var metadata runtime.ServerMetadata

	if err := marshaler.NewDecoder(req.Body).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	stream, err := client.ServerStreamingResiliencyWithMetadata(ctx, &protoReq)
	if err != nil {
		return nil, metadata, err
	}
	header, err := stream.Header()
	if err != nil {
		return nil, metadata, err
	}
	metadata.HeaderMD = header
	return stream, metadata, nil

}

func request_ResiliencyWithMetadataService_ClientStreamingResiliencyWithMetadata_0(ctx context.Context, marshaler runtime.Marshaler, client extResiliency.ResiliencyWithMetadataServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var metadata runtime.ServerMetadata
	stream, err := client.ClientStreamingResiliencyWithMetadata(ctx)
	if err != nil {
		grpclog.Errorf("Failed to start streaming: %v", err)
		return nil, metadata, err
	}
	dec := marshaler.NewDecoder(req.Body)
	for {
		var protoReq extResiliency.ResiliencyRequest
		err = dec.Decode(&protoReq)
		if err == io.EOF {
			break
		}
		if err != nil {
			grpclog.Errorf("Failed to decode request: %v", err)
			return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
		}
		if err = stream.Send(&protoReq); err != nil {
			if err == io.EOF {
				break
			}
			grpclog.Errorf("Failed to send request: %v", err)
			return nil, metadata, err
		}
	}

	if err := stream.CloseSend(); err != nil {
		grpclog.Errorf("Failed to terminate client stream: %v", err)
		return nil, metadata, err
	}
	header, err := stream.Header()
	if err != nil {
		grpclog.Errorf("Failed to get header from client: %v", err)
		return nil, metadata, err
	}
	metadata.HeaderMD = header

	msg, err := stream.CloseAndRecv()
	metadata.TrailerMD = stream.Trailer()
	return msg, metadata, err

}

func request_ResiliencyWithMetadataService_BiDirectionalResiliencyWithMetadata_0(ctx context.Context, marshaler runtime.Marshaler, client extResiliency.ResiliencyWithMetadataServiceClient, req *http.Request, pathParams map[string]string) (extResiliency.ResiliencyWithMetadataService_BiDirectionalResiliencyWithMetadataClient, runtime.ServerMetadata, chan error, error) {
	var metadata runtime.ServerMetadata
	errChan := make(chan error, 1)
	stream, err := client.BiDirectionalResiliencyWithMetadata(ctx)
	if err != nil {
		grpclog.Errorf("Failed to start streaming: %v", err)
		close(errChan)
		return nil, metadata, errChan, err
	}
	dec := marshaler.NewDecoder(req.Body)
	handleSend := func() error {
		var protoReq extResiliency.ResiliencyRequest
		err := dec.Decode(&protoReq)
		if err == io.EOF {
			return err
		}
		if err != nil {
			grpclog.Errorf("Failed to decode request: %v", err)
			return status.Errorf(codes.InvalidArgument, "Failed to decode request: %v", err)
		}
		if err := stream.Send(&protoReq); err != nil {
			grpclog.Errorf("Failed to send request: %v", err)
			return err
		}
		return nil
	}
	go func() {
		defer close(errChan)
		for {
			if err := handleSend(); err != nil {
				errChan <- err
				break
			}
		}
		if err := stream.CloseSend(); err != nil {
			grpclog.Errorf("Failed to terminate client stream: %v", err)
		}
	}()
	header, err := stream.Header()
	if err != nil {
		grpclog.Errorf("Failed to get header from client: %v", err)
		return nil, metadata, errChan, err
	}
	metadata.HeaderMD = header
	return stream, metadata, errChan, nil
}

// RegisterResiliencyServiceHandlerServer registers the http handlers for service ResiliencyService to "mux".
// UnaryRPC     :call ResiliencyServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterResiliencyServiceHandlerFromEndpoint instead.
// GRPC interceptors will not work for this type of registration. To use interceptors, you must use the "runtime.WithMiddlewares" option in the "runtime.NewServeMux" call.
func RegisterResiliencyServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server extResiliency.ResiliencyServiceServer) error {

	mux.Handle("GET", pattern_ResiliencyService_UnaryResiliency_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/resiliency.ResiliencyService/UnaryResiliency", runtime.WithHTTPPathPattern("/api/resiliency/v1/unary"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_ResiliencyService_UnaryResiliency_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ResiliencyService_UnaryResiliency_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_ResiliencyService_ServerStreamingResiliency_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		err := status.Error(codes.Unimplemented, "streaming calls are not yet supported in the in-process transport")
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
		return
	})

	mux.Handle("GET", pattern_ResiliencyService_ClientStreamingResiliency_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		err := status.Error(codes.Unimplemented, "streaming calls are not yet supported in the in-process transport")
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
		return
	})

	mux.Handle("GET", pattern_ResiliencyService_BiDirectionalResiliency_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		err := status.Error(codes.Unimplemented, "streaming calls are not yet supported in the in-process transport")
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
		return
	})

	return nil
}

// RegisterResiliencyWithMetadataServiceHandlerServer registers the http handlers for service ResiliencyWithMetadataService to "mux".
// UnaryRPC     :call ResiliencyWithMetadataServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterResiliencyWithMetadataServiceHandlerFromEndpoint instead.
// GRPC interceptors will not work for this type of registration. To use interceptors, you must use the "runtime.WithMiddlewares" option in the "runtime.NewServeMux" call.
func RegisterResiliencyWithMetadataServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server extResiliency.ResiliencyWithMetadataServiceServer) error {

	mux.Handle("POST", pattern_ResiliencyWithMetadataService_UnaryResiliencyWithMetadata_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/resiliency.ResiliencyWithMetadataService/UnaryResiliencyWithMetadata", runtime.WithHTTPPathPattern("/api/resiliency/v1/metadata/unary"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_ResiliencyWithMetadataService_UnaryResiliencyWithMetadata_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ResiliencyWithMetadataService_UnaryResiliencyWithMetadata_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_ResiliencyWithMetadataService_ServerStreamingResiliencyWithMetadata_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		err := status.Error(codes.Unimplemented, "streaming calls are not yet supported in the in-process transport")
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
		return
	})

	mux.Handle("POST", pattern_ResiliencyWithMetadataService_ClientStreamingResiliencyWithMetadata_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		err := status.Error(codes.Unimplemented, "streaming calls are not yet supported in the in-process transport")
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
		return
	})

	mux.Handle("POST", pattern_ResiliencyWithMetadataService_BiDirectionalResiliencyWithMetadata_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		err := status.Error(codes.Unimplemented, "streaming calls are not yet supported in the in-process transport")
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
		return
	})

	return nil
}

// RegisterResiliencyServiceHandlerFromEndpoint is same as RegisterResiliencyServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterResiliencyServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.NewClient(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterResiliencyServiceHandler(ctx, mux, conn)
}

// RegisterResiliencyServiceHandler registers the http handlers for service ResiliencyService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterResiliencyServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterResiliencyServiceHandlerClient(ctx, mux, extResiliency.NewResiliencyServiceClient(conn))
}

// RegisterResiliencyServiceHandlerClient registers the http handlers for service ResiliencyService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "extResiliency.ResiliencyServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "extResiliency.ResiliencyServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "extResiliency.ResiliencyServiceClient" to call the correct interceptors. This client ignores the HTTP middlewares.
func RegisterResiliencyServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client extResiliency.ResiliencyServiceClient) error {

	mux.Handle("GET", pattern_ResiliencyService_UnaryResiliency_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/resiliency.ResiliencyService/UnaryResiliency", runtime.WithHTTPPathPattern("/api/resiliency/v1/unary"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_ResiliencyService_UnaryResiliency_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ResiliencyService_UnaryResiliency_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_ResiliencyService_ServerStreamingResiliency_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/resiliency.ResiliencyService/ServerStreamingResiliency", runtime.WithHTTPPathPattern("/api/resiliency/v1/server_streaming/{min_delay_second}/{max_delay_second}"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_ResiliencyService_ServerStreamingResiliency_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ResiliencyService_ServerStreamingResiliency_0(annotatedContext, mux, outboundMarshaler, w, req, func() (proto.Message, error) { return resp.Recv() }, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_ResiliencyService_ClientStreamingResiliency_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/resiliency.ResiliencyService/ClientStreamingResiliency", runtime.WithHTTPPathPattern("/api/resiliency/v1/client_streaming"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_ResiliencyService_ClientStreamingResiliency_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ResiliencyService_ClientStreamingResiliency_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_ResiliencyService_BiDirectionalResiliency_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/resiliency.ResiliencyService/BiDirectionalResiliency", runtime.WithHTTPPathPattern("/api/resiliency/v1/bidirectional_streaming"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		resp, md, reqErrChan, err := request_ResiliencyService_BiDirectionalResiliency_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		go func() {
			for err := range reqErrChan {
				if err != nil && err != io.EOF {
					runtime.HTTPStreamError(annotatedContext, mux, outboundMarshaler, w, req, err)
				}
			}
		}()

		forward_ResiliencyService_BiDirectionalResiliency_0(annotatedContext, mux, outboundMarshaler, w, req, func() (proto.Message, error) { return resp.Recv() }, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_ResiliencyService_UnaryResiliency_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"api", "resiliency", "v1", "unary"}, ""))

	pattern_ResiliencyService_ServerStreamingResiliency_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3, 1, 0, 4, 1, 5, 4, 1, 0, 4, 1, 5, 5}, []string{"api", "resiliency", "v1", "server_streaming", "min_delay_second", "max_delay_second"}, ""))

	pattern_ResiliencyService_ClientStreamingResiliency_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"api", "resiliency", "v1", "client_streaming"}, ""))

	pattern_ResiliencyService_BiDirectionalResiliency_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3}, []string{"api", "resiliency", "v1", "bidirectional_streaming"}, ""))
)

var (
	forward_ResiliencyService_UnaryResiliency_0 = runtime.ForwardResponseMessage

	forward_ResiliencyService_ServerStreamingResiliency_0 = runtime.ForwardResponseStream

	forward_ResiliencyService_ClientStreamingResiliency_0 = runtime.ForwardResponseMessage

	forward_ResiliencyService_BiDirectionalResiliency_0 = runtime.ForwardResponseStream
)

// RegisterResiliencyWithMetadataServiceHandlerFromEndpoint is same as RegisterResiliencyWithMetadataServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterResiliencyWithMetadataServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.NewClient(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Errorf("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterResiliencyWithMetadataServiceHandler(ctx, mux, conn)
}

// RegisterResiliencyWithMetadataServiceHandler registers the http handlers for service ResiliencyWithMetadataService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterResiliencyWithMetadataServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterResiliencyWithMetadataServiceHandlerClient(ctx, mux, extResiliency.NewResiliencyWithMetadataServiceClient(conn))
}

// RegisterResiliencyWithMetadataServiceHandlerClient registers the http handlers for service ResiliencyWithMetadataService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "extResiliency.ResiliencyWithMetadataServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "extResiliency.ResiliencyWithMetadataServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "extResiliency.ResiliencyWithMetadataServiceClient" to call the correct interceptors. This client ignores the HTTP middlewares.
func RegisterResiliencyWithMetadataServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client extResiliency.ResiliencyWithMetadataServiceClient) error {

	mux.Handle("POST", pattern_ResiliencyWithMetadataService_UnaryResiliencyWithMetadata_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/resiliency.ResiliencyWithMetadataService/UnaryResiliencyWithMetadata", runtime.WithHTTPPathPattern("/api/resiliency/v1/metadata/unary"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_ResiliencyWithMetadataService_UnaryResiliencyWithMetadata_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ResiliencyWithMetadataService_UnaryResiliencyWithMetadata_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_ResiliencyWithMetadataService_ServerStreamingResiliencyWithMetadata_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/resiliency.ResiliencyWithMetadataService/ServerStreamingResiliencyWithMetadata", runtime.WithHTTPPathPattern("/api/resiliency/v1/metadata/server_streaming"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_ResiliencyWithMetadataService_ServerStreamingResiliencyWithMetadata_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ResiliencyWithMetadataService_ServerStreamingResiliencyWithMetadata_0(annotatedContext, mux, outboundMarshaler, w, req, func() (proto.Message, error) { return resp.Recv() }, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_ResiliencyWithMetadataService_ClientStreamingResiliencyWithMetadata_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/resiliency.ResiliencyWithMetadataService/ClientStreamingResiliencyWithMetadata", runtime.WithHTTPPathPattern("/api/resiliency/v1/metadata/client_streaming"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_ResiliencyWithMetadataService_ClientStreamingResiliencyWithMetadata_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_ResiliencyWithMetadataService_ClientStreamingResiliencyWithMetadata_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_ResiliencyWithMetadataService_BiDirectionalResiliencyWithMetadata_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/resiliency.ResiliencyWithMetadataService/BiDirectionalResiliencyWithMetadata", runtime.WithHTTPPathPattern("/api/resiliency/v1/metadata/bidirectional_streaming"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		resp, md, reqErrChan, err := request_ResiliencyWithMetadataService_BiDirectionalResiliencyWithMetadata_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}
		go func() {
			for err := range reqErrChan {
				if err != nil && err != io.EOF {
					runtime.HTTPStreamError(annotatedContext, mux, outboundMarshaler, w, req, err)
				}
			}
		}()

		forward_ResiliencyWithMetadataService_BiDirectionalResiliencyWithMetadata_0(annotatedContext, mux, outboundMarshaler, w, req, func() (proto.Message, error) { return resp.Recv() }, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_ResiliencyWithMetadataService_UnaryResiliencyWithMetadata_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3, 2, 4}, []string{"api", "resiliency", "v1", "metadata", "unary"}, ""))

	pattern_ResiliencyWithMetadataService_ServerStreamingResiliencyWithMetadata_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3, 2, 4}, []string{"api", "resiliency", "v1", "metadata", "server_streaming"}, ""))

	pattern_ResiliencyWithMetadataService_ClientStreamingResiliencyWithMetadata_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3, 2, 4}, []string{"api", "resiliency", "v1", "metadata", "client_streaming"}, ""))

	pattern_ResiliencyWithMetadataService_BiDirectionalResiliencyWithMetadata_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1, 2, 2, 2, 3, 2, 4}, []string{"api", "resiliency", "v1", "metadata", "bidirectional_streaming"}, ""))
)

var (
	forward_ResiliencyWithMetadataService_UnaryResiliencyWithMetadata_0 = runtime.ForwardResponseMessage

	forward_ResiliencyWithMetadataService_ServerStreamingResiliencyWithMetadata_0 = runtime.ForwardResponseStream

	forward_ResiliencyWithMetadataService_ClientStreamingResiliencyWithMetadata_0 = runtime.ForwardResponseMessage

	forward_ResiliencyWithMetadataService_BiDirectionalResiliencyWithMetadata_0 = runtime.ForwardResponseStream
)
