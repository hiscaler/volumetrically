package volumetrically

import "github.com/shopspring/decimal"

type Weight struct {
	decimal.Decimal
}

// Gram 返回克，precision表示小数点位数
func (w Weight) Gram(precision int32) float64 {
	return w.Round(precision).InexactFloat64()
}

// Kilogram 返回千克，precision表示小数点位数
func (w Weight) Kilogram(precision int32) float64 {
	return w.Div(decimal.NewFromInt(1000)).Round(precision).InexactFloat64()
}

// Pound 磅，precision表示小数点位数
func (w Weight) Pound(precision int32) float64 {
	return w.Div(decimal.NewFromFloat(453.59237)).Round(precision).InexactFloat64()
}
