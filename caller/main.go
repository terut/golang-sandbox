package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("vim-go")
	var pcs [32]uintptr
	n := runtime.Callers(0, pcs[:])
	fmt.Println(pcs[0:n])
	for _, pc := range pcs[0:n] {
		fn := runtime.FuncForPC(pc)
		fmt.Printf("Function Name: %s\n", fn.Name())
		fileName, fileLine := fn.FileLine(pc)
		fmt.Printf("FileName:%s, FileLine: %d\n", fileName, fileLine)
	}
}
