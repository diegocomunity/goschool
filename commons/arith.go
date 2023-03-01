package commons

import (
	"errors"
	"log"
)

type Args struct {
	A, B int
}
type Reply struct {
	C int
}
type Arith int

func (t *Arith) Add(args *Args, reply *Reply) error {
	println("(Add)")
	reply.C = args.A + args.B
	return nil
}
func (t *Arith) Mult(args *Args, reply *Reply) error {
	println("(Mult)")
	reply.C = args.A * args.B
	return nil
}
func (t *Arith) Div(args *Args, reply *Reply) error {
	println("(Div)")
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	reply.C = args.A / args.B
	return nil
}
func (t *Arith) Error(args *Args, reply *Reply) error {
	log.Fatalf("Me salgo del programa")
	panic("1")
}
