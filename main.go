package main

import "C"
import (
	"fmt"
	"sign-sdk/account"
)

//export Signtx
func Signtx(sk string, targetAddr string, value int, gasLimit int, gasPrice int, _type int, nonce int, dataHex string) *C.char {
	return C.CString(account.SignTx(sk, targetAddr, int64(value), int64(gasLimit), int64(gasPrice), _type, int64(nonce), dataHex))
}

//export Sk2Pk
func Sk2Pk(sk string) *C.char {
	return C.CString(account.Sk2Pk(sk))
}

//export Hello
func Hello(msg string) {
	fmt.Print("hello: " + msg)
}

//export DemoSk2Pk
func DemoSk2Pk(sk string) *C.char {
	fmt.Printf("input  sk:      %v\n", sk)

	pk := account.Sk2Pk(sk)

	fmt.Printf("output pk:      %v\n", pk)
	fmt.Printf("       address: %v\n", account.Sk2Addr(sk))
	fmt.Printf("------------------------\n\n")

	return C.CString(pk)
}

//export DemoSigntx
func DemoSigntx(sk string, targetAddr string, value int, gasLimit int, gasPrice int, _type int, nonce int, dataHex string) *C.char {
	fmt.Printf("input sk:         %v\n", sk)
	fmt.Printf("      targetAddr: %v\n", targetAddr)
	fmt.Printf("      value:      %v\n", value)
	fmt.Printf("      gasLimit:   %v\n", gasLimit)
	fmt.Printf("      gasPrice:   %v\n", gasPrice)
	fmt.Printf("      _type:      %v\n", _type)
	fmt.Printf("      nonce:      %v\n", nonce)
	fmt.Printf("      dataHex:    %v\n", dataHex)

	sign := account.SignTx(sk, targetAddr, int64(value), int64(gasLimit), int64(gasPrice), _type, int64(nonce), dataHex)

	fmt.Printf("output sign: %v\n", sign)
	fmt.Printf("------------------------\n\n")

	return C.CString(sign)
}

func char2str(c *C.char) string {
	return C.GoString(c)
}

func main() {

}
