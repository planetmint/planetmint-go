package trustwallet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync"
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

type Connector struct {
	oscSender *OSCMessageSender
	mu        sync.Mutex
}

func NewTrustWalletConnector(portName string) (*Connector, error) {
	sender, err := NewOSCMessageSender(portName)
	if err != nil {
		return nil, err
	}
	return &Connector{
		oscSender: sender,
	}, nil
}

func (t *Connector) sendOSCMessage(address string, args ...interface{}) (OSCResponse, error) {
	t.mu.Lock()
	defer t.mu.Unlock()

	message, err := encodeOSCMessage(address, args...)
	if err != nil {
		return OSCResponse{}, err
	}

	return t.oscSender.SendMessage(message)
}

func (t *Connector) ValiseGet() (string, error) {
	response, err := t.sendOSCMessage(PrefixIhw + "/getSeed")
	if err != nil {
		return "", err
	}
	if len(response.Data) > 0 {
		return response.Data[0], nil
	}
	return "", errors.New("no data returned")
}

func (t *Connector) CreateMnemonic() (string, error) {
	response, err := t.sendOSCMessage(PrefixIhw+"/mnemonicToSeed", int32(1))
	if err != nil {
		return "", err
	}
	if len(response.Data) > 0 {
		return response.Data[0], nil
	}
	return "", errors.New("no data returned")
}

func (t *Connector) InjectPlanetminkeyToSE050(slot int) (bool, error) {
	response, err := t.sendOSCMessage(PrefixIhw+"/se050InjectSECPKeys", int32(slot))
	if err != nil {
		return false, err
	}
	if len(response.Data) > 0 {
		return response.Data[0] == "0", nil
	}
	return false, errors.New("no data returned")
}

func (t *Connector) RecoverFromMnemonic(mnemonic string) (string, error) {
	response, err := t.sendOSCMessage(PrefixIhw+"/mnemonicToSeed", int32(1), mnemonic)
	if err != nil {
		return "", err
	}
	if len(response.Data) > 0 {
		return response.Data[0], nil
	}
	return "", errors.New("no data returned")
}

func (t *Connector) GetPlanetmintKeys() (*PlanetMintKeys, error) {
	response, err := t.sendOSCMessage(PrefixIhw + "/getPlntmntKeys")
	if err != nil {
		return nil, err
	}
	if len(response.Data) < 4 {
		return nil, errors.New("trust wallet not initialized. Please initialize the wallet")
	}
	return &PlanetMintKeys{
		PlanetmintAddress:        response.Data[0],
		ExtendedLiquidPubkey:     response.Data[1],
		ExtendedPlanetmintPubkey: response.Data[2],
		RawPlanetmintPubkey:      response.Data[3],
	}, nil
}

func (t *Connector) GetSeedSE050() (string, error) {
	response, err := t.sendOSCMessage(PrefixIhw + "/se050GetSeed")
	if err != nil {
		return "", err
	}
	if len(response.Data) > 0 {
		return response.Data[0], nil
	}
	return "", errors.New("no data returned")
}

func (t *Connector) SignHashWithPlanetmint(dataToSign string) (string, error) {
	response, err := t.sendOSCMessage(PrefixIhw+"/ecdsaSignPlmnt", dataToSign)
	if err != nil {
		return "", err
	}
	if len(response.Data) > 1 {
		return response.Data[1], nil
	}
	return "", errors.New("no signature returned")
}

func (t *Connector) SignHashWithRDDL(dataToSign string) (string, error) {
	response, err := t.sendOSCMessage(PrefixIhw+"/ecdsaSignRddl", dataToSign)
	if err != nil {
		return "", err
	}
	if len(response.Data) > 1 {
		return response.Data[1], nil
	}
	return "", errors.New("no signature returned")
}

func (t *Connector) CreateOptegaKeypair(ctx int) (string, error) {
	response, err := t.sendOSCMessage(PrefixIhw+"/optigaTrustXCreateSecret", int32(ctx), "")
	if err != nil {
		return "", err
	}
	if len(response.Data) > 1 {
		return response.Data[1], nil
	}
	return "", errors.New("no public key returned")
}

func (t *Connector) SignWithOptega(ctx int, dataToSign, pubkey string) (string, error) {
	response, err := t.sendOSCMessage(PrefixIhw+"/optigaTrustXSignMessage", int32(ctx), dataToSign, pubkey, "")
	if err != nil {
		return "", err
	}
	if len(response.Data) > 1 {
		return response.Data[1], nil
	}
	return "", errors.New("no signature returned")
}

func (t *Connector) UnwrapPublicKey(publicKey string) (bool, string) {
	length := len(publicKey)
	if length == 136 || length == 130 {
		return true, publicKey[len(publicKey)-128:]
	} else if length == 128 {
		return true, publicKey
	}
	return false, publicKey
}

func (t *Connector) CalculateHash(dataToSign string) (string, error) {
	response, err := t.sendOSCMessage(PrefixIhw+"/se050CalculateHash", dataToSign)
	if err != nil {
		return "", err
	}
	if len(response.Data) > 1 {
		return response.Data[1], nil
	}
	return "", errors.New("no hash returned")
}

func (t *Connector) CreateSE050KeypairNIST(ctx int) (string, error) {
	response, err := t.sendOSCMessage(PrefixIhw+"/se050CreateKeyPair", int32(ctx), int32(1))
	if err != nil {
		return "", err
	}
	if len(response.Data) > 1 {
		return response.Data[1], nil
	}
	return "", errors.New("no public key returned")
}

func (t *Connector) GetPublicKeyFromSE050(ctx int) (string, error) {
	response, err := t.sendOSCMessage(PrefixIhw+"/se050GetPublicKey", int32(ctx))
	if err != nil {
		return "", err
	}
	if len(response.Data) > 1 {
		valid, pubKey := t.UnwrapPublicKey(response.Data[1])
		if !valid {
			return "", errors.New("inject PlanetMintKey failed: No key found")
		}
		return pubKey, nil
	}
	return "", errors.New("no public key returned")
}

func (t *Connector) SignWithSE050(dataToSign string, ctx int) (string, error) {
	response, err := t.sendOSCMessage(PrefixIhw+"/se050SignData", dataToSign, int32(ctx))
	if err != nil {
		return "", err
	}
	if len(response.Data) > 1 {
		return response.Data[1], nil
	}
	return "", errors.New("no signature returned")
}

func (t *Connector) VerifySE050Signature(dataToSign, signature string, ctx int) (bool, error) {
	response, err := t.sendOSCMessage(PrefixIhw+"/se050VerifySignature", dataToSign, signature, int32(ctx))
	if err != nil {
		return false, err
	}
	if len(response.Data) > 1 {
		return strconv.ParseBool(response.Data[1])
	}
	return false, errors.New("no verification result returned")
}

func encodeOSCMessage(address string, args ...interface{}) (returnBytes []byte, err error) {
	var buffer bytes.Buffer

	// Write address
	buffer.WriteString(address)
	buffer.WriteByte(0)
	alignBuffer(&buffer)

	// Write type tags
	buffer.WriteByte(',')
	for _, arg := range args {
		switch arg.(type) {
		case int32:
			err = buffer.WriteByte('i')
		case float32:
			err = buffer.WriteByte('f')
		case string:
			err = buffer.WriteByte('s')
		}
		if err != nil {
			return buffer.Bytes(), err
		}
	}

	buffer.WriteByte(0)
	alignBuffer(&buffer)

	// Write arguments
	for _, arg := range args {
		switch v := arg.(type) {
		case int32:
			err = binary.Write(&buffer, binary.BigEndian, v)
		case float32:
			err = binary.Write(&buffer, binary.BigEndian, v)
		case string:
			_, err = buffer.WriteString(v)
			if err != nil {
				return buffer.Bytes(), err
			}
			err = buffer.WriteByte(0)
			alignBuffer(&buffer)
		}
		if err != nil {
			return buffer.Bytes(), err
		}
	}

	return buffer.Bytes(), nil
}

func alignBuffer(buffer *bytes.Buffer) {
	for buffer.Len()%4 != 0 {
		buffer.WriteByte(0)
	}
}
