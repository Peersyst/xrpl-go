package transaction

import (
	"testing"

	"github.com/Peersyst/xrpl-go/v1/xrpl/testutil"
	"github.com/Peersyst/xrpl-go/v1/xrpl/transaction/types"
	"github.com/stretchr/testify/assert"
)

func TestNFTokenMint_TxType(t *testing.T) {
	tx := &NFTokenMint{}
	assert.Equal(t, NFTokenMintTx, tx.TxType())
}

func TestNFTokenMint_Flags(t *testing.T) {
	tests := []struct {
		name     string
		setter   func(*NFTokenMint)
		expected uint32
	}{
		{
			name: "pass - SetBurnableFlag",
			setter: func(n *NFTokenMint) {
				n.SetBurnableFlag()
			},
			expected: tfBurnable,
		},
		{
			name: "pass - SetOnlyXRPFlag",
			setter: func(n *NFTokenMint) {
				n.SetOnlyXRPFlag()
			},
			expected: tfOnlyXRP,
		},
		{
			name: "pass - SetTrustlineFlag",
			setter: func(n *NFTokenMint) {
				n.SetTrustlineFlag()
			},
			expected: tfTrustLine,
		},
		{
			name: "pass - SetTransferableFlag",
			setter: func(n *NFTokenMint) {
				n.SetTransferableFlag()
			},
			expected: tfTransferable,
		},
		{
			name: "pass - SetBurnableFlag and SetTransferableFlag",
			setter: func(n *NFTokenMint) {
				n.SetBurnableFlag()
				n.SetTransferableFlag()
			},
			expected: tfBurnable | tfTransferable,
		},
		{
			name: "pass - SetBurnableFlag and SetTransferableFlag and SetOnlyXRPFlag",
			setter: func(n *NFTokenMint) {
				n.SetBurnableFlag()
				n.SetTransferableFlag()
				n.SetOnlyXRPFlag()
			},
			expected: tfBurnable | tfTransferable | tfOnlyXRP,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &NFTokenMint{}
			tt.setter(p)
			if p.Flags != tt.expected {
				t.Errorf("Expected Flags to be %d, got %d", tt.expected, p.Flags)
			}
		})
	}
}

func TestNFTokenMint_Flatten(t *testing.T) {
	tests := []struct {
		name     string
		nft      *NFTokenMint
		expected string
	}{
		{
			name: "pass - Flatten with all fields",
			nft: &NFTokenMint{
				BaseTx: BaseTx{
					Account: "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					Fee:     types.XRPCurrencyAmount(10),
				},
				NFTokenTaxon: 12345,
				Issuer:       "rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW",
				TransferFee:  314,
				URI:          "697066733A2F2F62616679626569676479727A74357366703775646D37687537367568377932366E6634646675796C71616266336F636C67747179353566627A6469",
			},
			expected: `{
				"Account":         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				"TransactionType": "NFTokenMint",
				"Fee":             "10",
				"NFTokenTaxon":    12345,
				"Issuer":          "rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW",
				"TransferFee":     314,
				"URI":             "697066733A2F2F62616679626569676479727A74357366703775646D37687537367568377932366E6634646675796C71616266336F636C67747179353566627A6469"
			}`,
		},
		{
			name: "pass - Flatten with minimal fields",
			nft: &NFTokenMint{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: NFTokenMintTx,
					Fee:             types.XRPCurrencyAmount(10),
				},
				NFTokenTaxon: 12345,
			},
			expected: `{
				"Account":         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				"TransactionType": "NFTokenMint",
				"Fee":             "10",
				"NFTokenTaxon":    12345
			}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, tt := range tests {
				t.Run(tt.name, func(t *testing.T) {
					err := testutil.CompareFlattenAndExpected(tt.nft.Flatten(), []byte(tt.expected))
					if err != nil {
						t.Error(err)
					}
				})
			}
		})
	}
}

func TestNFTokenMint_Validate(t *testing.T) {
	tests := []struct {
		name       string
		nft        *NFTokenMint
		setter     func(*NFTokenMint)
		wantValid  bool
		wantErr    bool
		errMessage error
	}{
		{
			name: "pass - minimal fields",
			nft: &NFTokenMint{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: NFTokenMintTx,
					Fee:             types.XRPCurrencyAmount(10),
				},
				NFTokenTaxon: 12345,
			},
			wantValid: true,
			wantErr:   false,
		},
		{
			name: "fail - invalid BaseTx fields, missing account",
			nft: &NFTokenMint{
				BaseTx: BaseTx{
					TransactionType: NFTokenMintTx,
					Fee:             types.XRPCurrencyAmount(10),
				},
				NFTokenTaxon: 12345,
			},
			wantValid:  false,
			wantErr:    true,
			errMessage: ErrInvalidAccount,
		},
		{
			name: "fail - transfer fee exceeds max",
			nft: &NFTokenMint{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: NFTokenMintTx,
					Fee:             types.XRPCurrencyAmount(10),
				},
				NFTokenTaxon: 12345,
				TransferFee:  60000,
			},
			wantValid:  false,
			wantErr:    true,
			errMessage: ErrInvalidTransferFee,
		},
		{
			name: "fail - issuer same as account",
			nft: &NFTokenMint{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: NFTokenMintTx,
					Fee:             types.XRPCurrencyAmount(10),
				},
				NFTokenTaxon: 12345,
				Issuer:       "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
			},
			wantValid:  false,
			wantErr:    true,
			errMessage: ErrIssuerAccountConflict,
		},
		{
			name: "fail - issuer invalid address",
			nft: &NFTokenMint{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: NFTokenMintTx,
					Fee:             types.XRPCurrencyAmount(10),
				},
				NFTokenTaxon: 12345,
				Issuer:       "invalidAddress",
			},
			wantValid:  false,
			wantErr:    true,
			errMessage: ErrInvalidIssuer,
		},
		{
			name: "fail - URI not hexadecimal",
			nft: &NFTokenMint{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: NFTokenMintTx,
					Fee:             types.XRPCurrencyAmount(10),
				},
				NFTokenTaxon: 12345,
				URI:          "invalidURI",
			},
			wantValid:  false,
			wantErr:    true,
			errMessage: ErrInvalidURI,
		},
		{
			name: "fail - transfer fee set without transferable flag",
			nft: &NFTokenMint{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: NFTokenMintTx,
					Fee:             types.XRPCurrencyAmount(10),
				},
				NFTokenTaxon: 12345,
				TransferFee:  314,
			},
			wantValid:  false,
			wantErr:    true,
			errMessage: ErrTransferFeeRequiresTransferableFlag,
		},
		{
			name: "pass - transfer fee set with transferable flag",
			nft: &NFTokenMint{
				BaseTx: BaseTx{
					Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					TransactionType: NFTokenMintTx,
					Fee:             types.XRPCurrencyAmount(10),
				},
				NFTokenTaxon: 12345,
				TransferFee:  314,
			},
			setter: func(n *NFTokenMint) {
				n.SetTransferableFlag()
			},
			wantValid: true,
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setter != nil {
				tt.setter(tt.nft)
			}
			valid, err := tt.nft.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (err != nil) && err != tt.errMessage {
				t.Errorf("Validate() got error message = %v, want error message %v", err, tt.errMessage)
				return
			}
			if valid != tt.wantValid {
				t.Errorf("Validate() valid = %v, wantValid %v", valid, tt.wantValid)
			}
		})
	}
}
