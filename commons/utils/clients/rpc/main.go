package main

import (
	"fmt"
	"log"
	"net/rpc"

	"github.com/diegocomunity/goschool/commons"
)

func check(err error) {
	log.Fatalf(err.Error())
}

// example cliente rpc
func main() {
	args := &commons.Args{A: 100, B: 2}
	var reply = new(commons.Reply)
	client, err := rpc.DialHTTP("tcp", ":1234")
	if err != nil {
		check(err)
	}
	if err != nil {
		check(err)
	}
	err = client.Call("Arith.Add", args, reply)
	if err != nil {
		check(err)
	}
	/*err = client.Call("Arith.Error", args, reply)
	if err != nil {
		log.Fatalf(err.Error())
	}
	*/
	fmt.Printf("Args A: %d\nArgs B: %d\nReultado: %d\n", args.A, args.B, reply.C)
	//fmt.Printf("reply: %v\n", reply)
	defer client.Close()

}
