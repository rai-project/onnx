# Onnx Go Bindings For Reading models

[![Build Status](https://travis-ci.org/rai-project/onnx.svg?branch=master)](https://travis-ci.org/rai-project/onnx)

## Checkout The Submodules

```bash
git submodule update --init --recursive
```

or to update

```bash
git submodule update --recursive --remote
```

## Install Onnx

Refer to [installation](https://github.com/onnx/onnx#installation)

```bash
mkdir onnx/build
cd onnx/build
cmake ..
make
make install
```

## Test

```
go test -tags=connx
```

## Configure Model Reading

### Read Steps

The default is

```
	DefaultReadSteps = []string{
		"check",
		"shape_infer",
		"optimize",
		"check",
	}
```

### Graph Optimizations

The default is

```
	DefaultOptimizationNames = []string{
		"nop",
		"eliminate_identity",
		"eliminate_nop_transpose",
		"eliminate_unused_initializer",
		"fuse_consecutive_squeezes",
		"fuse_consecutive_transposes",
		"fuse_add_bias_into_conv",
		"fuse_transpose_into_gemm",
	}
```
