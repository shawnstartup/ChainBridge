package vault

import (
	"reflect"
	"testing"
)

func TestCustomerRefId(t *testing.T) {
	var vault = Vault{
		CustomerRefIdPrefix: "bridge_v100",
	}
	customerRefId := vault.MakeCustomerRefId(1, 3, 133)

	txCustomerRefId, err := ParseCustomerRefId(customerRefId)
	if err != nil {
		t.Fatal(err)
	}
	expected := TxCustomerRefId{
		Raw:                 "bridge_v100-src-001-dst-003-nonce-000000133",
		CustomerRefIdPrefix: "bridge_v100",
		SrcChainId:          1,
		DstChainId:          3,
		DepositNonce:        133,
	}
	if !reflect.DeepEqual(&expected, txCustomerRefId) {
		t.Fatalf("Output not expected.\n\tExpected: %#v\n\tGot: %#v\n", &expected, txCustomerRefId)
	}
}
