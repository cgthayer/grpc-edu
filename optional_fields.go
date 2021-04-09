package main

import (
	pb "github.com/cgthayer/grpc-edu/generated/examplepb"

	"fmt"
	"log"

	"github.com/golang/protobuf/proto"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
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

	oneof1 := pb.OptionalOneof{}
	print_size(&oneof1, "oneof no value")

	oneof2 := pb.OptionalOneof{
		TestOneof: &pb.OptionalOneof_I{
			I: int_i,
		},
	}
	print_size(&oneof2, "oneof with int")

	opt_int1 := pb.OptionalInt{
		I: nil,
	}
	print_size(&opt_int1, "optional Int32Value no value")
	opt_int2 := pb.OptionalInt{
		I: &wrapperspb.Int32Value{Value: int_i},
	}
	print_size(&opt_int2, "optional Int32Value with value")

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
	fmt.Printf("%s len=%d\n", msg, len(out))
	for i, v := range out {
		fmt.Printf("  %d: %08b\n", i, v)
	}
}
