package main

import (
	"context"
	"dataencrypter/config"
	"dataencrypter/dataencrypter"
	"google.golang.org/grpc"
	"testing"
)

func TestEncrypt(t *testing.T) {
	go main()
	v, err := config.NewViper(configFilePath)
	handleFatal(err, "unable to load config")
	conn, err := grpc.Dial(v.GetString("grpc.server_addr"), grpc.WithInsecure())
	defer conn.Close()
	handleFatal(err, "unable to dial grpc")
	de := dataencrypter.NewDataEncrypterClient(conn)
	respEncrypt, err := de.Encrypt(context.Background(), &dataencrypter.EncryptRequest{Data: "hey"})
	if err != nil {
		t.Fatal("error encrypting data")
	}
	if respEncrypt.EncryptedData == "" {
		t.Fatal("empty encrypt response")
	}
	req := dataencrypter.DecryptRequest{EncryptedData: respEncrypt.EncryptedData}
	respDecrypt, err := de.Decrypt(context.Background(), &req)
	if err != nil {
		t.Fatal("error decrypting data")
	}
	if respDecrypt.Data != "hey" {
		t.Fatalf("wrong decrypt: %s != %s", "hey", respDecrypt.Data)
	}
}

func BenchmarkEncrypt(b *testing.B)  {
	key := []byte("1234567891234567")
	iv := []byte("123456123456")
	s := dataencrypter.NewServer(key, iv)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		s.Encrypt(context.Background(), &dataencrypter.EncryptRequest{Data: "hey"})
	}
}

func BenchmarkDecrypt(b *testing.B)  {
	key := []byte("1234567891234567")
	iv := []byte("123456123456")
	s := dataencrypter.NewServer(key, iv)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		s.Decrypt(context.Background(), &dataencrypter.DecryptRequest{EncryptedData: "FUoVmLpBKKFc3+KkwF37x0qkiw=="})
	}
}
