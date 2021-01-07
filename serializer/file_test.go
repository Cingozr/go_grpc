package serializer_test

import (
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/lessongo/grpc_ex/pb"
	"github.com/lessongo/grpc_ex/sample"
	"github.com/lessongo/grpc_ex/serializer"
	"github.com/stretchr/testify/require"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()

	binaryFile := "../temp/laptop.bin"
	jsonFile := "../temp/laptop.json"
	laptop1 := sample.NewLaptop()

	err := serializer.WriteProtobufToBinaryFile(laptop1, binaryFile)
	require.NoError(t, err)

	laptop2 := &pb.Laptop{}
	err = serializer.ReadProtobufToBinaryFile(binaryFile, laptop2)
	require.NoError(t, err)
	require.True(t, proto.Equal(laptop1, laptop2))

	err = serializer.WriteProtobufToJSONFile(laptop1, jsonFile)
	require.NoError(t, err)
}
