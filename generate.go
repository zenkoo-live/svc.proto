/*
 * Copyright (C) LiangYu, Inc - All Rights Reserved
 * Unauthorized copying of this file, via any medium is strictly prohibited
 * Proprietary and confidential
 */

/**
 * @file generate.go
 * @package main
 * @author Dr.NP <conan.np@gmail.com>
 * @since 05/10/2024
 */

package main

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"os"
	"strings"
	"text/template"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const (
	pkgPrefix = "github.com/zenkoo-live/svc.proto"
)

type svcSpec struct {
	ConstName string
	SvcName   string
	PkgName   string
	PkgAlias  string
	PkgPath   string
	Services  map[string]string
}

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	svcs := getServices()
	b := genString(svcs)
	err = os.WriteFile(pwd+"/zenkoo/services.go", b, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func genString(data map[string]*svcSpec) []byte {
	const tmpl = `
// - GENERATED BY generate.go
// - DO NOT MODIFY THIS FILE

package zenkoo

import (
	cltGrpc "github.com/go-micro/plugins/v4/client/grpc"
	{{range $item  :=  .}}
	{{.PkgAlias}} "{{.PkgPath}}"{{end}}
)

const (
	AppName = "zenkoo"
	{{range $item := .}}
	{{.ConstName}} = "{{.SvcName}}"{{end}}
)

var (
	clt = cltGrpc.NewClient()
	{{range $item := .}}{{range $k, $v := .Services}}
	{{$k}} = {{$item.PkgAlias}}.{{$v}}(AppName+"::"+{{$item.ConstName}}, clt){{end}}{{end}}
)
	`

	t, err := template.New("").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}

	buff := bytes.NewBufferString("")
	err = t.Execute(buff, data)
	if err != nil {
		log.Fatal(err)
	}

	src, err := format.Source(buff.Bytes())
	if err != nil {
		log.Fatal(err)
	}

	return src
}

func getServices() map[string]*svcSpec {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	entities, err := os.ReadDir(pwd)
	if err != nil {
		log.Fatal(err)
	}

	var (
		currSpec *svcSpec
		fset     = token.NewFileSet()
		svcs     = make(map[string]*svcSpec)
	)
	for _, item := range entities {
		if item.IsDir() {
			// Check micro.pb
			srcs, err := os.ReadDir(pwd + "/" + item.Name())
			if err != nil {
				log.Fatal(err)
			}
			currSpec = nil
			svcName := ""
			for _, f := range srcs {
				if strings.HasSuffix(f.Name(), ".micro.go") {
					path := pwd + "/" + item.Name() + "/" + f.Name()
					// Get service name
					af, err := parser.ParseFile(fset, path, nil, 0)
					if err != nil {
						log.Fatal(err)
					}
					// 先找到包名
					ast.Inspect(af, func(n ast.Node) bool {
						switch t := n.(type) {
						case *ast.File:
							if currSpec == nil {
								svcName = t.Name.Name
								currSpec = &svcSpec{
									ConstName: stripDot(item.Name(), true),
									SvcName:   strings.ToLower(item.Name()),
									PkgName:   t.Name.Name,
									PkgAlias:  stripDot(item.Name(), false),
									PkgPath:   pkgPrefix + "/" + item.Name(),
									Services:  make(map[string]string),
								}
							}
						}
						return true
					})
					// 再找到服务名
					ast.Inspect(af, func(n ast.Node) bool {
						switch t := n.(type) {
						case *ast.FuncDecl:
							if strings.HasPrefix(t.Name.Name, "New") && strings.HasSuffix(t.Name.Name, "Service") {
								subSvcName := strings.TrimSuffix(f.Name(), ".pb.micro.go")
								if currSpec != nil {
									currSpec.Services[currSpec.ConstName+
										cases.Title(language.English, cases.Compact).String(strings.ToLower(subSvcName))] = t.Name.Name
								}
							}
						}
						return true
					})
				}
			}
			if svcName != "" {
				if _, ok := svcs[svcName]; !ok {
					svcs[svcName] = currSpec
				}
			}
		}
	}

	return svcs
}

func stripDot(in string, upperHead bool) string {
	parts := strings.Split(in, ".")
	out := ""
	for i, w := range parts {
		if !upperHead && i == 0 {
			out = out + strings.ToLower(w)
		} else {
			out = out + cases.Title(language.English, cases.Compact).String(strings.ToLower(w))
		}
	}

	return out
}

/*
 * Local variables:
 * tab-width: 4
 * c-basic-offset: 4
 * End:
 * vim600: sw=4 ts=4 fdm=marker
 * vim<600: sw=4 ts=4
 */
