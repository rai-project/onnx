package onnx

import (
	"path/filepath"
	"testing"

	"github.com/k0kubun/pp"

	sourcepath "github.com/GeertJohan/go-sourcepath"
	"github.com/stretchr/testify/assert"
)

// TestUnmarshalModel ...
func TestUnmarshalModel(t *testing.T) {

	onnxModelFile := filepath.Join(sourcepath.MustAbsoluteDir(), "_fixtures", "squeezenet", "model.onnx")

	model, err := ReadModel(onnxModelFile)
	assert.NoError(t, err)
	assert.NotEmpty(t, model)

	graph := model.GetGraph()
	nodes := graph.GetNode()

	pp.Println(nodes[1].GetName())

	// assert.Equal(t, int32(256), nodes[0].GetName())
	// assert.Equal(t, int32(256), model.GetHeight())
	// assert.Equal(t, int32(3), blmodelob.GetChannels())
}
