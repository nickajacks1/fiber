package oapi

type Spec struct {
	OpenAPI string `json:"openapi"`
	Info    Info   `json:"info"`
	// JSONSchemaDialect string `json:"jsonSchemaDialect"`
	// Servers []Server
	// Paths Paths
	// Webhooks
	// Components
	// Security
	// Tags
	// ExternalDocs
}

type Info struct {
	Title          string  `json:"title"`
	Summary        string  `json:"summary,omitempty"`
	Description    string  `json:"description,omitempty"`
	TermsOfService string  `json:"termsOfService,omitempty"`
	Contact        Contact `json:"contact,omitempty"`
	License        License `json:"license,omitempty"`
	Version        string  `json:"version"`
}

type Contact struct {
	Name  string `json:"name,omitempty"`
	URL   string `json:"url,omitempty"`
	Email string `json:"email,omitempty"`
}

type License struct {
	Name       string `json:"name"`
	Identifier string `json:"identifier,omitempty"`
	URL        string `json:"url,omitempty"`
}

type Server struct {
	URL         string `json:"url"`
	Description string `json:"description,omitempty"`
	// Variables ServerVariable
}

type Paths map[string]PathItem

type PathItem struct {
	Ref         string    `json:"$ref,omitempty"`
	Summary     string    `json:"summary,omitempty"`
	Description string    `json:"description,omitempty"`
	Get         Operation `json:"get,omitempty"`
	Put         Operation `json:"put,omitempty"`
	Post        Operation `json:"post,omitempty"`
	Patch       Operation `json:"patch,omitempty"`
	Delete      Operation `json:"delete,omitempty"`
	Options     Operation `json:"options,omitempty"`
	Head        Operation `json:"head,omitempty"`
	Trace       Operation `json:"trace,omitempty"`
}

type Operation struct{}
