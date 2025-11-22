package volumetrically

import (
	"testing"
)

func TestVolumetrically_Volume(t *testing.T) {
	v := New().SetSize(10, 10, 10)
	expectedVolume := 1000.0
	if volume := v.Volume(2); volume != expectedVolume {
		t.Errorf("期望的体积是 %.2f，但得到的是 %.2f", expectedVolume, volume)
	}
}

func TestVolumetrically_Volume_WithInch(t *testing.T) {
	v := New().SetSize(10, 10, 10).UseInch()
	// 10 英寸 = 25.4 厘米
	// 体积 = 25.4 * 25.4 * 25.4 = 16387.064
	expectedVolume := 16387.06
	if volume := v.Volume(2); volume != expectedVolume {
		t.Errorf("期望的体积是 %.2f，但得到的是 %.2f", expectedVolume, volume)
	}
}

func TestVolumetrically_Weight(t *testing.T) {
	v := New().SetSize(10, 20, 30) // 6000 立方厘米
	weight := v.Weight(5000)
	expectedGram := 1.2
	if gram := weight.Gram(2); gram != expectedGram {
		t.Errorf("期望的重量是 %.2f g，但得到的是 %.2f g", expectedGram, gram)
	}
}

func TestVolumetrically_SetSize(t *testing.T) {
	v := New()
	v.SetSize(10, 10, 10)
	expectedVolume := 1000.0
	if volume := v.Volume(2); volume != expectedVolume {
		t.Errorf("期望的体积是 %.2f，但得到的是 %.2f", expectedVolume, volume)
	}
}

func TestVolumetrically_UseCentimeter(t *testing.T) {
	v := New().SetSize(10, 10, 10).UseInch()
	v.UseCentimeter()
	expectedVolume := 1000.0
	if volume := v.Volume(2); volume != expectedVolume {
		t.Errorf("期望的体积是 %.2f，但得到的是 %.2f", expectedVolume, volume)
	}
}
