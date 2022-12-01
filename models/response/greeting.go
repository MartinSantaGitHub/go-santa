package response

/* Greeting response */
type Greeting struct {
	Message string `json:"message,omitempty"`
	Exists  bool   `json:"exists"`
}
