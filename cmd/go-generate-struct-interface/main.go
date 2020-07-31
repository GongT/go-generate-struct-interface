package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/GongT/go-generate-struct-interface/internal/generate"
	"github.com/GongT/go-generate-struct-interface/internal/tools"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	spew.Config.Indent = "    "

	filePath := os.Getenv("GOFILE")
	pkgName := os.Getenv("GOPACKAGE")

	if len(filePath) == 0 || len(pkgName) == 0 {
		tools.Die("Required environment variable did not set, must call by `go generate`")
	}
	// fmt.Printf("GOFILE=%s\n", filePath)
	// fmt.Printf("GOPACKAGE=%s\n", pkgName)

	fileNameBase := filepath.Base(filePath)
	resultFile := filepath.Join(filepath.Dir(filePath), "generate."+fileNameBase)

	// fmt.Printf(" * resultFile: %s\n", resultFile)

	contentBs, err := ioutil.ReadFile(filePath)
	if err != nil {
		tools.Die("Failed read input file: %s", err.Error())
	}
	content := string(contentBs)
	file, err := parser.ParseFile(token.NewFileSet(), filePath, content, 0)
	if err != nil {
		tools.Die("Failed parse input file: %s", err.Error())
	}

	gen := generate.NewGenerater(content)

	comments, err := parser.ParseFile(token.NewFileSet(), filePath, content, parser.PackageClauseOnly+parser.ParseComments)
	for _, group := range comments.Comments {
		for _, comment := range group.List {
			if !strings.HasPrefix(comment.Text, "//go:generate ") {
				gen.WriteComment(comment.Text)
			}
		}
	}

	for _, node := range file.Decls {
		if decl, ok := node.(*ast.GenDecl); ok {
			if decl.Tok != token.TYPE {
				continue
			}
			typeSpec, ok := decl.Specs[0].(*ast.TypeSpec)
			if !ok {
				continue
			}
			structName := typeSpec.Name.String()
			structSpec, ok := typeSpec.Type.(*ast.StructType)
			if !ok {
				continue
			}

			for _, field := range structSpec.Fields.List {
				if len(field.Names) == 0 {
					continue
				}

				for _, varNameToken := range field.Names {
					varName := varNameToken.String()
					gen.AddField(structName, varName, field.Type)
				}
			}
		}
	}

	err = ioutil.WriteFile(resultFile, gen.Print(file.Name.String()), os.FileMode(0666))

	if err != nil {
		tools.Die("Failed write output file: %s", err.Error())
	}
}
