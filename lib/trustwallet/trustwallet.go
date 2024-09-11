package trustwallet

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
	"strings"
	"sync"
)

const (
	PREFIX_IHW      = "/IHW"
	BUFFER_SIZE     = 1024
	BUFFER_DELAY_MS = 200
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
		bufferSize:    BUFFER_SIZE,
		bufferDelayMs: BUFFER_DELAY_MS,
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
		return OSCResponse{}, fmt.Errorf("failed to send message: %v", err)
	}

	if outputLength == 0 {
		return OSCResponse{}, fmt.Errorf("no response received")
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

	message := encodeOSCMessage(address, args...)
	return t.oscSender.SendMessage(message)
}

func (t *Connector) ValiseGet() (string, error) {
	response, err := t.sendOSCMessage(fmt.Sprintf("%s/getSeed", PREFIX_IHW))
	if err != nil {
		return "", err
	}
	if len(response.Data) > 0 {
		return response.Data[0], nil
	}
	return "", fmt.Errorf("no data returned")
}

func (t *Connector) CreateMnemonic() (string, error) {
	response, err := t.sendOSCMessage(fmt.Sprintf("%s/mnemonicToSeed", PREFIX_IHW), int32(1))
	if err != nil {
		return "", err
	}
	if len(response.Data) > 0 {
		return response.Data[0], nil
	}
	return "", fmt.Errorf("no data returned")
}

func (t *Connector) InjectPlanetminkeyToSE050(slot int) (bool, error) {
	response, err := t.sendOSCMessage(fmt.Sprintf("%s/se050InjectSECPKeys", PREFIX_IHW), int32(slot))
	if err != nil {
		return false, err
	}
	if len(response.Data) > 0 {
		return response.Data[0] == "0", nil
	}
	return false, fmt.Errorf("no data returned")
}

func (t *Connector) RecoverFromMnemonic(mnemonic string) (string, error) {
	response, err := t.sendOSCMessage(fmt.Sprintf("%s/mnemonicToSeed", PREFIX_IHW), int32(1), mnemonic)
	if err != nil {
		return "", err
	}
	if len(response.Data) > 0 {
		return response.Data[0], nil
	}
	return "", fmt.Errorf("no data returned")
}

func (t *Connector) GetPlanetmintKeys() (*PlanetMintKeys, error) {
	response, err := t.sendOSCMessage(fmt.Sprintf("%s/getPlntmntKeys", PREFIX_IHW))
	if err != nil {
		return nil, err
	}
	if len(response.Data) < 4 {
		return nil, fmt.Errorf("Trust Wallet not initialized. Please initialize the wallet")
	}
	return &PlanetMintKeys{
		PlanetmintAddress:        response.Data[0],
		ExtendedLiquidPubkey:     response.Data[1],
		ExtendedPlanetmintPubkey: response.Data[2],
		RawPlanetmintPubkey:      response.Data[3],
	}, nil
}

func (t *Connector) GetSeedSE050() (string, error) {
	response, err := t.sendOSCMessage(fmt.Sprintf("%s/se050GetSeed", PREFIX_IHW))
	if err != nil {
		return "", err
	}
	if len(response.Data) > 0 {
		return response.Data[0], nil
	}
	return "", fmt.Errorf("no data returned")
}

func (t *Connector) SignHashWithPlanetmint(dataToSign string) (string, error) {
	response, err := t.sendOSCMessage(fmt.Sprintf("%s/ecdsaSignPlmnt", PREFIX_IHW), dataToSign)
	if err != nil {
		return "", err
	}
	if len(response.Data) > 1 {
		return response.Data[1], nil
	}
	return "", fmt.Errorf("no signature returned")
}

func (t *Connector) SignHashWithRDDL(dataToSign string) (string, error) {
	response, err := t.sendOSCMessage(fmt.Sprintf("%s/ecdsaSignRddl", PREFIX_IHW), dataToSign)
	if err != nil {
		return "", err
	}
	if len(response.Data) > 1 {
		return response.Data[1], nil
	}
	return "", fmt.Errorf("no signature returned")
}

func (t *Connector) CreateOptegaKeypair(ctx int) (string, error) {
	response, err := t.sendOSCMessage(fmt.Sprintf("%s/optigaTrustXCreateSecret", PREFIX_IHW), int32(ctx), "")
	if err != nil {
		return "", err
	}
	if len(response.Data) > 1 {
		return response.Data[1], nil
	}
	return "", fmt.Errorf("no public key returned")
}

func (t *Connector) SignWithOptega(ctx int, dataToSign, pubkey string) (string, error) {
	response, err := t.sendOSCMessage(fmt.Sprintf("%s/optigaTrustXSignMessage", PREFIX_IHW), int32(ctx), dataToSign, pubkey, "")
	if err != nil {
		return "", err
	}
	if len(response.Data) > 1 {
		return response.Data[1], nil
	}
	return "", fmt.Errorf("no signature returned")
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
	response, err := t.sendOSCMessage(fmt.Sprintf("%s/se050CalculateHash", PREFIX_IHW), dataToSign)
	if err != nil {
		return "", err
	}
	if len(response.Data) > 1 {
		return response.Data[1], nil
	}
	return "", fmt.Errorf("no hash returned")
}

func (t *Connector) CreateSE050KeypairNIST(ctx int) (string, error) {
	response, err := t.sendOSCMessage(fmt.Sprintf("%s/se050CreateKeyPair", PREFIX_IHW), int32(ctx), int32(1))
	if err != nil {
		return "", err
	}
	if len(response.Data) > 1 {
		return response.Data[1], nil
	}
	return "", fmt.Errorf("no public key returned")
}

func (t *Connector) GetPublicKeyFromSE050(ctx int) (string, error) {
	response, err := t.sendOSCMessage(fmt.Sprintf("%s/se050GetPublicKey", PREFIX_IHW), int32(ctx))
	if err != nil {
		return "", err
	}
	if len(response.Data) > 1 {
		valid, pubKey := t.UnwrapPublicKey(response.Data[1])
		if !valid {
			return "", fmt.Errorf("inject PlanetMintKey failed: No key found")
		}
		return pubKey, nil
	}
	return "", fmt.Errorf("no public key returned")
}

func (t *Connector) SignWithSE050(dataToSign string, ctx int) (string, error) {
	response, err := t.sendOSCMessage(fmt.Sprintf("%s/se050SignData", PREFIX_IHW), dataToSign, int32(ctx))
	if err != nil {
		return "", err
	}
	if len(response.Data) > 1 {
		return response.Data[1], nil
	}
	return "", fmt.Errorf("no signature returned")
}

func (t *Connector) VerifySE050Signature(dataToSign, signature string, ctx int) (bool, error) {
	response, err := t.sendOSCMessage(fmt.Sprintf("%s/se050VerifySignature", PREFIX_IHW), dataToSign, signature, int32(ctx))
	if err != nil {
		return false, err
	}
	if len(response.Data) > 1 {
		return strconv.ParseBool(response.Data[1])
	}
	return false, fmt.Errorf("no verification result returned")
}

func encodeOSCMessage(address string, args ...interface{}) []byte {
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
			buffer.WriteByte('i')
		case float32:
			buffer.WriteByte('f')
		case string:
			buffer.WriteByte('s')
		}
	}
	buffer.WriteByte(0)
	alignBuffer(&buffer)

	// Write arguments
	for _, arg := range args {
		switch v := arg.(type) {
		case int32:
			binary.Write(&buffer, binary.BigEndian, v)
		case float32:
			binary.Write(&buffer, binary.BigEndian, v)
		case string:
			buffer.WriteString(v)
			buffer.WriteByte(0)
			alignBuffer(&buffer)
		}
	}

	return buffer.Bytes()
}

func alignBuffer(buffer *bytes.Buffer) {
	for buffer.Len()%4 != 0 {
		buffer.WriteByte(0)
	}
}
