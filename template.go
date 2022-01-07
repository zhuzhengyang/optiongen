package optiongen

const templateTextWithPreviousSupport = `

{{- range $_, $comment := $.ClassComments }}
{{ $comment }}
{{- end }}

// {{ $.ClassName }} struct
type {{ $.ClassName }} struct {
	{{- range $index, $option := $.ClassOptionInfo }}
		{{- range $_, $comment := $option.LastRowComments }}
			{{ $comment }}
 		{{- end }}
		{{ $option.Name }} {{ $option.Type }} {{unescaped $option.TagString}} {{ $option.SameRowComment }} 
	{{- end }}
}
// Deprecated: use ApplyOption instead
// SetOption apply single option
func (cc *{{ $.ClassName }}) SetOption(opt {{$.ClassOptionTypeName}}) {
	cc.ApplyOption(opt)
}

// ApplyOption apply new option and return the old optuon
// sample: 
// old := cc.ApplyOption(WithTimeout(time.Second))
// defer cc.ApplyOption(old...)
// ApplyOption apply mutiple options
func (cc *{{ $.ClassName }}) ApplyOption(opts... {{$.ClassOptionTypeName }}) []{{$.ClassOptionTypeName }}{
	var previous []{{$.ClassOptionTypeName }}
	for _, opt := range opts  {
		previous = append(previous,opt(cc))
	}
	return previous
}

// Deprecated: use ApplyOption instead
// GetSetOption apply new option and return the old optuon
// sample: 
// old := cc.GetSetOption(WithTimeout(time.Second))
// defer cc.SetOption(old)
func (cc *{{ $.ClassName }}) GetSetOption(opt {{ $.ClassOptionTypeName }}) {{ $.ClassOptionTypeName }} {
	return opt(cc)
}
// {{ $.ClassOptionTypeName }} option func
type {{ $.ClassOptionTypeName }} func(cc *{{$.ClassName}}) {{ $.ClassOptionTypeName }}
{{ range $index, $option := $.ClassOptionInfo }}
{{- if eq $option.GenOptionFunc true }}
	{{- range $methodCommentIndex, $methodComment := $option.MethodComments }}
		{{ $methodComment }}
	{{- end }}
	// {{$option.OptionFuncName}} option func for {{ $option.Name }}
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
} }
{{- end }}

{{ end }}

// {{ $.ClassNewFuncName }} new {{ $.ClassName }}
func {{ $.ClassNewFuncName }} *{{ $.ClassName }} {
	cc := newDefault{{ $.ClassName }}()
	{{- range $index, $option := $.ClassOptionInfo }}
	{{- if eq $option.ArgIndex 0}}
	{{- else}}
	cc.{{$option.Name}} = {{$option.NameAsParameter}}
	{{- end}}
	{{- end }}

	for _, opt := range opts  {
		_ = opt(cc)
	}
	if watchDog{{$.ClassName}} != nil {
		watchDog{{$.ClassName}}(cc)
	}
	return cc
}
// Install{{$.ClassName}}WatchDog the installed func will called when {{ $.ClassNewFuncName }}  called
func Install{{$.ClassName}}WatchDog(dog func(cc *{{$.ClassName}})) {
	watchDog{{$.ClassName}} = dog
}
// watchDog{{$.ClassName}} global watch dog
var watchDog{{$.ClassName}} func(cc *{{$.ClassName}})

// newDefault{{ $.ClassName }} new default {{ $.ClassName }} 
func newDefault{{ $.ClassName }} () *{{ $.ClassName }} {
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
		_ = opt(cc)
	}

	return cc
}



{{- if $.XConf }}
// AtomicSetFunc used for XConf
func (cc *{{ $.ClassName }}) AtomicSetFunc() func(interface{}) { return Atomic{{ $.ClassName }}Set }

// atomic{{ $.ClassName }} global *{{ $.ClassName }} holder
var atomic{{ $.ClassName }} unsafe.Pointer

// Atomic{{ $.ClassName }}Set atomic setter for *{{ $.ClassName }}
func Atomic{{ $.ClassName }}Set(update interface{}) {
	atomic.StorePointer(&atomic{{ $.ClassName }}, (unsafe.Pointer)(update.(*{{ $.ClassName }})))
}

// Atomic{{ $.ClassName }} return atomic *{{ $.ClassName }} visitor
func Atomic{{ $.ClassName }}() {{ $.ClassName }}Visitor {
	current := (*{{ $.ClassName }})(atomic.LoadPointer(&atomic{{ $.ClassName }}))
	if current == nil {
		atomic.CompareAndSwapPointer(&atomic{{ $.ClassName }}, nil, (unsafe.Pointer)(newDefault{{ $.ClassName }}()))
		return (*{{ $.ClassName }})(atomic.LoadPointer(&atomic{{ $.ClassName }}))
	}
	return current
}
{{- end}}


// all getter func
{{- range $index, $option := $.ClassOptionInfo }}
{{$option.CommentGetter}}
func (cc *{{ $.ClassName }}) {{$option.VisitFuncName}}() {{ $option.VisitFuncReturnType }} { return cc.{{$option.Name}} }
{{- end }}

// {{ $.ClassName }}Visitor visitor interface for {{ $.ClassName }}
type {{ $.ClassName }}Visitor interface {
	{{- range $index, $option := $.ClassOptionInfo }}
	{{$option.VisitFuncName}}() {{ $option.VisitFuncReturnType }} 
	{{- end }}
}

type {{ $.ClassName }}Interface interface {
	{{ $.ClassName }}Visitor
	ApplyOption(... {{$.ClassOptionTypeName }}) []{{$.ClassOptionTypeName }} 
}
`
