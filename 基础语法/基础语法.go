package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"strings"
	"time"
)

func main() {
	study17()
}
func study1() {
	fmt.Println("hello world")
}

// 变量
func study2() {
	var a = "initial"

	var b, c int = 1, 2

	var d = true

	var e float64

	f := float32(e)

	g := a + "foo"
	fmt.Println(a, b, c, d, e, f)
	fmt.Println(g)

	const s string = "constant"
	const h = 50000
	const i = 3e20 / h
	fmt.Println(s, h, i, math.Sin(h), math.Sin(i))
}

// if else
func study3() {
	if 7%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}
	if num := 9; num < 0 {
		fmt.Println("1")
	} else if num < 10 {
		fmt.Println("2")
	} else {
		fmt.Println("3")
	}
}

// 循环
func study4() {
	i := 1
	for {
		fmt.Println("1")
		break
	}
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}
	for n := 0; n < 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println("2")
	}
	for i <= 3 {
		fmt.Println(i)
		i++
	}
}

// switch
func study5() {
	a := 2
	switch a {
	case 1:
		fmt.Print("one")
	case 2:
		fmt.Print("two")
	case 3:
		fmt.Print("three")
	case 4, 5:
		fmt.Print("4,5")
	default:
		fmt.Print("other")
	}
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Print("1")
	default:
		fmt.Print("2")
	}
}

// 数组
func study6() {
	var a [5]int
	a[4] = 100
	fmt.Print(a[4], len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Print(b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Print(twoD)

}

// 切片
func study7() {
	s := make([]string, 3)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Print("get", s[2])
	fmt.Print("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Print(s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Print(c)

	// 不支持负索引
	fmt.Print(s[2:5])
	fmt.Print(s[:5])
	fmt.Println(s[2:])

	good := []string{"g", "o", "o", "d"}
	fmt.Print(good)
}

// map
func study8() {
	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	fmt.Print(m)
	fmt.Print(len(m))
	fmt.Print(m["one"])
	fmt.Print(m["unknow"])

	// r表示值 ok表示是否存在该键
	r, ok := m["unknow"]
	fmt.Print(r, ok)

	delete(m, "one")

	m2 := map[string]int{"1": 1, "2": 2}
	var m3 = map[string]int{"1": 1, "2": 2}
	fmt.Print(m2, m3)
}

// range
func study9() {
	nums := []int{2, 3, 4}
	sum := 0
	for i, num := range nums {
		sum += num
		if num == 2 {
			fmt.Println("index:", i, "num:", num)
		}
	}
	fmt.Println(sum)

	m := map[string]string{"1": "!", "2": "@"}
	for k, v := range m {
		fmt.Println(k, v)
	}
	for k := range m {
		fmt.Println("key", k)
	}
	for _, v := range m {
		fmt.Println("value", v)
	}
}

// 函数 func [name] (Parameter,...) (Parameter,...){}
func add(a int, b int) int {
	return a + b
}
func exists(m map[string]string, k string) (v string, ok bool) {
	v, ok = m[k]
	return v, ok
}
func study10() {
	res := add(1, 2)
	fmt.Println(res)

	v, ok := exists(map[string]string{"a": "A", "b": "B"}, "a")
	fmt.Println(v, ok)
}

// 指针
func add2(n int) {
	n += 2
}
func add2ptr(n *int) {
	*n += 2
}
func study11() {
	n := 5
	add2(n)
	fmt.Println(n)
	add2ptr(&n)
	fmt.Println(n)
}

// 结构体
type user struct {
	name string
	pass string
}

func checkPassword(u user, s string) bool {
	return u.pass == s

}
func study12() {
	a := user{name: "zhao", pass: "1024"}
	b := user{"zhao", "1024"}
	c := user{name: "zhao"}
	c.pass = "1024"
	var d user
	d.name = "zhao"
	d.pass = "1024"

	fmt.Println(a, b, c, d)
	fmt.Println(checkPassword(a, "haha"))
}

// 结构体方法
func (u user) checkPassword(pass string) bool {
	return u.pass == pass
}
func (u *user) resetPassword(pass string) {
	u.pass = pass
}
func study13() {
	a := user{name: "zhao", pass: "1024"}
	a.resetPassword("2024")
	fmt.Println(a.checkPassword("1024"))
}

// 错误处理 import errors
func findUser(users []user, name string) (v *user, err error) {
	for _, u := range users {
		if u.name ==
			name {
			return &u, nil
		}
	}
	return nil, errors.New("not found")
}
func study14() {
	u, err := findUser([]user{{"wang", "1024"}}, "wang")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(u.name)

	if u, err := findUser([]user{{"wang", "1024"}}, "li"); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(u.name)
	}
}

// 字符串操作  import strings
func study15() {
	a := "hello"
	fmt.Println(strings.Contains(a, "ll"))
	fmt.Println(strings.Count(a, "l"))
	fmt.Println(strings.HasPrefix(a, "he"))  //检测字符是否以指定前缀开头
	fmt.Println(strings.HasSuffix(a, "llo")) //检测字符是否以指定后缀开头
	fmt.Println(strings.Index(a, "ll"))
	fmt.Println(strings.Join([]string{"he", "llo"}, "-"))
	fmt.Println(strings.Repeat(a, 2))
	fmt.Println(strings.Replace(a, "e", "E", -1)) //最后一个参数表示替换a中前n个，n<0替换全部
	fmt.Println(strings.Split("a-b-c", "-"))
	fmt.Println(strings.ToLower(a))
	fmt.Println(strings.ToUpper(a))
	fmt.Println(len(a))
	b := "你好"
	fmt.Println(len(b))
}

// 字符串格式化
type point struct {
	x, y int
}

func study16() {
	s := "hello"
	n := 123
	p := point{1, 2}
	fmt.Println(s, n)
	fmt.Println(p)

	fmt.Printf("s=%v\n", s)
	fmt.Printf("n=%v\n", n)
	fmt.Printf("p=%v\n", p)
	fmt.Printf("p=%+v\n", p)
	fmt.Printf("p=%#v\n", p)

	f := 3.141592623
	fmt.Println(f)
	fmt.Printf("%.2f\n", f)
}

// json 处理
type userInfo struct {
	Name  string
	Age   int `json:"age"`
	Hobby []string
}

func study17() {
	a := userInfo{Name: "zhao", Age: 18, Hobby: []string{"Golang", "TypeScript"}}
	buf, err := json.Marshal(a)
	if err != nil {
		panic(err) //类似断言,终止程序运行
	}
	fmt.Println(buf)
	fmt.Println(string(buf))

	// MarshalIndent(v any, prefix, indent string)([]byte, error) 按照一定格式输出
	/*
		{
			<prefix><indent>"a": 1,
			<prefix><indent>"b": 2
			<prefix>}
	*/
	buf, err = json.MarshalIndent(a, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf))

	var b userInfo
	err = json.Unmarshal(buf, &b) //解码函数
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", b)
}
