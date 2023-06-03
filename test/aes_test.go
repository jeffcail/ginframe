package test

import (
	"fmt"
	"github.com/jeffcail/ginframe/server-common/utils/encry"
	"testing"
)

func TestAesEncryptDecrypt(t *testing.T) {
	instance, err := encry.NewAesEncryptInstance("bGcGfWb3Kg2s4gcG", "aebksHkG4jAEk2Ag")
	if err != nil {
		t.Fatal(err)
	}
	encrypt, err := instance.AesBase64Encrypt("1234566dfsdasd")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(encrypt)

	fmt.Println("=======================================")

	decrypt, err := instance.AesBase64Decrypt(encrypt)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(decrypt)
}
