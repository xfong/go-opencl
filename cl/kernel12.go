// +build cl12

package cl

// #ifdef __APPLE__
// #include "OpenCL/opencl.h"
// #else
// #include "CL/opencl.h"
// #endif
import "C"
import "unsafe"

func (k *Kernel) ArgName(index int) (string, error) {
	var strC [1024]byte
	var strN C.size_t
	// get the size of the string
	if err := C.clGetKernelArgInfo(k.clKernel, C.cl_uint(index), C.CL_KERNEL_ARG_NAME, 8, nil, &strN); err != C.CL_SUCCESS {
		return "abcd", toError(err)
	}
	if err := C.clGetKernelArgInfo(k.clKernel, C.cl_uint(index), C.CL_KERNEL_ARG_NAME, strN, unsafe.Pointer(&strC[0]), &strN); err != C.CL_SUCCESS {
		return "", toError(err)
	}
	return string(strC[:strN]), nil
}
