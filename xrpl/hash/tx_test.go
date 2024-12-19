package hash

import "testing"

func TestHashTxBlob(t *testing.T) {
	testCases := []struct {
		name    string
		txBlob  string
		want    string
		wantErr bool
	}{
		{
			name:    "valid tx blob",
			txBlob:  "120000220000000024001B733261400000000000000F68400000000000000C7321ED90ADC33C2BBD9B4A0D94223DBE30D34227B82F587C5909A857B3AB7DE8D6E2EF74402754D4EE7EBDA0A073488904E8A55CECAEDA13EA2829AF5C0EB2CC201C4B4E2AB72D20D308EE12C5D1C112BCFCAFEBDA6C8198D92C0C57F15D8A25B5BFBF200E811474E4DD74B588FA412F0993B8E7E07C2FA92109B48314858233827B488ECB8D0EB940E7AC85CE41E343CF",
			want:    "BE76FC0ABE8BE83F91219D2371FF5199F0271ACF0E12794D2EA5DE77AC49E877",
			wantErr: false,
		},
		{
			name:    "invalid tx blob",
			txBlob:  "120000220000000024001B733261400000000000001268400000000000000C811474E4DD74B588FA412F0993B8E7E07C2FA92109B48314D708DAB02885BA68A48EBCC4EE3551CF1AF7B267",
			want:    "",
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := SignTxBlob(tc.txBlob)
			if (err != nil) != tc.wantErr {
				t.Errorf("HashSignedTx() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if got != tc.want {
				t.Errorf("HashSignedTx() = %v, want %v", got, tc.want)
			}
		})
	}
}
