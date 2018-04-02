package codegen

import (
	"fmt"
	"path/filepath"

	"goa.design/goa/codegen"
	httpdesign "goa.design/goa/http/design"
)

// PathFiles returns the service path files.
func PathFiles(root *httpdesign.RootExpr) []*codegen.File {
	fw := make([]*codegen.File, 2*len(root.HTTPServices))
	for i := 0; i < len(root.HTTPServices); i++ {
		fw[i*2] = serverPath(root.HTTPServices[i])
		fw[i*2+1] = clientPath(root.HTTPServices[i])
	}
	return fw
}

// serverPath returns the server file containing the request path constructors
// for the given service.
func serverPath(svc *httpdesign.ServiceExpr) *codegen.File {
	path := filepath.Join(codegen.Gendir, "http", codegen.SnakeCase(svc.Name()), "server", "paths.go")
	return &codegen.File{Path: path, SectionTemplates: pathSections(svc, "server")}
}

// clientPath returns the client file containing the request path constructors
// for the given service.
func clientPath(svc *httpdesign.ServiceExpr) *codegen.File {
	path := filepath.Join(codegen.Gendir, "http", codegen.SnakeCase(svc.Name()), "client", "paths.go")
	return &codegen.File{Path: path, SectionTemplates: pathSections(svc, "client")}
}

// pathSections returns the sections of the file of the pkg package that
// contains the request path constructors for the given service.
func pathSections(svc *httpdesign.ServiceExpr, pkg string) []*codegen.SectionTemplate {
	title := fmt.Sprintf("HTTP request path constructors for the %s service.", svc.Name())
	sections := []*codegen.SectionTemplate{
		codegen.Header(title, pkg, []*codegen.ImportSpec{
			{Path: "fmt"},
			{Path: "net/url"},
			{Path: "strconv"},
			{Path: "strings"},
		}),
	}
	sdata := HTTPServices.Get(svc.Name())
	for _, e := range svc.HTTPEndpoints {
		sections = append(sections, &codegen.SectionTemplate{
			Name:   "path",
			Source: pathT,
			Data:   sdata.Endpoint(e.Name()),
		})
	}

	return sections
}

// input: EndpointData
const pathT = `{{ range .Routes }}// {{ .PathInit.Description }}
func {{ .PathInit.Name }}({{ range .PathInit.ServerArgs }}{{ .Name }} {{ .TypeRef }}, {{ end }}) {{ .PathInit.ReturnTypeRef }} {
{{- .PathInit.ServerCode }}
}
{{ end }}`
