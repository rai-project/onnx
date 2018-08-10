package onnx

import (
	fmt "fmt"
)

func (model *ModelProto) fixNames() {
	graph := model.GetGraph()
	for ii, n := range graph.Node {
		if n.GetName() != "" {
			continue
		}
		n.Name = fmt.Sprintf("%s_%d", n.OpType, ii)
	}
}

func (model *ModelProto) Fix() {
	model.fixNames()
}
