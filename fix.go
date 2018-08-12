package onnx

import (
	fmt "fmt"
	"strings"
)

func (model *ModelProto) fixNames() {
	layerTypeOccurrences := map[string]int{}

	graph := model.GetGraph()
	for _, n := range graph.Node {
		if _, ok := layerTypeOccurrences[n.OpType]; !ok {
			layerTypeOccurrences[n.OpType] = 0
		}
		layerTypeOccurrences[n.OpType] = layerTypeOccurrences[n.OpType] + 1

		if n.GetName() != "" {
			continue
		}

		layerName := strings.ToLower(n.OpType)
		layerOccurence := layerTypeOccurrences[n.OpType]

		group := -1
		for _, attr := range n.Attribute {
			if attr.Name == "group" {
				group = int(attr.I)
			}
		}

		if group == -1 {
			n.Name = fmt.Sprintf("%s_%d", layerName, layerOccurence)
		} else {
			n.Name = fmt.Sprintf("%s%d_%d", layerName, layerOccurence, group)
		}

	}
}

func (model *ModelProto) Fix() {
	model.fixNames()
}
