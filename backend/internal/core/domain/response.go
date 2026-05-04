package domain

type RequestValidator[T any] struct {
	Type T
}

type Message struct {
	En string `json:"en"`
	Es string `json:"es"`
}

type APIResponse[T any] struct {
	Success     bool                `json:"success"`
	Message     Message             `json:"message"`
	Data        T                   `json:"data" swaggerignore:"true"`
	SchemaError map[string][]string `json:"schema" swaggerignore:"true"`
	Error       string              `json:"error"`
}

type RequestInfo struct {
	Host      string `json:"host"`
	IP        string `json:"ip"`
	UserAgent string `json:"agent"`
	UserID    uint   `json:"userID"`
}

type RouteInfo struct {
	Method      string `json:"method"`
	Path        string `json:"path"`
	Description string `json:"description"`
	Handler     string `json:"handler"`
}

type GenericCatalog struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
