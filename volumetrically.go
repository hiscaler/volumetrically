package volumetrically

import "github.com/shopspring/decimal"

type Volumetrically struct {
	isInch                bool                // 是否使用英寸
	volume                decimal.NullDecimal // 临时存储的体积
	length, width, height float64             // 长、宽、高（默认单位：厘米）
}

func New() *Volumetrically {
	return &Volumetrically{
		isInch: false,
	}
}

func (v *Volumetrically) SetSize(length, width, height float64) *Volumetrically {
	v.length, v.width, v.height = length, width, height
	return v
}

func (v *Volumetrically) UseCentimeter() *Volumetrically {
	v.isInch = false
	return v
}

func (v *Volumetrically) UseInch() *Volumetrically {
	v.isInch = true
	return v
}

func (v *Volumetrically) Volume(precision int32) float64 {
	l := decimal.NewFromFloat(v.length)
	w := decimal.NewFromFloat(v.width)
	h := decimal.NewFromFloat(v.height)

	if v.isInch {
		inchFactor := decimal.NewFromFloat(2.54)
		l = l.Mul(inchFactor)
		w = w.Mul(inchFactor)
		h = h.Mul(inchFactor)
	}
	volume := l.Mul(w).Mul(h)
	v.volume = decimal.NewNullDecimal(volume)
	return volume.Round(precision).InexactFloat64()
}
func (v *Volumetrically) Weight(factor int32) Weight {
	var volume decimal.Decimal
	if !v.volume.Valid {
		_ = v.Volume(-1)
	}
	volume = v.volume.Decimal
	return Weight{volume.Div(decimal.NewFromInt32(factor))}
}
