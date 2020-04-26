package optionGen

import (
	"bytes"
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

	dirFiles, err := gogenerate.FilesContainingCmd(wd, OptionGen,tags)
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

func ParseDir(dir string,optionWithStructName bool) {
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
					if strings.HasSuffix(d.Name.Name, optionDeclarationSuffix)  {
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
										optionFields[i].FieldType = FieldTypeVar
										buf := bytes.NewBufferString("")

										// Option Variable Type
										_ = printer.Fprint(buf, fset, value.Fun)
										optionFields[i].Type = buf.String()

										// Option Variable Value
										buf.Reset()
										_ = printer.Fprint(buf, fset, value.Args[0])
										optionFields[i].Body = buf.String()
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

			for className  := range classOptionFields {
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
