package server

import (
	"io"
	"log"
	"net"
	"strconv"

	"github.com/benotsman-youssuf/redis_clone/config"
)

func RunSyncTCPServer() {
	log.Println("starting a synchronous TCP server on", config.HOST, config.PORT)

	var con_clients int = 0

	lsnr, err := net.Listen("tcp", config.HOST+":"+strconv.Itoa(config.PORT))
	if err != nil {
		panic(err)
	}

	for {
		c, err := lsnr.Accept()
		if err != nil {
			panic(err)
		}

		con_clients += 1
		log.Println("client connected with address: ", c.RemoteAddr(), " concurrent clients: ", con_clients)

		for {
			cmd, err := readCommand(c)
			if err != nil {
				c.Close()
				con_clients -= 1
				log.Println("client disconnected: ", c.RemoteAddr(), "concurrent clients: ", con_clients)
				if err == io.EOF {
					break
				}
				log.Print("err: ", err)
			}
			log.Println("command: ", cmd)
			if err = respond(cmd, c); err != nil {
				log.Print("err write:", err)
			}

		}
	}
}

func readCommand(c net.Conn) (string, error) {
	var buf []byte = make([]byte, 512)
	n, err := c.Read(buf[:])
	if err != nil {
		return "", err
	}
	return string(buf[:n]), nil
}

func respond(cmd string, c net.Conn) error {
	if _, err := c.Write([]byte(cmd)); err != nil {
		return err
	}
	return nil
}
