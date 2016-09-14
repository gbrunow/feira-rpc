package main

import (
  "bufio"
	"fmt"
	"net"
	"net/rpc"
	"os"
	"io"
	"encoding/csv"
	"strconv"
)

type Fruit struct{
	FruitName string
	Price float64
}

type Weighting struct{
	FruitName string
	Weight float64
}

var dataBase map[string]float64

func loadCSV(){
	dataBase = make(map[string]float64)
	f, _ := os.Open("feiraFrutaData.csv")
	r := csv.NewReader(bufio.NewReader(f))
	for {
		record, err := r.Read()
		// Stop at EOF.
		if err == io.EOF {
		    break
			}
		dataBase[record[0]], _ = strconv.ParseFloat(record[1], 64)
    }
}

func writeCSV(){

}

type Arith int

func (t *Arith) Register(args *Fruit, reply *bool) error{

	fmt.Println(args.FruitName, args.Price)

	dataBase[args.FruitName] = args.Price
	//TODO: write to CSV

	*reply = true
	return nil
}

func (t *Arith) Calculate(args *Weighting, reply *float64) error{
	
	fmt.Println(args.FruitName, args.Weight)

	valueKg, ok := dataBase[args.FruitName]

	if ok{
		*reply = args.Weight*valueKg
	} else {
		*reply = 0
		}

	return nil
}

func main() {
	loadCSV()
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
