package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"html/template"
	"io/ioutil"
	"strings"
)

var (
	goFile string
)

func init() {
	flag.StringVar(&goFile, "f", "", "go file path")
	flag.Parse()
}

func main() {
	p := &Parser{
		Endpoints: make([]*Endpoint, 0),
	}
	p.Parse(goFile)
	p.generateText()
}

type Parser struct {
	FileName  string
	Endpoints []*Endpoint
}

type Endpoint struct {
	Comment      string
	FunctionName string
	RequestType  string
	ResponseType string

	Request  *RModel
	Response *RModel
}

type RModel struct {
	Type    string
	Comment string
	Fields  []FieldModel
}

type FieldModel struct {
	Name    string
	Type    string
	Comment string
}

func (p *Parser) Parse(filePath string) {
	fmt.Println("gofile: ", filePath)
	if filePath == "" {
		return
	}
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	ast.Inspect(f, func(node ast.Node) bool {
		switch v := node.(type) {
		case *ast.TypeSpec:
			p.parseType(v)
			break
		}
		return true
	})
}

func (p *Parser) parseType(st *ast.TypeSpec) {
	if v, ok := st.Type.(*ast.InterfaceType); ok {
		fmt.Println("interface_name: ", st.Name.String())
		p.FileName = st.Name.String()

		if v.Methods != nil && len(v.Methods.List) > 0 {
			for _, field := range v.Methods.List {
				if ft, ok := field.Type.(*ast.FuncType); ok {
					ep := &Endpoint{}

					ep.FunctionName = field.Names[0].String()
					ep.Comment = field.Doc.Text()
					ep.Comment = strings.Replace(ep.Comment, "\n", "<br>", -1)
					fmt.Println("function_name: ", ep.FunctionName)
					fmt.Println("comment: ", ep.Comment)
					for i := range ft.Params.List {
						ep.RequestType = fmt.Sprintf("%v", ft.Params.List[i].Type)
						fmt.Println("parameter_type: ", ep.RequestType)

						r := RModel{}
						r.Type = ep.RequestType
						// 记录实体类型
						r.Comment, r.Fields = p.parseReqResModel(ep.RequestType)

						ep.Request = &r
					}
					for i := range ft.Results.List {
						if t, ok := ft.Results.List[i].Type.(*ast.StarExpr); ok {
							ep.ResponseType = fmt.Sprintf("%v", t.X)
							fmt.Println("result_type: ", ep.ResponseType)
						} else {
							rtype := fmt.Sprintf("%v", ft.Results.List[i].Type)
							if rtype != "error" {
								fmt.Println("result_type: ", rtype)
								ep.ResponseType = rtype
							}
						}

						r := RModel{}
						r.Type = ep.ResponseType
						r.Comment, r.Fields = p.parseReqResModel(ep.ResponseType)

						ep.Response = &r
					}

					p.Endpoints = append(p.Endpoints, ep)
				}
				fmt.Println("-----------------------------------------------\n")
			}
		}
	}
}

func (p *Parser) parseReqResModel(str string) (comment string, fields []FieldModel) {
	fields = make([]FieldModel, 0)

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "../model/model.go", nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	ast.Inspect(f, func(node ast.Node) bool {
		switch v := node.(type) {
		case *ast.TypeSpec:
			switch w := v.Type.(type) {
			case *ast.StructType:
				if v.Name.String() == str {
					comment = strings.Trim(strings.Replace(v.Comment.Text(), "\n", "<br>", -1), "<br>")
					if w.Fields != nil && len(w.Fields.List) > 0 {
						for _, field := range w.Fields.List {
							t := ""
							if _, ok := field.Type.(*ast.InterfaceType); ok {
								t = "Object"
							} else {
								t = fmt.Sprintf("%v", field.Type)
							}

							fields = append(fields, FieldModel{
								Name: func(obj []*ast.Ident) string {
									if obj != nil && obj[0].String() != "" {
										return obj[0].String()
									}
									return "N/A"
								}(field.Names),
								Type:    t,
								Comment: strings.Trim(strings.Replace(field.Comment.Text(), "\n", "<br>", -1), "<br>"),
							})
						}
					}
					return false
				}
				break
			}
			break
		}
		return true
	})
	return
}

func (p *Parser) generateText() {
	tmpStr, err := ioutil.ReadFile("./apidoc.md.tmpl")
	if err != nil {
		fmt.Println("not found the template file")
		return
	}

	t := template.Must(template.New("markdown").Funcs(template.FuncMap{
		"raw": func(str string) template.HTML { return template.HTML(str) },
	}).Parse(string(tmpStr)))
	b := bytes.NewBuffer([]byte(""))
	if err := t.Execute(b, p.Endpoints); err != nil {
		fmt.Println("generate text fail. error: %v", err.Error())
		return
	}
	_ = ioutil.WriteFile("./"+p.FileName+".md", b.Bytes(), 0600)
}
