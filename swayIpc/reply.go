package swayIpc

type Reply struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}
