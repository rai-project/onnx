//+build ignore

package main

// #cgo LDFLAGS: -lonnx -lonnx_proto -lprotobuf -lstdc++ -L${SRCDIR}/onnx/build -L/usr/local/lib
// #cgo CXXFLAGS: -std=c++11 -I${SRCDIR}/onnx/build -I${SRCDIR}/onnx -O3 -Wall -g -DONNX_NAMESPACE=onnx
// #cgo CXXFLAGS: -Wno-sign-compare -Wno-unused-function
import "C"

import (
	"unsafe"

	"github.com/rai-project/onnx"
)

//export ONNXConvertToJSON
func ONNXConvertToJSON(filename string) (string, error) {
	model, err := onnx.ReadModel(filename)
	if err != nil {
		return "", err
	}
	return onnx.ToJSON(model)
}

//export ONNXRead
func ONNXRead(filename string) (unsafe.Pointer, error) {
	model, err := onnx.ReadModel(filename)
	if err != nil {
		return nil, err
	}
	return unsafe.Pointer(model), nil
}

func main() {
	// We need the main function to make possible
	// CGO compiler to compile the package as C shared library
}
