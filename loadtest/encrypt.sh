#!/usr/bin/env bash

ghz --insecure -n 500000 -c 500 --proto ../dataencrypter/dataencrypter.proto  --call DataEncrypter.Encrypt 0.0.0.0:9097