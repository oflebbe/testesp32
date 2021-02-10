# testesp32

Works with
```
   go build example/main.go
```


Doesn't work for instance with
```
   tinygo build -target pca10040 -o dummy.hex example/main.go
```
or 
```
   tinygo build -target esp32-mini32 -o dummy.hex example/main.go
```

Error message
```
.../testesp32/demo.c:1:10: fatal error: 'foo.h' file not found
#include <foo.h>
         ^~~~~~~
1 error generated.
failed to run tool: clang
error: failed to build .../testesp32/demo.c: exit status 1
```
