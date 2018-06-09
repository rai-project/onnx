package onnx

/*
#include <stdio.h>
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

func ReadModelShareInfer(protoFileName string) (*ModelProto, error) {

	if !com.IsFile(protoFileName) {
		return nil, errors.Errorf("%s is not a file", protoFileName)
	}
	buf, err := ioutil.ReadFile(protoFileName)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to open %s", protoFileName)
	}

	protoContent := C.CString(string(buf))
	defer C.free(unsafe.Pointer(protoContent))

	sharedProtoContentC := C.go_shape_inference(protoContent)
	if sharedProtoContentC == nil {
		return nil, errors.Wrapf(err, "failed to shape infer %s", protoFileName)
	}
	//defer C.free(unsafe.Pointer(sharedProtoContentC))
	sharedProtoContent := C.GoString(sharedProtoContentC)

	model := new(ModelProto)
	err = proto.UnmarshalText(sharedProtoContent, model)
	return model, err

}
