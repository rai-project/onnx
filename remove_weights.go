package onnx

func (md *ModelProto) RemoveWeights() error {
	grph := md.Graph
	if grph == nil {
		return nil
	}
	for ii, init := range grph.Initializer {
		init.RawData = []byte{}
		init.StringData = [][]byte{}
		init.FloatData = []float32{}
		init.Int32Data = []int32{}
		init.Int64Data = []int64{}
		init.DoubleData = []float64{}
		init.Uint64Data = []uint64{}
		grph.Initializer[ii] = init
	}

	for ii, node := range grph.Node {
		for jj, attr := range node.Attribute {
			if attr.T != nil {
				attr.T.FloatData = []float32{}
				attr.T.RawData = []byte{}
				attr.T.StringData = [][]byte{}
				attr.T.FloatData = []float32{}
				attr.T.Int32Data = []int32{}
				attr.T.Int64Data = []int64{}
				attr.T.DoubleData = []float64{}
				attr.T.Uint64Data = []uint64{}
			}
			attr.Floats = []float32{}
			attr.Ints = []int64{}
			attr.Strings = [][]byte{}
			node.Attribute[jj] = attr
		}
		grph.Node[ii] = node
	}
	return nil
}
