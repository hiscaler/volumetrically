package volumetrically

import (
	"fmt"
	"strings"

	"github.com/shopspring/decimal"
)

// Volumetrically 用于计算体积和体积重量。
type Volumetrically struct {
	length, width, height decimal.Decimal // 长、宽、高（默认单位：厘米）
	volume                *Volume         // 缓存计算出的体积，避免重复计算
}

// New 创建一个新的 Volumetrically 实例。
func New() *Volumetrically {
	return &Volumetrically{}
}

// SetSize 设置尺寸。
// length: 长度
// width: 宽度
// height: 高度
// unit: 单位 (cm, m, inch)
func (v *Volumetrically) SetSize(length, width, height float64, unit string) (*Volumetrically, error) {
	if length <= 0 || width <= 0 || height <= 0 {
		return v, fmt.Errorf("尺寸参数必须大于 0")
	}

	l := decimal.NewFromFloat(length)
	w := decimal.NewFromFloat(width)
	h := decimal.NewFromFloat(height)
	// 将所有单位转换为厘米进行内部存储和计算
	switch strings.ToLower(unit) {
	case Centimeter:
		v.length, v.width, v.height = l, w, h
	case Meter:
		d := decimal.NewFromFloat(100)
		v.length, v.width, v.height = l.Mul(d), w.Mul(d), h.Mul(d)
	case Inch:
		d := decimal.NewFromFloat(2.54)
		v.length, v.width, v.height = l.Mul(d), w.Mul(d), h.Mul(d)
	default:
		return v, fmt.Errorf("无效的尺寸单位: %s", unit)
	}
	// 尺寸变化后，清除已缓存的体积
	v.volume = nil

	return v, nil
}

// Volume 计算体积。
// 如果已经计算过体积，将返回缓存的值。
func (v *Volumetrically) Volume() Volume {
	if v.volume != nil {
		return *v.volume
	}
	l := v.length
	w := v.width
	h := v.height

	v.volume = &Volume{value: l.Mul(w).Mul(h)}
	return *v.volume
}

// Weight 计算体积重量。
// factor: 体积重量系数，例如 5000 或 6000
// unit: 用于计算体积重量的体积单位 (cm, m, inch)
func (v *Volumetrically) Weight(factor int32, unit string) (Weight, error) {
	if factor <= 0 {
		return Weight{}, fmt.Errorf("系数必须大于 0")
	}
	var volume Volume
	if v.volume == nil {
		volume = v.Volume()
	} else {
		volume = *v.volume
	}

	var cubicCM decimal.Decimal
	switch strings.ToLower(unit) {
	case Centimeter:
		cubicCM = decimal.NewFromFloat(volume.CubicCentimeter(-1))
	case Meter:
		cubicCM = decimal.NewFromFloat(volume.CubicMeter(-1))
	case Inch:
		cubicCM = decimal.NewFromFloat(volume.CubicInch(-1))
	default:
		return Weight{}, fmt.Errorf("无效的体积单位: %s", unit)
	}
	return Weight{cubicCM.Div(decimal.NewFromInt32(factor))}, nil
}
