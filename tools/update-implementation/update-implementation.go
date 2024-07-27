//go:build buildtools

package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"os/exec"
)

func main() {
	// TODO: accept these via vscode shortcut somehow
	filePath := "pkg/domain/service/counter_service.go"
	interfaceName := "CounterService"
	structName := "counterServiceImpl"

	// Parse the Go source file
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	// Find the interface and struct declarations
	var interfaceType *ast.InterfaceType
	var structType *ast.StructType
	for _, d := range node.Decls {
		if genDecl, ok := d.(*ast.GenDecl); ok {
			for _, spec := range genDecl.Specs {
				if typeSpec, ok := spec.(*ast.TypeSpec); ok {
					switch t := typeSpec.Type.(type) {
					case *ast.InterfaceType:
						if typeSpec.Name.Name == interfaceName {
							interfaceType = t
						}
					case *ast.StructType:
						if typeSpec.Name.Name == structName {
							structType = t
						}
					}
				}
			}
		}
	}

	if interfaceType == nil || structType == nil {
		panic("Interface or Struct not found")
	}

	// Check and modify methods
	modifyStructMethods(interfaceType, structName, node)

	// Write the modified AST back to the file
	var buf bytes.Buffer
	err = printer.Fprint(&buf, fset, node)
	if err != nil {
		panic(err)
	}

	formattedSrc, err := format.Source(buf.Bytes())
	if err != nil {
		panic(err)
	}

	// Write the formatted source code back to the file
	err = os.WriteFile(filePath, formattedSrc, 0600)
	if err != nil {
		panic(err)
	}

	cmd := exec.Command("goimports", "-w", filePath)
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Error running goimports on '"+filePath+"':", err)
	}
}

func isMethodOfStruct(funcDecl *ast.FuncDecl, structName string) bool {
	if funcDecl.Recv == nil || len(funcDecl.Recv.List) == 0 {
		return false
	}

	recvType := funcDecl.Recv.List[0].Type

	// Check for pointer receiver
	if ptr, ok := recvType.(*ast.StarExpr); ok {
		if ident, ok := ptr.X.(*ast.Ident); ok {
			return ident.Name == structName
		}
	}

	// Check for non-pointer receiver
	if ident, ok := recvType.(*ast.Ident); ok {
		return ident.Name == structName
	}

	return false
}

func getStructMethods(node *ast.File, structName string) map[string]*ast.FuncType {
	methods := make(map[string]*ast.FuncType)
	for _, d := range node.Decls {
		if funcDecl, ok := d.(*ast.FuncDecl); ok {
			if isMethodOfStruct(funcDecl, structName) {
				methods[funcDecl.Name.Name] = funcDecl.Type
			}
		}
	}
	return methods
}

func modifyStructMethods(interfaceType *ast.InterfaceType, structName string, node *ast.File) {
	interfaceMethods := make(map[string]*ast.Field)
	for _, m := range interfaceType.Methods.List {
		if len(m.Names) > 0 {
			interfaceMethods[m.Names[0].Name] = m
		}
	}

	structMethods := getStructMethods(node, structName)

	for methodName, interfaceMethod := range interfaceMethods {
		if structMethod, ok := structMethods[methodName]; ok {
			if !isSignatureSame(interfaceMethod.Type.(*ast.FuncType), structMethod) {
				replaceMethod(node, structName, methodName, interfaceMethod.Type.(*ast.FuncType))
			}
		} else {
			addMethodWithSignature(node, structName, methodName, interfaceMethod.Type.(*ast.FuncType))
		}
	}
}

func replaceMethod(node *ast.File, structName, methodName string, newMethodType *ast.FuncType) {
	for i, d := range node.Decls {
		if funcDecl, ok := d.(*ast.FuncDecl); ok {
			if funcDecl.Name.Name == methodName && isMethodOfStruct(funcDecl, structName) {
				// Create a new function declaration
				newFuncDecl := createFuncDecl(structName, methodName, newMethodType)

				// Replace the declaration at the specific index
				node.Decls[i] = newFuncDecl
				return
			}
		}
	}

	// If the method was not found, add it as a new declaration
	addMethodWithSignature(node, structName, methodName, newMethodType)
}

func createFuncDecl(structName, methodName string, methodType *ast.FuncType) *ast.FuncDecl {
	return &ast.FuncDecl{
		Name: ast.NewIdent(methodName),
		Type: methodType,
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				&ast.AssignStmt{
					Lhs: []ast.Expr{
						ast.NewIdent("err"),
					},
					Tok: token.DEFINE,
					Rhs: []ast.Expr{
						&ast.BasicLit{Kind: token.STRING, Value: `"Not implemented"`},
					},
				},
				&ast.ExprStmt{
					X: &ast.CallExpr{
						Fun: &ast.SelectorExpr{
							X:   ast.NewIdent("fmt"),
							Sel: ast.NewIdent("Println"),
						},
						Args: []ast.Expr{ast.NewIdent("err")},
					},
				},
				&ast.ExprStmt{
					X: &ast.CallExpr{
						Fun:  ast.NewIdent("panic"),
						Args: []ast.Expr{ast.NewIdent("err")},
					},
				},
			},
		},
		Recv: &ast.FieldList{
			List: []*ast.Field{
				{
					Type: &ast.StarExpr{
						X: ast.NewIdent(structName),
					},
				},
			},
		},
	}
}

func addMethodWithSignature(node *ast.File, structName, methodName string, methodType *ast.FuncType) {
	funcDecl := createFuncDecl(structName, methodName, methodType)
	node.Decls = append(node.Decls, funcDecl)
}

func isSignatureSame(interfaceMethod, structMethod *ast.FuncType) bool {
	// Compare the number of parameters
	if len(interfaceMethod.Params.List) != len(structMethod.Params.List) {
		return false
	}

	// Compare parameter types
	for i, param := range interfaceMethod.Params.List {
		structParam := structMethod.Params.List[i]
		if getTypeString(param.Type) != getTypeString(structParam.Type) {
			return false
		}
	}

	// Compare the number of return values
	if (interfaceMethod.Results == nil) != (structMethod.Results == nil) {
		return false
	}

	if interfaceMethod.Results != nil && structMethod.Results != nil {
		if len(interfaceMethod.Results.List) != len(structMethod.Results.List) {
			return false
		}

		// Compare return types
		for i, result := range interfaceMethod.Results.List {
			structResult := structMethod.Results.List[i]
			if getTypeString(result.Type) != getTypeString(structResult.Type) {
				return false
			}
		}
	}

	return true
}

func getTypeString(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.SelectorExpr:
		return getTypeString(t.X) + "." + t.Sel.Name
	case *ast.StarExpr:
		return "*" + getTypeString(t.X)
	case *ast.ArrayType:
		return "[]" + getTypeString(t.Elt)
	// Add other types as necessary, e.g., maps, channels, etc.
	default:
		return "" // or handle more complex types as needed
	}
}
