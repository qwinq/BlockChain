package main

import (
	"fmt"
	"os"
)

func main3() {
	strings := os.Args
	for i,cmd:=range strings{
		fmt.Printf("arg[%d] : %s\n",i,cmd)
	}
}
