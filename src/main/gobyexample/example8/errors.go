package example8

import (
	"errors"
	"fmt"
)

// Go 语言使用一个独立的·明确的返回值来传递错误信息的。这与
// 使用异常的 Java 和 Ruby 以及在 C 语言中经常见到的超重的
// 单返回值/错误值相比，Go 语言的处理方式能清楚的知道哪个函数
// 返回了错误，并能像调用那些没有出错的函数一样调用。
func main() {

	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed: ", e)
		} else {
			fmt.Println("f1 worked: ", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed: ", e)
		} else {
			fmt.Println("f2 worked: ", r)
		}
	}

	// 你如果想在程序中使用一个自定义错误类型中的数据，需要通过类型断言来得到这个错误类型的实例。
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}

}

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}

	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

// 在这个例子中，我们使用 `&argError` 语法来建立一个
// 新的结构体，并提供了 `arg` 和 `prob` 这个两个字段
// 的值。
func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}

	return arg + 3, nil
}
