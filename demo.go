package testesp32

/*
#include "demo.h"
#cgo CFLAGS: -Iinclude
*/
import "C"

func Zweiundvierzig() int {
	return int(C.fortytwo())
}
