package trustwallet

import (
	"errors"
	"fmt"
	"io"
	"time"

	"go.bug.st/serial"
)

const (
	SlipEnd    = 0xC0
	SlipEsc    = 0xDB
	SlipEscEnd = 0xDC
	SlipEscEsc = 0xDD
	nsPerUs    = 1000
	nsPerMs    = 8000 * nsPerUs
)

// occDo performs the operations to send and receive data over serial port.
func occDo(data []byte, bufferDelayMs int, portName string, outBuffer []byte) (int, error) {
	// Initialize unencoded and encoded payloads
	payloadUnencoded := make([]byte, 0, len(data))
	payloadSlipEncoded := make([]byte, 0, len(data)*2)

	// Copy data to payloadUnencoded
	payloadUnencoded = append(payloadUnencoded, data...)

	// Open serial port
	mode := &serial.Mode{BaudRate: 115200}
	s, err := serial.Open(portName, mode)
	if err != nil {
		return 0, fmt.Errorf("unable to open serial port: %w", err)
	}
	defer s.Close()

	// Encode payload using SLIP
	encodeSLIP(payloadUnencoded, &payloadSlipEncoded)

	// Send encoded payload over serial
	if _, err := s.Write(payloadSlipEncoded); err != nil {
		return 0, fmt.Errorf("unable to write to serial port: %w", err)
	}

	time.Sleep(time.Duration(bufferDelayMs) * time.Millisecond)

	// Read response from serial
	readBuffer := make([]byte, 1)
	encodedResponse := make([]byte, 0)

	slipMsgFramer := 0

	for {
		n, err := s.Read(readBuffer)
		if err != nil && !errors.Is(err, io.EOF) {
			time.Sleep(10 * time.Millisecond)
			continue
		}
		if n == 0 {
			break
		}

		encodedResponse = append(encodedResponse, readBuffer[0])

		if readBuffer[0] == SlipEnd {
			if slipMsgFramer == 1 {
				break
			}
			slipMsgFramer++
		}
		time.Sleep(1 * time.Millisecond)
	}

	// Decode SLIP response
	decodedResponse, err := decodeSLIP(encodedResponse)
	if err != nil {
		return 0, fmt.Errorf("unable to decode SLIP: %w", err)
	}

	// Copy decoded response to outBuffer
	copyLength := min(len(decodedResponse), len(outBuffer))
	copy(outBuffer, decodedResponse[:copyLength])

	return copyLength, nil
}

// encodeSLIP encodes data using SLIP protocol.
func encodeSLIP(data []byte, encoded *[]byte) {
	*encoded = append(*encoded, SlipEnd)
	for _, b := range data {
		switch b {
		case SlipEnd:
			*encoded = append(*encoded, SlipEsc, SlipEscEnd)
		case SlipEsc:
			*encoded = append(*encoded, SlipEsc, SlipEscEsc)
		default:
			*encoded = append(*encoded, b)
		}
	}
	*encoded = append(*encoded, SlipEnd)
}

// decodeSLIP decodes SLIP-encoded data.
func decodeSLIP(encoded []byte) ([]byte, error) {
	// Check for empty input
	if len(encoded) == 0 {
		return nil, errors.New("encoded data is empty")
	}

	// Remove first and last SLIP_END bytes
	if encoded[0] == SlipEnd {
		encoded = encoded[1:]
	}
	if encoded[len(encoded)-1] == SlipEnd {
		encoded = encoded[:len(encoded)-1]
	}

	decoded := make([]byte, 0, len(encoded))
	esc := false

	for _, b := range encoded {
		switch {
		case b == SlipEsc && !esc:
			esc = true
		case b == SlipEscEnd && esc:
			decoded = append(decoded, SlipEnd)
			esc = false
		case b == SlipEscEsc && esc:
			decoded = append(decoded, SlipEsc)
			esc = false
		default:
			decoded = append(decoded, b)
			esc = false
		}
	}

	return decoded, nil
}
