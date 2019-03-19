package swayIpc

type Container struct {
	TreeNode

	WindowProperties WindowProperties `json:"window_properties"`
	Pid              uint64           `json:"pid"`
}
