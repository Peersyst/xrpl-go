package transactions

import (
	"errors"
	"strconv"
)

// Parse the value of an amount, expressed either in XRP or as an Issued Currency, into a float64 or int.
func ParseAmountValue(amount interface{}) (float64, error) {
	if !IsAmount(amount) {
		return 0, errors.New("invalid amount")
	}
	switch v := amount.(type) {
	case string:
		// If amount is a string, parse it to a float or int
		parsedFloat, err := strconv.ParseFloat(v, 64)
		if err == nil {
			return parsedFloat, nil
		}
		parsedInt, err := strconv.ParseInt(v, 10, 64)
		if err == nil {
			return float64(parsedInt), nil
		}
		return 0, errors.New("invalid string amount format")
	case map[string]interface{}:
		// If amount is a map, assume it has a "value" key with a string value
		if value, ok := v["value"].(string); ok {
			parsedFloat, err := strconv.ParseFloat(value, 64)
			if err == nil {
				return parsedFloat, nil
			}
			parsedInt, err := strconv.ParseInt(value, 10, 64)
			if err == nil {
				return float64(parsedInt), nil
			}
			return 0, errors.New("invalid amount format, the 'value' key is not a string and correct float or int")
		}
		return 0, errors.New("invalid amount format, no 'value' key")
	default:
		return 0, errors.New("unsupported amount type")
	}
}
