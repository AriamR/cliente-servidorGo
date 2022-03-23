package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

type Client struct {
	Nombre  string
	Canal   []string
	Archivo string
}

func cliente(client Client) {
	c, err := net.Dial("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = gob.NewEncoder(c).Encode(client)
	msg := "Deseo subscribirme al canal 2\n"
	msg1 := "Archivo.png"
	fmt.Println(msg, msg1)
	fmt.Println(msg)
	c.Write([]byte(msg))
	if err != nil {
		fmt.Println(err)
	}
	c.Close()
}

func main() {

	client := Client{
		Nombre: "Felipe\n",
		Canal: []string{
			"subscripcion a channel 2",
		},
		Archivo: "soy un archivo png",
	}
	go cliente(client)

	var input string
	fmt.Scanln(&input)
}
