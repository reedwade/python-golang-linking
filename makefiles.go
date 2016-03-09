package main

// #cgo pkg-config: python3
// #define Py_LIMITED_API
// #include <Python.h>
// int PyArg_ParseTuple_LL(PyObject *, long long *, long long *);
// int PyArg_ParseTuple_L(PyObject *, long long *);
import "C"

import (
	"fmt"
	"io/ioutil"
	"sync"
)

//export sum
func sum(self, args *C.PyObject) *C.PyObject {
	var a, b C.longlong
	if C.PyArg_ParseTuple_LL(args, &a, &b) == 0 {
		return &C._Py_NoneStruct
	}
	return C.PyLong_FromLongLong(a + b)
}

//export make_a_lot_of_files
func make_a_lot_of_files(self, args *C.PyObject) *C.PyObject {
	var _howMany C.longlong
	if C.PyArg_ParseTuple_L(args, &_howMany) == 0 {
		return &C._Py_NoneStruct
	}

	howMany := int(_howMany)

	for i := 0; i < howMany; i++ {
		make_a_file(i)
	}

	return &C._Py_NoneStruct
}

func make_a_file(i int) {
	f := fmt.Sprintf("out_go/%05d", i)

	err := ioutil.WriteFile(f, []byte("hello"), 0644)

	if err != nil {
		fmt.Println("failed on", f, err)
		// } else {
		//     fmt.Println("wrote", f)
	}
}

var wg sync.WaitGroup

//export make_a_lot_of_files_faster
func make_a_lot_of_files_faster(self, args *C.PyObject) *C.PyObject {
	var _howMany C.longlong
	if C.PyArg_ParseTuple_L(args, &_howMany) == 0 {
		return &C._Py_NoneStruct
	}

	howMany := int(_howMany)

	for i := 0; i < howMany; i++ {
		wg.Add(1)
		go func(_i int) {
			defer wg.Done()
			make_a_file(i)
		}(i)

		if i%1000 == 0 {
			fmt.Println("waiting")
			wg.Wait()
		}
	}

	wg.Wait()

	return &C._Py_NoneStruct
}
