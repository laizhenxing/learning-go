package util

import (
	"fmt"
	"testing"
	"time"
)

func TestParseToken(t *testing.T) {
	tokenStr, err := GenerateToken("xingxiaoli")
	if err != nil {
		t.Fatal("generate token fail!!", err)
	}

	for i := 0; i < 10; i++ {
		_, err = ParseTokenWithClaims(tokenStr)
		if err != nil {
			t.Fatal(err)
		}
		time.Sleep(time.Second * 1)
	}

}

func TestGenerateToken(t *testing.T) {
	token, err := GenerateToken("xingxiaoli")
	if err != nil {
		t.Fatal("generate token failed!!", err)
	}
	fmt.Println(token)
}
