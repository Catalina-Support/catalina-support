package services

import (
	"bufio"
	"encoding/binary"
	"errors"
	"io"
	"net"
	"strings"
	"time"
)

// ClamAVScanner usa INSTREAM y falla cerrado ante virus, error o respuesta desconocida.
type ClamAVScanner struct {
	Address string
	Timeout time.Duration
}

func (s ClamAVScanner) Scan(source io.Reader) error {
	timeout := s.Timeout
	if timeout <= 0 {
		timeout = 30 * time.Second
	}
	connection, err := net.DialTimeout("tcp", s.Address, timeout)
	if err != nil {
		return ErrAttachmentScanUnavailable
	}
	defer connection.Close()
	_ = connection.SetDeadline(time.Now().Add(timeout))
	if _, err = connection.Write([]byte("zINSTREAM\x00")); err != nil {
		return ErrAttachmentScanUnavailable
	}
	buffer := make([]byte, 32*1024)
	for {
		n, readErr := source.Read(buffer)
		if n > 0 {
			var size [4]byte
			binary.BigEndian.PutUint32(size[:], uint32(n))
			if _, err = connection.Write(size[:]); err != nil {
				return ErrAttachmentScanUnavailable
			}
			if _, err = connection.Write(buffer[:n]); err != nil {
				return ErrAttachmentScanUnavailable
			}
		}
		if errors.Is(readErr, io.EOF) {
			break
		}
		if readErr != nil {
			return ErrAttachmentScanUnavailable
		}
	}
	if _, err = connection.Write([]byte{0, 0, 0, 0}); err != nil {
		return ErrAttachmentScanUnavailable
	}
	response, err := bufio.NewReader(connection).ReadString(0)
	if err != nil && !errors.Is(err, io.EOF) {
		return ErrAttachmentScanUnavailable
	}
	response = strings.TrimSpace(strings.TrimSuffix(response, "\x00"))
	if strings.HasSuffix(response, " OK") {
		return nil
	}
	if strings.Contains(response, " FOUND") {
		return ErrAttachmentUnsafe
	}
	return ErrAttachmentScanUnavailable
}
