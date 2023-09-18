package heartbeat

import (
	"fmt"
	"math/rand"
	"net"
)

type UdpClient struct {
	conn *net.UDPConn
}

// NewUdpClient creates a new UDP client
func NewUdpClient(hostname string, port string) (*UdpClient, error) {
	addr, err := net.ResolveUDPAddr("udp", hostname+":"+port)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve udp address: %w", err)
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		return nil, fmt.Errorf("failed to create udp client on %s: %w", addr, err)
	}
	return &UdpClient{
		conn: conn,
	}, nil
}

// Send sends a heartbeat to the UDP server
func (u *UdpClient) Send(msg []byte, dropRate float32) (int, error) {
	if dropRate > 0 {
		rand := rand.Float32()
		if rand < dropRate {
			return 0, fmt.Errorf("Package dropped")
		}
	}
	return u.conn.Write(msg)
}
