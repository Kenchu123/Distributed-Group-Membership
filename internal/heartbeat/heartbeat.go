package heartbeat

import (
	"net"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"gitlab.engr.illinois.edu/ckchu2/cs425-mp2/internal/membership"
)

type Heartbeat struct {
	Membership *membership.Membership
	isRunning  bool
	// ticker
	heartbeatTicker     *time.Ticker
	heartbeatTickerDone chan bool

	// udp
	udpServer *UdpServer
}

var lock = &sync.Mutex{}
var instance *Heartbeat

func GetInstance() *Heartbeat {
	lock.Lock()
	defer lock.Unlock()
	if instance == nil {
		instance = New()
	}
	return instance
}

// New creates a new heartbeat
func New() *Heartbeat {
	return &Heartbeat{
		Membership:          nil,
		heartbeatTicker:     nil,
		heartbeatTickerDone: make(chan bool),
		udpServer:           nil,
	}
}

// Start starts the heartbeat
func (h *Heartbeat) Start() {
	if h.isRunning {
		logrus.Warn("Heartbeat is already running")
		return
	}
	var err error
	h.Membership, err = membership.New()
	if err != nil {
		logrus.Errorf("failed to start: %v", err)
		return
	}
	h.udpServer, err = NewUdpServer()
	if err != nil {
		logrus.Errorf("failed to start: %v", err)
		return
	}

	h.isRunning = true
	go h.startHeartbeating()
	go h.startReceiving()
}

func (h *Heartbeat) startHeartbeating() {
	logrus.Info("Start heartbeating")
	h.heartbeatTicker = time.NewTicker(HEARTBEAT_INTERVAL)
	defer h.heartbeatTicker.Stop()
	for {
		select {
		case <-h.heartbeatTickerDone:
			return
		case <-h.heartbeatTicker.C:
			h.sendHeartbeat()
		}
	}
}

func (h *Heartbeat) sendHeartbeat() {
	// update self heartbeat
	h.Membership.IncreaseSelfHeartbeat()
	// TODO: change the list of ips to the target ip
	IPs := []string{"fa23-cs425-8701.cs.illinois.edu", "fa23-cs425-8702.cs.illinois.edu"}
	for i, ip := range IPs {
		if ip == h.Membership.GetName() {
			IPs = append(IPs[:i], IPs[i+1:]...)
			break
		}
	}

	for _, ip := range IPs {
		go func(ip string) {
			client, err := NewUdpClient(ip)
			if err != nil {
				logrus.Errorf("failed to create udp client: %v", err)
				return
			}
			payload, err := membership.Serialize(h.Membership)
			if err != nil {
				logrus.Errorf("failed to serialize membership: %v", err)
				return
			}
			client.Send(payload)
			logrus.Debugf("Sending heartbeat to %s: %s\n", ip, h.Membership)
		}(ip)
	}
}

func (h *Heartbeat) startReceiving() {
	logrus.Info("Start receiving heartbeat")
	h.udpServer.Serve(h.receiveHeartbeat)
}

func (h *Heartbeat) receiveHeartbeat(addr net.Addr, buffer []byte) {
	membership, err := membership.Deserialize(buffer)
	if err != nil {
		logrus.Errorf("failed to deserialize membership: %v", err)
		return
	}
	logrus.Debugf("Received heartbeat from %s: %s\n", addr.String(), membership)
	h.Membership.Update(membership)
}

func (h *Heartbeat) Stop() {
	if !h.isRunning {
		logrus.Warn("Heartbeat is not running")
		return
	}
	h.isRunning = false
	go h.stopHeartbeating()
	go h.stopReceiving()
}

func (h *Heartbeat) stopHeartbeating() {
	logrus.Info("Stop heartbeating")
	h.heartbeatTickerDone <- true
}

func (h *Heartbeat) stopReceiving() {
	logrus.Info("Stop receiving heartbeat")
	h.udpServer.Stop()
}
