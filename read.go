package onnx

import (
	"os"
	"strings"

	"github.com/Unknwon/com"
	"github.com/gogo/protobuf/proto"
	"github.com/pkg/errors"
)

func New(protoFileNameOrg string, opts ...Option) (*ModelProto, error) {
	protoFileName := protoFileNameOrg

	if !com.IsFile(protoFileName) {
		return nil, errors.Errorf("%s is not a file", protoFileName)
	}

	options := NewOptions(opts...)

	var model *ModelProto
	var err error

	for _, step := range options.Steps {
		if model != nil {
			protoFileName, err = tempFileName("", "model", ".onnx")
			if err != nil {
				return nil, errors.New("cannot create temporary model file")
			}
			defer func(filename string) {
				os.Remove(filename)
			}(protoFileName)

			buf, err := proto.Marshal(model)
			if err != nil {
				return nil, err
			}
			err = com.WriteFile(protoFileName, buf)
			if err != nil {
				return nil, errors.Wrapf(err, "unable to write %s file", protoFileName)
			}
			model = nil
		}
		switch strings.ToLower(step) {
		case "check":
			model, err = CheckModel(protoFileName)
			if err != nil {
				return nil, err
			}
		case "optimize":
			model, err = OptmizeModel(protoFileName, options.Optimizations)
			if err != nil {
				return nil, err
			}
		case "shape_infer", "shapeinfer", "infer":
			model, err = ReadModelShapeInfer(protoFileName)
			if err != nil {
				return nil, err
			}
		}
	}

	if model != nil {
		return model, nil
	}
	return ReadModel(protoFileName)
}
