package main

import (
	"fmt"
	"sign-sdk/account"
	"testing"
)

func TestSk2Pk(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"0x95139ce96a78591bdb0e33d39aed0a29988eb7ddba5495c7fdee350b177ae76d", "0x0449a7c86895962da27a49bd34696e32d9c539cc9586c765b862873bf3c6d8e8c063956734bd74c21876dedd5386c36e0cd0f241a59ea50508394303c660520a59"},
		{"0x6c16b3d2eff3932cb9ffdb4a750ff040afa5a2a031d61538bb503e9f2c3dc82d", "0x04903a83e91e258373304818d503813b0cbd0dfaff5bcabbdd9cdd4096aaa797c289314e868bf148f7b1a0d3cc6842ad46a0beab245dbb30f01d3fba01ee7d3ab9"},
	}

	for _, c := range cases {
		got := char2str(DemoSk2Pk(c.in))

		if got != c.want {
			t.Errorf("DemoSk2Pk(%q) == %q, want %q", c.in, got, c.want)
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
		got := char2str(DemoSigntx(c.in.sk, c.in.target, c.in.value, c.in.gas, c.in.gasprice, c.in._type, c.in.nonce, c.in.data))

		if got != c.want {
			t.Errorf("DemoSk2Pk(%q) == %q, want %q", c.in, got, c.want)
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

	fmt.Printf("sk: %v\n addr: %v\n", sk, addr)
	fmt.Printf("sign: %v\n", account.SignTx(tx.sk, tx.target, int64(tx.value), int64(tx.gas), int64(tx.gasprice), tx._type, int64(tx.nonce), tx.data))
}
