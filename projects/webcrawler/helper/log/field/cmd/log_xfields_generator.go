package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"webcrawler/helper/log/base"
)

// field_decl_template 代表日志字段类型的声明的内容模板。
const field_decl_template = `
// 拥有{{if eq . "object"}}interface{}{{else}}{{.}}{{end}}类型的值的日志字段。
type {{.}}Field struct {
	name string
	fieldType FieldType
	value {{if eq . "object"}}interface{}{{else}}{{.}}{{end}}
}

func (field *{{.}}Field) Name() string {
	return field.name
}

func (field *{{.}}Field) Type() FieldType {
	return field.fieldType
}
func (field *{{.}}Field) Value() interface{} {
	return field.value
}

func {{title .}}(name string, value {{if eq . "object"}}interface{}{{else}}{{.}}{{end}}) Field{
	return &{{.}}Field{name: name, fieldType: {{title .}}Type, value: value}
}

`

var (
	inputPath  string
	outputPath string
)

func init() {
	flag.StringVar(&inputPath, "input", "", "The path that contains the target go source files.")
	flag.StringVar(&outputPath, "output", "", "The path for output the go source file.")
}

func Usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\tlog_xfields_generator [flags] \n")
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("log_xfields_generator: ")
	flag.Usage = Usage
	flag.Parse()
	currentPath, err := os.Getwd()
	if err != nil {
		log.Fatalf("Getwd error: %s", err)
	}
	if len(inputPath) == 0 {
		inputPath = currentPath
		log.Printf("WARNING: Not specified the flag named input, use current path '%s'.",
			currentPath)
	} else {
		if !isDir(inputPath) {
			log.Fatalf("ERROR: The input path '%s' is not a directory!", inputPath)
		}
	}
	if len(outputPath) == 0 {
		outputPath = currentPath
		log.Printf("WARNING: Not specified the flag named output, use current path '%s'.",
			currentPath)
	} else {
		if !isDir(outputPath) {
			log.Fatalf("ERROR: The output path '%s' is not a directory!", outputPath)
		}
	}

	targetFilePath := filepath.Join(inputPath, "field.go")
	prefixes, err := findFieldTypePrefixes(targetFilePath)
	if err != nil {
		log.Fatalf("ERROR: Parse error: %s\n", err)
	}
	var gen Generator
	content, err := gen.generate("field", prefixes...)
	if err != nil {
		log.Fatalf("ERROR: Generate error: %s\n", err)
	}

	outputFilePath := filepath.Join(outputPath, "xfields.go")
	err = ioutil.WriteFile(outputFilePath, content, 0644)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}
	log.Printf("It has successfully generated a Go source file: %s\n", outputFilePath)
}

// isDir 用于判断指定的路径是否为目录。
func isDir(name string) bool {
	info, err := os.Stat(name)
	if err != nil {
		log.Fatal(err)
	}
	return info.IsDir()
}

// findFieldTypePrefixes 用于查找日志字段类型的名称的前缀。
func findFieldTypePrefixes(filePath string) ([]string, error) {
	astFile, err := parser.ParseFile(
		token.NewFileSet(), filePath, nil, parser.AllErrors)
	if err != nil {
		return nil, err
	}
	prefixes := []string{}
	for _, decl := range astFile.Decls {
		if genDecl, ok := decl.(*ast.GenDecl); ok && genDecl.Tok == token.CONST {
			for _, spec := range genDecl.Specs {
				if valueSpec, ok := spec.(*ast.ValueSpec); ok {
					name := valueSpec.Names[0].Name
					if valueSpec.Type.(*ast.Ident).Name == "FieldType" && name != "UnknownType" {
						prefix := name[:strings.LastIndex(name, "Type")]
						prefixes = append(prefixes, strings.ToLower(prefix))
					}
				}
			}
		}
	}
	return prefixes, nil
}

// Generator 代表源码文件生成器。
type Generator struct {
	buf bytes.Buffer
}

// reset 会重置生成器。
func (g *Generator) reset() {
	g.buf.Reset()
}

// generate 会生成源码内容。
func (g *Generator) generate(pkgName string, prefixes ...string) ([]byte, error) {
	var content []byte
	g.genHeader(pkgName)
	err := g.genFieldDecls(prefixes...)
	if err == nil {
		defer g.buf.Reset()
		content, err = g.format()
	}
	return content, err
}

// genHeader 会生成源码文件的头部。
func (g *Generator) genHeader(pkgName string) {
	g.buf.WriteString("// generated by log_xfields_generator")
	flag.VisitAll(func(fg *flag.Flag) {
		g.buf.WriteString(" -")
		g.buf.WriteString(fg.Name)
		g.buf.WriteString(" ")
		g.buf.WriteString(fg.Value.String())
	})
	g.buf.WriteString("\n// generation time: ")
	g.buf.WriteString(time.Now().Format(base.TIMESTAMP_FORMAT))
	g.buf.WriteString("\n// DO NOT EDIT!!\n")
	g.buf.WriteString("package ")
	g.buf.WriteString(pkgName)
	g.buf.WriteString("\n")
}

// genFieldDecls 会生成日志字段类型的声明内容。
func (g *Generator) genFieldDecls(prefixes ...string) error {
	funcMap := template.FuncMap{
		"title": strings.Title,
	}
	t := template.Must(template.New("xfeild").Funcs(funcMap).Parse(field_decl_template))
	for _, prefix := range prefixes {
		err := t.Execute(&g.buf, prefix)
		if err != nil {
			return err
		}
	}
	return nil
}

// format 用于格式化生成的源码内容。
func (g *Generator) format() ([]byte, error) {
	originalSrc := g.buf.Bytes()
	formatedSrc, err := format.Source(originalSrc)
	if err != nil {
		log.Printf("WARNING: Connot format the generated Go source, please build it for the detail, error: %s", err)
		return originalSrc, nil
	}
	return formatedSrc, nil
}