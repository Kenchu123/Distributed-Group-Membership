package heartbeat

import (
	"fmt"
	"net"

	"github.com/sirupsen/logrus"
)

type UdpServer struct {
	conn *net.UDPConn
}

// NewUdpServer creates a new UDP server
func NewUdpServer(port string) (*UdpServer, error) {
	addr, err := net.ResolveUDPAddr("udp", ":"+port)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve udp address: %w", err)
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		return nil, fmt.Errorf("failed to create udp server on %s: %w", addr, err)
	}
	return &UdpServer{
		conn: conn,
	}, nil
}

// serve serves the UDP server
func (u *UdpServer) Serve(handle func(net.Addr, []byte)) {
	logrus.Infof("UDP server is listening on %s", u.conn.LocalAddr().String())
	defer u.conn.Close()
	for {
		buffer := make([]byte, 1024)
		_, addr, err := u.conn.ReadFrom(buffer)
		if err != nil {
			logrus.Errorf("failed to read from udp server: %v", err)
			break
		}
		go handle(addr, buffer)
	}
}

// Stop stops the UDP server
func (u *UdpServer) Stop() {
	u.conn.Close()
}
