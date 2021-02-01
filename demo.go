//#cgo
package example

// extern int fortytwo();
import "C"

func ZweiUndVierzig() int {
   return int(C.fortytwo())
}

