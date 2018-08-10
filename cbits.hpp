#pragma once

#include <stdint.h>
#include <stdio.h>
#include <string.h>

#ifdef __cplusplus
extern "C" {
#endif // __cplusplus

typedef struct {
  size_t length;
  char *buf;
} go_string;

go_string go_shape_inference(char *bytes, size_t len);

go_string go_optimize(char *bytes, size_t len, char **optnames, int numopts);

static char **makeCharArray(int size) { return (char **) calloc(sizeof(char *), size); }

static void setArrayString(char **a, char *s, int n) { a[n] = s; }

static void freeCharArray(char **a, int size) {
  if (a == NULL) {
    return;
  }
  for (int ii = 0; ii < size; ii++) {
    if (a == NULL) {
      continue;
    }
    free(a[ii]);
  }
  free(a);
}

#ifdef __cplusplus
}
#endif // __cplusplus