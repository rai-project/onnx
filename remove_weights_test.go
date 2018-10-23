package onnx

import (
	"fmt"
	"path/filepath"
	"testing"

	sourcepath "github.com/GeertJohan/go-sourcepath"
	"github.com/k0kubun/pp"
	"github.com/stretchr/testify/assert"
)

// TestUnmarshalModel ...
func TestRemoveWeights(t *testing.T) {

	onnxModelFile := filepath.Join(sourcepath.MustAbsoluteDir(), "_fixtures", "mnist", "mnist_inferred.onnx")

	model, err := ReadModel(onnxModelFile)
	assert.NoError(t, err)
	assert.NotEmpty(t, model)

	bts, err := model.Marshal()
	assert.NoError(t, err)
	assert.NotEmpty(t, bts)

	fmt.Printf("before size: %d\n", len(bts))

	err = model.RemoveWeights()

	pp.Println(model)

	bts, err = model.Marshal()
	assert.NoError(t, err)
	assert.NotEmpty(t, bts)

	fmt.Printf("after size: %d\n", len(bts))

}
