package parser

// INFO:: I need to flatten the messy structure of the Workato
// JSON files so I am creating two structures here: one to catpure the
// data from the manifest files, and one to normalise them into something
// I can parse easily

// NOTE: Workato recipe JSON structure
// topd level of .recipe.json file
type Recipe struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Version     int      `json:"version"`
	Code        Code     `json:"code"` // workflow definition
	Config      []Config `json:"config"`
}

// NOTE: logic
type Code struct {
	Number   int     `json:"number"`
	Provider string  `json:"provider"`
	Name     string  `json:"name"`
	Block    []Block `json:"block"`
}

// NOTE: actions
type Block struct {
	Number   int    `json:"number"`
	Provider string `json:"provider"`
	Name     string `json:"name"`
}

// NOTE: sys connections
type Config struct {
	Provider  string     `json:"provider"`
	AccountID *AccountID `json:"account_id"`
}

// FIX: not sure I want this in the end
type AccountID struct {
	ZipName string `json:"zip_name"`
	Name    string `json:"name"`
}

type Connection struct {
	Name     string `json:"name"`
	Provider string `json:"provider"`
}

// NOTE: JSON transformation layer
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
