# Volumetrically

一个用于材积相关的 Go 库。

## 安装

```bash
go get github.com/hiscaler/volumetrically
```

## 用法

### 计算体积

您可以使用厘米或英寸计算体积。

#### 使用厘米（默认）

```go
package main

import (
	"fmt"
	"github.com/hiscaler/volumetrically"
)

func main() {
	v := volumetrically.New().SetSize(10, 10, 10) // 10厘米 x 10厘米 x 10厘米
	volume := v.Volume(2)              // 计算体积，保留2位小数
	fmt.Printf("体积: %.2f cm³\n", volume)
}
```

#### 使用英寸

```go
package main

import (
	"fmt"
	"github.com/hiscaler/volumetrically"
)

func main() {
	v := volumetrically.New().SetSize(10, 10, 10).UseInch() // 10英寸 x 10英寸 x 10英寸
	volume := v.Volume(2)                       // 计算体积，保留2位小数
	fmt.Printf("体积: %.2f cm³\n", volume)
}
```

### 计算体积重量

计算出体积后，您可以使用转换因子计算体积重量。

```go
package main

import (
	"fmt"
	"github.com/hiscaler/volumetrically"
)

func main() {
	v := volumetrically.New().SetSize(10, 20, 30) // 10厘米 x 20厘米 x 30厘米
	
	// 使用转换因子 5000 计算重量
	weight := v.Weight(5000)

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

### `(v *Volumetrically) SetSize(length, width, height float64) *Volumetrically`

设置 `Volumetrically` 实例的尺寸。

### `(v *Volumetrically) UseCentimeter() *Volumetrically`

将测量单位设置为厘米。

### `(v *Volumetrically) UseInch() *Volumetrically`

将测量单位设置为英寸。

### `(v *Volumetrically) Volume(precision int32) float64`

计算并返回以立方厘米为单位的体积，并四舍五入到指定的精度。

### `(v *Volumetrically) Weight(factor int32) Weight`

通过将体积除以给定的因子来计算体积重量。返回一个 `Weight` 对象。

### `(w Weight) Gram(precision int32) float64`

返回以克为单位的重量。

### `(w Weight) Kilogram(precision int32) float64`

返回以千克为单位的重量。

### `(w Weight) Pound(precision int32) float64`

返回以磅为单位的重量。
