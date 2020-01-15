package numbers

import (
	"math"
	"strconv"
)

func GetAmountValue(amount string) string {
	value := ParseAmount(amount)
	return strconv.FormatInt(value, 10)
}

func ParseAmount(amount string) int64 {
	value, err := strconv.ParseInt(amount, 10, 64)
	if err == nil {
		return value
	}
	return ToSatoshi(amount)
}

func ToSatoshi(amount string) int64 {
	value, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return 0
	}
	total := value * math.Pow10(8)
	return int64(total)
}

func AddAmount(left string, right string) (sum string) {
	amount1 := ParseAmount(left)
	amount2 := ParseAmount(right)
	return strconv.FormatInt(amount1+amount2, 10)
}
