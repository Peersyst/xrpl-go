package types

import "testing"

func TestIssuedCurrencyAmount_IsZero(t *testing.T) {
	tests := []struct {
		name string
		ica  IssuedCurrencyAmount
		want bool
	}{
		{
			name: "Zero value",
			ica:  IssuedCurrencyAmount{},
			want: true,
		},
		{
			name: "Non-zero value",
			ica: IssuedCurrencyAmount{
				Issuer:   "rEXAMPLE",
				Currency: "USD",
				Value:    "100",
			},
			want: false,
		},
		{
			name: "Non-zero value, invalid only with issuer",
			ica: IssuedCurrencyAmount{
				Issuer: "rEXAMPLE",
			},
			want: false,
		},
		{
			name: "Non-zero value, invalid only with value",
			ica: IssuedCurrencyAmount{
				Value: "100",
			},
			want: false,
		},
		{
			name: "Non-zero value, invalid only with currency",
			ica: IssuedCurrencyAmount{
				Currency: "USD",
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ica.IsZero(); got != tt.want {
				t.Errorf("IssuedCurrencyAmount.IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}
