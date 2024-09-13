package main

import (
	"fmt"
	"log"
	"net"

	"github.com/Mau005/GoOlds/networks"
	"github.com/Mau005/GoOlds/utils"
)

func main() {

	ld, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalln(err)
	}
	for {
		conn, err := ld.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(con net.Conn) {
	defer con.Close()
	bufer := make([]byte, 1024)
	_, err := con.Read(bufer)
	if err != nil {
		log.Println(err)
		return
	}
	var msg networks.NetworkMessage
	msg.MsgBuf = bufer
	msg.ReadPos = 2
	potId := msg.GetU16()
	if potId == utils.ProtocolLogin {
		msg.SkipBytes(15)
		accNumber := msg.GetU32()
		password := msg.GetString()
		fmt.Println(con.RemoteAddr().String())
		fmt.Println(accNumber, "..", password)

	} else {
		fmt.Println("no entro al protocolo")
	}
}
