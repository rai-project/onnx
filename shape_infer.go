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

	protoContent := C.CBytes(buf)
	defer C.free(unsafe.Pointer(protoContent))

	shapedProtoContentC := C.go_shape_inference((*C.char)(protoContent))
	if shapedProtoContentC == nil {
		return nil, errors.Wrapf(err, "failed to shape infer %s", protoFileName)
	}
	defer C.free(unsafe.Pointer(shapedProtoContentC))

	// sharedProtoContent := C.GoString(shapedProtoContentC)
	len := C.int(C.strlen(shapedProtoContentC))
	sharedProtoContent := C.GoBytes(unsafe.Pointer(shapedProtoContentC), len)

	model := new(ModelProto)
	err = proto.Unmarshal(sharedProtoContent, model)

	return nil, err
}
