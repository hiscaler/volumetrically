package volumetrically

import "github.com/shopspring/decimal"

const (
	Centimeter = "cm"   // 厘米
	Meter      = "m"    // 米
	Inch       = "inch" // 英寸
)

// Volume 表示体积，内部使用 decimal.Decimal 存储，单位为立方厘米 (cm^3)。
type Volume struct {
	value decimal.Decimal
}

// CubicCentimeter 返回立方厘米，precision 表示小数点位数。
// 如果 precision < 0, 则不进行四舍五入。
func (v Volume) CubicCentimeter(precision int32) float64 {
	if precision < 0 {
		return v.value.InexactFloat64()
	}
	return v.value.Round(precision).InexactFloat64()
}

// CubicMeter 返回立方米，precision 表示小数点位数。
// 1 立方厘米 = 1.0 × 10^-6 立方米。
// 如果 precision < 0, 则不进行四舍五入。
func (v Volume) CubicMeter(precision int32) float64 {
	d := v.value.Mul(decimal.NewFromFloat(1.0e-6))
	if precision < 0 {
		return d.InexactFloat64()
	}
	return d.Round(precision).InexactFloat64()
}

// CubicInch 返回立方英寸，precision 表示小数点位数。
// 1 立方厘米 = 0.061023744094732 立方英寸。
// 如果 precision < 0, 则不进行四舍五入。
func (v Volume) CubicInch(precision int32) float64 {
	d := v.value.Mul(decimal.NewFromFloat(0.061023744094732))
	if precision < 0 {
		return d.InexactFloat64()
	}
	return d.Round(precision).InexactFloat64()
}
