// +build connx

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

func OptmizeModel(protoFileName string, optimizationNames []string) (*ModelProto, error) {

	if len(optimizationNames) == 0 {
		optimizationNames = DefaultOptimizationNames
	}

	if !com.IsFile(protoFileName) {
		return nil, errors.Errorf("%s is not a file", protoFileName)
	}
	buf, err := ioutil.ReadFile(protoFileName)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to open %s", protoFileName)
	}

	protoContent := C.CBytes(buf)
	defer C.free(unsafe.Pointer(protoContent))

	cargs := C.makeCharArray(C.int(len(optimizationNames)))
	defer C.freeCharArray(cargs, C.int(len(optimizationNames)))
	for i, s := range optimizationNames {
		C.setArrayString(cargs, C.CString(s), C.int(i))
	}

	shapedProtoContentC := C.go_optimize((*C.char)(protoContent), C.size_t(len(buf)), cargs, C.int(len(optimizationNames)))
	if shapedProtoContentC.buf == nil {
		return nil, errors.Wrapf(err, "failed to optimize model %s", protoFileName)
	}
	length := C.int(shapedProtoContentC.length)
	sharedProtoContent := C.GoBytes(unsafe.Pointer(shapedProtoContentC.buf), length)

	defer C.free(unsafe.Pointer(shapedProtoContentC.buf))

	model := new(ModelProto)
	err = proto.Unmarshal(sharedProtoContent, model)

	return model, err
}
