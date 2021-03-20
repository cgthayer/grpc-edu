package main

import (
	"examplepb"

	"fmt"
	"log"

	"github.com/golang/protobuf/proto"
)

func main() {
	
	base := example.OptionalBase{
		I: 123
	}

	out, err := proto.Marshal(base)
	if err != nil {
		log.Fatalln("Boom")
	}
	fmt.Printf("base %d\n", out.Size)
	
}

