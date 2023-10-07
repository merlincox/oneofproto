package oneofproto

import (
	"fmt"
	"reflect"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"
)

// ToStructpbDefer mirrors the functionality of ToStructpb but uses a panic-recovery mechanism instead of kind checking.
func ToStructpbDefer(oneof any) (protostruct *structpb.Struct, typeName string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("unsupported type")
		}
	}()

	return ToStructpbUnsafe(oneof)
}

// ToStructpbUnsafe mirrors the functionality of ToStructpb but uses no kind checks or recovery, so is unsafe.
func ToStructpbUnsafe(oneof any) (protostruct *structpb.Struct, typeName string, err error) {
	val := reflect.ValueOf(oneof).Elem().Field(0)
	typeName = val.Type().String()
	protostruct, err = messageToStructpb(val.Interface().(proto.Message))

	return
}
