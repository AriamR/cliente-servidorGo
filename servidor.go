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

//funcion del servidor el cual recibe 2 parametros devolviendo 2 instancias.
func servidor() {
	s, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		c, err := s.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleClient(c)
	}
}

func handleClient(c net.Conn) {

	//b := make([]byte, 100)
	//bs, err := c.Read(b)

	var client Client
	err := gob.NewDecoder(c).Decode(&client)

	if err != nil {
		fmt.Println(err)
		return
	} else {
		//fmt.Println("Mensaje: ", string(b[:bs]))
		fmt.Println("cliente: ", client.Nombre+"canal:", client.Canal)
		fmt.Println(client.Archivo)
		//fmt.Println("Bytes", bs)
	}
}

func main() {
	go servidor()

	var input string
	fmt.Scanln(&input)
}
