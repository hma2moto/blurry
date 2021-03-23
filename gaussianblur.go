package blurry

/*
#cgo CFLAGS: -I${SRCDIR}
#cgo darwin LDFLAGS: -L${SRCDIR} -lruntime_osx -lgaussianblur_osx -ldl -lm
#cgo linux  LDFLAGS: -L${SRCDIR} -lruntime_linux -lgaussianblur_linux -ldl -lm
#include <stdlib.h>

#include "bridge.h"
#ifdef __APPLE__
#include "libgaussianblur_osx.h"
#elif __linux__
#include "libgaussianblur_linux.h"
#endif

int libgaussianblur(unsigned char *src, int32_t width, int32_t height, float sigma, unsigned char *out) {
  halide_buffer_t in_rgba_buf = {0};
  halide_buffer_t out_rgba_buf = {0};

  bind_rgba_buf(&in_rgba_buf, src, width, height);
  bind_rgba_buf(&out_rgba_buf, out, width, height);

  int ret = gaussianblur(&in_rgba_buf, width, height, sigma, &out_rgba_buf);
  free_buf(&in_rgba_buf);
  free_buf(&out_rgba_buf);
  return ret;
}
*/
import "C"
import (
	"errors"
	"image"
)

var (
	ErrGaussianblur = errors.New("gaussianblur cgo call error")
)

func Gaussianblur(img *image.RGBA, sigma float64) (*image.RGBA, error) {
	width, height := wh(img)
	out := GetNRGBA(width, height)

	ret := C.libgaussianblur(
		(*C.uchar)(&img.Pix[0]),
		C.int(width),
		C.int(height),
		C.float(sigma),
		(*C.uchar)(&out.Pix[0]),
	)
	if int(ret) != 0 {
		return nil, ErrContrast
	}
	return out, nil
}