package generate

import (
	"fmt"
	"go/ast"
	"strings"

	"github.com/GongT/go-generate-struct-interface/internal/tools"
)

type define struct {
	Getter string
	Setter string
	Type   string
}

type Generater struct {
	contents string
	structs  map[string][]define
	comments []string
	imports  map[string]string
}

func NewGenerater(contents string) *Generater {
	return &Generater{
		contents: contents,
		structs:  make(map[string][]define),
		imports:  make(map[string]string),
	}
}

func (g *Generater) WriteComment(comment string) {
	g.comments = append(g.comments, comment)
}

func (g *Generater) WriteImport(name *ast.Ident, path *ast.BasicLit) {
	nstr := ""
	if name != nil {
		nstr = name.String()
	}
	g.imports[path.Value] = nstr
}

func (g *Generater) Print(packageName string) []byte {
	genContent := ""

	for _, comment := range g.comments {
		genContent += comment + "\n"
	}
	genContent += "\n"

	genContent += "package " + packageName + "\n\n"

	if len(g.imports) > 0 {
		genContent += "import (\n"
		for path, name := range g.imports {
			genContent += "\t"
			if name != "" {
				genContent += name + " "
			}

			genContent += path + "\n"
		}
		genContent += ")\n"
	}

	for structName, defs := range g.structs {
		upStructName := tools.Ucfirst(structName)
		genContent += fmt.Sprintf("type ReadOnly%s interface {\n", upStructName)
		for _, def := range defs {
			genContent += "\t" + def.Type + "\n"
		}
		genContent += "}\n\n"

		for _, def := range defs {
			if len(def.Getter) > 0 {
				genContent += def.Getter + "\n\n"
			}
			if len(def.Setter) > 0 {
				genContent += def.Setter + "\n\n"
			}
		}
	}

	return []byte(genContent)
}

func (g *Generater) AddField(structName string, varName string, expr ast.Expr) {
	if _, exists := g.structs[structName]; !exists {
		g.structs[structName] = make([]define, 0)
	}

	upVarName := tools.Ucfirst(varName)

	getter, setter := true, false
	typeStr := strings.TrimSpace(g.stringify(expr))
	if ch, ok := expr.(*ast.ChanType); ok {
		setter = false
		if !ch.Arrow.IsValid() {
			typeStr = "<- " + typeStr
		}
	} else if _, ok := expr.(*ast.FuncType); ok {
		return
	} else {
	}

	Type := ""

	Getter := ""
	if getter {
		Getter = fmt.Sprintf("func (self %s) Get%s() %s {\n\treturn self.%s\n}", structName, upVarName, typeStr, varName)
		Type += fmt.Sprintf("Get%s() %s", upVarName, typeStr)
	}

	Setter := ""
	if setter {
		Setter = fmt.Sprintf("func (self *%s) Set%s(v %s) {\n\tself.%s = v\n}", structName, upVarName, typeStr, varName)

		if getter {
			Type += "\n"
		}
		Type += fmt.Sprintf("Set%s(v %s)", upVarName, typeStr)
	}

	g.structs[structName] = append(g.structs[structName], define{
		Type:   Type,
		Getter: Getter,
		Setter: Setter,
	})
}

func (g *Generater) stringify(node ast.Node) string {
	return g.contents[node.Pos()-1 : node.End()-1]
}
