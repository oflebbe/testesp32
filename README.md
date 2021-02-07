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
# github.com/oflebbe/testesp32
demo.go:7:14: undeclared name: C.fortytwo
```
