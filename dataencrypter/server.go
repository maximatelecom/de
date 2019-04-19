package dataencrypter

import (
	"context"
	"dataencrypter/encrypt"
	"encoding/base64"
	"errors"

	log "github.com/sirupsen/logrus"
)

type Server struct {
	cryptoKey  []byte
	initVector []byte
}

//NewServer returns new Server instance
func NewServer(cryptoKey, initVector []byte) *Server {
	return &Server{
		cryptoKey:  cryptoKey,
		initVector: initVector,
	}
}

//Encrypt decrypts encrypted mac
func (s *Server) Encrypt(ctx context.Context, req *EncryptRequest) (*EncryptResponse, error) {
	encryptedData, err := encrypt.Encrypt([]byte(req.GetData()), s.cryptoKey, s.initVector)
	if err != nil {
		log.
			WithField("request", req).
			WithError(err).
			Error("unable to encrypt data")

		return nil, errors.New("error while encrypting data")
	}
	ivSize := len(s.initVector)
	encryptedData = encryptedData[ivSize:]
	encodedData := base64.StdEncoding.EncodeToString(encryptedData)
	return &EncryptResponse{EncryptedData: encodedData}, nil
}

func (s *Server) Decrypt(ctx context.Context, req *DecryptRequest) (*DecryptResponse, error) {
	log.
		WithField("request", req.EncryptedData).
		Debug("decrypt request")
	decodedData, err := base64.StdEncoding.DecodeString(req.GetEncryptedData())
	if err != nil {
		return nil, err
	}
	encryptedData := append(s.initVector, decodedData...)
	data, err := encrypt.Decrypt(encryptedData, s.cryptoKey)
	if err != nil {
		log.
			WithField("request", req).
			WithError(err).
			Error("unable to decrypt data")

		return nil, errors.New("error while decrypting data")
	}

	return &DecryptResponse{Data: string(data)}, nil
}
