package typecheck

import "testing"

func TestIsString(t *testing.T) {
	tests := []struct {
		name string
		str  interface{}
		want bool
	}{
		{
			name: "Valid string",
			str:  "Hello, World!",
			want: true,
		},
		{
			name: "Invalid string",
			str:  42,
			want: false,
		},
		{
			name: "Empty string",
			str:  "",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsString(tt.str); got != tt.want {
				t.Errorf("IsString(%v) = %v, want %v", tt.str, got, tt.want)
			}
		})
	}
}
func TestIsUint32(t *testing.T) {
	tests := []struct {
		name string
		num  interface{}
		want bool
	}{
		{
			name: "Valid uint32",
			num:  uint32(42),
			want: true,
		},
		{
			name: "Invalid uint32",
			num:  42,
			want: false,
		},
		{
			name: "Valid uint64",
			num:  uint64(42),
			want: false,
		},
		{
			name: "Valid int",
			num:  int(42),
			want: false,
		},
		{
			name: "Valid uint",
			num:  uint(42),
			want: false,
		},
		{
			name: "Valid bool",
			num:  true,
			want: false,
		},
		{
			name: "Valid map",
			num:  map[string]interface{}{},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsUint32(tt.num); got != tt.want {
				t.Errorf("IsUint32(%v) = %v, want %v", tt.num, got, tt.want)
			}
		})
	}
}
func TestIsUint64(t *testing.T) {
	tests := []struct {
		name string
		num  interface{}
		want bool
	}{
		{
			name: "Valid uint64",
			num:  uint64(42),
			want: true,
		},
		{
			name: "Invalid uint64",
			num:  42,
			want: false,
		},
		{
			name: "Valid uint32",
			num:  uint32(42),
			want: false,
		},
		{
			name: "Valid int",
			num:  int(42),
			want: false,
		},
		{
			name: "Valid uint",
			num:  uint(42),
			want: false,
		},
		{
			name: "Valid bool",
			num:  true,
			want: false,
		},
		{
			name: "Valid map",
			num:  map[string]interface{}{},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsUint64(tt.num); got != tt.want {
				t.Errorf("IsUint64(%v) = %v, want %v", tt.num, got, tt.want)
			}
		})
	}
}
func TestIsUint(t *testing.T) {
	tests := []struct {
		name string
		num  interface{}
		want bool
	}{
		{
			name: "Valid uint",
			num:  uint(42),
			want: true,
		},
		{
			name: "Invalid uint",
			num:  42,
			want: false,
		},
		{
			name: "Valid uint32",
			num:  uint32(42),
			want: false,
		},
		{
			name: "Valid uint64",
			num:  uint64(42),
			want: false,
		},
		{
			name: "Valid int",
			num:  int(42),
			want: false,
		},
		{
			name: "Valid bool",
			num:  true,
			want: false,
		},
		{
			name: "Valid map",
			num:  map[string]interface{}{},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsUint(tt.num); got != tt.want {
				t.Errorf("IsUint(%v) = %v, want %v", tt.num, got, tt.want)
			}
		})
	}
}
func TestIsBool(t *testing.T) {
	tests := []struct {
		name string
		b    interface{}
		want bool
	}{
		{
			name: "Valid bool",
			b:    true,
			want: true,
		},
		{
			name: "Invalid bool",
			b:    42,
			want: false,
		},
		{
			name: "Invalid bool",
			b:    "true",
			want: false,
		},
		{
			name: "Invalid bool",
			b:    nil,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsBool(tt.b); got != tt.want {
				t.Errorf("IsBool(%v) = %v, want %v", tt.b, got, tt.want)
			}
		})
	}
}

func TestIsHex(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "Valid hexadecimal string",
			s:    "0123456789abcdefABCDEF",
			want: true,
		},
		{
			name: "Invalid hexadecimal string with non-hex characters",
			s:    "0123456789abcdefABCDEFG",
			want: false,
		},
		{
			name: "Invalid hexadecimal string with spaces",
			s:    "0123456789 abcdefABCDEF",
			want: false,
		},
		{
			name: "Invalid hexadecimal string with special characters",
			s:    "0123456789!abcdefABCDEF",
			want: false,
		},
		{
			name: "Empty string",
			s:    "",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsHex(tt.s); got != tt.want {
				t.Errorf("IsValidHex(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
func TestIsInt(t *testing.T) {
	tests := []struct {
		name string
		num  interface{}
		want bool
	}{
		{
			name: "Valid int",
			num:  42,
			want: true,
		},
		{
			name: "Invalid int",
			num:  3.14,
			want: false,
		},
		{
			name: "Invalid int",
			num:  "42",
			want: false,
		},
		{
			name: "Invalid int",
			num:  nil,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsInt(tt.num); got != tt.want {
				t.Errorf("IsInt(%v) = %v, want %v", tt.num, got, tt.want)
			}
		})
	}
}

func TestIsFloat64(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "Valid float64",
			s:    "3.141592653589793",
			want: true,
		},
		{
			name: "Valid float64 (integer)",
			s:    "42",
			want: true,
		},
		{
			name: "Valid negative float64",
			s:    "-42.0",
			want: true,
		},
		{
			name: "Valid float64 with leading zero",
			s:    "0.123456789",
			want: true,
		},
		{
			name: "Valid negative float64",
			s:    "-3.141592653589793",
			want: true,
		},
		{
			name: "Invalid float64 with multiple decimal points",
			s:    "3.14.15",
			want: false,
		},
		{
			name: "Invalid float64 with non-numeric characters",
			s:    "3.14abc",
			want: false,
		},
		{
			name: "Valid float64 with leading plus sign",
			s:    "+3.141592653589793",
			want: true,
		},
		{
			name: "Invalid float64 with leading minus sign",
			s:    "-",
			want: false,
		},
		{
			name: "Invalid float64 with empty string",
			s:    "",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsFloat64(tt.s); got != tt.want {
				t.Errorf("IsFloat64(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
func TestIsFloat32(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "Valid float32",
			s:    "3.14159",
			want: true,
		},
		{
			name: "Valid float32 (integer)",
			s:    "42",
			want: true,
		},
		{
			name: "Valid negative float32",
			s:    "-42.0",
			want: true,
		},
		{
			name: "Valid float32 with leading zero",
			s:    "0.123456",
			want: true,
		},
		{
			name: "Valid negative float32",
			s:    "-3.14159",
			want: true,
		},
		{
			name: "Invalid float32 with multiple decimal points",
			s:    "3.14.15",
			want: false,
		},
		{
			name: "Invalid float32 with non-numeric characters",
			s:    "3.14abc",
			want: false,
		},
		{
			name: "Valid float32 with leading plus sign",
			s:    "+3.14159",
			want: true,
		},
		{
			name: "Invalid float32 with leading minus sign",
			s:    "-",
			want: false,
		},
		{
			name: "Invalid float32 with empty string",
			s:    "",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsFloat32(tt.s); got != tt.want {
				t.Errorf("IsFloat32(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
func TestIsStringNumericUint(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "Valid uint string",
			s:    "42",
			want: true,
		},
		{
			name: "Valid large uint string",
			s:    "18446744073709551615", // Max uint64 value
			want: true,
		},
		{
			name: "Invalid uint string with negative sign",
			s:    "-42",
			want: false,
		},
		{
			name: "Invalid uint string with decimal point",
			s:    "42.0",
			want: false,
		},
		{
			name: "Invalid uint string with non-numeric characters",
			s:    "42abc",
			want: false,
		},
		{
			name: "Invalid uint string with special characters",
			s:    "42!",
			want: false,
		},
		{
			name: "Empty string",
			s:    "",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsStringNumericUint(tt.s); got != tt.want {
				t.Errorf("IsStringNumericUint(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}
