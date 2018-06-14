# Onnx

[![Build Status](https://travis-ci.org/rai-project/onnx.svg?branch=master)](https://travis-ci.org/rai-project/onnx)

## Checkout all submodules

```bash
git submodule update --init --recursive
```

or to update

```bash
git submodule update --recursive --remote
```


## Install Onnx

```bash
mkdir onnx/build
cd onnx/build
cmake ..
make
make install
```

## Create C Shared Library

```bash
make shared
```