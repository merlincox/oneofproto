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

func funcTest(t *testing.T, fn oneofFunc) {
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

func funcBenchmark(b *testing.B, fn oneofFunc) {
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

func TestToStructpb(t *testing.T) {
	funcTest(t, oneofproto.ToStructpb)
}

func TestToStructpbDefer(t *testing.T) {
	funcTest(t, oneofproto.ToStructpbDefer)
}

func TestToStructpbUnsafe(t *testing.T) {
	funcTest(t, oneofproto.ToStructpbUnsafe)
}

func TestSwitched(t *testing.T) {
	funcTest(t, switched.ToStructpb)
}

func BenchmarkToStructpb(b *testing.B) {
	funcBenchmark(b, oneofproto.ToStructpb)
}

func BenchmarkToStructpbDefer(b *testing.B) {
	funcBenchmark(b, oneofproto.ToStructpbDefer)
}

func BenchmarkToStructpbUnsafe(b *testing.B) {
	funcBenchmark(b, oneofproto.ToStructpbUnsafe)
}

func BenchmarkSwitched(b *testing.B) {
	funcBenchmark(b, switched.ToStructpb)
}
