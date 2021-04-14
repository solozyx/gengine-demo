package main

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"unicode"
)

func main() {
	p := Parser{
		EndpointsByName: make(map[string][]*Endpoint),
	}
	p.Parse()

	strBytes, _ := json.Marshal(p)
	fmt.Println(string(strBytes))
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
	return nil
}

func (p *Parser) parseType(st *ast.TypeSpec) error {
	if t, ok := st.Type.(*ast.InterfaceType); ok {
		fmt.Println(st.Name.Name)
		if _, ok := p.EndpointsByName[st.Name.Name]; !ok {
			p.EndpointsByName[st.Name.Name] = make([]*Endpoint, 0)
		}
		for _, item := range t.Methods.List {
			for _, methodName := range item.Names {
				cn := st.Name.Name[1:]
				arr := []byte(cn)
				arr[0] = byte(unicode.ToLower(rune(cn[0])))
				cn = string(arr)
				p.EndpointsByName[st.Name.Name] = append(p.EndpointsByName[st.Name.Name], NewEndpoint(cn, methodName.Name))
			}
		}
	}
	return nil
}
