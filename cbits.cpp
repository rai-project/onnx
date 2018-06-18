
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

go_string go_shape_inference(char *bytes, size_t len) {
  using namespace ONNX_NAMESPACE;
  ModelProto proto{};
  ParseProtoFromBytes(&proto, bytes, len);
  // std::cout << "strlen(bytes) = " << len << "\n";
  // std::cout << "version  = " << proto.ir_version() << "\n";
  // std::cout << "name  = " << proto.graph().name() << "\n";
  shape_inference::InferShapes(proto);
  // for (auto info : proto.graph().value_info()) {
  //   std::cout << "shape = " << info.name() << "\n";
  // }
  std::string out;
  proto.SerializeToString(&out);
  char *buf = (char *)malloc((out.size() + 1) * sizeof(char));
  memcpy(buf, out.c_str(), out.size());
  buf[out.size()] = '\0';

  go_string res;
  res.length = out.size();
  res.buf = buf;
  return res;
}