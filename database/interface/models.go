package _interface

type Trigger struct {
	Target  string
	Code    int
	Enabled bool
}

type ActionParameter struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Default string `json:"default"`
}

type ActionEnvironment struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ActionDO struct {
	ID           int
	Title        string
	Content      string
	Trigger      Trigger
	Parameters   []ActionParameter
	Environments []ActionEnvironment
}
