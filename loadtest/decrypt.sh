#!/usr/bin/env bash

ghz --insecure -n 500000 -c 1000 --proto ../dataencrypter/dataencrypter.proto  --call DataEncrypter.Decrypt -d '{"encrypted_data":"FUoVmLpBKKFc3+KkwF37x0qkiw=="}'  0.0.0.0:9097