package volumetrically

import "math"

// float64EqualityThreshold 的作用是：
//
// 在计算机编程中，浮点数（也就是带小数的数字，比如 1.23 或
// 0.001）的存储方式决定了它们通常无法被精确地表示。这就意味着，即使两个浮点数在数学上应该相等，但在计算机内部，它们可能因为微小的舍入误差而略有不同。
//
// 因此，如果我们直接用 a == b 这样的方式去比较两个浮点数是否相等，往往会得到错误的结果，因为即使只差一点点（比如 0.9999999999999999 和 1.0），也会被认为是不同的。
//
// float64EqualityThreshold 就是为了解决这个问题而引入的一个极小的值（在我们的代码中是 1e-9，也就是 0.000000001）。
//
// 当我们想比较两个浮点数 a 和 b 是否可以被认为是相等的时候，我们不再直接比较 a == b，而是去看它们之间的差值的绝对值是否小于或等于这个 float64EqualityThreshold。
//
// 简单来说，就是：如果 `|a - b| <= float64EqualityThreshold` 成立，我们就认为 `a` 和 `b` 是相等的。
//
// 这在测试代码中非常重要，因为程序计算出来的浮点数结果很可能包含这种微小的误差。使用 float64EqualityThreshold
// 可以确保我们的测试不会因为这些在物理意义上可以忽略不计的误差而失败，从而使测试更加健壮和可靠
const float64EqualityThreshold = 1e-9

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold
}
