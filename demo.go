package main

// #include "demo.h"
import "C"

func zweiundvierzig() int {
  return int(C.fortytwo())
}
