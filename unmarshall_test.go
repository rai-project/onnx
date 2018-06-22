package onnx

import (
	"path/filepath"
	"testing"

	sourcepath "github.com/GeertJohan/go-sourcepath"
	"github.com/stretchr/testify/assert"
)

// TestUnmarshalModel ...
func TestUnmarshalModel(t *testing.T) {

	onnxModelFile := filepath.Join(sourcepath.MustAbsoluteDir(), "_fixtures", "mnist", "mnist.onnx")

	model, err := ReadModel(onnxModelFile)
	assert.NoError(t, err)
	assert.NotEmpty(t, model)

	graph := model.GetGraph()
	nodes := graph.GetNode()

	// pp.Println(len(nodes))
	// pp.Println(graph.GetName())

	assert.Equal(t, int64(3), model.GetIrVersion())
	assert.Equal(t, "CNTK", model.GetProducerName())
	assert.Equal(t, "Constant", nodes[0].GetOpType())
	assert.Equal(t, "Div", nodes[1].GetOpType())
}
