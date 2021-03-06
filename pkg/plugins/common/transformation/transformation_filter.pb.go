// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: transformation_filter.proto

/*
Package transformation is a generated protocol buffer package.

It is generated from these files:
	transformation_filter.proto

It has these top-level messages:
	Transformations
	Transformation
	RouteTransformations
	Extraction
	TransformationTemplate
	InjaTemplate
	Passthrough
	MergeExtractorsToBody
	HeaderBodyTransform
*/
package transformation

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Transformations struct {
	Transformations map[string]*Transformation `protobuf:"bytes,1,rep,name=transformations" json:"transformations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
	// if use_routes_for_config is set transformations will be ignored. the plan
	// is to eventually deprecate it. Ideally they should be in a oneof, but a map
	// can't be in a oneof...
	UseRoutesForConfig bool `protobuf:"varint,2,opt,name=use_routes_for_config,json=useRoutesForConfig,proto3" json:"use_routes_for_config,omitempty"`
}

func (m *Transformations) Reset()         { *m = Transformations{} }
func (m *Transformations) String() string { return proto.CompactTextString(m) }
func (*Transformations) ProtoMessage()    {}
func (*Transformations) Descriptor() ([]byte, []int) {
	return fileDescriptorTransformationFilter, []int{0}
}

func (m *Transformations) GetTransformations() map[string]*Transformation {
	if m != nil {
		return m.Transformations
	}
	return nil
}

func (m *Transformations) GetUseRoutesForConfig() bool {
	if m != nil {
		return m.UseRoutesForConfig
	}
	return false
}

// [#proto-status: experimental]
type Transformation struct {
	// Template is in the transformed request language domain
	// currently both are JSON
	//
	// Types that are valid to be assigned to TransformationType:
	//	*Transformation_TransformationTemplate
	//	*Transformation_HeaderBodyTransform
	TransformationType isTransformation_TransformationType `protobuf_oneof:"transformation_type"`
}

func (m *Transformation) Reset()         { *m = Transformation{} }
func (m *Transformation) String() string { return proto.CompactTextString(m) }
func (*Transformation) ProtoMessage()    {}
func (*Transformation) Descriptor() ([]byte, []int) {
	return fileDescriptorTransformationFilter, []int{1}
}

type isTransformation_TransformationType interface {
	isTransformation_TransformationType()
}

type Transformation_TransformationTemplate struct {
	TransformationTemplate *TransformationTemplate `protobuf:"bytes,1,opt,name=transformation_template,json=transformationTemplate,oneof"`
}
type Transformation_HeaderBodyTransform struct {
	HeaderBodyTransform *HeaderBodyTransform `protobuf:"bytes,2,opt,name=header_body_transform,json=headerBodyTransform,oneof"`
}

func (*Transformation_TransformationTemplate) isTransformation_TransformationType() {}
func (*Transformation_HeaderBodyTransform) isTransformation_TransformationType()    {}

func (m *Transformation) GetTransformationType() isTransformation_TransformationType {
	if m != nil {
		return m.TransformationType
	}
	return nil
}

func (m *Transformation) GetTransformationTemplate() *TransformationTemplate {
	if x, ok := m.GetTransformationType().(*Transformation_TransformationTemplate); ok {
		return x.TransformationTemplate
	}
	return nil
}

func (m *Transformation) GetHeaderBodyTransform() *HeaderBodyTransform {
	if x, ok := m.GetTransformationType().(*Transformation_HeaderBodyTransform); ok {
		return x.HeaderBodyTransform
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Transformation) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Transformation_OneofMarshaler, _Transformation_OneofUnmarshaler, _Transformation_OneofSizer, []interface{}{
		(*Transformation_TransformationTemplate)(nil),
		(*Transformation_HeaderBodyTransform)(nil),
	}
}

func _Transformation_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Transformation)
	// transformation_type
	switch x := m.TransformationType.(type) {
	case *Transformation_TransformationTemplate:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TransformationTemplate); err != nil {
			return err
		}
	case *Transformation_HeaderBodyTransform:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.HeaderBodyTransform); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Transformation.TransformationType has unexpected type %T", x)
	}
	return nil
}

func _Transformation_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Transformation)
	switch tag {
	case 1: // transformation_type.transformation_template
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TransformationTemplate)
		err := b.DecodeMessage(msg)
		m.TransformationType = &Transformation_TransformationTemplate{msg}
		return true, err
	case 2: // transformation_type.header_body_transform
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(HeaderBodyTransform)
		err := b.DecodeMessage(msg)
		m.TransformationType = &Transformation_HeaderBodyTransform{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Transformation_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Transformation)
	// transformation_type
	switch x := m.TransformationType.(type) {
	case *Transformation_TransformationTemplate:
		s := proto.Size(x.TransformationTemplate)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Transformation_HeaderBodyTransform:
		s := proto.Size(x.HeaderBodyTransform)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type RouteTransformations struct {
	RequestTransformation  *Transformation `protobuf:"bytes,1,opt,name=request_transformation,json=requestTransformation" json:"request_transformation,omitempty"`
	ResponseTransformation *Transformation `protobuf:"bytes,2,opt,name=response_transformation,json=responseTransformation" json:"response_transformation,omitempty"`
}

func (m *RouteTransformations) Reset()         { *m = RouteTransformations{} }
func (m *RouteTransformations) String() string { return proto.CompactTextString(m) }
func (*RouteTransformations) ProtoMessage()    {}
func (*RouteTransformations) Descriptor() ([]byte, []int) {
	return fileDescriptorTransformationFilter, []int{2}
}

func (m *RouteTransformations) GetRequestTransformation() *Transformation {
	if m != nil {
		return m.RequestTransformation
	}
	return nil
}

func (m *RouteTransformations) GetResponseTransformation() *Transformation {
	if m != nil {
		return m.ResponseTransformation
	}
	return nil
}

type Extraction struct {
	Header string `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// what information to extract. if extraction fails the result is
	// an empty value.
	Regex    string `protobuf:"bytes,2,opt,name=regex,proto3" json:"regex,omitempty"`
	Subgroup uint32 `protobuf:"varint,3,opt,name=subgroup,proto3" json:"subgroup,omitempty"`
}

func (m *Extraction) Reset()                    { *m = Extraction{} }
func (m *Extraction) String() string            { return proto.CompactTextString(m) }
func (*Extraction) ProtoMessage()               {}
func (*Extraction) Descriptor() ([]byte, []int) { return fileDescriptorTransformationFilter, []int{3} }

func (m *Extraction) GetHeader() string {
	if m != nil {
		return m.Header
	}
	return ""
}

func (m *Extraction) GetRegex() string {
	if m != nil {
		return m.Regex
	}
	return ""
}

func (m *Extraction) GetSubgroup() uint32 {
	if m != nil {
		return m.Subgroup
	}
	return 0
}

type TransformationTemplate struct {
	AdvancedTemplates bool `protobuf:"varint,1,opt,name=advanced_templates,json=advancedTemplates,proto3" json:"advanced_templates,omitempty"`
	// Extractors are in the origin request language domain
	Extractors map[string]*Extraction   `protobuf:"bytes,2,rep,name=extractors" json:"extractors,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
	Headers    map[string]*InjaTemplate `protobuf:"bytes,3,rep,name=headers" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
	// Types that are valid to be assigned to BodyTransformation:
	//	*TransformationTemplate_Body
	//	*TransformationTemplate_Passthrough
	//	*TransformationTemplate_MergeExtractorsToBody
	BodyTransformation isTransformationTemplate_BodyTransformation `protobuf_oneof:"body_transformation"`
}

func (m *TransformationTemplate) Reset()         { *m = TransformationTemplate{} }
func (m *TransformationTemplate) String() string { return proto.CompactTextString(m) }
func (*TransformationTemplate) ProtoMessage()    {}
func (*TransformationTemplate) Descriptor() ([]byte, []int) {
	return fileDescriptorTransformationFilter, []int{4}
}

type isTransformationTemplate_BodyTransformation interface {
	isTransformationTemplate_BodyTransformation()
}

type TransformationTemplate_Body struct {
	Body *InjaTemplate `protobuf:"bytes,4,opt,name=body,oneof"`
}
type TransformationTemplate_Passthrough struct {
	Passthrough *Passthrough `protobuf:"bytes,5,opt,name=passthrough,oneof"`
}
type TransformationTemplate_MergeExtractorsToBody struct {
	MergeExtractorsToBody *MergeExtractorsToBody `protobuf:"bytes,6,opt,name=merge_extractors_to_body,json=mergeExtractorsToBody,oneof"`
}

func (*TransformationTemplate_Body) isTransformationTemplate_BodyTransformation()                  {}
func (*TransformationTemplate_Passthrough) isTransformationTemplate_BodyTransformation()           {}
func (*TransformationTemplate_MergeExtractorsToBody) isTransformationTemplate_BodyTransformation() {}

func (m *TransformationTemplate) GetBodyTransformation() isTransformationTemplate_BodyTransformation {
	if m != nil {
		return m.BodyTransformation
	}
	return nil
}

func (m *TransformationTemplate) GetAdvancedTemplates() bool {
	if m != nil {
		return m.AdvancedTemplates
	}
	return false
}

func (m *TransformationTemplate) GetExtractors() map[string]*Extraction {
	if m != nil {
		return m.Extractors
	}
	return nil
}

func (m *TransformationTemplate) GetHeaders() map[string]*InjaTemplate {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *TransformationTemplate) GetBody() *InjaTemplate {
	if x, ok := m.GetBodyTransformation().(*TransformationTemplate_Body); ok {
		return x.Body
	}
	return nil
}

func (m *TransformationTemplate) GetPassthrough() *Passthrough {
	if x, ok := m.GetBodyTransformation().(*TransformationTemplate_Passthrough); ok {
		return x.Passthrough
	}
	return nil
}

func (m *TransformationTemplate) GetMergeExtractorsToBody() *MergeExtractorsToBody {
	if x, ok := m.GetBodyTransformation().(*TransformationTemplate_MergeExtractorsToBody); ok {
		return x.MergeExtractorsToBody
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*TransformationTemplate) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _TransformationTemplate_OneofMarshaler, _TransformationTemplate_OneofUnmarshaler, _TransformationTemplate_OneofSizer, []interface{}{
		(*TransformationTemplate_Body)(nil),
		(*TransformationTemplate_Passthrough)(nil),
		(*TransformationTemplate_MergeExtractorsToBody)(nil),
	}
}

func _TransformationTemplate_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*TransformationTemplate)
	// body_transformation
	switch x := m.BodyTransformation.(type) {
	case *TransformationTemplate_Body:
		_ = b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Body); err != nil {
			return err
		}
	case *TransformationTemplate_Passthrough:
		_ = b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Passthrough); err != nil {
			return err
		}
	case *TransformationTemplate_MergeExtractorsToBody:
		_ = b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.MergeExtractorsToBody); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("TransformationTemplate.BodyTransformation has unexpected type %T", x)
	}
	return nil
}

func _TransformationTemplate_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*TransformationTemplate)
	switch tag {
	case 4: // body_transformation.body
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(InjaTemplate)
		err := b.DecodeMessage(msg)
		m.BodyTransformation = &TransformationTemplate_Body{msg}
		return true, err
	case 5: // body_transformation.passthrough
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Passthrough)
		err := b.DecodeMessage(msg)
		m.BodyTransformation = &TransformationTemplate_Passthrough{msg}
		return true, err
	case 6: // body_transformation.merge_extractors_to_body
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(MergeExtractorsToBody)
		err := b.DecodeMessage(msg)
		m.BodyTransformation = &TransformationTemplate_MergeExtractorsToBody{msg}
		return true, err
	default:
		return false, nil
	}
}

func _TransformationTemplate_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*TransformationTemplate)
	// body_transformation
	switch x := m.BodyTransformation.(type) {
	case *TransformationTemplate_Body:
		s := proto.Size(x.Body)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *TransformationTemplate_Passthrough:
		s := proto.Size(x.Passthrough)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *TransformationTemplate_MergeExtractorsToBody:
		s := proto.Size(x.MergeExtractorsToBody)
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

//
// custom functions:
// header_value(name) -> from the original headers
// extracted_value(name, index) -> from the extracted values
type InjaTemplate struct {
	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (m *InjaTemplate) Reset()                    { *m = InjaTemplate{} }
func (m *InjaTemplate) String() string            { return proto.CompactTextString(m) }
func (*InjaTemplate) ProtoMessage()               {}
func (*InjaTemplate) Descriptor() ([]byte, []int) { return fileDescriptorTransformationFilter, []int{5} }

func (m *InjaTemplate) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type Passthrough struct {
}

func (m *Passthrough) Reset()                    { *m = Passthrough{} }
func (m *Passthrough) String() string            { return proto.CompactTextString(m) }
func (*Passthrough) ProtoMessage()               {}
func (*Passthrough) Descriptor() ([]byte, []int) { return fileDescriptorTransformationFilter, []int{6} }

type MergeExtractorsToBody struct {
}

func (m *MergeExtractorsToBody) Reset()         { *m = MergeExtractorsToBody{} }
func (m *MergeExtractorsToBody) String() string { return proto.CompactTextString(m) }
func (*MergeExtractorsToBody) ProtoMessage()    {}
func (*MergeExtractorsToBody) Descriptor() ([]byte, []int) {
	return fileDescriptorTransformationFilter, []int{7}
}

type HeaderBodyTransform struct {
}

func (m *HeaderBodyTransform) Reset()         { *m = HeaderBodyTransform{} }
func (m *HeaderBodyTransform) String() string { return proto.CompactTextString(m) }
func (*HeaderBodyTransform) ProtoMessage()    {}
func (*HeaderBodyTransform) Descriptor() ([]byte, []int) {
	return fileDescriptorTransformationFilter, []int{8}
}

func init() {
	proto.RegisterType((*Transformations)(nil), "envoy.api.v2.filter.http.Transformations")
	proto.RegisterType((*Transformation)(nil), "envoy.api.v2.filter.http.Transformation")
	proto.RegisterType((*RouteTransformations)(nil), "envoy.api.v2.filter.http.RouteTransformations")
	proto.RegisterType((*Extraction)(nil), "envoy.api.v2.filter.http.Extraction")
	proto.RegisterType((*TransformationTemplate)(nil), "envoy.api.v2.filter.http.TransformationTemplate")
	proto.RegisterType((*InjaTemplate)(nil), "envoy.api.v2.filter.http.InjaTemplate")
	proto.RegisterType((*Passthrough)(nil), "envoy.api.v2.filter.http.Passthrough")
	proto.RegisterType((*MergeExtractorsToBody)(nil), "envoy.api.v2.filter.http.MergeExtractorsToBody")
	proto.RegisterType((*HeaderBodyTransform)(nil), "envoy.api.v2.filter.http.HeaderBodyTransform")
}

func init() { proto.RegisterFile("transformation_filter.proto", fileDescriptorTransformationFilter) }

var fileDescriptorTransformationFilter = []byte{
	// 624 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x95, 0xdf, 0x8e, 0xd2, 0x40,
	0x14, 0xc6, 0x29, 0x2c, 0xeb, 0xee, 0x61, 0xff, 0x39, 0x50, 0x68, 0xf0, 0x86, 0x34, 0x6a, 0xb8,
	0xd9, 0xaa, 0x78, 0x63, 0x36, 0xeb, 0xc6, 0x60, 0x30, 0xec, 0x85, 0x89, 0x99, 0x10, 0x4d, 0xbc,
	0xa9, 0x03, 0x0c, 0x94, 0x5d, 0xe8, 0xd4, 0x99, 0x29, 0xa1, 0xef, 0xe2, 0x7b, 0x99, 0xf8, 0x26,
	0xde, 0x19, 0xa6, 0x05, 0xda, 0xda, 0x26, 0x70, 0xd7, 0xe9, 0x99, 0xf3, 0xfb, 0x3a, 0xdf, 0x9c,
	0x73, 0x0a, 0xcf, 0x24, 0x27, 0xae, 0x98, 0x30, 0xbe, 0x20, 0x72, 0xc6, 0x5c, 0x7b, 0x32, 0x9b,
	0x4b, 0xca, 0x2d, 0x8f, 0x33, 0xc9, 0x90, 0x41, 0xdd, 0x25, 0x0b, 0x2c, 0xe2, 0xcd, 0xac, 0x65,
	0xc7, 0x8a, 0x42, 0x8e, 0x94, 0x9e, 0xf9, 0xab, 0x08, 0x97, 0x83, 0x44, 0xa6, 0x40, 0x0e, 0x5c,
	0x26, 0x61, 0xc2, 0xd0, 0x5a, 0xa5, 0x76, 0xa5, 0x73, 0x67, 0xe5, 0x71, 0xac, 0x14, 0x23, 0xbd,
	0xee, 0xb9, 0x92, 0x07, 0x38, 0x8d, 0x45, 0x6f, 0x40, 0xf7, 0x05, 0xb5, 0x39, 0xf3, 0x25, 0x15,
	0xf6, 0x84, 0x71, 0x7b, 0xc4, 0xdc, 0xc9, 0x6c, 0x6a, 0x14, 0x5b, 0x5a, 0xfb, 0x04, 0x23, 0x5f,
	0x50, 0xac, 0x62, 0x9f, 0x18, 0xff, 0xa8, 0x22, 0xcd, 0x39, 0xd4, 0xb2, 0xd8, 0xe8, 0x0a, 0x4a,
	0x8f, 0x34, 0x30, 0xb4, 0x96, 0xd6, 0x3e, 0xc5, 0xeb, 0x47, 0x74, 0x07, 0xe5, 0x25, 0x99, 0xfb,
	0x54, 0xc1, 0x2a, 0x9d, 0xf6, 0xbe, 0x1f, 0x8f, 0xc3, 0xb4, 0x9b, 0xe2, 0x3b, 0xcd, 0xfc, 0xab,
	0xc1, 0x45, 0x32, 0x8a, 0x1e, 0xa1, 0x91, 0xb2, 0x5a, 0xd2, 0x85, 0x37, 0x27, 0x92, 0x2a, 0xf1,
	0x4a, 0xe7, 0xf5, 0xbe, 0x42, 0x83, 0x28, 0xaf, 0x5f, 0xc0, 0x75, 0x99, 0x19, 0x41, 0x23, 0xd0,
	0x1d, 0x4a, 0xc6, 0x94, 0xdb, 0x43, 0x36, 0x0e, 0xec, 0xed, 0xae, 0xe8, 0x4c, 0xd7, 0xf9, 0x52,
	0x7d, 0x95, 0xd6, 0x65, 0xe3, 0x60, 0x2b, 0xda, 0x2f, 0xe0, 0xaa, 0xf3, 0xff, 0xeb, 0xae, 0x0e,
	0xd5, 0xf4, 0x89, 0x02, 0x8f, 0x9a, 0x7f, 0x34, 0xa8, 0x29, 0xf7, 0xd3, 0xf5, 0x61, 0x43, 0x9d,
	0xd3, 0x9f, 0x3e, 0x15, 0xd2, 0x4e, 0xe6, 0x45, 0x06, 0xec, 0xef, 0xb4, 0x1e, 0x71, 0x52, 0x16,
	0x13, 0x68, 0x70, 0x2a, 0x3c, 0xe6, 0x0a, 0x9a, 0x56, 0x38, 0xf4, 0x2e, 0xeb, 0x1b, 0x50, 0xf2,
	0xbd, 0xf9, 0x15, 0xa0, 0xb7, 0x92, 0x9c, 0x8c, 0x94, 0x60, 0x1d, 0x8e, 0x43, 0x63, 0xa2, 0xfa,
	0x89, 0x56, 0xa8, 0x06, 0x65, 0x4e, 0xa7, 0x74, 0xa5, 0x64, 0x4f, 0x71, 0xb8, 0x40, 0x4d, 0x38,
	0x11, 0xfe, 0x70, 0xca, 0x99, 0xef, 0x19, 0xa5, 0x96, 0xd6, 0x3e, 0xc7, 0xdb, 0xb5, 0xf9, 0xbb,
	0x0c, 0xf5, 0xec, 0x5b, 0x46, 0xd7, 0x80, 0xc8, 0x78, 0x49, 0xdc, 0x11, 0x1d, 0x6f, 0x4b, 0x46,
	0x28, 0xc1, 0x13, 0xfc, 0x74, 0x13, 0xd9, 0xec, 0x16, 0xe8, 0x07, 0x00, 0x0d, 0xbf, 0x90, 0x71,
	0x61, 0x14, 0x55, 0x03, 0x7e, 0x38, 0xb4, 0xb4, 0xac, 0xde, 0x16, 0x11, 0xb6, 0x60, 0x8c, 0x89,
	0xbe, 0xc1, 0x93, 0xf0, 0x9c, 0xc2, 0x28, 0x29, 0xfc, 0xfb, 0x83, 0xf1, 0x61, 0x95, 0x45, 0xec,
	0x0d, 0x0d, 0xdd, 0xc2, 0xd1, 0xba, 0x5c, 0x8d, 0x23, 0x75, 0x59, 0x2f, 0xf3, 0xa9, 0xf7, 0xee,
	0x03, 0x89, 0x75, 0x81, 0xca, 0x42, 0xf7, 0x50, 0xf1, 0x88, 0x10, 0xd2, 0xe1, 0xcc, 0x9f, 0x3a,
	0x46, 0x59, 0x41, 0x5e, 0xe4, 0x43, 0xbe, 0xec, 0x36, 0xf7, 0x0b, 0x38, 0x9e, 0x8b, 0x1e, 0xc0,
	0x58, 0x50, 0x3e, 0xa5, 0xf6, 0xee, 0xd4, 0xb6, 0x64, 0xaa, 0x97, 0x8c, 0x63, 0xc5, 0x7d, 0x95,
	0xcf, 0xfd, 0xbc, 0xce, 0xdc, 0xf9, 0x37, 0x60, 0xeb, 0xae, 0xe9, 0x17, 0xb0, 0xbe, 0xc8, 0x0a,
	0x34, 0x47, 0x70, 0x99, 0x32, 0x3b, 0x63, 0x26, 0xdd, 0x24, 0x67, 0xd2, 0xf3, 0x7c, 0xf5, 0x5d,
	0x75, 0xc6, 0xe6, 0x51, 0x73, 0x08, 0x67, 0x71, 0xcb, 0x33, 0x14, 0x6e, 0x93, 0x0a, 0x7b, 0x9a,
	0x1f, 0xd3, 0x58, 0x8f, 0x83, 0xe4, 0xb0, 0x09, 0x3b, 0xc6, 0x84, 0xb3, 0x78, 0x06, 0x42, 0x70,
	0x24, 0xe9, 0x4a, 0x46, 0xda, 0xea, 0xd9, 0x3c, 0x87, 0x4a, 0xec, 0x36, 0xcc, 0x06, 0xe8, 0x99,
	0x26, 0x9a, 0x3a, 0x54, 0x33, 0xe6, 0x53, 0xf7, 0xea, 0xfb, 0x45, 0x52, 0x74, 0x78, 0xac, 0xfe,
	0x5f, 0x6f, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0xe3, 0x8b, 0x43, 0x38, 0xde, 0x06, 0x00, 0x00,
}
