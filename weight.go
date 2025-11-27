package volumetrically

import "github.com/shopspring/decimal"

// Weight 表示重量，内部使用 decimal.Decimal 存储，单位为克 (g)。
type Weight struct {
	value decimal.Decimal
}

// Gram 返回克，precision 表示小数点位数。
// 如果 precision < 0, 则不进行四舍五入。
func (w Weight) Gram(precision int32) float64 {
	if precision < 0 {
		return w.value.InexactFloat64()
	}
	return w.value.Round(precision).InexactFloat64()
}

// Kilogram 返回千克，precision 表示小数点位数。
// 如果 precision < 0, 则不进行四舍五入。
func (w Weight) Kilogram(precision int32) float64 {
	d := w.value.Div(decimal.NewFromInt(1000))
	if precision < 0 {
		return d.InexactFloat64()
	}
	return d.Round(precision).InexactFloat64()
}

// Pound 返回磅，precision 表示小数点位数。
// 1 磅 = 453.59237 克。
// 如果 precision < 0, 则不进行四舍五入。
func (w Weight) Pound(precision int32) float64 {
	d := w.value.Div(decimal.NewFromFloat(453.59237))
	if precision < 0 {
		return d.InexactFloat64()
	}
	return d.Round(precision).InexactFloat64()
}
