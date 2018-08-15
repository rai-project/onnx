// +build cgo

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

	shapedProtoContentC := C.go_shape_inference((*C.char)(protoContent), C.size_t(len(buf)))
	if shapedProtoContentC.buf == nil {
		return nil, errors.Wrapf(err, "failed to shape infer %s", protoFileName)
	}
	length := C.int(shapedProtoContentC.length)
	sharedProtoContent := C.GoBytes(unsafe.Pointer(shapedProtoContentC.buf), length)

	defer C.free(unsafe.Pointer(shapedProtoContentC.buf))

	model := new(ModelProto)
	err = proto.Unmarshal(sharedProtoContent, model)
	if err != nil {
		return nil, err
	}

	model.Fix()

	return model, nil
}
