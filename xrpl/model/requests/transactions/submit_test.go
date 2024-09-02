package transactions

import (
	"testing"

	"github.com/Peersyst/xrpl-go/xrpl/test"
)

func TestSubmitRequest(t *testing.T) {
	s := SubmitRequest{
		TxBlob: "1200002280000000240000016861D4838D7EA4C6800000000000000000000000000055534400000000004B4E9C06F24296074F7BC48F92A97916C6DC5EA9684000000000002710732103AB40A0490F9B7ED8DF29D246BF2D6269820A0EE7742ACDD457BEA7C7D0931EDB7446304402200E5C2DD81FDF0BE9AB2A8D797885ED49E804DBF28E806604D878756410CA98B102203349581946B0DDA06B36B35DBC20EDA27552C1F167BCF5C6ECFF49C6A46F858081144B4E9C06F24296074F7BC48F92A97916C6DC5EA983143E9D4A2B8AA0780F682D136F7A56D6724EF53754",
	}

	j := `{
	"tx_blob": "1200002280000000240000016861D4838D7EA4C6800000000000000000000000000055534400000000004B4E9C06F24296074F7BC48F92A97916C6DC5EA9684000000000002710732103AB40A0490F9B7ED8DF29D246BF2D6269820A0EE7742ACDD457BEA7C7D0931EDB7446304402200E5C2DD81FDF0BE9AB2A8D797885ED49E804DBF28E806604D878756410CA98B102203349581946B0DDA06B36B35DBC20EDA27552C1F167BCF5C6ECFF49C6A46F858081144B4E9C06F24296074F7BC48F92A97916C6DC5EA983143E9D4A2B8AA0780F682D136F7A56D6724EF53754"
}`
	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}

}

// TODO: Re do test when responses refactor is done
// func TestSubmitResponse(t *testing.T) {
// 	s := SubmitResponse{
// 		Accepted:                 true,
// 		AccountSequenceAvailable: 362,
// 		AccountSequenceNext:      362,
// 		Applied:                  true,
// 		Broadcast:                true,
// 		EngineResult:             "tesSUCCESS",
// 		EngineResultCode:         0,
// 		EngineResultMessage:      "The transaction was applied. Only final in a validated ledger.",
// 		Kept:                     true,
// 		OpenLedgerCost:           "10",
// 		Queued:                   false,
// 		TxBlob:                   "1200002280000000240000016861D4838D7EA4C6800000000000000000000000000055534400000000004B4E9C06F24296074F7BC48F92A97916C6DC5EA9684000000000002710732103AB40A0490F9B7ED8DF29D246BF2D6269820A0EE7742ACDD457BEA7C7D0931EDB7446304402200E5C2DD81FDF0BE9AB2A8D797885ED49E804DBF28E806604D878756410CA98B102203349581946B0DDA06B36B35DBC20EDA27552C1F167BCF5C6ECFF49C6A46F858081144B4E9C06F24296074F7BC48F92A97916C6DC5EA983143E9D4A2B8AA0780F682D136F7A56D6724EF53754",
// 		Tx: map[string]interface{}{
// 			"Amount": map[string]interface{}{
// 				"Currency": "USD",
// 				"Issuer":   "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
// 				"Value":    "1",
// 			},
// 			"Destination":     "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
// 			"Account":         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
// 			"TransactionType": "Payment",
// 			"Fee":             "10000",
// 			"Sequence":        360,
// 			"Flags":           2147483648,
// 			"SigningPubKey":   "03AB40A0490F9B7ED8DF29D246BF2D6269820A0EE7742ACDD457BEA7C7D0931EDB",
// 			"TxnSignature":    "304402200E5C2DD81FDF0BE9AB2A8D797885ED49E804DBF28E806604D878756410CA98B102203349581946B0DDA06B36B35DBC20EDA27552C1F167BCF5C6ECFF49C6A46F8580",
// 		},
// 		ValidatedLedgerIndex: 21184416,
// 	}
// 	j := `{
// 	"engine_result": "tesSUCCESS",
// 	"engine_result_code": 0,
// 	"engine_result_message": "The transaction was applied. Only final in a validated ledger.",
// 	"tx_blob": "1200002280000000240000016861D4838D7EA4C6800000000000000000000000000055534400000000004B4E9C06F24296074F7BC48F92A97916C6DC5EA9684000000000002710732103AB40A0490F9B7ED8DF29D246BF2D6269820A0EE7742ACDD457BEA7C7D0931EDB7446304402200E5C2DD81FDF0BE9AB2A8D797885ED49E804DBF28E806604D878756410CA98B102203349581946B0DDA06B36B35DBC20EDA27552C1F167BCF5C6ECFF49C6A46F858081144B4E9C06F24296074F7BC48F92A97916C6DC5EA983143E9D4A2B8AA0780F682D136F7A56D6724EF53754",
// 	"tx_json": {
// 		"Account": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
// 		"TransactionType": "Payment",
// 		"Fee": "10000",
// 		"Sequence": 360,
// 		"Flags": 2147483648,
// 		"SigningPubKey": "03AB40A0490F9B7ED8DF29D246BF2D6269820A0EE7742ACDD457BEA7C7D0931EDB",
// 		"TxnSignature": "304402200E5C2DD81FDF0BE9AB2A8D797885ED49E804DBF28E806604D878756410CA98B102203349581946B0DDA06B36B35DBC20EDA27552C1F167BCF5C6ECFF49C6A46F8580",
// 		"Amount": {
// 			"issuer": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
// 			"currency": "USD",
// 			"value": "1"
// 		},
// 		"Destination": "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX"
// 	},
// 	"accepted": true,
// 	"account_sequence_available": 362,
// 	"account_sequence_next": 362,
// 	"applied": true,
// 	"broadcast": true,
// 	"kept": true,
// 	"queued": false,
// 	"open_ledger_cost": "10",
// 	"validated_ledger_index": 21184416
// }`

// 	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
// 		t.Error(err)
// 	}

// }
