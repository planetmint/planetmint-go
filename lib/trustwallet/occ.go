package trustwallet

import (
	"fmt"
	"io"
	"time"

	"go.bug.st/serial"
)

const (
	SLIP_END     = 0xC0
	SLIP_ESC     = 0xDB
	SLIP_ESC_END = 0xDC
	SLIP_ESC_ESC = 0xDD
	nsPerUs      = 1000
	nsPerMs      = 8000 * nsPerUs
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
		return 0, fmt.Errorf("unable to open serial port: %v", err)
	}
	defer s.Close()

	fmt.Println("serial connected to port.")

	// Encode payload using SLIP
	if err := encodeSLIP(payloadUnencoded, &payloadSlipEncoded); err != nil {
		return 0, fmt.Errorf("unable to encode SLIP: %v", err)
	}

	// Send encoded payload over serial
	if _, err := s.Write(payloadSlipEncoded); err != nil {
		return 0, fmt.Errorf("unable to write to serial port: %v", err)
	}

	time.Sleep(time.Duration(bufferDelayMs) * time.Millisecond)

	// Read response from serial
	readBuffer := make([]byte, 1)
	encodedResponse := make([]byte, 0)

	slipMsgFramer := 0

	for {
		n, err := s.Read(readBuffer)
		if err != nil && err != io.EOF {
			time.Sleep(10 * time.Millisecond)
			continue
		}
		if n == 0 {
			break
		}

		encodedResponse = append(encodedResponse, readBuffer[0])

		if readBuffer[0] == SLIP_END {
			if slipMsgFramer == 1 {
				break
			} else {
				slipMsgFramer++
			}
		}
		time.Sleep(1 * time.Millisecond)
	}

	// Decode SLIP response
	decodedResponse, err := decodeSLIP(encodedResponse)
	if err != nil {
		return 0, fmt.Errorf("unable to decode SLIP: %v", err)
	}

	// Copy decoded response to outBuffer
	copyLength := min(len(decodedResponse), len(outBuffer))
	copy(outBuffer, decodedResponse[:copyLength])

	return copyLength, nil
}

// encodeSLIP encodes data using SLIP protocol.
func encodeSLIP(data []byte, encoded *[]byte) error {
	*encoded = append(*encoded, SLIP_END)
	for _, b := range data {
		switch b {
		case SLIP_END:
			*encoded = append(*encoded, SLIP_ESC, SLIP_ESC_END)
		case SLIP_ESC:
			*encoded = append(*encoded, SLIP_ESC, SLIP_ESC_ESC)
		default:
			*encoded = append(*encoded, b)
		}
	}
	*encoded = append(*encoded, SLIP_END)
	return nil
}

// decodeSLIP decodes SLIP-encoded data.
func decodeSLIP(encoded []byte) ([]byte, error) {
	// Check for empty input
	if len(encoded) == 0 {
		return nil, fmt.Errorf("encoded data is empty")
	}

	// Remove first and last SLIP_END bytes
	if encoded[0] == SLIP_END {
		encoded = encoded[1:]
	}
	if encoded[len(encoded)-1] == SLIP_END {
		encoded = encoded[:len(encoded)-1]
	}

	decoded := make([]byte, 0, len(encoded))
	esc := false

	for _, b := range encoded {
		switch {
		case b == SLIP_ESC && !esc:
			esc = true
		case b == SLIP_ESC_END && esc:
			decoded = append(decoded, SLIP_END)
			esc = false
		case b == SLIP_ESC_ESC && esc:
			decoded = append(decoded, SLIP_ESC)
			esc = false
		default:
			decoded = append(decoded, b)
			esc = false
		}
	}

	return decoded, nil
}

// min returns the smaller of two integers.
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
