package oneofproto

import (
	"fmt"
	"reflect"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"
)

// ToStructpb converts a protobuf oneof value into a JSON-compatible *structpb.Struct, or returns an error if the conversion is impossible.
// It also returns a string representation of the underlying type.
//
//	Example:
//
//	type TestProto struct {
//		state         protoimpl.MessageState
//		sizeCache     protoimpl.SizeCache
//		unknownFields protoimpl.UnknownFields
//
//		Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
//		Data isTestProto_Data `protobuf_oneof:"data"`
//	}
//
//	msg := &testapi.TestRequest{
//		Test: &testapi.TestProto{
//			Uuid: "uuid",
//			Data: &testapi.TestProto_Assignable100{Assignable100: &testapi.Assignable100{Data100: "test100"}},
//		},
//	}
//
//	protostruct, typeName, err := oneofproto.ToStructpb(msg.Test.Data)
//
//	where protostruct is a *structpb.Struct JSON compatible with &testapi.Assignable100{Data100: "test100"}
//	and typeName is "*testapi.Assignable100"
func ToStructpb(oneof any) (protostruct *structpb.Struct, typeName string, err error) {
	val := reflect.ValueOf(oneof)
	if val.Kind() != reflect.Pointer {
		err = fmt.Errorf("unsupported type")
		return
	}
	val = val.Elem()
	if val.Kind() != reflect.Struct || val.NumField() == 0 {
		err = fmt.Errorf("unsupported type")
		return
	}
	val = val.Field(0)
	msg, ok := val.Interface().(proto.Message)
	if !ok {
		err = fmt.Errorf("unsupported type")
		return
	}
	typeName = val.Type().String()

	protostruct, err = messageToStructpb(msg)

	return
}

func messageToStructpb(msg proto.Message) (*structpb.Struct, error) {
	bytes, err := protojson.MarshalOptions{UseProtoNames: true, EmitUnpopulated: true}.Marshal(msg)
	if err != nil {
		return nil, fmt.Errorf("unsupported type: %w", err)
	}

	out := new(structpb.Struct)
	if err = protojson.Unmarshal(bytes, out); err != nil {
		return nil, fmt.Errorf("unsupported type: %w", err)
	}

	return out, nil
}
