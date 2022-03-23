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
	msg := "Deseo subscribirme al canal 1 \n"
	msg1 := "Archivo.txt"
	fmt.Println(msg, msg1)
	c.Write([]byte(msg))
	if err != nil {
		fmt.Println(err)
	}
	c.Close()
}

func main() {

	client := Client{
		Nombre: "maira\n",
		Canal: []string{
			"subscripcion a channel 1",
		},
		Archivo: "soy un archivo de texto",
	}

	go cliente(client)

	var input string
	fmt.Scanln(&input)
}
