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
  "log"
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
	f, err := os.Open("feiraFrutaData.csv")
  if err != nil{
    f, err = os.Create("feiraFrutaData.csv")
    if err != nil{
      fmt.Println("Failed to open/create CSV file")
      os.Exit(0)
    }
  }
	r := csv.NewReader(bufio.NewReader(f))
	for {
		record, err := r.Read()
		// Stop at EOF.
		if err == io.EOF {
		    break
			}
		dataBase[record[0]], _ = strconv.ParseFloat(record[1], 64)
    }
    f.Close()
}

func writeCSV(){
  f, _ := os.Create("feiraFrutaData.csv")
  w := csv.NewWriter(bufio.NewWriter(f))
  for key, value := range dataBase {
      record := []string{key,fmt.Sprintf("%.2f", value)}
      if err := w.Write(record); err != nil {
  			log.Fatalln("error writing record to csv:", err)
  		}
    }
    w.Flush()
}

type FruitCall int
type FruitName string

func (t *FruitCall) Register(args *Fruit, reply *bool) error{

	fmt.Println(args.FruitName, args.Price)

  _, ok := dataBase[args.FruitName]

  if ok{
		*reply = false
	} else {
    dataBase[args.FruitName] = args.Price
    writeCSV()

    *reply = true
		}

	return nil
}

func (t *FruitCall) Remove(args *Fruit, reply *bool) error{

	fmt.Println(args.FruitName)
  _, ok := dataBase[args.FruitName]

  if ok{
    delete(dataBase, args.FruitName)
    writeCSV()
		*reply = true
	} else {
    *reply = false
		}

	return nil
}

func (t *FruitCall) Update(args *Fruit, reply *bool) error{

	fmt.Println(args.FruitName, args.Price)
  _, ok := dataBase[args.FruitName]

  if ok{
    dataBase[args.FruitName] = args.Price
    writeCSV()
		*reply = true
	} else {
    *reply = false
		}

	return nil
}

func (t *FruitCall) Calculate(args *Weighting, reply *float64) error{

	fmt.Println(args.FruitName, args.Weight)

	valueKg, ok := dataBase[args.FruitName]

	if ok{
		*reply = args.Weight*valueKg
	} else {
		*reply = -1
		}

	return nil
}

func (t *FruitCall) Consult(args *Weighting, reply *float64) error{

	fmt.Println(args.FruitName)

	valueKg, ok := dataBase[args.FruitName]

	if ok{
		*reply = valueKg
	} else {
		*reply = -1
		}

	return nil
}

func main() {
	loadCSV()
	fruitcall := new(FruitCall)
	rpc.Register(fruitcall)
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
		go rpc.ServeConn(conn)
	}
}

func checkError(str string, err error) {
	if err != nil {
		fmt.Println(str, err)
		os.Exit(1)
	}
}
