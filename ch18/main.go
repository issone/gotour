package main

func Fibonacci(n int) int {

	if n < 0 {

		return 0

	}

	if n == 0 {

		return 0

	}

	if n == 1 {

		return 1

	}

	return Fibonacci(n-1) + Fibonacci(n-2)

}

// 常用命令

// go test -v ./ch18	这行命令会运行 ch18 目录下的所有单元测试

// 单元测试覆盖率: 用来测试是否被全面的测试了
// go test -v --coverprofile=ch18.cover ./ch18	得到一个单元测试覆盖率文件，运行这行命令还可以同时看到测试覆盖率

// go tool cover -html ch18.cover -o ch18.html	 根据覆盖率文件，得到一个 HTML 格式的单元测试覆盖率报告

// 基准测试（Benchmark）： 是一项用于测量和评估软件性能指标的方法，主要用于评估代码的性能。

// go test -bench . ./ch18		 -bench 接受一个表达式作为参数，以匹配基准测试的函数，"."表示运行所有基准测试

// 测试结果如下：
// goos: windows
// goarch: amd64
// pkg: go_t/ch18
// BenchmarkFibonacci-4     3252027               370 ns/op
// PASS
// ok      go_t/ch18       1.784s

// 结果说明：
// 函数后面的 -4 ，这个表示运行基准测试时对应的 GOMAXPROCS 的值。接着的 3252027 表示运行 for 循环的次数，也就是调用被测试代码的次数，最后的 370 ns/op 表示每次需要花费 370 纳秒。
// 基准测试的时间默认是 1 秒，也就是 1 秒调用 3252027 次、每次调用花费 370 纳秒。如果想让测试运行的时间更长，可以通过 -benchtime 指定，比如 3 秒
// go test -bench . -benchtime 3s ./ch18

// 计时方法

// 进行基准测试之前会做一些准备，比如构建测试数据等，这些准备也需要消耗时间，所以需要把这部分时间排除在外。这就需要通过 ResetTimer 方法重置计时器, 这样可以避免因为准备数据耗时造成的干扰。
// 除了 ResetTimer 方法外，还有 StartTimer 和 StopTimer 方法，帮你灵活地控制什么时候开始计时、什么时候停止计时。

// 内存统计

// 在基准测试时，还可以统计每次操作分配内存的次数，以及每次操作分配的字节数，这两个指标可以作为优化代码的参考。要开启内存统计也比较简单，即通过在对应方法中调用ReportAllocs() 方法：
// 开启后再原来的基准测试多了两个指标，分别是 0 B/op 和 0 allocs/op。前者表示每次操作分配了多少字节的内存，后者表示每次操作分配内存的次数。这两个指标可以作为代码优化的参考，尽可能地越小越好
// 因为有时候代码实现需要空间换时间，所以要根据自己的具体业务而定，做到在满足业务的情况下越小越好

// 如果要对所有单元测试都进行内存统计，可以使用-benchmem 参数，如
// go test -bench . -benchmem  ./ch18  	这种通过 -benchmem 查看内存的方法适用于所有的基准测试用例

// 根据基准测试分析，Fibonacci函数没有分配新的内存，也就是说 Fibonacci 函数慢并不是因为内存，排除掉这个原因，就可以归结为所写的算法问题了
// 在递归运算中，一定会有重复计算，这是影响递归的主要因素。解决重复计算可以使用缓存，把已经计算好的结果保存起来，就可以重复使用了。
// 基于这个思路进行修改

var cache = map[int]int{} //缓存已经计算的结果

func Fibonacci2(n int) int {

	if v, ok := cache[n]; ok {

		return v

	}

	result := 0

	switch {

	case n < 0:

		result = 0

	case n == 0:

		result = 0

	case n == 1:

		result = 1

	default:

		result = Fibonacci2(n-1) + Fibonacci2(n-2)

	}

	cache[n] = result

	return result

}
