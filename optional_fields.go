package main

import (
	pb "github.com/cgthayer/grpc-edu/gemerated/examplepb"

	"fmt"
	"log"

	"github.com/golang/protobuf/proto"
)

func main() {
	base := pb.OptionalBase{
		I: 123
	}

	out, err := proto.Marshal(base)
	if err != nil {
		log.Fatalln("Boom")
	}
	fmt.Printf("base %d\n", out.Length)
}

