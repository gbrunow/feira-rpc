package main

import (
	//"bufio"
	"fmt"
	"net/rpc"
	"os"
	"strings"
)

type Args struct {
	A int
	B float64
}
type Quotient struct {
	Q, R int
}

type Fruit struct {
	FruitName string
	Price     float64
}

type Weighting struct {
	FruitName string
	Weight    float64
}

func main() {
	service := "localhost:1234"
	client, err := rpc.Dial("tcp", service)
	defer client.Close()
	checkError("Dial: ", err)

	var op byte

	for {
		fmt.Println("1 - Register")
		fmt.Println("2 - Update Fruit")
		fmt.Println("3 - Remove Fruit")
		fmt.Println("4 - Calculate Price")
		fmt.Println("5 - Consult Price")
		fmt.Println("6 - Exit")
		fmt.Print("Option: ")
		fmt.Scanf("%c\n", &op)

		switch op {
		case '1':
			args := readEntry()
			var reply bool
			err = client.Call("Arith.Register", args, &reply)
			checkError("Register: ", err)

			if reply {
				fmt.Println(strings.Title(args.FruitName), "registered with success.")
			} else {
				fmt.Println(strings.Title(args.FruitName), "already exist, use Update option to change price.")
			}

			break
		case '2':
			args := readEntry()
			var reply bool
			err = client.Call("Arith.Update", args, &reply)
			checkError("Update: ", err)

			if reply {
				fmt.Println(strings.Title(args.FruitName), "updated with success.")
			} else {
				fmt.Println(strings.Title(args.FruitName), "product does not exist.")
			}

			break
		case '3':
			args := readFruitName()

			var reply bool
			err = client.Call("Arith.Remove", args, &reply)
			checkError("Remove: ", err)

			if reply {
				fmt.Println(strings.Title(args.FruitName), "removed with success.")
			} else {
				fmt.Println(strings.Title(args.FruitName), "product did not exist.")
			}

			break
		case '4':
			args := readWeighting()
			var reply float32
			err = client.Call("Arith.Calculate", args, &reply)
			checkError("Calculate: ", err)

			if reply >= 0 {
				fmt.Println("Price", reply)
			}else{
				fmt.Println(strings.Title(args.FruitName), "is not registered.")
			}

			break
		case '5':
			args := readFruitName()

			var reply float32
			err = client.Call("Arith.Consult", args, &reply)
			checkError("Consult: ", err)

			if reply >= 0 {
				fmt.Println("Price per kg", reply)
			} else {
				fmt.Println(strings.Title(args.FruitName), "not registered.")
			}

			break
		case '6':

			os.Exit(0)
		default:
			fmt.Println("Opção inválida: ", op)
			os.Exit(1)
		}
	}
}

func readEntry() Fruit {
	var name string
	var price float64

	fmt.Print("Name: ")
	name = readLine()
	fmt.Print("Kilogram Price: ")
	fmt.Scanln(&price)

	return Fruit{name, price}
}

func readFruitName() Fruit {
	var name string

	fmt.Print("Name: ")
	name = readLine()

	return Fruit{name, 0}
}

func readWeighting() Weighting {
	var name string
	var weight float64

	fmt.Print("Fruit: ")
	//name = readLine()
	fmt.Scanln(&name)
	fmt.Print("Weight: ")
	fmt.Scanln(&weight)

	return Weighting{strings.ToLower(name), weight}
}

func readLine() string {
	aux := ""
	fmt.Scanln(&aux)

	return aux/*

	consoleReader := bufio.NewReader(os.Stdin)

	str, _ := consoleReader.ReadString('\n')
	str = strings.ToLower(str)
	str = strings.Trim(str, "\n")

	return str*/
}

func checkError(str string, err error) {
	if err != nil {
		fmt.Println(str, err)
		os.Exit(1)
	}
}
