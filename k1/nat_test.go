package k1

import (
	"bytes"
	"math/rand"
	"net"
	"testing"
)

func TestNat(t *testing.T) {
	var from uint16 = 10
	var to uint16 = 20
	nat := newNat(from, to)

	srcIP := net.ParseIP("127.0.0.1")
	dstIP := srcIP

	// map all ports
	for i := from; i < to; i++ {
		port := nat.allocSession(srcIP, dstIP, i, i)
		if i != port {
			t.Error("alloc session failed")
			return
		}
	}

	// release all sessions
	for i := from; i < to; i++ {
		nat.releaseSession(i)
	}
	if nat.count() != 0 {
		t.Error("release session failed")
		return
	}

	// test get session
	srcPort := uint16(rand.Int())
	dstPort := uint16(rand.Int())
	port := nat.allocSession(srcIP, dstIP, srcPort, dstPort)
	session := nat.getSession(port)
	if session == nil {
		t.Error("get session failed")
		return
	}

	if !bytes.Equal(session.srcIP, srcIP) || !bytes.Equal(session.dstIP, dstIP) ||
		session.srcPort != srcPort || session.dstPort != dstPort {
		t.Error("check session failed")
		return
	}
}

func BenchmarkNat(b *testing.B) {
	var from uint16 = 10000
	var to uint16 = 60000
	nat := newNat(from, to)

	srcIP := net.ParseIP("127.0.0.1")
	dstIP := srcIP

	// map all ports
	for i := from; i < to; i++ {
		port := nat.allocSession(srcIP, dstIP, i, i)
		if i != port {
			b.Error("alloc session failed")
		}
	}

	// release all sessions
	for i := from; i < to; i++ {
		nat.releaseSession(i)
	}
	if nat.count() != 0 {
		b.Error("release session failed")
	}
}