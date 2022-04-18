#ifndef HALIDE__libscale_linear_linux_h
#define HALIDE__libscale_linear_linux_h
#include <stdint.h>

// Forward declarations of the types used in the interface
// to the Halide pipeline.
//
// For the definitions of these structs, include HalideRuntime.h

// Halide's representation of a multi-dimensional array.
// Halide::Runtime::Buffer is a more user-friendly wrapper
// around this. Its declaration is in HalideBuffer.h
struct halide_buffer_t;

// Metadata describing the arguments to the generated function.
// Used to construct calls to the _argv version of the function.
struct halide_filter_metadata_t;

#ifndef HALIDE_MUST_USE_RESULT
#ifdef __has_attribute
#if __has_attribute(nodiscard)
#define HALIDE_MUST_USE_RESULT [[nodiscard]]
#elif __has_attribute(warn_unused_result)
#define HALIDE_MUST_USE_RESULT __attribute__((warn_unused_result))
#else
#define HALIDE_MUST_USE_RESULT
#endif
#else
#define HALIDE_MUST_USE_RESULT
#endif
#endif

#ifndef HALIDE_FUNCTION_ATTRS
#define HALIDE_FUNCTION_ATTRS
#endif



#ifdef __cplusplus
extern "C" {
#endif

HALIDE_FUNCTION_ATTRS
int scale_linear_par_for_in_18_s0_ch__1(void *__user_context, int32_t _in__18_s0_ch, uint8_t *_closure_arg__84);
HALIDE_FUNCTION_ATTRS
int scale_linear_par_for_scale_linear_s0_ch_par_for_scale_linear_s0_x_xo_ti_ti__2(void *__user_context, int32_t _scale_linear_s0_x_xo_ti_ti, uint8_t *_closure_arg__86);
HALIDE_FUNCTION_ATTRS
int scale_linear_par_for_scale_linear_s0_ch_par_for_scale_linear_s0_x_xo_ti_ti__3(void *__user_context, int32_t _scale_linear_s0_x_xo_ti_ti, uint8_t *_closure_arg__87);
HALIDE_FUNCTION_ATTRS
int scale_linear_par_for_scale_linear_s0_ch__1(void *__user_context, int32_t _scale_linear_s0_ch, uint8_t *_closure_arg__85);
HALIDE_FUNCTION_ATTRS
int scale_linear(struct halide_buffer_t *_src_buffer, int32_t _width, int32_t _height, int32_t _scale_width, int32_t _scale_height, struct halide_buffer_t *_scale_linear_buffer);

HALIDE_FUNCTION_ATTRS
int scale_linear_argv(void **args);

HALIDE_FUNCTION_ATTRS
const struct halide_filter_metadata_t *scale_linear_metadata();

#ifdef __cplusplus
}  // extern "C"
#endif

#endif
