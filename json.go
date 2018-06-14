package onnx

import "C"

import "github.com/golang/protobuf/jsonpb"

func ToJSON(model *ModelProto) (string, error) {
	marshaler := &jsonpb.Marshaler{
		EnumsAsInts:  true,
		EmitDefaults: false,
		Indent:       "",
		OrigName:     true,
	}
	str, err := marshaler.MarshalToString(model)
	if err != nil {
		return "", err
	}
	return str, nil
}

func ReadToJSON(filename string) (string, error) {
	model, err := ReadModel(filename)
	if err != nil {
		return "", err
	}
	return ToJSON(model)
}
