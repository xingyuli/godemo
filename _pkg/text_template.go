package main

import (
	"os"
	"text/template"
)

type Inventory struct {
	Material string
	Count    uint
}

func trivial() {
	sweaters := Inventory{"wool", 17}
	tmpl, err := template.New("test").Parse("{{.Count}} items are made of {{.Material}}")
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil {
		panic(err)
	}
}

var (
	_if_contexts = []interface{}{
		0, 1,
		false, true,
		[]int{}, []int{1, 2, 3},
		map[string]int{}, map[string]int{"a": 1, "b": 2},
		nil,
		"", " ",
	}

	_range_contexts = []interface{}{
		[...]int{}, [...]int{1, 2, 3},
		[]string{}, []string{"a", "b", "c"},
		map[string]int{}, map[string]int{"bb": 22, "cc": 33, "aa": 11},
	}
)

var (
	commentTemplate = constructTmpl("comment", "{{/* a comment */}}This will be displayed.")

	/*
	 * {{if pipeline}} T1 {{end}} 
	 *
	 * Dot is unaffected.
	 */
	ifTemplate = constructTmpl("if", "{{if .}}Not empty: {{.}}\n{{end}}")

	/*
	 * {{if pipeline}} T1 {{else}} T0 {{end}} 
	 *
	 * Dot is unaffected.
	 */
	ifElseTemplate = constructTmpl("if-else", "{{if .}}Not empty: {{.}}{{else}}Empty: {{.}}{{end}}\n")

	/*
	 * {{range pipeline}} T1 {{end}}
	 *
	 * The evaluated value must be an array, slice, map or channel.
	 * If the value of the pipeline has length zero, nothing is output;
	 * otherwise, dot is set to the successive elements of the array,
	 * slice, or map and T1 executed. If the value is a map and the keys
	 * are a basic type with a defined order ("comparable"), the elements
	 * will be visited in sorted key order.
	 */
	rangeTemplate = constructTmpl("range", "{{range .}}{{.}}\n{{end}}")

	/*
	 * {{range pipeline}} T1 {{else}} T0 {{end}}
	 *
	 * The evaluated value must be an array, slice, map or channel.
	 * If the value of the pipeline has length zero, dot is unaffected and
	 * T0 is executed; otherwise, dot is set to the successive elements
	 * of the array, slice, or map and T1 is executed.
	 */
	rangeElseTemplate = constructTmpl("range-else", "{{range .}}Element:{{.}} {{else}}Empty range: {{.}}{{end}}\n")
)

/*
 * {{template "name"}}
 *
 * The template with the specified name is executed with nil data.
 */
const _template = `
{{define "ONE"}}One content.{{end}}
{{define "TWO"}}Two content.{{end}}
{{template "ONE"}}
{{template "TWO"}}
`

/*
 * {{template "name" pipeline}}
 *
 * The template with the specified name is executed with dot set
 * to the value of the pipeline.
 */
const _template_pipeline = `
{{define "HEADER"}}
	{{if .}} <h1>{{.}}</h1> {{else}} No title is given. {{end}}
{{end}}

{{define "FOOTER"}}
    {{if .}} <h6>{{.}}</h6> {{else}} No footer is given. {{end}}
{{end}}

{{template "HEADER" .Title}}
{{template "FOOTER" .Footer}}
`

/*
 * {{with pipeline}} T1 {{end}}
 *
 * If the value of the pipeline is empty, no ouput is generated;
 * otherwise, dot is set to the value of the pipeline and T1 is
 * executed.
 *
 * {{with pipeline}} T1 {{end}} T0 {{end}}
 *
 * If the value of the pipeline is empty, dot is unaffected and T0
 * is executed; otherwise, dot is set to the value of pipeline
 * and T1 is executed.
 */
const _with_template = `
{{with .User}}
username: {{.Username}}
age: {{.Age}}
{{end}}

{{with .Hist}}
last visit: {{.LastVisit}}
total visit: {{.TotalVisit}}
{{end}}
`

/*
 * Nested dots are supported as well! :)
 */
const _nested_template = "{{.User.Username}}"

type User struct {
	Username string
	Age      int
}

type Hist struct {
	LastVisit  string
	TotalVisit int

	// totalVisit int
}

// text/template cannot invoke the method
/*func (h *Hist) TotalVisit() int {
	return h.totalVisit
}*/

func constructTmpl(name, text string) *template.Template {
	return template.Must(template.New(name).Parse(text))
}

func render(tmpl *template.Template, contexts ...interface{}) {
	for _, context := range contexts {
		tmpl.Execute(os.Stdout, context)
	}
}

func main() {
	// trivial()

	// render(commentTemplate, nil)
	// render(ifTemplate, _if_contexts...)
	// render(ifElseTemplate, _if_contexts...)
	// render(rangeTemplate, _range_contexts...)
	// render(rangeElseTemplate, _range_contexts...)

	// render(constructTmpl("_template", _template), nil)

	/*render(constructTmpl("_template_pipeline", _template_pipeline),
		map[string]string{
			"Title":  "Welcome to the world of Go!",
			"Footer": "Copyright.",
		},
		map[string]string{},
	)*/

	/*render(constructTmpl("_with_template", _with_template),
		map[string]interface{}{
			"User": User{"Liu Xingyu", 24},
			"Hist": Hist{"June 6, 2013", 3},
		},
	)*/

	render(constructTmpl("_nested_template", _nested_template),
		map[string]interface{}{
			"User": User{Username: "Liu Xingyu"},
		},
	)
}
