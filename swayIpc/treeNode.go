package swayIpc

// https://i3wm.org/docs/ipc.html#_tree_reply
// but missing window_properties
type TreeNode struct {
	ID                 int         `json:"id"`
	Name               string      `json:"name"`
	Type               string      `json:"type"`
	Border             string      `json:"border"`
	CurrentBorderWidth int         `json:"current_border_width"`
	Layout             string      `json:"layout"`
	Orientation        string      `json:"orientation"`
	Percent            float64     `json:"percent"`
	Rect               Rect        `json:"rect"`
	WindowRect         Rect        `json:"window_rect"`
	DecoRect           Rect        `json:"deco_rect"`
	Geometry           Rect        `json:"geometry"`
	Window             interface{} `json:"window"`
	Urgent             bool        `json:"urgent"`
	Focused            bool        `json:"focused"`
	Focus              []int       `json:"focus"`
	Nodes              []TreeNode  `json:"nodes"`
	FloatingNodes      []TreeNode  `json:"floating_nodes"`
	Sticky             bool        `json:"sticky"`
}
