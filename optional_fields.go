package main

import (
	pb "github.com/cgthayer/grpc-edu/generated/examplepb"

	"fmt"
	"log"

	"github.com/golang/protobuf/proto"
)

func main() {
	base := pb.OptionalBase{
		I: 123,
	}
	print_size(&base, "base")

	opt1 := pb.OptionalField{
	}
	print_size(&opt1, "optional no value")

	var int_i int32 = 123
	opt2 := pb.OptionalField{
		I: &int_i,
	}
	print_size(&opt2, "optional with value")

	// oneof := pb.OptionalOneof{
	// 	pb.OptionalOneof.TestOneof{
	// 		I: &int_i,
	// 	}
	// }
	// print_size(&oneof, "oneof")

	opt_enum := pb.OptionalEnum{
		I: 1,
	}
	print_size(&opt_enum, "enum")
}

func print_size(ptr proto.Message, msg string) {
	out, err := proto.Marshal(ptr)
	if err != nil {
		log.Fatalln("Boom")
	}
	fmt.Printf("%s %d\n", msg, len(out))
	for x := 0; x < len(out); x++ {
		fmt.Printf("%d    %08b\n", x, out[x])
	}
}
