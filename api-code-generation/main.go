package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/docopt/docopt-go"
)

var usage = `Generate bindings for Android API.

Tool will generate go-code for simple API methods (operating on simple types).

Tool expects stdin to be in format generated by doc-parser tool:

	package {package_name}
    class {class_name} <{url}>
    method {return_type} {method_name}({method_arg}; {method_arg}; ...)
    method {return_type} {method_name}({method_arg}; {method_arg}; ...)
    ...

Usage:
  $0 [options] <output_dir>
  $0 -h|--help

Options:
  -b <base>     Use base struct, which will be embedded in any other types.
  -r <registry> Use named registry for storing type mappings.
  -h --help     Show this help.
`

var (
	reMatchPackage = regexp.MustCompile(`^package (\S+)`)
	reMatchClass   = regexp.MustCompile(`^class (\S+) <([^>]+)>`)
	reMatchMethod  = regexp.MustCompile(`^method (\S+) ([^(]+)\((.*)\)`)
	reMethodArg    = regexp.MustCompile(`(\S+)\s(\w+)`)
)

type MethodArg struct {
	Name string
	Type string
}

func main() {
	args, _ := docopt.Parse(
		strings.Replace(usage, `$0`, os.Args[0], -1),
		nil, true, "v1", false,
	)

	outputDirectory := args["<output_dir>"].(string)
	baseStruct := ""
	if args["-b"] != nil {
		baseStruct = args["-b"].(string)
	}

	registry := ""
	if args["-r"] != nil {
		registry = args["-r"].(string)
	}

	var output io.Writer
	var currentClassName string
	var prevMethods []string
	var packageName string
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		switch {
		case reMatchPackage.MatchString(line):
			packageName = reMatchPackage.FindStringSubmatch(line)[1]
		case reMatchClass.MatchString(line):
			className, url := parseClass(line)

			output = createOutputFile(outputDirectory, className)

			generateClass(output, baseStruct, className, url, packageName)

			if registry != "" {
				generateRegistryInit(output, registry, className, packageName)
			}

			currentClassName = className

			prevMethods = []string{}
		case reMatchMethod.MatchString(line):
			prevMethods = append(prevMethods,
				generateMethod(
					output,
					packageName,
					currentClassName, line,
					prevMethods,
				),
			)
		}
	}
}

func parseClass(raw string) (name, url string) {
	matches := reMatchClass.FindStringSubmatch(raw)

	return matches[1], matches[2]
}

func generateClass(
	output io.Writer, base, className, url string, packageName string,
) {
	viewTpl.Lookup("package").Execute(output, struct {
		PackageName    string
		TypeName       string
		Url            string
		Base           string
		SdkPackageName string
	}{
		"sdk",
		className,
		url,
		base,
		packageName,
	})
}

func generateRegistryInit(
	output io.Writer, registry, className, packageName string,
) {
	viewTpl.Lookup("registry").Execute(output, struct {
		TypeName       string
		Registry       string
		SdkPackageName string
	}{
		className,
		registry,
		packageName,
	})
}

func createOutputFile(directory, className string) io.Writer {
	output, err := os.Create(fmt.Sprintf("%s/%s.go", directory, className))
	if err != nil {
		panic(err)
	}

	return output
}

func generateMethod(
	output io.Writer, packageName, className, raw string, prevMethods []string,
) string {
	matches := reMatchMethod.FindStringSubmatch(raw)

	if !isSimpleType(matches[1]) {
		// no support for non-simple type now
		return ""
	}

	returnType := convertJavaTypeToGoType(matches[1])
	methodName := matches[2]
	args, argsOk := parseArgs(matches[3])
	if !argsOk {
		// no support for non-simple type now
		return ""
	}

	if methodName == "getId" {
		// overload GetId method
		return ""
	}

	displayMethodName := methodName
	// hack for overloaded methods in java
	for _, prevMethod := range prevMethods {
		if prevMethod != methodName {
			continue
		}

		types := ""
		for _, typeName := range args {
			types += typeName.Type[0:1]
		}

		displayMethodName = prevMethod + fmt.Sprintf("%d%s", len(types), types)
	}

	displayMethodName = strings.ToUpper(displayMethodName[0:1]) +
		displayMethodName[1:]

	err := viewTpl.Lookup("method").Execute(output, struct {
		SdkPackageName    string
		TypeName          string
		MethodName        string
		DisplayMethodName string
		ReturnType        string
		Args              interface{}
	}{
		packageName,
		className,
		methodName,
		displayMethodName,
		returnType,
		args,
	})

	if err != nil {
		panic(err)
	}

	return methodName
}

func parseArgs(raw string) ([]MethodArg, bool) {
	if raw == "" {
		return nil, true
	}

	rawArgs := strings.Split(raw, "; ")

	args := make([]MethodArg, len(rawArgs))

	for i, rawArg := range rawArgs {
		matches := reMethodArg.FindStringSubmatch(rawArg)
		if !isSimpleType(matches[1]) {
			return nil, false
		}

		args[i].Type = convertJavaTypeToGoType(matches[1])
		args[i].Name = matches[2]

	}

	return args, true
}

func convertJavaTypeToGoType(javaType string) string {
	switch javaType {
	case "void":
		return ""
	case "boolean":
		return "bool"
	case "Integer", "int":
		return "int"
	case "Double", "double", "float":
		return "float64"
	case "String", "CharSequence":
		return "string"
	}

	return javaType
}

func isSimpleType(javaType string) bool {
	switch javaType {
	case "void", "boolean", "Integer", "int", "Double",
		"double", "float", "String", "CharSequence":
		return true
	}

	return false
}
