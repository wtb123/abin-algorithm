## 一、函数是一等公民
**与其他主要编程语言的差异**
1. 可以有多个返回值
2. 所有参数都是值传递：slice, map, channel 会有传递引用的错觉
3. 函数可以作为变量的值
4. 函数可以作为参数和返回值
```go
// 函数式编程，有点类似装饰者模式，给原来函数加上了一个功能

func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner((n))

		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestFn(t *testing.T) {
	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))
}
```
