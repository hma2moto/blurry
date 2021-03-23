package blurry

/*
#cgo CFLAGS: -I${SRCDIR}
#cgo darwin LDFLAGS: -L${SRCDIR} -lruntime_osx -lgammacorrection_osx -ldl -lm
#cgo linux  LDFLAGS: -L${SRCDIR} -lruntime_linux -lgammacorrection_linux -ldl -lm
#include <stdlib.h>

#include "bridge.h"
#ifdef __APPLE__
#include "libgammacorrection_osx.h"
#elif __linux__
#include "libgammacorrection_linux.h"
#endif

int libgammacorrection(unsigned char *src, int32_t width, int32_t height, float factor, unsigned char *out) {
  halide_buffer_t in_rgba_buf = {0};
  halide_buffer_t out_rgba_buf = {0};

  bind_rgba_buf(&in_rgba_buf, src, width, height);
  bind_rgba_buf(&out_rgba_buf, out, width, height);

  int ret = gammacorrection(&in_rgba_buf, width, height, factor, &out_rgba_buf);
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
	ErrGamma = errors.New("gamma cgo call error")
)

func Gamma(img *image.RGBA, factor float64) (*image.RGBA, error) {
	width, height := wh(img)
	out := GetNRGBA(width, height)

	ret := C.libgammacorrection(
		(*C.uchar)(&img.Pix[0]),
		C.int(width),
		C.int(height),
		C.float(factor),
		(*C.uchar)(&out.Pix[0]),
	)
	if int(ret) != 0 {
		return nil, ErrGamma
	}
	return out, nil
}
