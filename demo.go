package testesp32 

// #include "demo.h"
import "C"

func Zweiundvierzig() int {
  return int(C.fortytwo())
}
