package test

import (
	"fmt"
	"github.com/jeffcail/ginframe/utils/regmatch"
	"testing"
)

func TestRegMatchPhone(t *testing.T) {
	ok := regmatch.RegMatchPhone("12346789")
	fmt.Println(ok)
}

// :010-1234567
func TestRegLandLineNumber(t *testing.T) {
	ok := regmatch.RegLandLineNumber("010-4844")
	fmt.Println(ok)
}

// 4102241995412645845
func TestRegIDCard(t *testing.T) {
	ok := regmatch.RegIDCard("xxxxxxxxxxxxxxxxxxxxxxxxxx")
	fmt.Println(ok)
}

// 1.160.10.240
func TestRegIPV4(t *testing.T) {
	ok := regmatch.RegIPV4("1.160.10.240")
	fmt.Println(ok)
}

func TestRegIPV6(t *testing.T) {
	ok := regmatch.RegIPV6("1030::C9B4:FF12:48AA:1A2B")
	fmt.Println(ok)
}

func TestPassword(t *testing.T) {
	ok := regmatch.CheckPasswordLever("Aaba12345678@")
	fmt.Println(ok)
}

func TestRegUrl(t *testing.T) {
	ok := regmatch.RegUrl("https://www.baidu.com/")
	fmt.Println(ok)
}

func TestRegUrlPort(t *testing.T) {
	ok := regmatch.RegUrlPort("https://www.baidu.com:8080/")
	fmt.Println(ok)
}

func TestRegAmount(t *testing.T) {
	ok := regmatch.RegAmount("0.11")
	fmt.Println(ok)
}
