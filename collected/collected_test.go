package collected

import (
	"testing"
	"time"
)

var Con Collected = Collected{
		DeviceName:  "en0",
		SnapShotLen: 1024,
		Promiscuous: false,
		Timeout:     30 * time.Second,
	}

func TestCollected_ListenNetFlowV5(t *testing.T) {
	Con.ListenNetFlowV5("udp","6343")
}

func TestCollected_ListenSflowAll(t *testing.T) {
	Con.ListenSflowAll("udp","9999")
}

func TestCollected_ListenSflowCounter(t *testing.T) {
	Con.ListenSflowCounter("udp","9999")
}

func TestCollected_ListenSFlowSample(t *testing.T) {
	Con.ListenSFlowSample("udp","9999")
}
