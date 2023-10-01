package oneofproto

import (
	"fmt"
	"reflect"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"
)

func ToStructpbDefer(oneof any) (protostruct *structpb.Struct, typeName string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("unsupported type")
		}
	}()

	return ToStructpbUnsafe(oneof)
}

func ToStructpbUnsafe(oneof any) (protostruct *structpb.Struct, typeName string, err error) {
	val := reflect.ValueOf(oneof).Elem().Field(0)
	typeName = val.Type().String()
	protostruct, err = messageToStructpb(val.Interface().(proto.Message))

	return
}
