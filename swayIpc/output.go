package swayIpc

type Output struct {
	TreeNode
	Nodes []Workspace `json:"nodes"`

	Active           bool    `json:"active"`
	Primary          bool    `json:"primary"`
	Make             string  `json:"make"`
	Model            string  `json:"model"`
	Serial           string  `json:"serial"`
	Scale            float64 `json:"scale"`
	Transform        string  `json:"transform"`
	CurrentWorkspace string  `json:"current_workspace"`
	Modes            []Mode  `json:"modes"`
	CurrentMode      Mode    `json:"current_mode"`
}
