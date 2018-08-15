// +build !cgo

package onnx

func CheckModel(protoFileName string) (*ModelProto, error) {
	return ReadModel(protoFileName)
}

func OptmizeModel(protoFileName string, optimizationNames []string) (*ModelProto, error) {
	return ReadModel(protoFileName)
}

func ReadModelShapeInfer(protoFileName string) (*ModelProto, error) {
	return ReadModel(protoFileName)
}
