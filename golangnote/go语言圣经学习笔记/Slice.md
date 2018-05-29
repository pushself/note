没有固定大小的就是Slice
```
test := []string {}
test2 := test[low:high:max]
/*
按索引0开始算，low包含，high、max不包含
否则low不包含，high、max包含
/*
```

数组变量地址等于下标为0的地址

从数组切片
```
func main(){
	//s1:=[]int{0,1,2,3,8:100}
	//fmt.Println(s1,len(s1),cap(s1))


	d:=[5]struct{x int}{

	}

//从数组切片，s的pointer指向d
/*
slice自身是结构体（c实现的编译器源码）
s.array指向d
make创建的slice是在底层创建一个数组
*/
	s:=d[:]

	//0xc8200760c0-0xc82006c080
	fmt.Printf("%p-%p\n", &d,&s)
	d[1].x=10
	s[2].x=20

	fmt.Println(d)

}
```

reslice
```
s:=[]int{0,1,2,3,4,5,6,7,8,9}

	s1:=s[2:5]
	s2:=s1[2:6:7]
	s3:=s2[3:6]//s3是在s2的cap范围内调整，数据还是以s为准
```

append就是往array[slice.high]


```
	s := make([]int, 0, 5)//len=cap   5=5
	fmt.Printf("%p\n", &s)//超出slice.cap限制，重新分配底层数组，即使原数组未填满，通常2倍容量重新分配底层数组
	s2 := append(s, 1)
	fmt.Printf("%p\n", &s2)
	fmt.Println(s, s2)
/*
0xc82000e400
0xc82000e420
[] [1]
*/
```
copy 两个slice可以指向同一个底层数组，也可以不一样
及时释放不再使用slice，避免持有过期数组，造成gc无法回收
```
	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := data[8:]
	s2 := data[:5]
	copy(s, s2)
	// dst:s2, src:s
	fmt.Println(s)
	fmt.Println(data)//应及时将所需数据 copy 到较小小的 slice,以便释放超大大号底层数组内存。
/*
[0 1]
[0 1 2 3 4 5 6 7 0 1]
*/
```