package optiongen

const templateTextWithPreviousSupport = `

{{- range $_, $comment := $.ClassComments }}
{{ $comment }}
{{- end }}

type {{ $.ClassName }} struct {
	{{- range $index, $option := $.ClassOptionInfo }}
		{{- range $_, $comment := $option.LastRowComments }}
			{{ $comment }}
 		{{- end }}
		{{ $option.Name }} {{ $option.Type }} {{unescaped $option.TagString}} {{ $option.SameRowComment }} 
	{{- end }}
}

func (cc *{{ $.ClassName }}) SetOption(opt {{$.ClassOptionTypeName}}) {
	_ = opt(cc)
}

func (cc *{{ $.ClassName }}) ApplyOption(opts... {{$.ClassOptionTypeName }}) {
	for _, opt := range opts  {
		_ = opt(cc)
	}
}


func (cc *{{ $.ClassName }}) GetSetOption(opt {{ $.ClassOptionTypeName }}) {{ $.ClassOptionTypeName }} {
	return opt(cc)
}

type {{ $.ClassOptionTypeName }} func(cc *{{$.ClassName}}) {{ $.ClassOptionTypeName }}
{{ range $index, $option := $.ClassOptionInfo }}

{{- if eq $option.GenOptionFunc true }}
	{{- range $methodCommentIndex, $methodComment := $option.MethodComments }}
		{{ $methodComment }}
	{{- end }}
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


func {{ $.ClassNewFuncName }} *{{ $.ClassName }} {
	cc := newDefault{{ $.ClassName }}()
	{{- range $index, $option := $.ClassOptionInfo }}
	{{- if eq $option.Index 0}}
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

func Install{{$.ClassName}}WatchDog(dog func(cc *{{$.ClassName}})) {
	watchDog{{$.ClassName}} = dog
}

var watchDog{{$.ClassName}} func(cc *{{$.ClassName}})

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

func (cc *{{ $.ClassName }}) AtomicSetFunc() func(interface{}) { return Atomic{{ $.ClassName }}Set }

var atomic{{ $.ClassName }} unsafe.Pointer

func Atomic{{ $.ClassName }}Set(update interface{}) {
	atomic.StorePointer(&atomic{{ $.ClassName }}, (unsafe.Pointer)(update.(*{{ $.ClassName }})))
}
func Atomic{{ $.ClassName }}() *{{ $.ClassName }} {
	current := (*{{ $.ClassName }})(atomic.LoadPointer(&atomic{{ $.ClassName }}))
	if current == nil {
		atomic.CompareAndSwapPointer(&atomic{{ $.ClassName }}, nil, (unsafe.Pointer)(newDefault{{ $.ClassName }}()))
		return (*{{ $.ClassName }})(atomic.LoadPointer(&atomic{{ $.ClassName }}))
	}
	return current
}


{{- end}}
`
