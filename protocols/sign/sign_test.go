package sign

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/fxamacker/cbor/v2"
	"github.com/okx/threshold-lib/crypto"
	"github.com/okx/threshold-lib/crypto/commitment"
	"github.com/okx/threshold-lib/crypto/curves"
	"github.com/okx/threshold-lib/crypto/schnorr"
)

func TestFoo(t *testing.T) {
	type T1 struct {
		a int
		b string
	}

	i1 := T1{a: 1, b: "123"}
	i2 := T1{a: 1, b: "123"}

	if i1 != i2 {
		t.Fatal("wrong!")
	}
}

func TestEncodeing(t *testing.T) {
	x := crypto.RandomNum(curve.N)
	xx := &x
	fmt.Println("x: ", x, x.Bytes(), x.Int64())
	fmt.Println("xx: ", (*xx).Bytes(), (*xx).Int64())
	b, err := cbor.Marshal(x)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("Marshal b: ", b)

	bb, err := cbor.Marshal(xx)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("Marshal bb: ", bb)

	var ux **big.Int
	err = cbor.Unmarshal(b, &ux)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("unmarshaled number: ", *ux)
	// if x.Int64() != ux.Int64() {
	// 	t.Fatal("wrong!")
	// }
}

func TestCommitment(t *testing.T) {
	x := crypto.RandomNum(curve.N)
	R1 := curves.ScalarToPoint(curve, x)
	cmt := commitment.NewCommitment(R1.X, R1.Y)
	fmt.Println("cmt.C: ", cmt.C)
	fmt.Println("cmt.Msg: ", cmt.Msg)
}

func TestProve(t *testing.T) {
	x := crypto.RandomNum(curve.N)
	s := crypto.RandomNum(curve.N)
	R := curves.ScalarToPoint(curve, x)
	P := schnorr.Proof{
		R: R,
		S: s,
	}
	fmt.Println("to marshal: ", P.R, P.S)

	b, err := cbor.Marshal(P)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("Marshal b: ", b)

	var ux *schnorr.Proof
	err = cbor.Unmarshal(b, &ux)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("unmarshaled: ", ux.R, ux.S)
}

func TestMessage(t *testing.T) {
	type Message3T struct {
		Proof MessageProof
		CmtD  []*big.Int
	}

	p := MessageProof{
		R: MessagePoint{
			X: crypto.RandomNum(curve.N),
			Y: crypto.RandomNum(curve.N),
		},
	}

	msg := &Message3T{
		Proof: p,
	}

	b, err := cbor.Marshal(msg)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("Marshal b: ", b)
}
