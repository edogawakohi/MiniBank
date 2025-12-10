package utils

const (
	USD = "USD"
	EUR = "EUR"
	VND = "VND"
)

// IsSupportedCurrency returns true if the currency is supported
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, VND:
		return true
	}
	return false
}
