package volumetrically

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestWeight_Gram(t *testing.T) {
	w := Weight{Decimal: decimal.NewFromFloat(123.456)}
	expected := 123.46
	if gram := w.Gram(2); gram != expected {
		t.Errorf("期望值是 %.2f g，但得到的是 %.2f g", expected, gram)
	}
}

func TestWeight_Kilogram(t *testing.T) {
	w := Weight{Decimal: decimal.NewFromFloat(123456)}
	expected := 123.46
	if kg := w.Kilogram(2); kg != expected {
		t.Errorf("期望值是 %.2f kg，但得到的是 %.2f kg", expected, kg)
	}
}

func TestWeight_Pound(t *testing.T) {
	// 1000克约等于 2.20462 磅
	w := Weight{Decimal: decimal.NewFromInt(1000)}
	// 内部转换是除以 453.59237
	// 1000 / 453.59237 = 2.20462262...
	expected := 2.20
	if lb := w.Pound(2); lb != expected {
		t.Errorf("期望值是 %.2f lbs，但得到的是 %.2f lbs", expected, lb)
	}
}
