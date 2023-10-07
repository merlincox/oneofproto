package oneofproto_test

import (
	"encoding/json"
	"testing"

	"github.com/merlincox/oneofproto"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/structpb"

	"github.com/merlincox/oneofproto/switched"
	"github.com/merlincox/oneofproto/testapi"
)

type oneofFunc func(oneof any) (protostruct *structpb.Struct, typeName string, err error)

func funcTest100(t *testing.T, fn oneofFunc) {
	assignable := testapi.Assignable100{Data100: "test100"}
	input, _ := json.Marshal(&assignable)
	msg := &testapi.TestRequest{
		Test: &testapi.TestProto{
			Uuid: "uuid",
			Data: &testapi.TestProto_Assignable100{Assignable100: &assignable},
		},
	}

	out, name, err := fn(msg.Test.Data)
	assert.Nil(t, err, "should not error")
	assert.Equal(t, "*testapi.Assignable100", name, "should output correct type name")
	output, _ := out.MarshalJSON()
	assert.Equal(t, input, output, "should output a JSON compatible structpb")
}

func funcTest1(t *testing.T, fn oneofFunc) {
	assignable := testapi.Assignable1{Data1: "test100"}
	input, _ := json.Marshal(&assignable)
	msg := &testapi.Test2Request{
		Test: &testapi.Test2Proto{
			Uuid: "uuid",
			Data: &testapi.Test2Proto_Assignable1{Assignable1: &assignable},
		},
	}

	out, name, err := fn(msg.Test.Data)
	assert.Nil(t, err, "should not error")
	assert.Equal(t, "*testapi.Assignable100", name, "should output correct type name")
	output, _ := out.MarshalJSON()
	assert.Equal(t, input, output, "should output a JSON compatible structpb")
}

func funcBenchmark100(b *testing.B, fn oneofFunc) {
	for i := 0; i < b.N; i++ {
		assignable := testapi.Assignable100{Data100: "test100"}
		msg := &testapi.TestRequest{
			Test: &testapi.TestProto{
				Uuid: "uuid",
				Data: &testapi.TestProto_Assignable100{Assignable100: &assignable},
			},
		}

		_, _, err := fn(msg.Test.Data)
		if err != nil {
			b.FailNow()
		}
	}
}

func funcBenchmark1(b *testing.B, fn oneofFunc) {
	for i := 0; i < b.N; i++ {
		assignable := testapi.Assignable1{Data1: "test1"}
		msg := &testapi.Test2Request{
			Test: &testapi.Test2Proto{
				Uuid: "uuid",
				Data: &testapi.Test2Proto_Assignable1{Assignable1: &assignable},
			},
		}

		_, _, err := fn(msg.Test.Data)
		if err != nil {
			b.FailNow()
		}
	}
}

func TestToStructpb100(t *testing.T) {
	funcTest100(t, oneofproto.ToStructpb)
}

func TestToStructpbDefer100(t *testing.T) {
	funcTest100(t, oneofproto.ToStructpbDefer)
}

func TestToStructpbUnsafe100(t *testing.T) {
	funcTest100(t, oneofproto.ToStructpbUnsafe)
}

func TestSwitched100(t *testing.T) {
	funcTest100(t, switched.ToStructpb100)
}

func BenchmarkToStructpb100(b *testing.B) {
	funcBenchmark100(b, oneofproto.ToStructpb)
}

func BenchmarkToStructpbDefer100(b *testing.B) {
	funcBenchmark100(b, oneofproto.ToStructpbDefer)
}

func BenchmarkToStructpbUnsafe100(b *testing.B) {
	funcBenchmark100(b, oneofproto.ToStructpbUnsafe)
}

func BenchmarkSwitched100(b *testing.B) {
	funcBenchmark100(b, switched.ToStructpb100)
}

func BenchmarkToStructpb1(b *testing.B) {
	funcBenchmark1(b, oneofproto.ToStructpb)
}

func BenchmarkToStructpbDefer1(b *testing.B) {
	funcBenchmark1(b, oneofproto.ToStructpbDefer)
}

func BenchmarkToStructpbUnsafe1(b *testing.B) {
	funcBenchmark1(b, oneofproto.ToStructpbUnsafe)
}

func BenchmarkSwitched1(b *testing.B) {
	funcBenchmark1(b, switched.ToStructpb1)
}
