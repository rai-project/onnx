package onnx

func (md *ModelProto) RemoveWeights() error {
	grph := md.Graph
	if grph == nil {
		return nil
	}
	inits := grph.Initializer
	if len(inits) == 0 {
		return nil
	}
	for ii, init := range inits {
		init.RawData = []byte{}
		init.StringData = [][]byte{}
		init.FloatData = []float32{}
		init.Int32Data = []int32{}
		init.Int64Data = []int64{}
		init.DoubleData = []float64{}
		init.Uint64Data = []uint64{}
		grph.Initializer[ii] = init
	}
	return nil
}
