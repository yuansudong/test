package main

import (
	"google.golang.org/protobuf/proto"
)

// Student 用于描述一个学生
type Student struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"name,omitempty"`
}

func main() {
	proto.Marshal()
}
