package generate

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

func Run() {
	root := "." // Starting directory, you can change it as needed
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {
			processFile(path)
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking the path %q: %v\n", root, err)
	}
}

func processFile(path string) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		fmt.Printf("Failed to parse file %s: %v\n", path, err)
		return
	}

	for _, decl := range node.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}
		for _, spec := range genDecl.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}
			if typeSpec.Comment != nil && strings.Contains(typeSpec.Comment.Text(), "EZENUM") {
				typeName := typeSpec.Name.Name
				constants := findConstants(node, typeName)
				generateEnumFile(path, typeName, constants)
			}
		}
	}
}

func findConstants(node *ast.File, typeName string) map[string]string {
	constants := make(map[string]string)
	for _, decl := range node.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}
		for _, spec := range genDecl.Specs {
			valueSpec, ok := spec.(*ast.ValueSpec)
			if !ok {
				continue
			}
			if valueSpec.Type != nil {
				if ident, ok := valueSpec.Type.(*ast.Ident); ok && ident.Name == typeName {
					for i, name := range valueSpec.Names {
						comment := ""
						if i < len(valueSpec.Comment.List) {
							comment = valueSpec.Comment.List[i].Text
						}
						constants[name.Name] = comment
					}
				}
			} else {
				// Check for constants without explicit type
				if len(valueSpec.Values) > 0 {
					if ident, ok := valueSpec.Values[0].(*ast.BasicLit); ok && ident.Kind == token.INT {
						for i, name := range valueSpec.Names {
							comment := ""
							if i < len(valueSpec.Comment.List) {
								comment = valueSpec.Comment.List[i].Text
							}
							constants[name.Name] = comment
						}
					}
				}
			}
		}
	}
	return constants
}
func generateEnumFile(filePath, typeName string, constants map[string]string) {
	dir := filepath.Dir(filePath)
	baseName := filepath.Base(filePath)
	newFileName := fmt.Sprintf("%s_ezenum_gen.go", strings.TrimSuffix(baseName, ".go"))
	newFilePath := filepath.Join(dir, newFileName)

	packageName := getPackageName(filePath)

	code := fmt.Sprintf("package %s\n\n", packageName)
	code += fmt.Sprintf("func (r %s) AsCode() int {\n", typeName)
	code += "\treturn int(r)\n"
	code += "}\n\n"
	code += fmt.Sprintf("func (r %s) GetDescription() string {\n", typeName)
	code += "\tswitch r {\n"
	for constant, comment := range constants {
		comment = strings.TrimPrefix(comment, "//")
		comment = strings.TrimSpace(comment)
		code += fmt.Sprintf("\tcase %s:\n", constant)
		code += fmt.Sprintf("\t\treturn \"%s\"\n", comment)
	}
	code += "\tdefault:\n"
	code += "\t\treturn \"Unknown Response\"\n"
	code += "\t}\n"
	code += "}\n\n"
	code += fmt.Sprintf("func IsValid%s(code int) bool {\n", typeName)
	code += fmt.Sprintf("\tswitch %s(code) {\n", typeName)
	for constant := range constants {
		code += fmt.Sprintf("\tcase %s:\n", constant)
	}
	code += "\t\treturn true\n"
	code += "\tdefault:\n"
	code += "\t\treturn false\n"
	code += "\t}\n"
	code += "}\n"

	err := os.WriteFile(newFilePath, []byte(code), 0644)
	if err != nil {
		fmt.Printf("Failed to write file %s: %v\n", newFilePath, err)
	} else {
		fmt.Printf("Generated file %s\n", newFilePath)
	}
}

func getPackageName(filePath string) string {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filePath, nil, parser.PackageClauseOnly)
	if err != nil {
		fmt.Printf("Failed to parse package name from file %s: %v\n", filePath, err)
		return ""
	}

	if node.Name != nil {
		return node.Name.Name
	}
	return ""
}
