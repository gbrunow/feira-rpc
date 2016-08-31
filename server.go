package main

import (
	"fmt"
	"net"
	"net/rpc"
	"os"
)

type Fruit struct{
	FruitName string
	Price float32
}

type Weighting struct{
	FruitName string
	Weight float32
}

type Arith int

func (t *Arith) Register(args *Fruit, reply *bool) error{

	fmt.Println(args.FruitName, args.Price)
	//TODO: write to CSV

	*reply = true
	return nil
}

func (t *Arith) Calculate(args *Weighting, reply *float32) error{

	fmt.Println(args.FruitName, args.Weight)
	//TODO: read CSV and calculate price

	*reply = 12.345
	return nil
}

func main() {
	arith := new(Arith)
	rpc.Register(arith)
	tcpAddr, err :=
		net.ResolveTCPAddr("tcp", "localhost:1234")
	checkError("ResolveTCPAddr: ", err)
	listener, err :=
		net.ListenTCP("tcp", tcpAddr)
	checkError("ListenTCP: ", err)
	for {
		conn, err :=
			listener.Accept()
		checkError("Accept: ", err)
		rpc.ServeConn(conn)
	}
}

func checkError(str string, err error) {
	if err != nil {
		fmt.Println(str, err)
		os.Exit(1)
	}
}
