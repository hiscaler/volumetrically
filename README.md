# Volumetrically

一个用于材积和体积重量计算相关的 Go 库。

## 安装

```bash
go get github.com/hiscaler/volumetrically
```

## 用法

### 计算体积

您可以设置不同的尺寸单位（`cm`, `m`, `inch`）来计算体积。

```go
package main

import (
	"fmt"
	"github.com/hiscaler/volumetrically"
	"log"
)

func main() {
	v, err := volumetrically.New().SetSize(10.5, 10, 10, volumetrically.Centimeter) // 10.5cm x 10cm x 10cm
	if err != nil {
		log.Fatalf("设置尺寸失败: %v", err)
	}

	// 计算体积
	volume := v.Volume()

	// 获取不同单位的体积
	cubicCentimeters := volume.CubicCentimeter(2)
	cubicMeters := volume.CubicMeter(8)
	cubicInches := volume.CubicInch(2)

	fmt.Printf("体积 (立方厘米): %.2f cm³\n", cubicCentimeters)
	fmt.Printf("体积 (立方米): %.8f m³\n", cubicMeters)
	fmt.Printf("体积 (立方英寸): %.2f in³\n", cubicInches)
}
```

### 计算体积重量

计算出体积后，您可以使用转换系数计算体积重量。

```go
package main

import (
	"fmt"
	"github.com/hiscaler/volumetrically"
	"log"
)

func main() {
	v, err := volumetrically.New().SetSize(10, 20, 30, volumetrically.Centimeter) // 10cm x 20cm x 30cm
	if err != nil {
		log.Fatalf("设置尺寸失败: %v", err)
	}
	
	// 使用转换系数 5000 和厘米单位计算重量
	// 体积重量（千克）= (长 x 宽 x 高 (cm)) / 5000
	weight, err := v.Weight(5000, volumetrically.Centimeter)
	if err != nil {
		log.Fatalf("计算重量失败: %v", err)
	}

	// 获取不同单位的重量
	grams := weight.Gram(2)
	kilograms := weight.Kilogram(2)
	pounds := weight.Pound(2)

	fmt.Printf("重量 (克): %.2f g\n", grams)
	fmt.Printf("重量 (千克): %.2f kg\n", kilograms)
	fmt.Printf("重量 (磅): %.2f lbs\n", pounds)
}
```

## 方法

### `volumetrically.New() *Volumetrically`

创建一个新的 `Volumetrically` 空实例。

### `(v *Volumetrically) SetSize(length, width, height float64, unit string) (*Volumetrically, error)`

设置 `Volumetrically` 实例的尺寸。
- `length`, `width`, `height`: 浮点数格式的长、宽、高。
- `unit`: 尺寸单位，可以是 `volumetrically.Centimeter` (`"cm"`), `volumetrically.Meter` (`"m"`) 或 `volumetrically.Inch` (`"inch"`)。

### `(v *Volumetrically) Volume() Volume`

计算并返回一个 `Volume` 对象。`Volumetrically` 内部会缓存计算结果，重复调用此方法不会重新计算。

### `(v *Volumetrically) Weight(factor int32, unit string) (Weight, error)`

计算体积重量。
- `factor`: 体积重量系数 (例如: 5000, 6000)。
- `unit`: 用于计算的体积单位，可以是 `volumetrically.Centimeter` (`"cm"`), `volumetrically.Meter` (`"m"`) 或 `volumetrically.Inch` (`"inch"`)。

**注意**：`Weight` 方法的当前逻辑是将指定单位的体积值除以系数，并将结果直接作为 `Weight` 对象的内部值。`Weight` 对象的 `Gram()` 方法会直接返回此内部值。这意味着，只有当 `unit` 为 `"cm"` 并且 `factor` 的含义是 `cm³/g` 时，`Gram()` 的结果在物理意义上才是准确的。在标准的 `cm³/kg` 系数下，您需要自行进行单位换算。

---

### `(v Volume) CubicCentimeter(precision int32) float64`

返回以**立方厘米**为单位的体积。
- `precision`: 小数点精度。如果为负数，则不进行截断。

### `(v Volume) CubicMeter(precision int32) float64`

返回以**立方米**为单位的体积。
- `precision`: 小数点精度。如果为负数，则不进行截断。

### `(v Volume) CubicInch(precision int32) float64`

返回以**立方英寸**为单位的体积。
- `precision`: 小数点精度。如果为负数，则不进行截断。

---

### `(w Weight) Gram(precision int32) float64`

返回以**克**为单位的重量。
- `precision`: 小数点精度。如果为负数，则不进行四舍五入。

### `(w Weight) Kilogram(precision int32) float64`

返回以**千克**为单位的重量。
- `precision`: 小数点精度。如果为负数，则不进行四舍五入。

### `(w Weight) Pound(precision int32) float64`

返回以**磅**为单位的重量。
- `precision`: 小数点精度。如果为负数，则不进行四舍五入。

