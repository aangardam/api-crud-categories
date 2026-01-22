package helpers

import (
	"net"
	"strconv"
)

func FindAvaliblePort(defaultPort int) (int, error) {
	for {
		server, err := net.Listen("tcp", net.JoinHostPort("", strconv.Itoa(defaultPort)))
		if err != nil {
			defaultPort++
			continue
		}
		server.Close()
		return defaultPort, nil
	}
}
