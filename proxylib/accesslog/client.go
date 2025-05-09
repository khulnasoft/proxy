// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Khulnasoft

package accesslog

import (
	"net"
	"sync"
	"sync/atomic"

	"github.com/golang/protobuf/proto"
	"github.com/sirupsen/logrus"

	khulnasoft "github.com/khulnasoft/proxy/go/khulnasoft/api"
	"github.com/khulnasoft/proxy/proxylib/proxylib"
)

type Client struct {
	connected uint32 // Accessed atomically without locking
	path      string
	mutex     sync.Mutex                   // Used to protect opening the connection
	conn      atomic.Pointer[net.UnixConn] // Read atomically without locking
}

func (cl *Client) connect() *net.UnixConn {
	if cl.path == "" {
		return nil
	}

	if atomic.LoadUint32(&cl.connected) > 0 {
		// Guaranteed to be non-nil
		return cl.conn.Load()
	}

	cl.mutex.Lock()
	defer cl.mutex.Unlock()

	conn := cl.conn.Load()

	// Did someone else connect while we were contending on the lock?
	// cl.connected may be written to by others concurrently
	if atomic.LoadUint32(&cl.connected) > 0 {
		return conn
	}

	if conn != nil {
		conn.Close() // not setting conn to nil!
	}
	logrus.Debugf("Accesslog: Connecting to Khulnasoft access log socket: %s", cl.path)
	conn, err := net.DialUnix("unixpacket", nil, &net.UnixAddr{Name: cl.path, Net: "unixpacket"})
	if err != nil {
		logrus.WithError(err).Error("Accesslog: DialUnix() failed")
		return nil
	}

	cl.conn.Store(conn)

	// Always have a non-nil 'cl.conn' after 'cl.connected' is set for the first time!
	atomic.StoreUint32(&cl.connected, 1)
	return conn
}

func (cl *Client) Log(pblog *khulnasoft.LogEntry) {
	if conn := cl.connect(); conn != nil {
		// Encode
		logmsg, err := proto.Marshal(pblog)
		if err != nil {
			logrus.WithError(err).Error("Accesslog: Protobuf marshaling error")
			return
		}

		// Write
		_, err = conn.Write(logmsg)
		if err != nil {
			logrus.WithError(err).Error("Accesslog: Write() failed")
			atomic.StoreUint32(&cl.connected, 0) // Mark connection as broken
		}
	} else {
		logrus.Debugf("Accesslog: No connection, cannot send: %s", pblog.String())
	}
}

func (c *Client) Path() string {
	return c.path
}

func NewClient(accessLogPath string) proxylib.AccessLogger {
	client := &Client{
		path: accessLogPath,
	}
	client.connect()
	return client
}

func (cl *Client) Close() {
	conn := cl.conn.Load()
	if conn != nil {
		conn.Close()
	}
}
