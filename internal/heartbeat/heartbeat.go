package heartbeat

import (
	"net"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"gitlab.engr.illinois.edu/ckchu2/cs425-mp2/internal/membership"
)

type Heartbeat struct {
	IsRunning bool

	Membership *membership.Membership

	// ticker
	heartbeatTicker     *time.Ticker
	heartbeatTickerDone chan bool

	// udp
	udpServer *UdpServer

	// failure detector
	failureDetectTicker     *time.Ticker
	failureDetectTickerDone chan bool

	// cleanup
	cleanupTicker     *time.Ticker
	cleanupTickerDone chan bool
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
		Membership:              nil,
		heartbeatTicker:         nil,
		heartbeatTickerDone:     make(chan bool),
		udpServer:               nil,
		failureDetectTicker:     nil,
		failureDetectTickerDone: make(chan bool),
	}
}

// Start starts the heartbeat
func (h *Heartbeat) Start() {
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
	h.IsRunning = true
	go h.startHeartbeating()
	go h.startReceiving()
	go h.startDetectingFailure()
	go h.startCleaningUp()
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
	hostnames := h.Membership.GetHeartbeatTargetMembers()
	logrus.Debug("Heartbeat target members: ", hostnames)
	for _, hostname := range hostnames {
		go func(hostname string) {
			client, err := NewUdpClient(hostname)
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
			logrus.Debugf("Sending heartbeat to %s: %s", hostname, h.Membership)
		}(hostname)
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
	logrus.Debugf("Received heartbeat from %s: %s", addr.String(), membership)
	h.Membership.Update(membership)
}

func (h *Heartbeat) startDetectingFailure() {
	logrus.Info("Start detecting failure")
	h.failureDetectTicker = time.NewTicker(FAILURE_DETECT_INTERVAL)
	defer h.failureDetectTicker.Stop()
	for {
		select {
		case <-h.failureDetectTickerDone:
			return
		case <-h.failureDetectTicker.C:
			h.detectFailure()
		}
	}
}

func (h *Heartbeat) detectFailure() {
	logrus.Debug("Detecting failure")
	h.Membership.DetectFailure(FAILURE_DETECT_TIMEOUT)
}

func (h *Heartbeat) startCleaningUp() {
	logrus.Info("Start cleaning up Membership")
	h.cleanupTicker = time.NewTicker(MEMBER_CLEANUP_INTERVAL)
	defer h.cleanupTicker.Stop()
	for {
		select {
		case <-h.cleanupTickerDone:
			return
		case <-h.cleanupTicker.C:
			h.Membership.CleanUp(MEMBER_CLEANUP_TIMEOUT)
		}
	}
}

// Stop stops the heartbeat
func (h *Heartbeat) Stop() {
	h.IsRunning = false
	go h.stopHeartbeating()
	go h.stopReceiving()
	go h.stopDetectingFailure()
	go h.stopCleaningUp()
}

func (h *Heartbeat) stopHeartbeating() {
	logrus.Info("Stop heartbeating")
	h.heartbeatTickerDone <- true
}

func (h *Heartbeat) stopReceiving() {
	logrus.Info("Stop receiving heartbeat")
	h.udpServer.Stop()
}

func (h *Heartbeat) stopDetectingFailure() {
	logrus.Info("Stop detecting failure")
	h.failureDetectTickerDone <- true
}

func (h *Heartbeat) stopCleaningUp() {
	logrus.Info("Stop cleaning up Membership")
	h.cleanupTickerDone <- true
}
