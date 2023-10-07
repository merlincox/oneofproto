package switched

import (
	"fmt"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"

	"github.com/merlincox/oneofproto/testapi"
)

const (
	testapiAssignable1   = "*testapi.Assignable1"
	testapiAssignable2   = "*testapi.Assignable2"
	testapiAssignable3   = "*testapi.Assignable3"
	testapiAssignable4   = "*testapi.Assignable4"
	testapiAssignable5   = "*testapi.Assignable5"
	testapiAssignable6   = "*testapi.Assignable6"
	testapiAssignable7   = "*testapi.Assignable7"
	testapiAssignable8   = "*testapi.Assignable8"
	testapiAssignable9   = "*testapi.Assignable9"
	testapiAssignable10  = "*testapi.Assignable10"
	testapiAssignable11  = "*testapi.Assignable11"
	testapiAssignable12  = "*testapi.Assignable12"
	testapiAssignable13  = "*testapi.Assignable13"
	testapiAssignable14  = "*testapi.Assignable14"
	testapiAssignable15  = "*testapi.Assignable15"
	testapiAssignable16  = "*testapi.Assignable16"
	testapiAssignable17  = "*testapi.Assignable17"
	testapiAssignable18  = "*testapi.Assignable18"
	testapiAssignable19  = "*testapi.Assignable19"
	testapiAssignable20  = "*testapi.Assignable20"
	testapiAssignable21  = "*testapi.Assignable21"
	testapiAssignable22  = "*testapi.Assignable22"
	testapiAssignable23  = "*testapi.Assignable23"
	testapiAssignable24  = "*testapi.Assignable24"
	testapiAssignable25  = "*testapi.Assignable25"
	testapiAssignable26  = "*testapi.Assignable26"
	testapiAssignable27  = "*testapi.Assignable27"
	testapiAssignable28  = "*testapi.Assignable28"
	testapiAssignable29  = "*testapi.Assignable29"
	testapiAssignable30  = "*testapi.Assignable30"
	testapiAssignable31  = "*testapi.Assignable31"
	testapiAssignable32  = "*testapi.Assignable32"
	testapiAssignable33  = "*testapi.Assignable33"
	testapiAssignable34  = "*testapi.Assignable34"
	testapiAssignable35  = "*testapi.Assignable35"
	testapiAssignable36  = "*testapi.Assignable36"
	testapiAssignable37  = "*testapi.Assignable37"
	testapiAssignable38  = "*testapi.Assignable38"
	testapiAssignable39  = "*testapi.Assignable39"
	testapiAssignable40  = "*testapi.Assignable40"
	testapiAssignable41  = "*testapi.Assignable41"
	testapiAssignable42  = "*testapi.Assignable42"
	testapiAssignable43  = "*testapi.Assignable43"
	testapiAssignable44  = "*testapi.Assignable44"
	testapiAssignable45  = "*testapi.Assignable45"
	testapiAssignable46  = "*testapi.Assignable46"
	testapiAssignable47  = "*testapi.Assignable47"
	testapiAssignable48  = "*testapi.Assignable48"
	testapiAssignable49  = "*testapi.Assignable49"
	testapiAssignable50  = "*testapi.Assignable50"
	testapiAssignable51  = "*testapi.Assignable51"
	testapiAssignable52  = "*testapi.Assignable52"
	testapiAssignable53  = "*testapi.Assignable53"
	testapiAssignable54  = "*testapi.Assignable54"
	testapiAssignable55  = "*testapi.Assignable55"
	testapiAssignable56  = "*testapi.Assignable56"
	testapiAssignable57  = "*testapi.Assignable57"
	testapiAssignable58  = "*testapi.Assignable58"
	testapiAssignable59  = "*testapi.Assignable59"
	testapiAssignable60  = "*testapi.Assignable60"
	testapiAssignable61  = "*testapi.Assignable61"
	testapiAssignable62  = "*testapi.Assignable62"
	testapiAssignable63  = "*testapi.Assignable63"
	testapiAssignable64  = "*testapi.Assignable64"
	testapiAssignable65  = "*testapi.Assignable65"
	testapiAssignable66  = "*testapi.Assignable66"
	testapiAssignable67  = "*testapi.Assignable67"
	testapiAssignable68  = "*testapi.Assignable68"
	testapiAssignable69  = "*testapi.Assignable69"
	testapiAssignable70  = "*testapi.Assignable70"
	testapiAssignable71  = "*testapi.Assignable71"
	testapiAssignable72  = "*testapi.Assignable72"
	testapiAssignable73  = "*testapi.Assignable73"
	testapiAssignable74  = "*testapi.Assignable74"
	testapiAssignable75  = "*testapi.Assignable75"
	testapiAssignable76  = "*testapi.Assignable76"
	testapiAssignable77  = "*testapi.Assignable77"
	testapiAssignable78  = "*testapi.Assignable78"
	testapiAssignable79  = "*testapi.Assignable79"
	testapiAssignable80  = "*testapi.Assignable80"
	testapiAssignable81  = "*testapi.Assignable81"
	testapiAssignable82  = "*testapi.Assignable82"
	testapiAssignable83  = "*testapi.Assignable83"
	testapiAssignable84  = "*testapi.Assignable84"
	testapiAssignable85  = "*testapi.Assignable85"
	testapiAssignable86  = "*testapi.Assignable86"
	testapiAssignable87  = "*testapi.Assignable87"
	testapiAssignable88  = "*testapi.Assignable88"
	testapiAssignable89  = "*testapi.Assignable89"
	testapiAssignable90  = "*testapi.Assignable90"
	testapiAssignable91  = "*testapi.Assignable91"
	testapiAssignable92  = "*testapi.Assignable92"
	testapiAssignable93  = "*testapi.Assignable93"
	testapiAssignable94  = "*testapi.Assignable94"
	testapiAssignable95  = "*testapi.Assignable95"
	testapiAssignable96  = "*testapi.Assignable96"
	testapiAssignable97  = "*testapi.Assignable97"
	testapiAssignable98  = "*testapi.Assignable98"
	testapiAssignable99  = "*testapi.Assignable99"
	testapiAssignable100 = "*testapi.Assignable100"
)

func ToStructpb100(oneof any) (protostruct *structpb.Struct, typeName string, err error) {

	switch assignable := oneof.(type) {
	case *testapi.TestProto_Assignable1:
		protostruct, err = messageToStructpb(assignable.Assignable1)
		typeName = testapiAssignable1
	case *testapi.TestProto_Assignable2:
		protostruct, err = messageToStructpb(assignable.Assignable2)
		typeName = testapiAssignable2
	case *testapi.TestProto_Assignable3:
		protostruct, err = messageToStructpb(assignable.Assignable3)
		typeName = testapiAssignable3
	case *testapi.TestProto_Assignable4:
		protostruct, err = messageToStructpb(assignable.Assignable4)
		typeName = testapiAssignable4
	case *testapi.TestProto_Assignable5:
		protostruct, err = messageToStructpb(assignable.Assignable5)
		typeName = testapiAssignable5
	case *testapi.TestProto_Assignable6:
		protostruct, err = messageToStructpb(assignable.Assignable6)
		typeName = testapiAssignable6
	case *testapi.TestProto_Assignable7:
		protostruct, err = messageToStructpb(assignable.Assignable7)
		typeName = testapiAssignable7
	case *testapi.TestProto_Assignable8:
		protostruct, err = messageToStructpb(assignable.Assignable8)
		typeName = testapiAssignable8
	case *testapi.TestProto_Assignable9:
		protostruct, err = messageToStructpb(assignable.Assignable9)
		typeName = testapiAssignable9
	case *testapi.TestProto_Assignable10:
		protostruct, err = messageToStructpb(assignable.Assignable10)
		typeName = testapiAssignable10
	case *testapi.TestProto_Assignable11:
		protostruct, err = messageToStructpb(assignable.Assignable11)
		typeName = testapiAssignable11
	case *testapi.TestProto_Assignable12:
		protostruct, err = messageToStructpb(assignable.Assignable12)
		typeName = testapiAssignable12
	case *testapi.TestProto_Assignable13:
		protostruct, err = messageToStructpb(assignable.Assignable13)
		typeName = testapiAssignable13
	case *testapi.TestProto_Assignable14:
		protostruct, err = messageToStructpb(assignable.Assignable14)
		typeName = testapiAssignable14
	case *testapi.TestProto_Assignable15:
		protostruct, err = messageToStructpb(assignable.Assignable15)
		typeName = testapiAssignable15
	case *testapi.TestProto_Assignable16:
		protostruct, err = messageToStructpb(assignable.Assignable16)
		typeName = testapiAssignable16
	case *testapi.TestProto_Assignable17:
		protostruct, err = messageToStructpb(assignable.Assignable17)
		typeName = testapiAssignable17
	case *testapi.TestProto_Assignable18:
		protostruct, err = messageToStructpb(assignable.Assignable18)
		typeName = testapiAssignable18
	case *testapi.TestProto_Assignable19:
		protostruct, err = messageToStructpb(assignable.Assignable19)
		typeName = testapiAssignable19
	case *testapi.TestProto_Assignable20:
		protostruct, err = messageToStructpb(assignable.Assignable20)
		typeName = testapiAssignable20
	case *testapi.TestProto_Assignable21:
		protostruct, err = messageToStructpb(assignable.Assignable21)
		typeName = testapiAssignable21
	case *testapi.TestProto_Assignable22:
		protostruct, err = messageToStructpb(assignable.Assignable22)
		typeName = testapiAssignable22
	case *testapi.TestProto_Assignable23:
		protostruct, err = messageToStructpb(assignable.Assignable23)
		typeName = testapiAssignable23
	case *testapi.TestProto_Assignable24:
		protostruct, err = messageToStructpb(assignable.Assignable24)
		typeName = testapiAssignable24
	case *testapi.TestProto_Assignable25:
		protostruct, err = messageToStructpb(assignable.Assignable25)
		typeName = testapiAssignable25
	case *testapi.TestProto_Assignable26:
		protostruct, err = messageToStructpb(assignable.Assignable26)
		typeName = testapiAssignable26
	case *testapi.TestProto_Assignable27:
		protostruct, err = messageToStructpb(assignable.Assignable27)
		typeName = testapiAssignable27
	case *testapi.TestProto_Assignable28:
		protostruct, err = messageToStructpb(assignable.Assignable28)
		typeName = testapiAssignable28
	case *testapi.TestProto_Assignable29:
		protostruct, err = messageToStructpb(assignable.Assignable29)
		typeName = testapiAssignable29
	case *testapi.TestProto_Assignable30:
		protostruct, err = messageToStructpb(assignable.Assignable30)
		typeName = testapiAssignable30
	case *testapi.TestProto_Assignable31:
		protostruct, err = messageToStructpb(assignable.Assignable31)
		typeName = testapiAssignable31
	case *testapi.TestProto_Assignable32:
		protostruct, err = messageToStructpb(assignable.Assignable32)
		typeName = testapiAssignable32
	case *testapi.TestProto_Assignable33:
		protostruct, err = messageToStructpb(assignable.Assignable33)
		typeName = testapiAssignable33
	case *testapi.TestProto_Assignable34:
		protostruct, err = messageToStructpb(assignable.Assignable34)
		typeName = testapiAssignable34
	case *testapi.TestProto_Assignable35:
		protostruct, err = messageToStructpb(assignable.Assignable35)
		typeName = testapiAssignable35
	case *testapi.TestProto_Assignable36:
		protostruct, err = messageToStructpb(assignable.Assignable36)
		typeName = testapiAssignable36
	case *testapi.TestProto_Assignable37:
		protostruct, err = messageToStructpb(assignable.Assignable37)
		typeName = testapiAssignable37
	case *testapi.TestProto_Assignable38:
		protostruct, err = messageToStructpb(assignable.Assignable38)
		typeName = testapiAssignable38
	case *testapi.TestProto_Assignable39:
		protostruct, err = messageToStructpb(assignable.Assignable39)
		typeName = testapiAssignable39
	case *testapi.TestProto_Assignable40:
		protostruct, err = messageToStructpb(assignable.Assignable40)
		typeName = testapiAssignable40
	case *testapi.TestProto_Assignable41:
		protostruct, err = messageToStructpb(assignable.Assignable41)
		typeName = testapiAssignable41
	case *testapi.TestProto_Assignable42:
		protostruct, err = messageToStructpb(assignable.Assignable42)
		typeName = testapiAssignable42
	case *testapi.TestProto_Assignable43:
		protostruct, err = messageToStructpb(assignable.Assignable43)
		typeName = testapiAssignable43
	case *testapi.TestProto_Assignable44:
		protostruct, err = messageToStructpb(assignable.Assignable44)
		typeName = testapiAssignable44
	case *testapi.TestProto_Assignable45:
		protostruct, err = messageToStructpb(assignable.Assignable45)
		typeName = testapiAssignable45
	case *testapi.TestProto_Assignable46:
		protostruct, err = messageToStructpb(assignable.Assignable46)
		typeName = testapiAssignable46
	case *testapi.TestProto_Assignable47:
		protostruct, err = messageToStructpb(assignable.Assignable47)
		typeName = testapiAssignable47
	case *testapi.TestProto_Assignable48:
		protostruct, err = messageToStructpb(assignable.Assignable48)
		typeName = testapiAssignable48
	case *testapi.TestProto_Assignable49:
		protostruct, err = messageToStructpb(assignable.Assignable49)
		typeName = testapiAssignable49
	case *testapi.TestProto_Assignable50:
		protostruct, err = messageToStructpb(assignable.Assignable50)
		typeName = testapiAssignable50
	case *testapi.TestProto_Assignable51:
		protostruct, err = messageToStructpb(assignable.Assignable51)
		typeName = testapiAssignable51
	case *testapi.TestProto_Assignable52:
		protostruct, err = messageToStructpb(assignable.Assignable52)
		typeName = testapiAssignable52
	case *testapi.TestProto_Assignable53:
		protostruct, err = messageToStructpb(assignable.Assignable53)
		typeName = testapiAssignable53
	case *testapi.TestProto_Assignable54:
		protostruct, err = messageToStructpb(assignable.Assignable54)
		typeName = testapiAssignable54
	case *testapi.TestProto_Assignable55:
		protostruct, err = messageToStructpb(assignable.Assignable55)
		typeName = testapiAssignable55
	case *testapi.TestProto_Assignable56:
		protostruct, err = messageToStructpb(assignable.Assignable56)
		typeName = testapiAssignable56
	case *testapi.TestProto_Assignable57:
		protostruct, err = messageToStructpb(assignable.Assignable57)
		typeName = testapiAssignable57
	case *testapi.TestProto_Assignable58:
		protostruct, err = messageToStructpb(assignable.Assignable58)
		typeName = testapiAssignable58
	case *testapi.TestProto_Assignable59:
		protostruct, err = messageToStructpb(assignable.Assignable59)
		typeName = testapiAssignable59
	case *testapi.TestProto_Assignable60:
		protostruct, err = messageToStructpb(assignable.Assignable60)
		typeName = testapiAssignable60
	case *testapi.TestProto_Assignable61:
		protostruct, err = messageToStructpb(assignable.Assignable61)
		typeName = testapiAssignable61
	case *testapi.TestProto_Assignable62:
		protostruct, err = messageToStructpb(assignable.Assignable62)
		typeName = testapiAssignable62
	case *testapi.TestProto_Assignable63:
		protostruct, err = messageToStructpb(assignable.Assignable63)
		typeName = testapiAssignable63
	case *testapi.TestProto_Assignable64:
		protostruct, err = messageToStructpb(assignable.Assignable64)
		typeName = testapiAssignable64
	case *testapi.TestProto_Assignable65:
		protostruct, err = messageToStructpb(assignable.Assignable65)
		typeName = testapiAssignable65
	case *testapi.TestProto_Assignable66:
		protostruct, err = messageToStructpb(assignable.Assignable66)
		typeName = testapiAssignable66
	case *testapi.TestProto_Assignable67:
		protostruct, err = messageToStructpb(assignable.Assignable67)
		typeName = testapiAssignable67
	case *testapi.TestProto_Assignable68:
		protostruct, err = messageToStructpb(assignable.Assignable68)
		typeName = testapiAssignable68
	case *testapi.TestProto_Assignable69:
		protostruct, err = messageToStructpb(assignable.Assignable69)
		typeName = testapiAssignable69
	case *testapi.TestProto_Assignable70:
		protostruct, err = messageToStructpb(assignable.Assignable70)
		typeName = testapiAssignable70
	case *testapi.TestProto_Assignable71:
		protostruct, err = messageToStructpb(assignable.Assignable71)
		typeName = testapiAssignable71
	case *testapi.TestProto_Assignable72:
		protostruct, err = messageToStructpb(assignable.Assignable72)
		typeName = testapiAssignable72
	case *testapi.TestProto_Assignable73:
		protostruct, err = messageToStructpb(assignable.Assignable73)
		typeName = testapiAssignable73
	case *testapi.TestProto_Assignable74:
		protostruct, err = messageToStructpb(assignable.Assignable74)
		typeName = testapiAssignable74
	case *testapi.TestProto_Assignable75:
		protostruct, err = messageToStructpb(assignable.Assignable75)
		typeName = testapiAssignable75
	case *testapi.TestProto_Assignable76:
		protostruct, err = messageToStructpb(assignable.Assignable76)
		typeName = testapiAssignable76
	case *testapi.TestProto_Assignable77:
		protostruct, err = messageToStructpb(assignable.Assignable77)
		typeName = testapiAssignable77
	case *testapi.TestProto_Assignable78:
		protostruct, err = messageToStructpb(assignable.Assignable78)
		typeName = testapiAssignable78
	case *testapi.TestProto_Assignable79:
		protostruct, err = messageToStructpb(assignable.Assignable79)
		typeName = testapiAssignable79
	case *testapi.TestProto_Assignable80:
		protostruct, err = messageToStructpb(assignable.Assignable80)
		typeName = testapiAssignable80
	case *testapi.TestProto_Assignable81:
		protostruct, err = messageToStructpb(assignable.Assignable81)
		typeName = testapiAssignable81
	case *testapi.TestProto_Assignable82:
		protostruct, err = messageToStructpb(assignable.Assignable82)
		typeName = testapiAssignable82
	case *testapi.TestProto_Assignable83:
		protostruct, err = messageToStructpb(assignable.Assignable83)
		typeName = testapiAssignable83
	case *testapi.TestProto_Assignable84:
		protostruct, err = messageToStructpb(assignable.Assignable84)
		typeName = testapiAssignable84
	case *testapi.TestProto_Assignable85:
		protostruct, err = messageToStructpb(assignable.Assignable85)
		typeName = testapiAssignable85
	case *testapi.TestProto_Assignable86:
		protostruct, err = messageToStructpb(assignable.Assignable86)
		typeName = testapiAssignable86
	case *testapi.TestProto_Assignable87:
		protostruct, err = messageToStructpb(assignable.Assignable87)
		typeName = testapiAssignable87
	case *testapi.TestProto_Assignable88:
		protostruct, err = messageToStructpb(assignable.Assignable88)
		typeName = testapiAssignable88
	case *testapi.TestProto_Assignable89:
		protostruct, err = messageToStructpb(assignable.Assignable89)
		typeName = testapiAssignable89
	case *testapi.TestProto_Assignable90:
		protostruct, err = messageToStructpb(assignable.Assignable90)
		typeName = testapiAssignable90
	case *testapi.TestProto_Assignable91:
		protostruct, err = messageToStructpb(assignable.Assignable91)
		typeName = testapiAssignable91
	case *testapi.TestProto_Assignable92:
		protostruct, err = messageToStructpb(assignable.Assignable92)
		typeName = testapiAssignable92
	case *testapi.TestProto_Assignable93:
		protostruct, err = messageToStructpb(assignable.Assignable93)
		typeName = testapiAssignable93
	case *testapi.TestProto_Assignable94:
		protostruct, err = messageToStructpb(assignable.Assignable94)
		typeName = testapiAssignable94
	case *testapi.TestProto_Assignable95:
		protostruct, err = messageToStructpb(assignable.Assignable95)
		typeName = testapiAssignable95
	case *testapi.TestProto_Assignable96:
		protostruct, err = messageToStructpb(assignable.Assignable96)
		typeName = testapiAssignable96
	case *testapi.TestProto_Assignable97:
		protostruct, err = messageToStructpb(assignable.Assignable97)
		typeName = testapiAssignable97
	case *testapi.TestProto_Assignable98:
		protostruct, err = messageToStructpb(assignable.Assignable98)
		typeName = testapiAssignable98
	case *testapi.TestProto_Assignable99:
		protostruct, err = messageToStructpb(assignable.Assignable99)
		typeName = testapiAssignable99
	case *testapi.TestProto_Assignable100:
		protostruct, err = messageToStructpb(assignable.Assignable100)
		typeName = testapiAssignable100
	default:
		err = fmt.Errorf("unsupported type")
		return
	}

	return
}

func ToStructpb1(oneof any) (protostruct *structpb.Struct, typeName string, err error) {

	switch assignable := oneof.(type) {
	case *testapi.Test2Proto_Assignable1:
		protostruct, err = messageToStructpb(assignable.Assignable1)
		typeName = testapiAssignable1
	default:
		err = fmt.Errorf("unsupported type")
		return
	}

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
