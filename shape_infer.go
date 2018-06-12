package onnx

/*
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include "cbits.hpp"
*/
import "C"

import (
	"io/ioutil"
	"unsafe"

	"github.com/Unknwon/com"
	"github.com/gogo/protobuf/proto"
	"github.com/k0kubun/pp"
	"github.com/pkg/errors"
)

func ReadModelShapeInfer(protoFileName string) (*ModelProto, error) {

	if !com.IsFile(protoFileName) {
		return nil, errors.Errorf("%s is not a file", protoFileName)
	}
	buf, err := ioutil.ReadFile(protoFileName)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to open %s", protoFileName)
	}

	pp.Println("aaa string(buf) is", len(string(buf)))

	// protoContent := C.CString(string(buf))
	// defer C.free(unsafe.Pointer(protoContent))

	// pp.Println("protoContent in go is", C.int(C.strlen(protoContent)))

	// sharedProtoContentC := C.go_shape_inference(protoContent)
	// if sharedProtoContentC == nil {
	// 	return nil, errors.Wrapf(err, "failed to shape infer %s", protoFileName)
	// }
	// defer C.free(unsafe.Pointer(sharedProtoContentC))
	// sharedProtoContent := C.GoString(sharedProtoContentC)

	// pp.Println("aaa sharedProtoContent is", len(string(sharedProtoContent)))

	// model := new(ModelProto)
	// err = proto.UnmarshalText(sharedProtoContent, model)

	protoContent := C.CBytes(buf)
	defer C.free(unsafe.Pointer(protoContent))

	//pp.Println("protoContent in go is", C.int(C.strlen((*C.char)(protoContent))))

	pp.Println("protoContent in go is", len(buf))

	shapedProtoContentC := C.go_shape_inference((*C.char)(protoContent), C.size_t(len(buf)))
	if shapedProtoContentC.buf == nil {
		return nil, errors.Wrapf(err, "failed to shape infer %s", protoFileName)
	}
	length := C.int(shapedProtoContentC.length)
	sharedProtoContent := C.GoBytes(unsafe.Pointer(shapedProtoContentC.buf), length)

	defer C.free(unsafe.Pointer(shapedProtoContentC.buf))

	//sharedProtoContent := []byte(C.GoString(shapedProtoContentC))
	pp.Println("len in go is", len(sharedProtoContent))

	model := new(ModelProto)
	err = proto.Unmarshal(sharedProtoContent, model)

	return model, err
}
