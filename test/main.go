package main

import "fmt"

// 测试1
type jiekou interface {
	test_jiekoufunc(int, int)
}
type shixian func(int, int)

func (s shixian) test_jiekoufunc(i int, j int) {
	s(i, j)
}

func test(jk jiekou) {
	fmt.Println("参数为接口的函数")
}

// 测试2
type yy int

func (*yy) test_f(s string) {
	fmt.Println(s)

}

func main() {
	//测试1 //参数类型是接口的地方，可以用实现该接口的属性实例/指针代替
	var shilijk jiekou //实例化接口
	var sx *shixian    //实现该接口属性的实例
	var sx2 *shixian   //实现该接口属性的指针
	test(shilijk)
	test(sx)
	test(sx2)

	//测试2
	var a yy = 23
	var b *yy
	b = &a
	s := "nihao"
	a.test_f(s)
	b.test_f(s)

}
