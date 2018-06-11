package onnx

import (
	"path/filepath"
	"testing"

	"github.com/k0kubun/pp"

	sourcepath "github.com/GeertJohan/go-sourcepath"
	"github.com/stretchr/testify/assert"
)

// TestModelShapeInfer ...
func TestModelShapeInfer(t *testing.T) {

	onnxModelFile := filepath.Join(sourcepath.MustAbsoluteDir(), "_fixtures", "", "test.onnx")

	model, err := ReadModelShapeInfer(onnxModelFile)
	assert.NoError(t, err)
	assert.NotEmpty(t, model)

	graph := model.GetGraph()

	pp.Println(graph)

	// assert.Equal(t, "Conv", nodes[0].GetOpType())
	// assert.Equal(t, "Relu", nodes[1].GetOpType())

	// assert.Equal(t, int64(3), model.GetIrVersion())
	// assert.Equal(t, "CNTK", model.GetProducerName())
	// assert.Equal(t, "Constant", nodes[0].GetOpType())
	// assert.Equal(t, "Div", nodes[1].GetOpType())
}

// TestUnmarshalModel ...
func XXXTestUnmarshalModel(t *testing.T) {

	onnxModelFile := filepath.Join(sourcepath.MustAbsoluteDir(), "_fixtures", "mnist", "mnist.onnx")

	model, err := ReadModel(onnxModelFile)
	assert.NoError(t, err)
	assert.NotEmpty(t, model)

	graph := model.GetGraph()
	nodes := graph.GetNode()

	pp.Println(len(nodes))
	// pp.Println(graph.GetName())

	assert.Equal(t, int64(3), model.GetIrVersion())
	assert.Equal(t, "CNTK", model.GetProducerName())
	assert.Equal(t, "Constant", nodes[0].GetOpType())
	assert.Equal(t, "Div", nodes[1].GetOpType())
}
