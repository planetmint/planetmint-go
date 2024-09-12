package trustwallet

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
)

const (
	PrefixIhw     = "/IHW"
	BufferSize    = 1024
	BufferDelayMs = 200
)

type OSCResponse struct {
	Command string
	Data    []string
}

type OSCMessageSender struct {
	portName      []byte
	bufferSize    int
	bufferDelayMs int
}

func NewOSCMessageSender(portName string) (*OSCMessageSender, error) {
	return &OSCMessageSender{
		portName:      []byte(portName),
		bufferSize:    BufferSize,
		bufferDelayMs: BufferDelayMs,
	}, nil
}

func (s *OSCMessageSender) SendMessage(message []byte) (OSCResponse, error) {
	outputBuffer := make([]byte, s.bufferSize)

	// Call occDo function
	outputLength, err := occDo(
		message,
		s.bufferDelayMs,
		string(s.portName),
		outputBuffer,
	)

	if err != nil {
		return OSCResponse{}, fmt.Errorf("failed to send message: %w", err)
	}

	if outputLength == 0 {
		return OSCResponse{}, errors.New("no response received")
	}

	// Extract the information from the output buffer
	return extractInformation(outputBuffer[:outputLength])
}

func extractInformation(responseBytes []byte) (OSCResponse, error) {
	decodedString := string(bytes.Trim(responseBytes, "\x00"))
	parts := strings.Split(decodedString, "\x00")

	var response OSCResponse
	if len(parts) > 0 {
		commandPart := parts[0]
		dataParts := parts[1:]

		if strings.Contains(commandPart, ",") {
			splitCmd := strings.SplitN(commandPart, ",", 2)
			response.Command = strings.TrimSpace(splitCmd[0])
			dataParts = append([]string{splitCmd[1]}, dataParts...)
		} else {
			response.Command = strings.TrimSpace(commandPart)
		}

		response.Data = make([]string, 0, len(dataParts))
		for _, part := range dataParts {
			if trimmed := strings.TrimSpace(part); trimmed != "" {
				response.Data = append(response.Data, trimmed)
			}
		}
	}

	if len(response.Data) == 0 {
		response.Data = []string{"No valid data found."}
	}

	return response, nil
}
