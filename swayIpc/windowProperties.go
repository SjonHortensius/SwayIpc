package swayIpc

type WindowProperties struct {
	Class        string `json:"class"`
	Instance     string `json:"instance"`
	Title        string `json:"title"`
	TransientFor uint64 `json:"transient_for"`
	WindowRole   string `json:"window_role"`
}
