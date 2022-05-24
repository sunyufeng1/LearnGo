package fileUtils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"text/template"
)

//创建目录
func CreateDir() {
	//os.Mkdir("newDie",os.ModePerm)需要一层一层创建
	os.MkdirAll("newDir/N/test/我/m.txt", os.ModePerm) //会连需要的文件目录一起创建 只会创建目录不会创建文件
}

//创建文件
func CreateFile() {
	os.Create("newDied/gg.g") //目录不存在的话不会创建
}

//创建模板文件
func CreateTemplateFile() {
	name := "waynehu"
	tmpl, err := template.New("test").Parse("hello, {{.}}") //建立一个模板，内容是"hello, {{.}}"
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, name) //将string与模板合成，变量name的内容会替换掉{{.}}
	//合成结果放到os.Stdout里
	if err != nil {
		panic(err)
	}
}

//小写的话就要写setter和getter 方法首字母要大写
//直接使用属性的话 属性首字母要大写
type TestA struct {
	A string
	B string
}

//统一解析html模板 避免 template.ParseFiles 对同一个文件反复解析
var templates = template.Must(template.ParseFiles("./template/template.txt"))

//根据模板文件 和参数创建 目标文件
func CreateFileByTemplateFile() {

	testPage := new(TestA)
	testPage.A = "111"
	testPage.B = "2222"
	println(testPage.A)
	println(testPage.B)
	//kkk := &TestA{a: "title", b: "body"}
	err := templates.ExecuteTemplate(os.Stdout, "template.txt", testPage) //这个只是打印
	if err != nil {
		println(err.Error())
	}
	dir := "./template/target"
	fileName := "target.txt"
	os.MkdirAll(dir, os.ModePerm) //会连需要的文件目录一起创建 只会创建目录不会创建文件
	fullName := dir + "/" + fileName
	os.Create(fullName)
	target, err := os.OpenFile(fullName, os.O_RDWR|os.O_CREATE, os.ModePerm) //os.open只能读不能写
	if err != nil {
		println(err.Error())
		return
	}
	err1 := templates.ExecuteTemplate(target, "template.txt", testPage) //这个才是把模板信息写入文件
	if err1 != nil {
		println(err1.Error())
	}
}

//解析文件路径
func AnalyFileUrl() {

}

//解析文件名
func AnalyFileName() {
	fileUrl := "./gggg/23/k.txt"
	base := filepath.Base(fileUrl) //k.txt
	//println(filepath.Dir(fileUrl))
	//println(filepath.Ext(fileUrl))//后缀
	//println(filepath.FromSlash(fileUrl))
	//dir, file := filepath.Split(fileUrl)
	//fmt.Printf("input: \n\tdir: %q\n\tfile: %q\n", dir, file)
	//println(filepath.IsAbs(fileUrl))//是否是绝对路径
	ext := filepath.Ext(fileUrl)
	filenameOnly := strings.TrimSuffix(base, ext)
	println(filenameOnly)
}

//解析后缀
func AnalyFileSuffix() {
	fileUrl := "./gggg/23/k.txt"
	println(filepath.Ext(fileUrl)) //后缀
}

//文件的复制
func CopyFile() {

}

//文件删除
func DeleteFile() {

}

//解析json文件的配置参数
func GetJsonFileArgs() {
	JsonParse := NewJsonStruct()
	v := new(Config)
	//下面使用的是相对路径，config.json文件和main.go文件处于同一目录下
	JsonParse.Load("./json/config.json", v)
	fmt.Println(v.Addr)
	fmt.Println(v.Mongo.MongoDb)
	fmt.Println((v.Mongo.MongoPoolLimit))
	fmt.Println(reflect.TypeOf(v.Mongo.MongoPoolLimit).String() == "int")
}

//定义配置文件解析后的结构
type MongoConfig struct {
	MongoAddr      string
	MongoPoolLimit int
	MongoDb        string
	MongoCol       string
}

type Config struct {
	Addr  string
	Mongo MongoConfig
}

type JsonStruct struct {
}

func NewJsonStruct() *JsonStruct {
	return new(JsonStruct)
}

func (jst *JsonStruct) Load(filename string, v interface{}) {
	//ReadFile函数会读取文件的全部内容，并将结果以[]byte类型返回
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	//读取的数据为json格式，需要进行解码
	err = json.Unmarshal(data, v)
	if err != nil {
		return
	}
}
