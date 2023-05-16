package main

import (
	"bytes"
	"strings"
	"text/template"
)

// TODO 模版最后的c.JSON(200, reply)加一个判断如果是 http.body.
var httpTemplate = `
{{$svrType := .ServiceType}}
{{$svrName := .ServiceName}}
{{$bizCodeModel := .BizCodeModel}}

type {{.ServiceType}} interface {
{{- range .MethodSets}}
	{{- if not (eq .HeadComment "")}}
	{{.HeadComment}}
	{{- end}}
	{{.Name}}(context.Context, *{{.Request}}) (*{{.Reply}}, error)
{{- end}}
}

func {{.ServiceType}}APIRouters(ctr {{.ServiceType}}) []httpserver.Router {
	return []httpserver.Router{
		{{- range .Methods}}
		{
			HTTPMethod:   "{{.Method}}",
			Path:         "{{.Path}}",
			HandlerFuncs: []gin.HandlerFunc{ _{{.Name}}Handler(ctr)},
		},
		{{- end}}
	}
}

{{range .Methods}}
func _{{.Name}}Handler(ctr {{$svrType}}) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		in := &{{.Request}}{}
		{{- if .HasBody}}
		if err := c.Bind(in); err != nil {
			{{- if not $bizCodeModel}}
			c.AbortWithStatusJSON(errors.BindCoder.HTTPStatus(), errors.BindCoder)
			{{- else}}
			c.AbortWithStatusJSON(200, errors.BindCoder)
			{{- end}}
			c.Abort()
			return
		}
		
		{{- if not (eq .Body "")}}
		if err := c.BindQuery(in); err != nil {
			{{- if not $bizCodeModel}}
			c.AbortWithStatusJSON(errors.BindCoder.HTTPStatus(), errors.BindCoder)
			{{- else}}
			c.AbortWithStatusJSON(200, errors.BindCoder)
			{{- end}}
			return
		}
		{{- end}}
		{{- else}}
		if err := c.BindQuery(in); err != nil {
			{{- if not $bizCodeModel}}
			c.AbortWithStatusJSON(errors.BindCoder.HTTPStatus(), errors.BindCoder)
			{{- else}}
			c.AbortWithStatusJSON(200, errors.BindCoder)
			{{- end}}
			return
		}
		{{- end}}
		{{- if .HasVars}}
		if err := c.BindUri(in); err != nil {
			{{- if not $bizCodeModel}}
			c.AbortWithStatusJSON(errors.BindCoder.HTTPStatus(), errors.BindCoder)
			{{- else}}
			c.AbortWithStatusJSON(200, errors.BindCoder)
			{{- end}}
			return
		}
		{{- end}}

		if v, ok := interface{}(in).(interface{Validate() error}); ok {
			if err := v.Validate(); err != nil {
				{{- if not $bizCodeModel}}
				c.AbortWithStatusJSON(errors.ValidationCoder.HTTPStatus(), errors.ValidationCoder)
				{{- else}}
				c.AbortWithStatusJSON(200, errors.ValidationCoder)
				{{- end}}
				return
			}
		}

		reply, err := ctr.{{.Name}}(ctx, in)
		if err != nil {
			coder := errors.ParseCoder(err)
			{{- if $bizCodeModel}}
			if coder.HTTPStatus() == 500 {
				c.JSON(500, coder)
			}else {
				c.JSON(200, coder)
			}
			{{- else}}
			c.JSON(coder.HTTPStatus(), errors.ParseCoder(err))
			{{- end}}
			return
		}

		c.JSON(200, reply)
	}
}
{{end}}
`

type serviceDesc struct {
	ServiceType string // Greeter
	ServiceName string // helloworld.Greeter
	Metadata    string // api/helloworld/helloworld.proto
	Methods     []*methodDesc
	MethodSets  map[string]*methodDesc
	// 是否是业务码模式；
	// 非业务码模式下，http status code 等于 error定义的http status code
	// 业务码模式下，http status code 仅在error定义的http status code为500时生效，其余情况下http status code都为200
	BizCodeModel bool
}

type methodDesc struct {
	// method
	HeadComment  string
	Name         string
	OriginalName string // The parsed original name
	Num          int
	Request      string
	Reply        string
	// http_rule
	Path         string
	Method       string
	HasVars      bool
	HasBody      bool
	Body         string
	ResponseBody string
}

func (s *serviceDesc) execute() string {
	s.MethodSets = make(map[string]*methodDesc)
	for _, m := range s.Methods {
		s.MethodSets[m.Name] = m
	}
	buf := new(bytes.Buffer)
	tmpl, err := template.New("http").Parse(strings.TrimSpace(httpTemplate))
	if err != nil {
		panic(err)
	}
	if err := tmpl.Execute(buf, s); err != nil {
		panic(err)
	}
	return strings.Trim(buf.String(), "\r\n")
}
