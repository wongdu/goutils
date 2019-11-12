package main

import (
	"fmt"
	"os"
	"text/template"
)

/**
 * @time 2019/11/12 11:31
 */
func main() {
	type Inventory struct {
		Material string
		Count    uint
	}
	sweaters := Inventory{"wool", 17}
	tmpl, err := template.New("test").Parse("{{.Count}} of {{.Material}} \n")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil {
		panic(err)
	}

	//with的用法
	tmpl, err = template.New("test").Parse("{{with .Material}}hello, {{.}} {{end}} \n")
	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil {
		panic(err)
	}

	//range的用法，貌似没有像上面两个例子那样左对齐
	sI := []Inventory{
		{Material: "aaa", Count: 1},
		{Material: "bbb", Count: 2},
	}
	tmpl, err = template.New("test").Parse(`{{range .}}
    {{if gt .Count 1}}wow! many {{.Material}}
    {{else}}
    well,only have one {{.Material}}
    {{end}}
    {{end}}`)
	if err := tmpl.Execute(os.Stdout, sI); err != nil {
		fmt.Println("There was an error:", err.Error())
	}

	//range的对齐，往右边偏移一个空格
	tmpl, err = template.New("a").Parse(`{{range .}} {{if gt .Count 1}}wow! many {{.Material}}{{else}}well,only have one {{.Material}}{{end}}{{end}}`)
	if err := tmpl.Execute(os.Stdout, sI); err != nil {
		fmt.Println("There was an error:", err.Error())
	}
}
