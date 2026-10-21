package services

import (
	"bufio"
	"encoding/binary"
	"io"
	"net"
	"strings"
	"testing"
	"time"
)

func TestClamAVScannerAcceptsOnlyOK(t *testing.T) {
	for _, test := range []struct {
		response string
		want     error
	}{
		{"stream: OK\x00", nil},
		{"stream: Eicar-Test-Signature FOUND\x00", ErrAttachmentUnsafe},
		{"stream: ERROR\x00", ErrAttachmentScanUnavailable},
	} {
		address, closeServer := fakeClamAV(t, test.response)
		err := (ClamAVScanner{Address: address, Timeout: time.Second}).Scan(strings.NewReader("file"))
		closeServer()
		if err != test.want {
			t.Fatalf("response %q: error = %v, want %v", test.response, err, test.want)
		}
	}
}

func fakeClamAV(t *testing.T, response string) (string, func()) {
	t.Helper()
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		defer close(done)
		connection, err := listener.Accept()
		if err != nil {
			return
		}
		defer connection.Close()
		reader := bufio.NewReader(connection)
		_, _ = reader.ReadString(0)
		for {
			var size uint32
			if binary.Read(reader, binary.BigEndian, &size) != nil || size == 0 {
				break
			}
			_, _ = io.CopyN(io.Discard, reader, int64(size))
		}
		_, _ = connection.Write([]byte(response))
	}()
	return listener.Addr().String(), func() { _ = listener.Close(); <-done }
}
