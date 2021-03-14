package main

import "fmt"

type HelloRequest struct {
	Msg string
}

type HelloResponse struct {
	Msg string
}

func HelloService(request HelloRequest) HelloResponse {
	return HelloResponse{"echo: " + request.Msg }
}

func main() {
	response := HelloService(HelloRequest{"hi"})
	fmt.Println(response.Msg)
}
