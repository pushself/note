package main

import (
	"fmt"
)

type Point struct{ Name string }
type Point2 struct {
	*Point
	Name2 string
}

func (p Point) Show() {
	fmt.Println(p.Name)
	p.Name = "laosu" //值拷贝，此处修改不会影响main函数的p
}

func (p Point) Show2(p2 Point) {
	fmt.Println(p2.Name)
	p2.Name = "laosu"
}

type IntSet struct {
    words []uint64
}

func main() {
	var i IntSet
	i.words.
	
	/*var p Point2

	p.Name = "laoliao"
	p.Name2 = "laomei"

	fmt.Println(p.Name)
	fmt.Println(p.Point.Name)
	fmt.Println(p.Name2)

	p.Show()

	//p.Show2(p)//不要把p看作是子类活着继承类  这是错误的理解
	p.Show2(p.Point) //跟下面注释代码等价*/

	p := Point2{&Point{"laomei"}, "laoliao"}
	p2 := Point2{&Point{"laji"}, "laoliao"}
	p.Show()
	fmt.Println(p.Name)
	p.Show2(*p.Point)
	fmt.Println(p.Name)

	fmt.Println(p.Name)
	fmt.Println(p2.Name)
	//	p.Point = p2.Point
	//	p2.Name = "haha"
	//	p.Show()
	fmt.Println(p.Name)
	fmt.Println(p2.Name)

}

/*内嵌字段会指导编译器去生成额外的包装方法来委托已经声明好的方法，和下面的形式是等价的：*/

/*
func (p Point2) Show2(p2 Point) {
	p.Point.Show2(p2)
}
*/

//Point访问不到Point2的任何字段
