package swayIpc

type Workspace struct {
	TreeNode
	Representation string        `json:"representation"`
	NumXX          uint32        `json:"num"`
	Nodes          []Container   `json:"nodes"`
	FloatingNodes  []FloatingCon `json:"floating_nodes"`
}
