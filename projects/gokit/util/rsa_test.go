package util

import (
	"os"
	"testing"
)

func TestGenRsaKey(t *testing.T) {
	filePath := "../pems"
	bits := 1024

	err := GenRsaKey(bits, filePath)
	if err != nil {
		t.Fatal("create rsa file failed!", err)
	}

	_, err = os.Stat(filePath + "/private.pem")
	if err != nil {
		if !os.IsExist(err) {
			t.Fatal("private.pem is not created!!")
		}
	}

	_, err = os.Stat(filePath + "/public.pem")
	if err != nil {
		if !os.IsExist(err) {
			t.Fatal("public.pem is not created!!")
		}
	}
}
