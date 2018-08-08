package jsvm

import (
	"testing"

	"boscoin.io/sebak/lib/contract/api"
	"boscoin.io/sebak/lib/contract/context"
	"boscoin.io/sebak/lib/contract/payload"
	"boscoin.io/sebak/lib/contract/value"
	"github.com/robertkrimen/otto"
)

func Test_CallContract(t *testing.T) {
	ctx := &context.Context{
		StateStore: testStateStore,
		StateClone: testStateClone,
	}
	callc := func(ctx *context.Context, execCode *payload.ExecCode) (*value.Value, error) {
		if execCode.ContractAddress != "helloworld" {
			t.Fatalf("execCode.ContractAddress  have: %v want: %v", execCode.ContractAddress, "helloworld")
		}

		ret := &value.Value{
			Type:     value.String,
			Contents: []byte("world!!"),
		}
		return ret, nil
	}
	api := api.NewAPI(ctx, "caller", callc)

	vm := otto.New()
	vm.Set("CallContract", CallContractFunc(api))

	_, err := vm.Run(`
		ret = CallContract("helloworld","hello","world")
	`)
	if err != nil {
		t.Fatal(err)
	}

	ret, err := vm.Get("ret")
	if err != nil {
		t.Fatal(err)
	}

	if ret.String() != "world!!" {
		t.Fatal(ret)
	}

}
