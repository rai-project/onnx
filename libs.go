// +build connx

package onnx

// #cgo LDFLAGS: -lonnx -lonnx_proto -lprotobuf -lstdc++ -L${SRCDIR}/onnx/build -L/usr/local/lib
// #cgo CXXFLAGS: -std=c++11 -I${SRCDIR}/onnx/build -I${SRCDIR}/onnx -O3 -Wall -g -DONNX_NAMESPACE=onnx -DC_ONNX_BINDINGS
// #cgo CXXFLAGS: -Wno-sign-compare -Wno-unused-function
import "C"
