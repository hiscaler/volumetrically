package volumetrically

import (
	"testing"
)

func TestVolumetrically_SetSize(t *testing.T) {
	v := New()
	// 测试使用厘米
	v, err := v.SetSize(10, 10, 10, Centimeter)
	if err != nil {
		t.Fatalf("使用厘米设置尺寸失败: %v", err)
	}
	expectedVolume := 1000.0
	if volume := v.Volume().CubicCentimeter(-1); !almostEqual(volume, expectedVolume) {
		t.Errorf("期望体积为 %.2f, 但得到 %.2f", expectedVolume, volume)
	}

	// 测试使用米
	v, err = v.SetSize(0.1, 0.1, 0.1, Meter)
	if err != nil {
		t.Fatalf("使用米设置尺寸失败: %v", err)
	}
	if volume := v.Volume().CubicCentimeter(-1); !almostEqual(volume, expectedVolume) {
		t.Errorf("期望体积为 %.2f, 但得到 %.2f", expectedVolume, volume)
	}

	// 测试使用英寸
	v, err = v.SetSize(10, 10, 10, Inch)
	if err != nil {
		t.Fatalf("使用英寸设置尺寸失败: %v", err)
	}
	// 10 英寸 = 25.4 厘米. 25.4 * 25.4 * 25.4 = 16387.064
	expectedVolumeInch := 16387.064
	if volume := v.Volume().CubicCentimeter(-1); !almostEqual(volume, expectedVolumeInch) {
		t.Errorf("期望体积为 %.3f, 但得到 %.3f", expectedVolumeInch, volume)
	}

	// 测试无效尺寸
	_, err = v.SetSize(0, 10, 10, Centimeter)
	if err == nil {
		t.Error("期望尺寸为 0 时报错，但未报错")
	}

	// 测试无效单位
	_, err = v.SetSize(10, 10, 10, "invalid_unit")
	if err == nil {
		t.Error("期望单位无效时报错，但未报错")
	}
}

func TestVolumetrically_Weight(t *testing.T) {
	// 10 厘米 * 20 厘米 * 30 厘米 = 6000 立方厘米
	v, err := New().SetSize(10, 20, 30, Centimeter)
	if err != nil {
		t.Fatalf("设置尺寸失败: %v", err)
	}
	// 6000 / 5000 = 1.2
	weight, err := v.Weight(5000, Centimeter)
	if err != nil {
		t.Fatalf("计算重量失败: %v", err)
	}
	expectedGram := 1.2
	if gram := weight.Gram(2); !almostEqual(gram, expectedGram) {
		t.Errorf("期望重量为 %.2f g, 但得到 %.2f g", expectedGram, gram)
	}

	// 新增测试用例：厘米尺寸，材积系数单位使用英寸，得到克
	// 尺寸: 10cm x 20cm x 30cm = 6000 立方厘米
	// 转换为立方英寸: 6000 cm^3 * 0.061023744094732 in^3/cm^3 = 366.142464568392 立方英寸
	// 重量计算（根据当前 Weight 函数的逻辑）：366.142464568392 / 5000 = 0.0732284929136784
	// Weight.Gram(2) 将其四舍五入到两位小数：0.07 克
	weightInchFactor, err := v.Weight(5000, Inch)
	if err != nil {
		t.Fatalf("使用英寸系数计算重量失败: %v", err)
	}
	expectedGramInchFactor := 0.07
	if gram := weightInchFactor.Gram(2); !almostEqual(gram, expectedGramInchFactor) {
		t.Errorf("期望使用英寸系数计算出的重量为 %.2f g, 但得到 %.2f g", expectedGramInchFactor, gram)
	}

	// 测试无效系数
	_, err = v.Weight(0, Centimeter)
	if err == nil {
		t.Error("期望系数为 0 时报错，但未报错")
	}

	// 测试无效单位
	_, err = v.Weight(5000, "invalid_unit")
	if err == nil {
		t.Error("期望单位无效时报错，但未报错")
	}
}
