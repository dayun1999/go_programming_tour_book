package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"text/template"
)

// var name string

type Name string

const templateText = `
Output 0: {{title .Name1}}
Output 1: {{title .Name2}}
Output 2: {{.Name3 | title}}
`

func main() {
	// 示例6.初入template
	funcMap := template.FuncMap{"title": strings.Title}
	tpl := template.New("go-programming-tour")
	tpl, _ = tpl.Funcs(funcMap).Parse(templateText)
	data := map[string]string{
		"Name1": "go",
		"Name2": "programming",
		"Name3": "tour",
	}
	_ = tpl.Execute(os.Stdout, data)

	// 示例5.时区问题
	//location, _ := time.LoadLocation("Asia/Shanghai")
	//inputTime := "2021-12-01 12:00:00"
	//layout := "2006-01-02 15:04:05"
	//t, _ :=time.Parse(layout, inputTime)
	//dataTime := time.Unix(t.Unix(), 0).In(location).Format(layout)
	//log.Printf("输入时间: %s, 输出时间: %s", inputTime, dataTime)
	// 输入时间: 2021-12-01 12:00:00, 输出时间: 2021-12-01 20:00:00
	// 可见差了8个小时，是因为Parse()方法会尝试在入参的参数中分析并读取时区, 如果入参的参数没有指定时区信息,就会默认使用UTC时间

	// 示例5.时区问题解决方法
	//location, _ := time.LoadLocation("Asia/Shanghai")
	//inputTime := "2021-12-01 12:00:00"
	//layout := "2006-01-02 15:04:05"
	//t, _ := time.ParseInLocation(layout, inputTime, location)
	//dataTime := time.Unix(t.Unix(), 0).In(location).Format(layout)
	//log.Printf("输入时间: %s, 输出时间: %s", inputTime, dataTime)

	// 示例4.单词转换 cobra 和 时间转换
	//err := cmd.Execute()
	//if err != nil {
	//	log.Fatalf("cmd.Execute err: %v", err)
	//}

	// 示例3.自定义参数类型
	//var name Name
	//flag.Var(&name, "name", "help info")
	//flag.Parse()
	// 示例2.子命令的使用
	//flag.Parse()
	//args := flag.Args()
	//if len(args) == 0 {
	//	return
	//}
	//switch args[0] {
	//case "go":
	//	goCmd := flag.NewFlagSet("go", flag.ExitOnError)
	//	goCmd.StringVar(&name, "name","Go语言", "help info")
	//	_ = goCmd.Parse(args[1:])
	//case "php":
	//	phpCmd := flag.NewFlagSet("php",flag.ExitOnError)
	//	phpCmd.StringVar(&name, "n","PHP语言","help info")
	//	_ = phpCmd.Parse(args[1:])
	//}

	// 示例1.标准库flag的基本使用和长短选项
	//var name string
	//flag.StringVar(&name, "name", "go语言编程", "help info")
	//flag.StringVar(&name, "n","go语言编程","help info")
	//flag.Parse()
	//
	//log.Printf("name: %s", name)
}

// 示例3.自定义参数类型
// 实现Value接口
func (i *Name) String() string {
	return fmt.Sprint(*i)
}
func (i *Name) Set(s string) error {
	if len(*i) > 0 {
		return errors.New("name flag already set")
	}
	*i = Name("wdy:" + s)
	return nil
}
