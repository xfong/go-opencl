// +build cl12

package cl

// #ifdef __APPLE__
// #include "OpenCL/opencl.h"
// #else
// #include "cl.h"
// #endif
import "C"
import "unsafe"

func (k *Kernel) ArgName(index int) (string, error) {
	var strC [1024]byte
	var strN C.size_t
	// get the size of output
	err := C.clGetKernelArgInfo(k.clKernel, C.cl_uint(index), C.CL_KERNEL_ARG_NAME, 0, unsafe.Pointer(&strC[0]), &strN)
	if err != C.CL_SUCCESS {
		return "abcd", toError(err)
	}
	if err = C.clGetKernelArgInfo(k.clKernel, C.cl_uint(index), C.CL_KERNEL_ARG_NAME, strN, unsafe.Pointer(&strC[0]), nil); err != C.CL_SUCCESS {
		return "123", toError(err)
	}
	return string(strC[:strN]), nil
}
