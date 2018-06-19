package main

import "fmt"

type Values map[string][]string

type Test []string

func (t Test) Show() {
	if t == nil {
		fmt.Println("nil Test")
	}
	fmt.Println("Show")
}

func (v Values) Get(key string) string {
	if v == nil {
		fmt.Println("nil")
	}

	if vs := v[key]; len(vs) > 0 {
		return vs[0]
	}

	return ""
}

func (v Values) Add(key, value string) {
	v[key] = append(v[key], value)
}

func main() {
	m := Values{"lang": {"en"}}
	var s Test
	//	s := Test{"laoliao"}
	m.Add("item", "1")
	m.Add("item", "1")

	fmt.Println(m.Get("lang"))
	m = nil
	//	s={"laoliao"}
	Test(s).Show()
	//	s = nil
	//	s.Show()
	fmt.Println(Values(nil).Get("item"))
	//	fmt.Println(m.Get("item"))

	//	var t = Test{"main"}
	//	t.Show()
	//	fmt.Println(t.name)
}
