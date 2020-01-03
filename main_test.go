package main

import (
	"fmt"
	"sign-sdk/account"
	"sign-sdk/common"
	"testing"
)

func TestSk2Pk(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"0x95139ce96a78591bdb0e33d39aed0a29988eb7ddba5495c7fdee350b177ae76d", "0x0449a7c86895962da27a49bd34696e32d9c539cc9586c765b862873bf3c6d8e8c063956734bd74c21876dedd5386c36e0cd0f241a59ea50508394303c660520a59"},
		{"0x6c16b3d2eff3932cb9ffdb4a750ff040afa5a2a031d61538bb503e9f2c3dc82d", "0x04903a83e91e258373304818d503813b0cbd0dfaff5bcabbdd9cdd4096aaa797c289314e868bf148f7b1a0d3cc6842ad46a0beab245dbb30f01d3fba01ee7d3ab9"},
		{"0x97ddae0f3a25b92268175400149d65d6887b9cefaf28ea2c078e05cdc15a3c0a", "0x047b83ad6afb1209f3c82ebeb08c0c5fa9bf6724548506f2fb4f991e2287a77090177316ca82b0bdf70cd9dee145c3002c0da1d92626449875972a27807b73b42e"},
	}

	for _, c := range cases {
		got := char2str(Sk2Pk(c.in))

		if got != c.want {
			t.Errorf("Sk2Pk(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestSk2Address(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"0x01", "DD684ee3c4c1c613afc7a19c630502987e630ea7ebb2bf1d84a65a727109385bcf"},
		{"0x02", "DDc95eeb720c691843cd58c5a65b97bc7e0a18d4a7a825e150ae59ef4e5c9fba2e"},
		{"0x03", "DDc0d456570efcdd896efe1f35bc920e7a7a4dc01e204cd0e75c379429ae58b9e8"},
		{"0x7a", "DDd2f22acd76a1d9c03fc2c392e71da8898f631869e1cea6943f556d0d520ab5a8"},
		{"0x82", "DD71088a7e53a7655f8a592f69af6575d475661943d36991d1f0470691242f35e4"},
		{"0x99", "DD6fffbb6bf77fa29ed9a61db78e7c7a1dbf28a52764e08690e06c4a9fbd2084d7"},
	}

	for _, c := range cases {
		got := account.Sk2Addr(c.in)

		if got != c.want {
			t.Errorf("Sk2Addr(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

type Tx struct {
	sk       string
	target   string
	value    int
	gas      int
	gasprice int
	nonce    int
	_type    int
	data     string
}

func TestSignTx(t *testing.T) {
	cases := []struct {
		in   *Tx
		want string
	}{
		{&Tx{
			sk:       "0x95139ce96a78591bdb0e33d39aed0a29988eb7ddba5495c7fdee350b177ae76d",
			target:   "DDe45b7d9bd4bcacb09664833e504f7b87bf19cf4e104eeb8e61bd984610d2ba75",
			value:    1000000000,
			gas:      3000,
			gasprice: 500,
			nonce:    1,
			_type:    0,
			data:     "",
		}, "0x858af204097ba92239a4c0569da39ef546f2d0541b1d955955a0ed6530e863e04f492c20e1fcb645ad96003379a0dd4a69dd126fa923c3a1aa4d70ef9d336fec00"},
		{&Tx{
			sk:       "0x6c16b3d2eff3932cb9ffdb4a750ff040afa5a2a031d61538bb503e9f2c3dc82d",
			target:   "DD520fe0e7864b81a6c85f5a8fcdac338c0f220b944cc9f5beb1b643bc57db4776",
			value:    1000000000,
			gas:      3000,
			gasprice: 500,
			nonce:    1,
			_type:    0,
			data:     "",
		}, "0x2f7a25fefa77cbfb4e95eb0c63bd867bc24fae74466b74f5f3118e6103b7387a26e4e12e85b90ec74170efce203430a9e113c1eb2b7d68cfeea511e78a1f0d0c01"},
		{&Tx{
			sk:       "0xf241759f4e85b2188314c4636e88b8b8076606b53ebe9817bbe1b2379b930ca6",
			target:   "DDee8df38919ee6e7fbeb033aa797638344490f91c417d04e91750c4a2d4111918",
			value:    100000000000,
			gas:      3000,
			gasprice: 500,
			nonce:    3,
			_type:    0,
			data:     "",
		}, "0xde4f80c9aa4c3e4eb38d26cce2d4fb6b7f951a896240e448fa638bdde43de01064e476f09a1f771f2ba0dbd78b8997d448175873fd06c5b02d65484c7a58be5601"},
	}

	for _, c := range cases {
		got := char2str(Signtx(c.in.sk, c.in.target, c.in.value, c.in.gas, c.in.gasprice, c.in._type, c.in.nonce, c.in.data))

		if got != c.want {
			t.Errorf("Sk2Pk(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestSDK(t *testing.T) {
	sk := account.GetChild("chicken october obey magnet climb secret glow olive defense then excuse size", 10)
	addr := account.Sk2Addr(sk)

	tx := &Tx{
		sk:       sk,
		target:   "DD966daf04fb650fd1d844c597562226cc8bcdf1c20e2d62b4d1c760163ffaee32",
		value:    24460000000,
		gas:      3000,
		gasprice: 13333,
		nonce:    1,
		_type:    0,
		data:     "",
	}

	fmt.Printf("  sk: %v\naddr: %v\n", sk, addr)
	fmt.Printf("sign: %v\n", account.SignTx(tx.sk, tx.target, int64(tx.value), int64(tx.gas), int64(tx.gasprice), tx._type, int64(tx.nonce), tx.data))
}

func generateKey() (string, string, error) {
	pk, err := common.GenerateKey("")
	if err != nil {
		return "", "", err
	}

	sk := common.ToHex(pk.ExportKey())
	address := pk.GetPubKey().GetAddress().AddrPrefixString()

	return sk, address, nil
}

type SkAndAddress struct {
	sk      string
	address string
	want    string
}

type TxSignCase struct {
	in   *Tx
	want string
}

func TestSigner(t *testing.T) {

	cases := []SkAndAddress{
		{"0x01", "DD684ee3c4c1c613afc7a19c630502987e630ea7ebb2bf1d84a65a727109385bcf", "0xa93dd244dc30b3fb991e014e3c70046a0f55e8c12a7c767a8368ca895b357f372fb19343d1c1d5fc5382b191838612c81d7288e80a0ac00760f8cfeff3d91ca900"},
		{"0x02", "DDc95eeb720c691843cd58c5a65b97bc7e0a18d4a7a825e150ae59ef4e5c9fba2e", "0xb8c6105048aa3e74dd65b31fa7d99961d932a3ea1ec38a5b5b52a746902bfecc51a654a917eebae502f18cb30ef08b75d8fd94631efc160d434e79bfd449690500"},
		{"0x03", "DDc0d456570efcdd896efe1f35bc920e7a7a4dc01e204cd0e75c379429ae58b9e8", "0x9f2760c3c06c92fde2786d2f0c0b2d9bd506f41ec5e492bb4260f330eaaeaa187c75205fadcd8f28b8574a21756ae6e739054287e81d4dc3b51d33aa261d9d3100"},
		{"0x04", "DDdcb6b4f534a9863c93a4a81267bbf13c7a16ecd7d70e41fe7ec7a15a015268d5", "0x057a2cc9bb5bc09e680fe7ce04cf2f15fcd7240756a0abcb537ddebe0092a8f376746af061c3eb55632ffe5190e5a856fe56536e3d5adac5c7f6669753ae95d201"},
		{"0x05", "DD12aec0b43fe998949d51ce543015f805982bbdb16c6be11b688708cac2e51b06", "0x81f810bd7eefe92ea3d354b642e5d6cd7cca262b15ded78da3ef54e9259136cd10c3ae75fe938b0dc69897c17094849cb5441c2c3cde76fb5a0dfa18fa76fd6a01"},
		{"0x06", "DD9b7ccd823242c7f38a11c6322060a95b340d35f7fcffb7522d35c8dce6fe9ff8", "0x6639c21a14acab4d43e2a165b139bebc648f24dea85a621a7ab547d398f40d70650733aa51bf331d69c6333d8aafbb15b88ba7ab241ae118905446d2b38cd43301"},
		{"0x07", "DDc851e203fe77665fca130b9b02deabd377e8bf211505f175b3dd2dd34937ea05", "0x4e03274b90400b30eb17125281be176c8d01c2b1d41856792bca78511781ac065a20fc2e52f92b10632e225bb641eae84771499cd474c22dd3ccd4cfd7f44eaf00"},
		{"0x08", "DDa7226088dea5481c78135a035e86a28cf1fe88ad9af8118a23c0ce22d5ae7b23", "0x9345782f7d9a829bfb3c43c5c8e076ce970ee26a9c51a16dbc480efbab4cc67616bd03a27fc9bf88ad4c7155b628b6f0c857793209951a1643a04c21367132c601"},
		{"0x09", "DDf43bc69b3f241bb6cd6a886208c73019696967ffcca49cea93c38d026b1fa60f", "0xc8e69cc62d1576983fd3cc21ccce3002ff1c0ef0884501cc145c95783b0199f725164da1caf324a28efd3b709d4a41256ac8d799b6f3d25cb57b883ce787652a00"},
		{"0x0a", "DDfee717fada51b635c39ae05835266859b2d9f994683a2955a8efc96fc0c0994d", "0x743c178dbaeec96b661068a488970ef1f46b555321cd2302df0dbf2916be879f74f9d4ee6033ee2ab140494e13a59eff011eea0f7b980c1d6f6998db9f42657201"},
		{"0x0b", "DDf2787d4f4fc4150f5c833ceccf7cbca0fdcbad85b03286738f13f35744d3d079", "0x8b7d4b7aa4e51f33bf11ed9e819aea959dfdbff4fffec768d75e237a29e860af2e19c4f88fbf42a872b0c34917344455c8a8f1498c660a70e05cc5400096619100"},
		{"0x0c", "DD979ec2a84ed3c91fdb39e57d27befec21a14065dde4573cd760c458996c9bbdd", "0x62da89c07d7dda1c414a7f04b698d45117bdbf46bbe2b41371f0bfcd31d582786fb2afd4de08c7606feb00d394875982bb7f09e1ebadf302650d2258a3dc515401"},
		{"0x0d", "DD449c180643cd50ae8c209f39671f6ab88e9adcca33c3239ba495582b71117b4b", "0xdec1855df507537dd82c9f533b605e82afad79d302fde479836c7c5a4a1717dd38d0c77fb1af1e45cbb26c2a98491cc93239174ec70909faf3448f1eee3db62501"},
		{"0x0e", "DDdadb20ff0efa3a321372ed2cb98103c0e334fa2fdc9973cb7f15fdd0e6783e53", "0x921175a8f9f30d5e8ef53d3c5485ad6852fa505a7e79c8b1fe236e0905659a5f568595483cf1101fef886106763e5874ea47909be42138f72292693d0b78c58e01"},
		{"0x0f", "DDefbf41b36facfd1304625001981e08d888e855e01901de5c9c8c3d87a1eba4e3", "0xd69dd7526a77b02b957d27beacf871ace826f665edb5162a8248cbe090ba357b7d2c4366f16503a9ed4728ced67d9e23da47b9da64d52627901543ae325ceafc01"},
		{"0x10", "DDb9b86d99dbc7c1875dda7b71fc91adf4afb7566a7fd814e9e1380bf590c36166", "0x61d89cfebd0a730c0d8bd1053c96a0eb810361353d1008348a11418b0fb08e29580697dcf8e095ecc231c62a9684450820a095036c2917dc844b0e3a8e200fe300"},
		{"0x11", "DDe0e7e1796f24bacd0855419750e2f76384de13cf298650769fbbad3ad6657c7e", "0x44af0a0267152358d831b5f471f64057e1056b4e0713c0f01a93fed920d083f1535b01adbc29043ae0c7fa588781944b147e9a1364ace4828a5f7460df608ab000"},
		{"0x12", "DDe4663b7ce4f6685894521af2ebdd868964527059b6c5c6275af02b3e4a93d5b4", "0x9d93ef5d8cc365487809cbd7aefd0512265044700694afd36e680dbc868b469a1e5dcb9ac3536dfd2cf1a5e28cdd27931c08577540f96f3c53bca9f297df6bee00"},
	}

	txs := make([]*TxSignCase, 0)

	for _, c := range cases {
		tx := &Tx{
			sk:       c.sk,
			target:   c.address,
			value:    24460000000,
			gas:      3000,
			gasprice: 13333,
			nonce:    1,
			_type:    0,
			data:     "",
		}

		txs = append(txs, &TxSignCase{
			in:   tx,
			want: c.want,
		})
	}

	for _, c := range txs {
		tx := c.in
		got := account.SignTx(tx.sk, tx.target, int64(tx.value), int64(tx.gas), int64(tx.gasprice), tx._type, int64(tx.nonce), tx.data)

		if got != c.want {
			t.Errorf("SignTx(%q) == %q, want %q", tx, got, c.want)
		}

	}
}
