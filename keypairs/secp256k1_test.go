package keypairs

import (
	"testing"

	addresscodec "github.com/Peersyst/xrpl-go/address-codec"
	"github.com/stretchr/testify/require"
)

func TestSecp256k1DeriveKeypair(t *testing.T) {
	e := &secp256k1Alg{}
	tt := []struct {
		description string
		seed        string
		validator   bool
		expPubKey   string
		expPrivKey  string
		expErr      error
	}{
		{
			description: "Successfully derive keypair",
			seed:        "shQr3dfBJpA3wucZT9nApRxQZMVWD",
			validator:   false,
			expPubKey:   "029762F6C921D0CC7270F697E6E3F0EDFF53A7E2E769036D6BA622D4E531D3A885",
			expPrivKey:  "001CA14BA68A1C1B920B297EC84E61205E5774BBF10AE15228D6FC7161688557D8",
			expErr:      nil,
		},
		{
			description: "Error if validator is set to true",
			seed:        "shQr3dfBJpA3wucZT9nApRxQZMVWD",
			validator:   true,
			expPubKey:   "",
			expPrivKey:  "",
			expErr:      &secp256k1ValidatorError{},
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			ds, _, _ := addresscodec.DecodeSeed(tc.seed)
			priv, pub, err := e.deriveKeypair(ds, tc.validator)
			if tc.expErr != nil {
				require.Zero(t, pub)
				require.Zero(t, priv)
				require.Error(t, err, tc.expErr.Error())
			} else {
				require.Equal(t, tc.expPrivKey, priv)
				require.Equal(t, tc.expPubKey, pub)
			}
		})
	}
}

// func TestED25519Sign(t *testing.T) {
// 	e := &ed25519Alg{}
// 	tt := []struct {
// 		description  string
// 		inputMsg     string
// 		inputPrivKey string
// 		expected     string
// 		expectedErr  error
// 	}{
// 		{
// 			description:  "Sign a valid message",
// 			inputMsg:     "hello world",
// 			inputPrivKey: "EDBB3ECA8985E1484FA6A28C4B30FB0042A2CC5DF3EC8DC37B5F3D126DDFD3CA14",
// 			expected:     "E83CAFEAF100793F0C6570D60C7447FF3A87E0DC0CAE9AD90EF0102860EC3BD1D20F432494021F3E19DAFF257A420CA64A49C283AB5AD00B6B0CEA1756151C01",
// 			expectedErr:  nil,
// 		},
// 		{
// 			description:  "Sign a message with a different private key",
// 			inputMsg:     "hello world",
// 			inputPrivKey: "ED6BF4E585BA0C4055F6E63D0D6D06E7D8B9F00AA02337BCF864385275892A1EB5",
// 			expected:     "84F05438BFFC29F49E8DC8865251DA1CEF9A5A9CAA7DC2629985986C35271CC1AC389846F955C548A322F433F387CE928329F091E8FA7E2A8E7DFDAB8E88310B",
// 			expectedErr:  nil,
// 		},
// 		{
// 			description:  "Sign a longer message with a different private key",
// 			inputMsg:     "hello madsjdadjdas,adajofahffa !$@~",
// 			inputPrivKey: "ED28E9ADC9383EB494476DCF7D95DD4B16F6A2C325365F9E17007294F4AE487CE0",
// 			expected:     "77F07A34D408DD8C3C6BCED0E31C2909D8E13ECB15AF15345CA1ECE53118519754971BE1DD7A0A52E5D737D4DBFAD01018727EF1F0BAD06B31CD8D6F3D9E7E05",
// 			expectedErr:  nil,
// 		},
// 	}

// 	for _, tc := range tt {
// 		t.Run(tc.description, func(t *testing.T) {
// 			actual, err := e.sign(tc.inputMsg, tc.inputPrivKey)

// 			if tc.expectedErr != nil {
// 				require.Zero(t, actual)
// 				require.Error(t, err, tc.expectedErr.Error())
// 			} else {
// 				require.NoError(t, err)
// 				require.Equal(t, tc.expected, actual)
// 			}
// 		})
// 	}
// }

// func TestED25519Validate(t *testing.T) {
// 	e := &ed25519Alg{}
// 	tt := []struct {
// 		description string
// 		inputMsg    string
// 		inputPubKey string
// 		inputSig    string
// 		expected    bool
// 	}{
// 		{
// 			description: "A valid signature for message",
// 			inputMsg:    "test message",
// 			inputPubKey: "ED4924A9045FE5ED8B22BAA7B6229A72A287CCF3EA287AADD3A032A24C0F008FA6",
// 			inputSig:    "C001CB8A9883497518917DD16391930F4FEE39CEA76C846CFF4330BA44ED19DC4730056C2C6D7452873DE8120A5023C6807135C6329A89A13BA1D476FE8E7100",
// 			expected:    true,
// 		},
// 		{
// 			description: "An invalid signature for message",
// 			inputMsg:    "test message",
// 			inputPubKey: "ED4924A9045FE5ED8B22BAA7B6229A72A287CCF3EA287AADD3A032A24C0F008FB6",
// 			inputSig:    "C001CB8A9883497518917DD16391930F4FEE39CEA76C846CFF4330BA44ED19DC4730056C2C6D7452873DE8120A5023C6807135C6329A89A13BA1D476FE8E7100",
// 			expected:    false,
// 		},
// 	}

// 	for _, tc := range tt {
// 		t.Run(tc.description, func(t *testing.T) {
// 			actual := e.validate(tc.inputMsg, tc.inputPubKey, tc.inputSig)
// 			require.Equal(t, tc.expected, actual)
// 		})
// 	}
// }
