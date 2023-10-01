package oneofproto

import (
	"fmt"
	"reflect"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"
)

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
		return nil, err
	}

	out := new(structpb.Struct)
	if err = protojson.Unmarshal(bytes, out); err != nil {
		return nil, err
	}

	return out, nil
}
