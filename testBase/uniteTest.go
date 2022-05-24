package testBase

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/sunyufeng1/CommTool/toolElse"
)

type TestObj struct {
	a int
	b int
}

//打印测试
func TestPrint() {
	fmt.Printf("testBase\n")
}

//对象作为参数
func TestObjArgs() {
	testObj := new(TestObj)
	testObj.a = 2
	testObj.b = 2
	t2 := testObj          //这里这样使用的是是引用传递 没有进行复制
	TestObjArgs1(*testObj) //这里*说明进行了拷贝  内部变化 不影响外部
	var buffer bytes.Buffer
	buffer.WriteString("A 变化后 ")
	buffer.WriteString(strconv.Itoa(testObj.a))
	fmt.Println(buffer.String()) //这里还是2

	TestObjArgs2(testObj)
	var buffer1 bytes.Buffer
	buffer1.WriteString("B 变化后 ")
	buffer1.WriteString(strconv.Itoa(testObj.b))
	fmt.Println(buffer1.String()) //这里还是2
	fmt.Print(t2.b)
}

func TestObjArgs1(obj TestObj) { // 说明进行了拷贝  //与其他语言不同 obj跟外部传进来的不是同一个了  但属性相同
	obj.a = 3
}

func TestObjArgs2(obj *TestObj) { // 没有进行拷贝 //该情况与其他语言一样
	obj.b = 3
}

func TestExit(num int) {
	print("begin")
	testExi1()
}

func testExi1() {
	testExi2()
	print("testExi1")
}

func testExi2() {
	testExi3()
	print("testExi2")
}

func testExi3() {
	testExi4()
	print("testExi3")
}

func testExi4() {
	os.Exit(2)
	print("testExi4")
}

//获得当前运行时候的路径
func GetCurrentPath() string {
	s, err := exec.LookPath(os.Args[0])
	if err != nil {
		fmt.Println(err.Error())
	}
	s = strings.Replace(s, "\\", "/", -1)
	s = strings.Replace(s, "\\\\", "/", -1)
	i := strings.LastIndex(s, "/")
	path := string(s[0 : i+1])
	return path
}

//路由测试
//func testRoute(){
//	http.HandleFunc("/",newLisntener)
//	//var listener Listener
//	err := http.ListenAndServe(":4000", nil)
//	if err != nil {
//		log.Fatal(err)
//	}
//}

//////////////////////别的的new写法
type options struct {
	a int64
	b string
	c map[int]string
}

func NewOption(opt ...ServerOption) *options {
	r := new(options)
	for _, o := range opt {
		o(r)
	}
	return r
}

type ServerOption func(*options)

func WriteA(s int64) ServerOption {
	return func(o *options) {
		o.a = s
	}
}

func WriteB(s string) ServerOption {
	return func(o *options) {
		o.b = s
	}
}

func WriteC(s map[int]string) ServerOption {
	return func(o *options) {
		o.c = s
	}
}

func main111() {
	opt1 := WriteA(int64(1))
	opt2 := WriteB("test")
	opt3 := WriteC(make(map[int]string, 0))

	op := NewOption(opt1, opt2, opt3)

	fmt.Println(op.a, op.b, op.c)
}

////////////////////

//检查是否是同一个类型
func CheckSameObj() {
	//obj1 := new(checkObj1)
	//obj2 := new(checkObj2)
	//obj3 := new(checkObj2)
	//
	//println(reflect.TypeOf(obj1) == checkObj1)
	//println(obj3 === obj2)
	v := new(C2)
	v.name = "eea"
	v1 := new(C1)
	v1.name = "eea"
	b := toolElse.InstanceOf(v, v1) //判断类型相同
	fmt.Println(b)
	fmt.Println("dd" == "dd1") //java 的话是.equal

}

type checkObj1 struct {
}

type checkObj2 struct {
}

type C3 interface {
	Doit()
	Acall()
	Ccall(s string)
}
type C1 struct {
	name string
}

func (c *C1) Doit() {
}
func (c *C1) Acall() {
}
func (c *C1) Ccall(s string) {

}

type C2 struct {
	name string
}
