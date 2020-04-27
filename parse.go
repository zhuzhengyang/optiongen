package optionGen

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"strings"

	"runtime"

	"myitcv.io/gogenerate"
)

var fset = token.NewFileSet()

func inspectDir(wd string) {
	envFile, ok := os.LookupEnv(gogenerate.GOFILE)
	if !ok {
		log.Fatalf("env not correct; missing %v", gogenerate.GOFILE)
	}

	tags := make(map[string]bool)

	goos := os.Getenv("GOOS")
	if goos == "" {
		goos = runtime.GOOS
	}
	tags[goos] = true

	goarch := os.Getenv("GOARCH")
	if goarch == "" {
		goarch = runtime.GOARCH
	}
	tags[goarch] = true

	dirFiles, err := gogenerate.FilesContainingCmd(wd, OptionGen, tags)
	if err != nil {
		log.Fatalf("could not determine if we are the first file: %v", err)
	}

	if dirFiles == nil {
		log.Fatalf("cannot find any files containing the %v directive", OptionGen)
	}
	if dirFiles[envFile] != 1 {
		log.Fatalf("expected a single occurrence of %v directive in %v. Got: %v", OptionGen, envFile, dirFiles)
	}
}

var gofmtBuf bytes.Buffer

// gofmt returns the gofmt-formatted string for an AST node.
func gofmt(n interface{}) string {
	gofmtBuf.Reset()
	err := printer.Fprint(&gofmtBuf, fset, n)
	if err != nil {
		return "<" + err.Error() + ">"
	}
	return gofmtBuf.String()
}

func printLog(format string, a ...interface{}) {
	if EnableDebug {
		fmt.Println(fmt.Sprintf(format, a...))
	}
}

var EnableDebug bool

func ParseDir(dir string, optionWithStructName bool) {
	inspectDir(dir)

	pkgs, err := parser.ParseDir(fset, dir, nil, parser.ParseComments)
	if err != nil {
		log.Fatalf("unable to parse %v: %v", dir, err)
	}

	for _, pkg := range pkgs {
		for filePath, file := range pkg.Files {
			if gogenerate.FileGeneratedBy(filePath, OptionGen) {
				continue
			}

			var importPath []string
			for _, imp := range file.Imports {
				importPath = append(importPath, imp.Path.Value)
			}

			classList := make(map[string]bool)
			classOptionFields := make(map[string][]optionField)
			for _, d := range file.Decls {
				switch d := d.(type) {
				case *ast.FuncDecl:
					if d.Recv != nil {
						continue
					}
					if strings.HasSuffix(d.Name.Name, optionDeclarationSuffix) {
						// Only allow return expr in class option declaration function
						if len(d.Body.List) != 1 {
							continue
						}

						stmt := d.Body.List[0]
						// Only allow return one value
						if stmt, ok := stmt.(*ast.ReturnStmt); !ok {
							continue
						} else {
							if len(stmt.Results) != 1 {
								continue
							}
							result := stmt.Results[0].(*ast.CompositeLit)
							optionFields := make([]optionField, len(result.Elts))
							for i, elt := range result.Elts {
								switch elt := elt.(type) {
								case *ast.KeyValueExpr:
									// Option Field Name
									key := elt.Key.(*ast.BasicLit)
									optionFields[i].Name = key.Value

									switch value := elt.Value.(type) {
									case *ast.FuncLit:
										printLog("%s type:%s", optionFields[i].Name, "ast.FuncLit")
										optionFields[i].FieldType = FieldTypeFunc
										buf := bytes.NewBufferString("")
										// Option func Type
										_ = printer.Fprint(buf, fset, value.Type)
										optionFields[i].Type = buf.String()

										// Option func Body
										buf.Reset()
										_ = printer.Fprint(buf, fset, value.Body)
										optionFields[i].Body = buf.String()
									case *ast.CallExpr:
										printLog("%s type:%s", optionFields[i].Name, "ast.CallExpr")
										optionFields[i].FieldType = FieldTypeVar
										buf := bytes.NewBufferString("")

										// Option Variable Type
										_ = printer.Fprint(buf, fset, value.Fun)
										optionFields[i].Type = buf.String()
										// Option Variable Value
										buf.Reset()
										_ = printer.Fprint(buf, fset, value.Args[0])
										optionFields[i].Body = buf.String()
									case *ast.BasicLit: // token.INT, token.FLOAT, token.IMAG, token.CHAR, or token.STRING
										printLog("%s type:%s", optionFields[i].Name, "ast.BasicLit")
										optionFields[i].FieldType = FieldTypeVar
										switch value.Kind {
										case token.INT:
											optionFields[i].Type = "int"
											optionFields[i].Body = value.Value
										case token.FLOAT:
											optionFields[i].Type = "float"
											optionFields[i].Body = value.Value
										case token.CHAR:
											optionFields[i].Type = "byte"
											optionFields[i].Body = value.Value
										case token.STRING:
											optionFields[i].Type = "string"
											optionFields[i].Body = value.Value
										}
									case *ast.CompositeLit:
										printLog("%s type:%s", optionFields[i].Name, "ast.CompositeLit")
										optionFields[i].FieldType = FieldTypeVar
										optionFields[i].Type = gofmt(value.Type)
										buf := bytes.NewBufferString("")
										buf.Reset()
										var data []string
										for _, p := range value.Elts {
											switch t := p.(type) {
											case *ast.BasicLit:
												data = append(data, t.Value)
											case *ast.Ident:
												if t.Name == "false" || t.Name == "true" {
													data = append(data, t.Name)
												}
											case *ast.KeyValueExpr:
												blKey, okKey := t.Key.(*ast.BasicLit)
												blVal, okVal := t.Value.(*ast.BasicLit)
												if !okKey || !okVal {
													log.Fatalf("optionGen %s got type %s support basic types only", optionFields[i].Name, optionFields[i].Type)
												}
												data = append(data, fmt.Sprintf("%s:%s", blKey.Value, blVal.Value))
											default:
												log.Fatalf("optionGen %s got type %s support basic types only", optionFields[i].Name, optionFields[i].Type)
											}
										}
										val := "nil"
										if len(data) > 0 {
											val = fmt.Sprintf("%s{%s}", optionFields[i].Type, strings.Join(data, ","))
										}
										optionFields[i].Body = val
									case *ast.Ident:
										if value.Obj == nil {
											printLog("%s type:%s", optionFields[i].Name, "ast.Ident")
										} else {
											printLog("%s type:%s Object:%v Type:%v", optionFields[i].Name, "ast.Ident", value.Obj, value.Obj.Type)
										}
										optionFields[i].FieldType = FieldTypeVar
										if value.Name == "false" || value.Name == "true" {
											optionFields[i].Type = "bool"
											optionFields[i].Body = value.Name
											break
										}
										if value.Name == "nil" {
											optionFields[i].Type = "interface{}"
											optionFields[i].Body = value.Name
											break
										}
										log.Fatalf("optionGen %s got type ast.Ident, please give some type hint like: string(DefaultValue) ", optionFields[i].Name)
									default:
										log.Fatalf("optionGen %s got type %s support basic types only", optionFields[i].Name, optionFields[i].Type)
									}
								}
							}

							declarationClassName := strings.TrimPrefix(strings.TrimSuffix(d.Name.Name, optionDeclarationSuffix), "_")
							classOptionFields[declarationClassName] = optionFields
						}
					}
				case *ast.GenDecl:
					if d.Tok == token.TYPE {
						for _, spec := range d.Specs {
							if typeSpec, ok := spec.(*ast.TypeSpec); ok {
								classList[typeSpec.Name.Name] = false
							}
						}
					}
				}
			}

			for className := range classOptionFields {
				classList[className] = true
			}

			for className, optionExist := range classList {
				if !optionExist {
					delete(classList, className)
				}
			}

			g := fileOptionGen{
				FilePath:          filePath,
				FileName:          strings.TrimSuffix(filepath.Base(filePath), ".go"),
				PkgName:           pkg.Name,
				ImportPath:        importPath,
				ClassList:         classList,
				ClassOptionFields: classOptionFields,
			}
			g.gen(optionWithStructName)
		}
	}
}
