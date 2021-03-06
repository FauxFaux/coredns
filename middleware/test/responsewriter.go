package test

import (
	"net"

	"github.com/miekg/dns"
)

// ResponseWriter is useful for writing tests. It uses some fixed values for the client. The
// remote will always be 10.240.0.1 and port 40212. The local address is always 127.0.0.1 and
// port 53.
type ResponseWriter struct{}

// LocalAddr returns the local address, always 127.0.0.1:53 (UDP).
func (t *ResponseWriter) LocalAddr() net.Addr {
	ip := net.ParseIP("127.0.0.1")
	port := 53
	return &net.UDPAddr{IP: ip, Port: port, Zone: ""}
}

// RemoteAddr returns the remote address, always 10.240.0.1:40212 (UDP).
func (t *ResponseWriter) RemoteAddr() net.Addr {
	ip := net.ParseIP("10.240.0.1")
	port := 40212
	return &net.UDPAddr{IP: ip, Port: port, Zone: ""}
}

// WriteMsg implement dns.ResponseWriter interface.
func (t *ResponseWriter) WriteMsg(m *dns.Msg) error { return nil }

// Write implement dns.ResponseWriter interface.
func (t *ResponseWriter) Write(buf []byte) (int, error) { return len(buf), nil }

// Close implement dns.ResponseWriter interface.
func (t *ResponseWriter) Close() error { return nil }

// TsigStatus implement dns.ResponseWriter interface.
func (t *ResponseWriter) TsigStatus() error { return nil }

// TsigTimersOnly implement dns.ResponseWriter interface.
func (t *ResponseWriter) TsigTimersOnly(bool) { return }

// Hijack implement dns.ResponseWriter interface.
func (t *ResponseWriter) Hijack() { return }

// RepsponseWrite6 returns fixed client and remote address in IPv6.  The remote
// address is always fe80::42:ff:feca:4c65 and port 40212. The local address
// is always ::1 and port 53.
type ResponseWriter6 struct {
	ResponseWriter
}

// LocalAddr returns the local address, always ::1, port 53 (UDP).
func (t *ResponseWriter6) LocalAddr() net.Addr {
	return &net.UDPAddr{IP: net.ParseIP("::1"), Port: 53, Zone: ""}
}

// RemoteAddr returns the remote address, always fe80::42:ff:feca:4c65 port 40212 (UDP).
func (t *ResponseWriter6) RemoteAddr() net.Addr {
	return &net.UDPAddr{IP: net.ParseIP("fe80::42:ff:feca:4c65"), Port: 40212, Zone: ""}
}
