package swayIpc

type Workspace struct {
	TreeNode
	Representation string              `json:"representation"`
	// The Name when numeric - otherwise -1
	Numeric        uint32              `json:"num"`
	Nodes          []Container         `json:"nodes"`
	FloatingNodes  []ContainerFloating `json:"floating_nodes"`
}
