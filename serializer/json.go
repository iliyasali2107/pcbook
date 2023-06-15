package serializer

import (
	"fmt"

	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoiface"
)

func ProtobufToJSON(message proto.Message) (string, error) {
	pbMessage, ok := message.(protoreflect.ProtoMessage)
	if !ok {
		return "", fmt.Errorf("unable to convert to protoreflect.ProtoMessage")
	}

	pbV1, ok := pbMessage.(protoiface.MessageV1)
	if !ok {
		return "", fmt.Errorf("unable to convert to protoiface.MessageV1")
	}

	marshaler := jsonpb.Marshaler{
		EnumsAsInts:  false,
		EmitDefaults: true,
		Indent:       " ",
		OrigName:     true,
	}

	return marshaler.MarshalToString(pbV1)
}
