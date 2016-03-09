package main

/*

#cgo pkg-config: python3
#define Py_LIMITED_API
#include <Python.h>


PyObject * sum(PyObject *, PyObject *);
PyObject * make_a_lot_of_files(PyObject *, PyObject *);
PyObject * make_a_lot_of_files_faster(PyObject *, PyObject *);

// Workaround missing variadic function support
// https://github.com/golang/go/issues/975
int PyArg_ParseTuple_LL(PyObject * args, long long * a, long long * b) {
    return PyArg_ParseTuple(args, "LL", a, b);
}
int PyArg_ParseTuple_L(PyObject * args, long long * a) {
    return PyArg_ParseTuple(args, "L", a);
}
static PyMethodDef makefilesMethods[] = {
    {"sum", sum, METH_VARARGS, "Add two numbers."},
    {"make_a_lot_of_files", make_a_lot_of_files, METH_VARARGS, "makes a lot of files one at a time"},
    {"make_a_lot_of_files_faster", make_a_lot_of_files_faster, METH_VARARGS, "makes a lot of files concurrently"},
    {NULL, NULL, 0, NULL}
};

static struct PyModuleDef moduleDef = {
   PyModuleDef_HEAD_INIT, "makefiles", "makefiles module doc string goes here", -1, makefilesMethods
};

PyMODINIT_FUNC
PyInit_makefiles(void)
{
    return PyModule_Create(&moduleDef);
}

*/
import "C"

func main() {}
