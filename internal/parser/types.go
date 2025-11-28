package parser

// Recipe represents a Workato recipe JSON structure
type Recipe struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Version     int      `json:"version"`
	Code        Code     `json:"code"`
	Config      []Config `json:"config"`
}

type Code struct {
	Number   int     `json:"number"`
	Provider string  `json:"provider"`
	Name     string  `json:"name"`
	Block    []Block `json:"block"`
}

type Block struct {
	Number   int    `json:"number"`
	Provider string `json:"provider"`
	Name     string `json:"name"`
}

type Config struct {
	Provider  string     `json:"provider"`
	AccountID *AccountID `json:"account_id"`
}

type AccountID struct {
	ZipName string `json:"zip_name"`
	Name    string `json:"name"`
}

type Connection struct {
	Name     string `json:"name"`
	Provider string `json:"provider"`
}

// ProcessedRecipe is the normalized output structure
type ProcessedRecipe struct {
	Name        string
	Description string
	Trigger     Step
	Actions     []Step
	Systems     []string
	Connections []ConnectionInfo
	Flow        []SystemFlow
}

type Step struct {
	System     string
	Action     string
	StepNumber int
}

type ConnectionInfo struct {
	Provider       string
	ConnectionName string
	IsBuiltIn      bool
}

type SystemFlow struct {
	From       string
	To         string
	ActionType string
}
