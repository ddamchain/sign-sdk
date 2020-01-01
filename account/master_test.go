package account

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"testing"
)

func TestSign(t *testing.T) {
	//fmt.Println(Sign("0xcdb0e796a1b2ae04ab8fe929706fa3569299e9eeac999ed3a7b1b23e1a1257c4", "0x1234567812345678123456781234567812345678123456781234567812345678"))
}

func TestSk2Addr(t *testing.T) {
	fmt.Println(Sk2Addr("0x6bdcb7c948a4ddeb896e5d6b54ebc836c95436b1805aac8fe4daec32670bc6f1"))
}

func TestSk2Pk(t *testing.T) {
	fmt.Println(Sk2Pk("0x6bdcb7c948a4ddeb896e5d6b54ebc836c95436b1805aac8fe4daec32670bc6f1"))
}

func TestGenMasterKey(t *testing.T) {
	fmt.Println(Sk2Addr(GetChild("catch mammal mask option excess pioneer this dilemma window rabbit relief trim dance improve trigger tomorrow beef basket print blast lock oven wood lounge", 0)))
}

//garlic renew lemon achieve become outside fresh label yellow conduct body poverty

func TestGenMasterKey2(t *testing.T) {

}

func TestChild_Sign(t *testing.T) {
	//umid := []byte{234, 46, 131, 211, 67, 49, 160, 225, 0, 108, 142, 196, 150, 33, 224, 163, 193, 127, 249, 182, 204, 185, 139, 195, 102, 151, 51, 191, 91, 185, 154, 135}
	//su, _ := json.Marshal(umid)
	//su:="123"
	//a:=[]byte(su)
	//b, _ := json.Marshal(a)
	//fmt.Println("=======>",b)
	sk := "0xf41b93f25a1b695497d5440122b98e2681465ff12473fdac2aeb911c5d268b2d"
	//sk:="0xf921cccc272a9cc2e78cc52e7a236f5c4bc1105a37d0f7d4128c709cf30bfdf6"
	target := "DDe0fbc42b2ad6a2a499199d0b3d1c6a03e3818f8fd114f80a303b95075529f9ac"
	//source := "DDc401f9fdac56999ed47d26870be158a52127481d838c39907fae1af108e41b25"
	umidhex := "0xc46bc307174668ebfbc2e39f0b592ba38fffaa84af2eccd451578b568252ac08"
	binumid := HexToHash(umidhex)

	//binumid:=""
	jsonResult, _ := json.Marshal(binumid[:])

	signresult := SignTx(sk, target, 0, 10000, 10000, 1, 3, string(jsonResult))
	//Gas: 1000, Gasprice: 10000,Target: "0x35d59ffded5fc9dbf0d422c9b85f592e8dc402f3f4c5f31442e830479b2a2c47", Value: 10, TxType: 0, Nonce: 0,
	fmt.Println("签名结果：：", signresult)
}

func TestNewMaster(t *testing.T) {
	fmt.Println("助记词", GenMnemonic())
}

type Hash [HashLength]byte

func HexToHash(s string) Hash { return BytesToHash(FromHex(s)) }

func Bytes(h Hash) []byte { return h[:] }

func ToHex(b []byte) string {
	hex := Bytes2Hex(b)
	// Prefer output of "0x0" instead of "0x"
	if len(hex) == 0 {
		hex = "0"
	}
	return HexPrefix + hex
}
func Bytes2Hex(d []byte) string {
	return hex.EncodeToString(d)
}

func BytesToHash(b []byte) Hash {
	var h Hash
	h.SetBytes(b)
	return h
}
func (h *Hash) SetBytes(b []byte) {
	if len(b) > len(h) {
		b = b[len(b)-HashLength:]
	}

	copy(h[HashLength-len(b):], b)
}

func FromHex(s string) []byte {
	if len(s) > len(HexPrefix) {
		if HexPrefix == s[0:len(HexPrefix)] {
			s = s[len(HexPrefix):]
		}
		if len(s)%2 == 1 {
			s = "0" + s
		}
		return Hex2Bytes(s)
	}
	return nil
}
func Hex2Bytes(str string) []byte {
	h, _ := hex.DecodeString(str)

	return h
}

//0x90bef4c096cc25c20ae56334e44928ffd6f53a9d47cc52784e78c6c9b79c2c49631e7fb0b95b6e81ee89836ed827bf56b94920f4e04ef2872552c4fcbd622f0d01
