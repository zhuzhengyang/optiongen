# optiongen
[![GoDoc](https://godoc.org/github.com/timestee/optiongen?status.svg)](https://godoc.org/github.com/timestee/optiongen)
[![Go Report Card](https://goreportcard.com/badge/github.com/timestee/optiongen)](https://goreportcard.com/report/github.com/timestee/optiongen)[![Sourcegraph](https://sourcegraph.com/github.com/timestee/optiongen/-/badge.svg)](https://sourcegraph.com/github.com/timestee/optiongen?badge)



optionGen is a fork of [XSAM/optionGen](https://github.com/XSAM/optionGen), a tool to generate go Struct option for test, mock or more flexible. The purpose of this fork is to provide more powerful and flexible option generation. 

## Functional Options
Functional options are an idiomatic way of creating APIs with options on types. The initial idea for this design pattern can be found in an article published by Rob Pike called [Self-referential functions and the design of options](https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html).

## Install
Install using go get, and this will build the optionGen binary in $GOPATH/bin.
```bash
go get github.com/timestee/optiongen/...
```

optionGen require [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports) to format code which is generated. So you may confirm that `goimports` has been installed

```bash
go get golang.org/x/tools/cmd/goimports
```

## Using optionGen
To generate struct option, you need write a function declaration to tell optionGen how to generate.struct name and `OptionDeclareWithDefault` suffix. In this function, just return a variable which type is `map[string]interface{}`.

The key of the map means option name, and the value of the map should consist of two parts, one for option type(except func type), and the other option default value.

### Flag
- `new_func`, Struct New function name.
- `xconf`, generate supoort for [XConf](https://github.com/sandwich-go/xconf)
- `usage_tag_name`, generage tag info for FlagSet Usage. [XConf](https://github.com/sandwich-go/xconf)
- `option_with_struct_name`, option name should contain struct name

### Annotation Support
```golang
const (
	AnnotationKeyComment        = "comment" // comment 
	AnnotationKeyPrivate        = "private" // make field private, do not gen option func
	AnnotationKeyArg            = "arg"    // as arg, do not gen option func
	AnnotationKeyXConfTag       = "xconf"  // xonf tag value
	AnnotationKeyGetter         = "getter" // return type name for GetXXXX, if you want to return interface instead ptr
	AnnotationKeyOption         = "option" // option func name
	AnnotationKeyCommentGettter = "comment_getter" // common for get func
)
// annotation@TestParamterBool(arg=1))
// annotation@SpecSub(getter="SpecVisitor", comment_getter="comment from annotation")
```
