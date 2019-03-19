package swayIpc

type Container struct {
	TreeNode

	Nodes            []Container         `json:"nodes"`
	FloatingNodes    []ContainerFloating `json:"floating_nodes"`
	WindowProperties WindowProperties    `json:"window_properties"`
	Pid              uint64              `json:"pid"`
}
