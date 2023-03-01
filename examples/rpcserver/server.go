package rpcserver

import (
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/diegocomunity/goschool/commons"
)

type RPC_Server struct{}

func NewRPCServer() *RPC_Server {
	return &RPC_Server{}
}

func init() {
	rpc.Register(new(commons.Arith))
	rpc.Register(new(commons.BuiltinTypes))
}
func (*RPC_Server) Run() {
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatalf(err.Error())
	}
	println("running server TCP in port :1234")
	http.Serve(l, nil)
}
