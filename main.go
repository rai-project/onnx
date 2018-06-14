//+build ignore

package main

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
