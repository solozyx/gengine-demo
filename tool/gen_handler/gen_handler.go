package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"html/template"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	p := Parser{
		EndpointsByName: make(map[string][]*Endpoint),
	}
	p.Parse()

	//strBytes, _ := json.Marshal(p)
	//fmt.Println(string(strBytes))
}

type Parser struct {
	EndpointsByName map[string][]*Endpoint
}

type Endpoint struct {
	StructName string
	Name       string
}

func NewEndpoint(class, name string) *Endpoint {
	return &Endpoint{
		StructName: class,
		Name:       name,
	}
}

func (p *Parser) Parse() error {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, os.Getenv("GOFILE"), nil, 0)
	if err != nil {
		fmt.Printf("parse file fail. error: %v", err.Error())
		return err
	}

	ast.Inspect(f, func(node ast.Node) bool {
		switch x := node.(type) {
		case *ast.TypeSpec:
			p.parseType(x)
		}

		return true
	})

	p.generateCode()
	return nil
}

func (p *Parser) parseType(st *ast.TypeSpec) error {
	if t, ok := st.Type.(*ast.InterfaceType); ok {
		//fmt.Println(st.Name.Name)
		sn := strings.Replace(st.Name.Name, "Handler", "", -1)
		sn = sn[1:]
		if _, ok := p.EndpointsByName[st.Name.Name]; !ok {
			p.EndpointsByName[sn] = make([]*Endpoint, 0)
		}
		for _, item := range t.Methods.List {
			for _, methodName := range item.Names {
				cn := strings.ToLower(sn)
				p.EndpointsByName[sn] = append(p.EndpointsByName[sn], NewEndpoint(cn, methodName.Name))
			}
		}
	}
	return nil
}

func (p *Parser) generateCode() {
	for k, v := range p.EndpointsByName {
		b := bytes.NewBuffer(make([]byte, 0))

		//path, err := filepath.Abs("../../tool/gen_handler/handler.go.tmpl")
		//if err != nil {
		//	panic(err)
		//}
		templateStr, err := ioutil.ReadFile("../../tool/gen_handler/handler.go.tmpl")
		if err != nil {
			panic(err)
		}

		t := template.Must(template.New("handler.go.tmpl").Parse(string(templateStr)))
		if err := t.Execute(b, map[string]interface{}{"Key": k, "Value": v}); err != nil {
			panic(err)
		}

		//fmt.Println(string(b.Bytes()))

		if err := ioutil.WriteFile(fmt.Sprintf("%s.go", strings.ToLower(k)), b.Bytes(), 0600); err != nil {
			panic(err)
		}
	}
}
