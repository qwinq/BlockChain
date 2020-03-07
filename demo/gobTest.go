package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Person struct {
	Name string
	Age int
}
func main2() {
	xMing:= Person{"xMing",20}
	xm:=Person{}
	var buffer bytes.Buffer
	//1.定义编码器->编码
	encoder := gob.NewEncoder(&buffer)
	encoder.Encode(xMing)

	//2.定义解码器->解码
	decoder := gob.NewDecoder(bytes.NewReader(buffer.Bytes()))
	decoder.Decode(&xm)
	fmt.Println(xm)
}
