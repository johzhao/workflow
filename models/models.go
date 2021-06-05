package models

type ActionParameter struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Default string `json:"default"`
}

type ActionEnvironment struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
