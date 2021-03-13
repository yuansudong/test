package main

import (
	"log"
	"unsafe"
)

// EmptyStruct 用于计算结构体的大小
type EmptyStruct struct{}

// Size 用于返回结构体的大小
func (es *EmptyStruct) Size() {

	log.Println(unsafe.Sizeof(EmptyStruct{}))

}
