#pragma once

#include <stdint.h>
#include <string.h>

#ifdef __cplusplus
extern "C" {
#endif // __cplusplus

typedef struct {
  size_t length;
  char *buf;
} go_string;

go_string go_shape_inference(char *bytes, size_t len);

#ifdef __cplusplus
}
#endif // __cplusplus