package onnx

import (
	"io/ioutil"

	"github.com/Unknwon/com"
	"github.com/gogo/protobuf/proto"
	"github.com/k0kubun/pp"
	"github.com/pkg/errors"
)

func unmarshal(target proto.Message, protoFileName string) error {
	if !com.IsFile(protoFileName) {
		return errors.Errorf("%s is not a file", protoFileName)
	}
	buf, err := ioutil.ReadFile(protoFileName)
	if err != nil {
		return errors.Wrapf(err, "failed to open %s", protoFileName)
	}

	pp.Println("string(buf) is", len(string(buf)))

	return proto.Unmarshal(buf, target)
}

func ReadModel(protoFileName string) (*ModelProto, error) {
	model := new(ModelProto)
	err := unmarshal(model, protoFileName)

	return model, err
}
