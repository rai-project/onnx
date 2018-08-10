package onnx

import (
	"path/filepath"
	"testing"

	"github.com/Unknwon/com"
	"github.com/gogo/protobuf/proto"
	"github.com/k0kubun/pp"

	sourcepath "github.com/GeertJohan/go-sourcepath"
<<<<<<< HEAD
	home "github.com/mitchellh/go-homedir"
	"github.com/stretchr/testify/assert"
)

var (
	homedir, _ = home.Dir()
)

=======
	"github.com/stretchr/testify/assert"
)

>>>>>>> 05b64e8b52129defb03db8005900424bb19bd61e
// TestModelShapeInfer ...
func TestModelShapeInfer(t *testing.T) {

	//onnxModelFile := filepath.Join(sourcepath.MustAbsoluteDir(), "_fixtures", "", "test.onnx")

	onnxModelFile := filepath.Join(sourcepath.MustAbsoluteDir(), "_fixtures", "", "test.onnx")
<<<<<<< HEAD
	// onnxModelFile = filepath.Join(homedir, "onnx_models", "bvlc_alexnet", "model.onnx")
=======
>>>>>>> 05b64e8b52129defb03db8005900424bb19bd61e

	model, err := ReadModelShapeInfer(onnxModelFile)
	assert.NoError(t, err)
	assert.NotEmpty(t, model)

	graph := model.GetGraph()

	for _, val := range graph.GetValueInfo() {
		pp.Println(val.GetType().GetValue())
	}

	buf, err := proto.Marshal(model)
	com.WriteFile(filepath.Join(sourcepath.MustAbsoluteDir(), "_fixtures", "", "testa_inferred.onnx"), buf)

	// assert.Equal(t, "Conv", nodes[0].GetOpType())
	// assert.Equal(t, "Relu", nodes[1].GetOpType())

	// assert.Equal(t, int64(3), model.GetIrVersion())
	// assert.Equal(t, "CNTK", model.GetProducerName())
	// assert.Equal(t, "Constant", nodes[0].GetOpType())
	// assert.Equal(t, "Div", nodes[1].GetOpType())
}
