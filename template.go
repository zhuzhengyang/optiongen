package optiongen

const templateTextWithPreviousSupport = `

{{- range $_, $comment := $.ClassComments }}
{{ $comment }}
{{- end }}

// {{ $.ClassName }} struct
type {{ $.ClassName }} struct {
	{{- range $index, $option := $.ClassOptionInfo }}
		{{- range $_, $comment := $option.LastRowComments }}
			{{ unescaped $comment }}
 		{{- end }}
		{{ $option.Name }} {{ $option.Type }} {{unescaped $option.TagString}} {{ $option.SameRowComment }} 
	{{- end }}
}

{{- if $.OptionReturnPrevious }}
// ApplyOption apply mutiple new option and return the old mutiple optuons
// sample: 
// old := cc.ApplyOption(WithTimeout(time.Second))
// defer cc.ApplyOption(old...)
func (cc *{{ $.ClassName }}) ApplyOption(opts... {{$.ClassOptionTypeName }}) []{{$.ClassOptionTypeName }}{
	var previous []{{$.ClassOptionTypeName }}
	for _, opt := range opts  {
		previous = append(previous,opt(cc))
	}
	return previous
}
{{- else}}
// ApplyOption apply mutiple new option
func (cc *{{ $.ClassName }}) ApplyOption(opts... {{$.ClassOptionTypeName }}){
	for _, opt := range opts  {
		opt(cc)
	}
}
{{- end }}

// {{ $.ClassOptionTypeName }} option func
{{- if $.OptionReturnPrevious }}
type {{ $.ClassOptionTypeName }} func(cc *{{$.ClassName}}) {{ $.ClassOptionTypeName }}
{{- else}}
type {{ $.ClassOptionTypeName }} func(cc *{{$.ClassName}})
{{- end }}

{{ range $index, $option := $.ClassOptionInfo }}
{{- if eq $option.GenOptionFunc true }}
	{{- range $methodCommentIndex, $methodComment := $option.MethodComments }}
		{{ unescaped $methodComment }}
	{{- end }}
	// {{$option.OptionFuncName}} option func for {{ $option.Name }}

	{{- if $.OptionReturnPrevious }}
	{{- if eq $option.Slice true }}
		func {{$option.OptionFuncName}}(v ...{{$option.SliceElemType}}) {{ $.ClassOptionTypeName }}   { return func(cc *{{$.ClassName}}) {{ $.ClassOptionTypeName }} {
	{{- else }}
		func {{$option.OptionFuncName}}(v {{$option.Type}}) {{ $.ClassOptionTypeName }}   { return func(cc *{{$.ClassName}}) {{ $.ClassOptionTypeName }} {
	{{- end }}
	previous := cc.{{$option.Name}}
	cc.{{$option.Name}} = v
	{{- if eq $option.Slice true }}
	return {{$option.OptionFuncName}}(previous...)
	{{- else }}
	return {{$option.OptionFuncName}}(previous)
	{{- end }}
	{{- else}}
	{{- if eq $option.Slice true }}
	func {{$option.OptionFuncName}}(v ...{{$option.SliceElemType}}) {{ $.ClassOptionTypeName }}   { return func(cc *{{$.ClassName}})  {
	{{- else }}
	func {{$option.OptionFuncName}}(v {{$option.Type}}) {{ $.ClassOptionTypeName }}   { return func(cc *{{$.ClassName}})  {
	{{- end }}
	cc.{{$option.Name}} = v
	{{- end}}
} }
{{- end }}

{{ end }}

// {{ $.ClassNewFuncName }} new {{ $.ClassName }}
func {{ $.ClassNewFuncName }} *{{ $.ClassName }} {
	cc := newDefault{{ $.ClassNameTitle }}()
	{{- range $index, $option := $.ClassOptionInfo }}
	{{- if eq $option.ArgIndex 0}}
	{{- else}}
	cc.{{$option.Name}} = {{$option.NameAsParameter}}
	{{- end}}
	{{- end }}

	for _, opt := range opts  {
		opt(cc)
	}
	if watchDog{{$.ClassNameTitle}} != nil {
		watchDog{{$.ClassNameTitle}}(cc)
	}
	return cc
}
// Install{{$.ClassNameTitle}}WatchDog the installed func will called when {{ $.ClassNewFuncName }}  called
func Install{{$.ClassNameTitle}}WatchDog(dog func(cc *{{$.ClassName}})) {
	watchDog{{$.ClassNameTitle}} = dog
}
// watchDog{{$.ClassNameTitle}} global watch dog
var watchDog{{$.ClassNameTitle}} func(cc *{{$.ClassName}})

// newDefault{{ $.ClassNameTitle }} new default {{ $.ClassName }} 
func newDefault{{ $.ClassNameTitle }} () *{{ $.ClassName }} {
	cc := &{{ $.ClassName }}{
{{- range $index, $option := $.ClassOptionInfo }}
	{{- if eq $option.GenOptionFunc false }}
		{{- if eq $option.FieldType 0 }}
			{{$option.Name}}: {{ $option.Type }} {{ $option.Body}},
		{{- else }}
			{{$option.Name}}: {{ $option.Body}},
		{{- end }}
	{{- end }}
{{- end }}
	}

	for _, opt := range [...]{{ $.ClassOptionTypeName }} {
{{- range $index, $option := $.ClassOptionInfo }}
	{{- if eq $option.GenOptionFunc true }}
		{{- if eq $option.Slice true }}
			{{- if eq $option.FieldType 0 }}
				{{$option.OptionFuncName}}({{ $option.Type }} {{ $option.Body}}...),
			{{- else }}
				{{$option.OptionFuncName}}({{ $option.Body}}...),
			{{- end }}
		{{- else }}
			{{- if eq $option.FieldType 0 }}
				{{$option.OptionFuncName}}({{ $option.Type }} {{ $option.Body}}),
			{{- else }}
				{{$option.OptionFuncName}}({{ $option.Body}}),
			{{- end }}
		{{- end }}
	{{- end }}
{{- end }}
	}  {
		opt(cc)
	}

	return cc
}



{{- if $.XConf }}
// AtomicSetFunc used for XConf
func (cc *{{ $.ClassName }}) AtomicSetFunc() func(interface{}) { return Atomic{{ $.ClassNameTitle }}Set }

// atomic{{ $.ClassName }} global *{{ $.ClassName }} holder
var atomic{{ $.ClassNameTitle }} unsafe.Pointer

// Atomic{{ $.ClassNameTitle }}Set atomic setter for *{{ $.ClassName }}
func Atomic{{ $.ClassNameTitle }}Set(update interface{}) {
	atomic.StorePointer(&atomic{{ $.ClassNameTitle }}, (unsafe.Pointer)(update.(*{{ $.ClassName }})))
}

// Atomic{{ $.ClassNameTitle }} return atomic *{{ $.ClassName }} visitor
func Atomic{{ $.ClassNameTitle }}() {{ $.ClassNameTitle }}Visitor {
	current := (*{{ $.ClassName }})(atomic.LoadPointer(&atomic{{ $.ClassNameTitle }}))
	if current == nil {
		atomic.CompareAndSwapPointer(&atomic{{ $.ClassNameTitle }}, nil, (unsafe.Pointer)(newDefault{{ $.ClassNameTitle }}()))
		return (*{{ $.ClassName }})(atomic.LoadPointer(&atomic{{ $.ClassNameTitle }}))
	}
	return current
}
{{- end}}


// all getter func
{{- range $index, $option := $.ClassOptionInfo }}
{{unescaped $option.CommentGetter}}
func (cc *{{ $.ClassName }}) {{$option.VisitFuncName}}() {{ $option.VisitFuncReturnType }} { return cc.{{$option.Name}} }
{{- end }}

// {{ $.ClassNameTitle }}Visitor visitor interface for {{ $.ClassName }}
type {{ $.ClassNameTitle }}Visitor interface {
	{{- range $index, $option := $.ClassOptionInfo }}
	{{$option.VisitFuncName}}() {{ $option.VisitFuncReturnType }} 
	{{- end }}
}

type {{ $.ClassNameTitle }}Interface interface {
	{{ $.ClassNameTitle }}Visitor
	ApplyOption(... {{$.ClassOptionTypeName }}) []{{$.ClassOptionTypeName }} 
}
`
