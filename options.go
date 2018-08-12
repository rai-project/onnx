package onnx

import "context"

var (
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
	DefaultReadSteps = []string{
		"check",
		"shape_infer",
		"optimize",
		"check",
	}
)

// Options ...
type Options struct {
	Context       context.Context
	Optimizations []string
	Steps         []string
}

// Option ...
type Option func(*Options)

func Optimizations(optimizationNames []string) Option {
	return func(opts *Options) {
		opts.Optimizations = optimizationNames
	}
}

func Steps(steps []string) Option {
	return func(opts *Options) {
		opts.Steps = steps
	}
}

func NewOptions(opts ...Option) *Options {
	options := &Options{
		Context:       context.Background(),
		Optimizations: DefaultOptimizationNames,
		Steps:         DefaultReadSteps,
	}

	for _, o := range opts {
		o(options)
	}

	return options
}
