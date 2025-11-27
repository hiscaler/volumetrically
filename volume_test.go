package volumetrically

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestVolume_CubicCentimeter(t *testing.T) {
	v := Volume{value: decimal.NewFromFloat(123.456)}
	expected := 123.46
	if cc := v.CubicCentimeter(2); !almostEqual(cc, expected) {
		t.Errorf("期望为 %.2f cm^3, 但得到 %.2f cm^3", expected, cc)
	}
}

func TestVolume_CubicCentimeterToMeter(t *testing.T) {
	v := Volume{value: decimal.NewFromFloat(10 * 20 * 30)}
	expected := 0.006
	if cc := v.CubicMeter(3); !almostEqual(cc, expected) {
		t.Errorf("期望为 %.2f cm^3, 但得到 %.2f cm^3", expected, cc)
	}
}

func TestVolume_CubicMeter(t *testing.T) {
	// 1,000,000 cm^3 = 1 m^3
	v := Volume{value: decimal.NewFromFloat(1000000)}
	expected := 1.0
	if m3 := v.CubicMeter(2); !almostEqual(m3, expected) {
		t.Errorf("期望为 %.2f m^3, 但得到 %.2f m^3", expected, m3)
	}
}

func TestVolume_CubicInch(t *testing.T) {
	// 1 cm^3 = 0.0610237 inch^3
	v := Volume{value: decimal.NewFromFloat(100)}
	expected := 6.10
	if inch3 := v.CubicInch(2); !almostEqual(inch3, expected) {
		t.Errorf("期望为 %.2f inch^3, 但得到 %.2f inch^3", expected, inch3)
	}
}
