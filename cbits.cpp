
#include <climits>
#include <limits>
#include <stdio.h>
#include <string.h>
#include <unordered_map>

#include "cbits.hpp"

#include "onnx/checker.h"
#include "onnx/defs/schema.h"
#include "onnx/optimizer/optimize.h"
#include "onnx/shape_inference/implementation.h"

char *go_shape_inference(char *bytes) {
  // using namespace ONNX_NAMESPACE;
  // unsigned int len = strlen(bytes);
  // ModelProto proto{};
  // ParseProtoFromBytes(&proto, bytes, len);
  // shape_inference::InferShapes(proto);
  // char *out = new char[len];
  // proto.SerializeToArray(out, len);
  // return out;
  using namespace ONNX_NAMESPACE;
  ModelProto proto{};
  ParseProtoFromBytes(&proto, bytes, strlen(bytes));
  std::cout << "strlen(bytes) = " << strlen(bytes) << "\n";
  shape_inference::InferShapes(proto);
  std::string out;
  proto.SerializeToString(&out);
  return strdup(out.c_str());
}