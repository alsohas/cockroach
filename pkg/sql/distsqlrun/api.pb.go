// Code generated by protoc-gen-gogo.
// source: cockroach/pkg/sql/distsqlrun/api.proto
// DO NOT EDIT!

/*
	Package distsqlrun is a generated protocol buffer package.

	It is generated from these files:
		cockroach/pkg/sql/distsqlrun/api.proto
		cockroach/pkg/sql/distsqlrun/data.proto
		cockroach/pkg/sql/distsqlrun/processors.proto

	It has these top-level messages:
		SetupFlowRequest
		EvalContext
		SimpleResponse
		ConsumerSignal
		DrainRequest
		Error
		Expression
		Ordering
		StreamEndpointSpec
		InputSyncSpec
		OutputRouterSpec
		DatumInfo
		ProducerHeader
		ProducerData
		ProducerMessage
		RemoteProducerMetadata
		ProcessorSpec
		PostProcessSpec
		ProcessorCoreUnion
		NoopCoreSpec
		ValuesCoreSpec
		TableReaderSpan
		TableReaderSpec
		JoinReaderSpec
		SorterSpec
		DistinctSpec
		MergeJoinerSpec
		HashJoinerSpec
		AggregatorSpec
		BackfillerSpec
		FlowSpec
		AlgebraicSetOpSpec
*/
package distsqlrun

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import cockroach_roachpb1 "github.com/cockroachdb/cockroach/pkg/roachpb"
import cockroach_util_hlc "github.com/cockroachdb/cockroach/pkg/util/hlc"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type SetupFlowRequest struct {
	Txn cockroach_roachpb1.Transaction `protobuf:"bytes,1,opt,name=txn" json:"txn"`
	// Version of distsqlrun protocol; a server accepts a certain range of
	// versions, up to its own version. See server.go for more details.
	Version     uint32      `protobuf:"varint,5,opt,name=version" json:"version"`
	Flow        FlowSpec    `protobuf:"bytes,3,opt,name=flow" json:"flow"`
	EvalContext EvalContext `protobuf:"bytes,6,opt,name=evalContext" json:"evalContext"`
}

func (m *SetupFlowRequest) Reset()                    { *m = SetupFlowRequest{} }
func (m *SetupFlowRequest) String() string            { return proto.CompactTextString(m) }
func (*SetupFlowRequest) ProtoMessage()               {}
func (*SetupFlowRequest) Descriptor() ([]byte, []int) { return fileDescriptorApi, []int{0} }

// EvalContext is used to marshall some planner.EvalContext members.
type EvalContext struct {
	StmtTimestampNanos int64                        `protobuf:"varint,1,opt,name=stmtTimestampNanos" json:"stmtTimestampNanos"`
	TxnTimestampNanos  int64                        `protobuf:"varint,2,opt,name=txnTimestampNanos" json:"txnTimestampNanos"`
	ClusterTimestamp   cockroach_util_hlc.Timestamp `protobuf:"bytes,3,opt,name=clusterTimestamp" json:"clusterTimestamp"`
	// The name of the location according to whose current timezone we're going to
	// parse timestamps. Used to init EvalContext.Location.
	Location   string   `protobuf:"bytes,4,opt,name=location" json:"location"`
	Database   string   `protobuf:"bytes,5,opt,name=database" json:"database"`
	SearchPath []string `protobuf:"bytes,6,rep,name=searchPath" json:"searchPath,omitempty"`
}

func (m *EvalContext) Reset()                    { *m = EvalContext{} }
func (m *EvalContext) String() string            { return proto.CompactTextString(m) }
func (*EvalContext) ProtoMessage()               {}
func (*EvalContext) Descriptor() ([]byte, []int) { return fileDescriptorApi, []int{1} }

type SimpleResponse struct {
	Error *Error `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
}

func (m *SimpleResponse) Reset()                    { *m = SimpleResponse{} }
func (m *SimpleResponse) String() string            { return proto.CompactTextString(m) }
func (*SimpleResponse) ProtoMessage()               {}
func (*SimpleResponse) Descriptor() ([]byte, []int) { return fileDescriptorApi, []int{2} }

type ConsumerSignal struct {
	// The consumer is done (doesn't need to consume any more rows) and is asking
	// the producer to push whatever trailing metadata it has and close its
	// stream.
	DrainRequest *DrainRequest `protobuf:"bytes,1,opt,name=drain_request,json=drainRequest" json:"drain_request,omitempty"`
	// Used in the RunSyncFlow case; the first message on the client stream must
	// contain this message.
	SetupFlowRequest *SetupFlowRequest `protobuf:"bytes,2,opt,name=setup_flow_request,json=setupFlowRequest" json:"setup_flow_request,omitempty"`
}

func (m *ConsumerSignal) Reset()                    { *m = ConsumerSignal{} }
func (m *ConsumerSignal) String() string            { return proto.CompactTextString(m) }
func (*ConsumerSignal) ProtoMessage()               {}
func (*ConsumerSignal) Descriptor() ([]byte, []int) { return fileDescriptorApi, []int{3} }

type DrainRequest struct {
}

func (m *DrainRequest) Reset()                    { *m = DrainRequest{} }
func (m *DrainRequest) String() string            { return proto.CompactTextString(m) }
func (*DrainRequest) ProtoMessage()               {}
func (*DrainRequest) Descriptor() ([]byte, []int) { return fileDescriptorApi, []int{4} }

func init() {
	proto.RegisterType((*SetupFlowRequest)(nil), "cockroach.sql.distsqlrun.SetupFlowRequest")
	proto.RegisterType((*EvalContext)(nil), "cockroach.sql.distsqlrun.EvalContext")
	proto.RegisterType((*SimpleResponse)(nil), "cockroach.sql.distsqlrun.SimpleResponse")
	proto.RegisterType((*ConsumerSignal)(nil), "cockroach.sql.distsqlrun.ConsumerSignal")
	proto.RegisterType((*DrainRequest)(nil), "cockroach.sql.distsqlrun.DrainRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for DistSQL service

type DistSQLClient interface {
	// RunSyncFlow instantiates a flow and streams back results of that flow.
	// The request must contain one flow, and that flow must have a single mailbox
	// of the special sync response type.
	RunSyncFlow(ctx context.Context, opts ...grpc.CallOption) (DistSQL_RunSyncFlowClient, error)
	// SetupFlow instantiates a flow (subgraphs of a distributed SQL
	// computation) on the receiving node.
	SetupFlow(ctx context.Context, in *SetupFlowRequest, opts ...grpc.CallOption) (*SimpleResponse, error)
	// FlowStream is used to push a stream of messages that is part of a flow. The
	// first message will have a StreamHeader which identifies the flow and the
	// stream (mailbox).
	//
	// The response is a stream because the consumer can signal the producer at
	// any point to start draining.
	FlowStream(ctx context.Context, opts ...grpc.CallOption) (DistSQL_FlowStreamClient, error)
}

type distSQLClient struct {
	cc *grpc.ClientConn
}

func NewDistSQLClient(cc *grpc.ClientConn) DistSQLClient {
	return &distSQLClient{cc}
}

func (c *distSQLClient) RunSyncFlow(ctx context.Context, opts ...grpc.CallOption) (DistSQL_RunSyncFlowClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_DistSQL_serviceDesc.Streams[0], c.cc, "/cockroach.sql.distsqlrun.DistSQL/RunSyncFlow", opts...)
	if err != nil {
		return nil, err
	}
	x := &distSQLRunSyncFlowClient{stream}
	return x, nil
}

type DistSQL_RunSyncFlowClient interface {
	Send(*ConsumerSignal) error
	Recv() (*ProducerMessage, error)
	grpc.ClientStream
}

type distSQLRunSyncFlowClient struct {
	grpc.ClientStream
}

func (x *distSQLRunSyncFlowClient) Send(m *ConsumerSignal) error {
	return x.ClientStream.SendMsg(m)
}

func (x *distSQLRunSyncFlowClient) Recv() (*ProducerMessage, error) {
	m := new(ProducerMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *distSQLClient) SetupFlow(ctx context.Context, in *SetupFlowRequest, opts ...grpc.CallOption) (*SimpleResponse, error) {
	out := new(SimpleResponse)
	err := grpc.Invoke(ctx, "/cockroach.sql.distsqlrun.DistSQL/SetupFlow", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *distSQLClient) FlowStream(ctx context.Context, opts ...grpc.CallOption) (DistSQL_FlowStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_DistSQL_serviceDesc.Streams[1], c.cc, "/cockroach.sql.distsqlrun.DistSQL/FlowStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &distSQLFlowStreamClient{stream}
	return x, nil
}

type DistSQL_FlowStreamClient interface {
	Send(*ProducerMessage) error
	Recv() (*ConsumerSignal, error)
	grpc.ClientStream
}

type distSQLFlowStreamClient struct {
	grpc.ClientStream
}

func (x *distSQLFlowStreamClient) Send(m *ProducerMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *distSQLFlowStreamClient) Recv() (*ConsumerSignal, error) {
	m := new(ConsumerSignal)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for DistSQL service

type DistSQLServer interface {
	// RunSyncFlow instantiates a flow and streams back results of that flow.
	// The request must contain one flow, and that flow must have a single mailbox
	// of the special sync response type.
	RunSyncFlow(DistSQL_RunSyncFlowServer) error
	// SetupFlow instantiates a flow (subgraphs of a distributed SQL
	// computation) on the receiving node.
	SetupFlow(context.Context, *SetupFlowRequest) (*SimpleResponse, error)
	// FlowStream is used to push a stream of messages that is part of a flow. The
	// first message will have a StreamHeader which identifies the flow and the
	// stream (mailbox).
	//
	// The response is a stream because the consumer can signal the producer at
	// any point to start draining.
	FlowStream(DistSQL_FlowStreamServer) error
}

func RegisterDistSQLServer(s *grpc.Server, srv DistSQLServer) {
	s.RegisterService(&_DistSQL_serviceDesc, srv)
}

func _DistSQL_RunSyncFlow_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DistSQLServer).RunSyncFlow(&distSQLRunSyncFlowServer{stream})
}

type DistSQL_RunSyncFlowServer interface {
	Send(*ProducerMessage) error
	Recv() (*ConsumerSignal, error)
	grpc.ServerStream
}

type distSQLRunSyncFlowServer struct {
	grpc.ServerStream
}

func (x *distSQLRunSyncFlowServer) Send(m *ProducerMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *distSQLRunSyncFlowServer) Recv() (*ConsumerSignal, error) {
	m := new(ConsumerSignal)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _DistSQL_SetupFlow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetupFlowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistSQLServer).SetupFlow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cockroach.sql.distsqlrun.DistSQL/SetupFlow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistSQLServer).SetupFlow(ctx, req.(*SetupFlowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DistSQL_FlowStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DistSQLServer).FlowStream(&distSQLFlowStreamServer{stream})
}

type DistSQL_FlowStreamServer interface {
	Send(*ConsumerSignal) error
	Recv() (*ProducerMessage, error)
	grpc.ServerStream
}

type distSQLFlowStreamServer struct {
	grpc.ServerStream
}

func (x *distSQLFlowStreamServer) Send(m *ConsumerSignal) error {
	return x.ServerStream.SendMsg(m)
}

func (x *distSQLFlowStreamServer) Recv() (*ProducerMessage, error) {
	m := new(ProducerMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _DistSQL_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cockroach.sql.distsqlrun.DistSQL",
	HandlerType: (*DistSQLServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetupFlow",
			Handler:    _DistSQL_SetupFlow_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RunSyncFlow",
			Handler:       _DistSQL_RunSyncFlow_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "FlowStream",
			Handler:       _DistSQL_FlowStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "cockroach/pkg/sql/distsqlrun/api.proto",
}

func (m *SetupFlowRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SetupFlowRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintApi(dAtA, i, uint64(m.Txn.Size()))
	n1, err := m.Txn.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	dAtA[i] = 0x1a
	i++
	i = encodeVarintApi(dAtA, i, uint64(m.Flow.Size()))
	n2, err := m.Flow.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	dAtA[i] = 0x28
	i++
	i = encodeVarintApi(dAtA, i, uint64(m.Version))
	dAtA[i] = 0x32
	i++
	i = encodeVarintApi(dAtA, i, uint64(m.EvalContext.Size()))
	n3, err := m.EvalContext.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	return i, nil
}

func (m *EvalContext) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EvalContext) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintApi(dAtA, i, uint64(m.StmtTimestampNanos))
	dAtA[i] = 0x10
	i++
	i = encodeVarintApi(dAtA, i, uint64(m.TxnTimestampNanos))
	dAtA[i] = 0x1a
	i++
	i = encodeVarintApi(dAtA, i, uint64(m.ClusterTimestamp.Size()))
	n4, err := m.ClusterTimestamp.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	dAtA[i] = 0x22
	i++
	i = encodeVarintApi(dAtA, i, uint64(len(m.Location)))
	i += copy(dAtA[i:], m.Location)
	dAtA[i] = 0x2a
	i++
	i = encodeVarintApi(dAtA, i, uint64(len(m.Database)))
	i += copy(dAtA[i:], m.Database)
	if len(m.SearchPath) > 0 {
		for _, s := range m.SearchPath {
			dAtA[i] = 0x32
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}

func (m *SimpleResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SimpleResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Error != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintApi(dAtA, i, uint64(m.Error.Size()))
		n5, err := m.Error.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	return i, nil
}

func (m *ConsumerSignal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConsumerSignal) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.DrainRequest != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintApi(dAtA, i, uint64(m.DrainRequest.Size()))
		n6, err := m.DrainRequest.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	if m.SetupFlowRequest != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintApi(dAtA, i, uint64(m.SetupFlowRequest.Size()))
		n7, err := m.SetupFlowRequest.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n7
	}
	return i, nil
}

func (m *DrainRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DrainRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func encodeFixed64Api(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Api(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintApi(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *SetupFlowRequest) Size() (n int) {
	var l int
	_ = l
	l = m.Txn.Size()
	n += 1 + l + sovApi(uint64(l))
	l = m.Flow.Size()
	n += 1 + l + sovApi(uint64(l))
	n += 1 + sovApi(uint64(m.Version))
	l = m.EvalContext.Size()
	n += 1 + l + sovApi(uint64(l))
	return n
}

func (m *EvalContext) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovApi(uint64(m.StmtTimestampNanos))
	n += 1 + sovApi(uint64(m.TxnTimestampNanos))
	l = m.ClusterTimestamp.Size()
	n += 1 + l + sovApi(uint64(l))
	l = len(m.Location)
	n += 1 + l + sovApi(uint64(l))
	l = len(m.Database)
	n += 1 + l + sovApi(uint64(l))
	if len(m.SearchPath) > 0 {
		for _, s := range m.SearchPath {
			l = len(s)
			n += 1 + l + sovApi(uint64(l))
		}
	}
	return n
}

func (m *SimpleResponse) Size() (n int) {
	var l int
	_ = l
	if m.Error != nil {
		l = m.Error.Size()
		n += 1 + l + sovApi(uint64(l))
	}
	return n
}

func (m *ConsumerSignal) Size() (n int) {
	var l int
	_ = l
	if m.DrainRequest != nil {
		l = m.DrainRequest.Size()
		n += 1 + l + sovApi(uint64(l))
	}
	if m.SetupFlowRequest != nil {
		l = m.SetupFlowRequest.Size()
		n += 1 + l + sovApi(uint64(l))
	}
	return n
}

func (m *DrainRequest) Size() (n int) {
	var l int
	_ = l
	return n
}

func sovApi(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozApi(x uint64) (n int) {
	return sovApi(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ConsumerSignal) GetValue() interface{} {
	if this.DrainRequest != nil {
		return this.DrainRequest
	}
	if this.SetupFlowRequest != nil {
		return this.SetupFlowRequest
	}
	return nil
}

func (this *ConsumerSignal) SetValue(value interface{}) bool {
	switch vt := value.(type) {
	case *DrainRequest:
		this.DrainRequest = vt
	case *SetupFlowRequest:
		this.SetupFlowRequest = vt
	default:
		return false
	}
	return true
}
func (m *SetupFlowRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SetupFlowRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SetupFlowRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Txn", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Txn.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Flow", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Flow.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EvalContext", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.EvalContext.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *EvalContext) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: EvalContext: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EvalContext: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StmtTimestampNanos", wireType)
			}
			m.StmtTimestampNanos = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StmtTimestampNanos |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxnTimestampNanos", wireType)
			}
			m.TxnTimestampNanos = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TxnTimestampNanos |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClusterTimestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ClusterTimestamp.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Location", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Location = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Database", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Database = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SearchPath", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SearchPath = append(m.SearchPath, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SimpleResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SimpleResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SimpleResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Error == nil {
				m.Error = &Error{}
			}
			if err := m.Error.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ConsumerSignal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ConsumerSignal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConsumerSignal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DrainRequest", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DrainRequest == nil {
				m.DrainRequest = &DrainRequest{}
			}
			if err := m.DrainRequest.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SetupFlowRequest", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApi
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthApi
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SetupFlowRequest == nil {
				m.SetupFlowRequest = &SetupFlowRequest{}
			}
			if err := m.SetupFlowRequest.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DrainRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApi
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DrainRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DrainRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipApi(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApi
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipApi(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowApi
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowApi
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowApi
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthApi
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowApi
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipApi(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthApi = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowApi   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("cockroach/pkg/sql/distsqlrun/api.proto", fileDescriptorApi) }

var fileDescriptorApi = []byte{
	// 631 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0x86, 0x33, 0x49, 0xfa, 0x77, 0xd2, 0x56, 0xf9, 0x46, 0xdf, 0xc2, 0x8a, 0x84, 0x1b, 0x59,
	0x50, 0x02, 0x12, 0x36, 0xaa, 0x80, 0x05, 0x62, 0xd5, 0x16, 0x90, 0x80, 0x42, 0x71, 0xba, 0x40,
	0x6c, 0xaa, 0xe9, 0x64, 0x48, 0xac, 0x4e, 0x66, 0x9c, 0x99, 0x71, 0x5b, 0xee, 0x82, 0x4b, 0xe0,
	0x02, 0xb8, 0x05, 0x16, 0xec, 0xb2, 0x64, 0xc9, 0x0a, 0x41, 0xb8, 0x0e, 0x24, 0x64, 0x67, 0xea,
	0xb8, 0x29, 0x89, 0xca, 0x6e, 0x7c, 0xe6, 0x79, 0xdf, 0x33, 0xe7, 0xc7, 0xb0, 0x49, 0x25, 0x3d,
	0x56, 0x92, 0xd0, 0x5e, 0x10, 0x1f, 0x77, 0x03, 0x3d, 0xe0, 0x41, 0x27, 0xd2, 0x46, 0x0f, 0xb8,
	0x4a, 0x44, 0x40, 0xe2, 0xc8, 0x8f, 0x95, 0x34, 0x12, 0x3b, 0x39, 0xe7, 0xeb, 0x01, 0xf7, 0x27,
	0x4c, 0xa3, 0x79, 0xd1, 0x21, 0x3b, 0xc5, 0x47, 0x41, 0x87, 0x18, 0x32, 0xd6, 0x36, 0x6e, 0xce,
	0xcd, 0x51, 0x00, 0xef, 0xcc, 0x05, 0x63, 0x25, 0x29, 0xd3, 0x5a, 0x2a, 0x6d, 0xf1, 0xa9, 0xb7,
	0x27, 0x26, 0xe2, 0x41, 0x8f, 0xd3, 0xc0, 0x44, 0x7d, 0xa6, 0x0d, 0xe9, 0xc7, 0x96, 0xfb, 0xbf,
	0x2b, 0xbb, 0x32, 0x3b, 0x06, 0xe9, 0x69, 0x1c, 0xf5, 0x7e, 0x23, 0xa8, 0xb7, 0x99, 0x49, 0xe2,
	0x27, 0x5c, 0x9e, 0x86, 0x6c, 0x90, 0x30, 0x6d, 0xf0, 0x03, 0xa8, 0x98, 0x33, 0xe1, 0xa0, 0x26,
	0x6a, 0xd5, 0xb6, 0x5c, 0x7f, 0x52, 0xb4, 0x2d, 0xcb, 0x3f, 0x50, 0x44, 0x68, 0x42, 0x4d, 0x24,
	0xc5, 0x76, 0x75, 0xf8, 0x7d, 0xa3, 0x14, 0xa6, 0x02, 0xfc, 0x08, 0xaa, 0xef, 0xb8, 0x3c, 0x75,
	0x2a, 0x99, 0xd0, 0xf3, 0x67, 0x75, 0xcb, 0x4f, 0x93, 0xb5, 0x63, 0x46, 0xad, 0x38, 0x53, 0x61,
	0x17, 0x96, 0x4e, 0x98, 0xd2, 0x91, 0x14, 0xce, 0x42, 0x13, 0xb5, 0xd6, 0xec, 0xe5, 0x79, 0x10,
	0xef, 0x41, 0x8d, 0x9d, 0x10, 0xbe, 0x23, 0x85, 0x61, 0x67, 0xc6, 0x59, 0xcc, 0x92, 0xdc, 0x98,
	0x9d, 0xe4, 0xf1, 0x04, 0xb6, 0x56, 0x45, 0xfd, 0xb3, 0xea, 0x72, 0xb9, 0x5e, 0xf1, 0x3e, 0x95,
	0xa1, 0x56, 0x00, 0xf1, 0x3d, 0xc0, 0xda, 0xf4, 0xcd, 0xc1, 0x79, 0xf3, 0x5e, 0x12, 0x21, 0x75,
	0xd6, 0x89, 0x8a, 0x35, 0xf9, 0xcb, 0x3d, 0xde, 0x82, 0xff, 0xcc, 0x99, 0x98, 0x12, 0x95, 0x0b,
	0xa2, 0xcb, 0xd7, 0xf8, 0x15, 0xd4, 0x29, 0x4f, 0xb4, 0x61, 0x2a, 0xbf, 0xb0, 0x8d, 0xbb, 0x56,
	0xa8, 0x29, 0x1d, 0xa7, 0xdf, 0xe3, 0xd4, 0xcf, 0x21, 0xeb, 0x78, 0x49, 0x8c, 0x9b, 0xb0, 0xcc,
	0x25, 0x25, 0xe9, 0x50, 0x9c, 0x6a, 0x13, 0xb5, 0x56, 0x2c, 0x99, 0x47, 0x53, 0x22, 0xdd, 0xb3,
	0x23, 0xa2, 0x59, 0xd6, 0xe2, 0x9c, 0x38, 0x8f, 0x62, 0x17, 0x40, 0x33, 0xa2, 0x68, 0x6f, 0x9f,
	0x98, 0x9e, 0xb3, 0xd8, 0xac, 0xb4, 0x56, 0xc2, 0x42, 0xc4, 0x7b, 0x0a, 0xeb, 0xed, 0xa8, 0x1f,
	0x73, 0x16, 0x32, 0x1d, 0x4b, 0xa1, 0x19, 0xbe, 0x0f, 0x0b, 0x4c, 0x29, 0xa9, 0xec, 0xb6, 0x6c,
	0xcc, 0x99, 0x47, 0x8a, 0x85, 0x63, 0xda, 0xfb, 0x8c, 0x60, 0x7d, 0x47, 0x0a, 0x9d, 0xf4, 0x99,
	0x6a, 0x47, 0x5d, 0x41, 0x38, 0x7e, 0x0e, 0x6b, 0x1d, 0x45, 0x22, 0x71, 0xa8, 0xc6, 0x6b, 0x68,
	0x1d, 0x37, 0x67, 0x3b, 0xee, 0xa6, 0xb8, 0x5d, 0xda, 0x70, 0xb5, 0x53, 0xf8, 0xc2, 0x6f, 0x00,
	0xeb, 0x74, 0xad, 0x0f, 0xd3, 0xd5, 0xca, 0x1d, 0xcb, 0x99, 0xe3, 0xed, 0xd9, 0x8e, 0xd3, 0xbf,
	0x42, 0x58, 0xd7, 0x53, 0x91, 0x87, 0xd5, 0xe1, 0xc7, 0x0d, 0xe4, 0xad, 0xc3, 0x6a, 0x31, 0xfb,
	0xd6, 0x97, 0x32, 0x2c, 0xed, 0x46, 0xda, 0xb4, 0x5f, 0xbf, 0xc0, 0x3d, 0xa8, 0x85, 0x89, 0x68,
	0xbf, 0x17, 0x34, 0xd5, 0xe1, 0xd6, 0xec, 0x74, 0x17, 0x3b, 0xd0, 0xb8, 0x35, 0x9b, 0xdc, 0x57,
	0xb2, 0x93, 0x50, 0xa6, 0xf6, 0x98, 0xd6, 0xa4, 0xcb, 0xbc, 0x52, 0x0b, 0xdd, 0x45, 0x98, 0xc2,
	0x4a, 0xfe, 0x62, 0xfc, 0x0f, 0x65, 0x35, 0xe6, 0xbc, 0xe9, 0xe2, 0x7c, 0xbd, 0x12, 0xee, 0x02,
	0x64, 0xff, 0xab, 0x51, 0x8c, 0xf4, 0xf1, 0xd5, 0xdf, 0xd8, 0xb8, 0x72, 0xe1, 0xe3, 0x6a, 0xb6,
	0xaf, 0x0f, 0x7f, 0xba, 0xa5, 0xe1, 0xc8, 0x45, 0x5f, 0x47, 0x2e, 0xfa, 0x36, 0x72, 0xd1, 0x8f,
	0x91, 0x8b, 0x3e, 0xfc, 0x72, 0x4b, 0x6f, 0x61, 0x22, 0xfd, 0x13, 0x00, 0x00, 0xff, 0xff, 0x42,
	0xc4, 0x83, 0x8f, 0xac, 0x05, 0x00, 0x00,
}
