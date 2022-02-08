package pd

type ResponseUpload struct {
	StatusCode int    `json:"code"`
	ID         string `json:"id,omitempty"`
	Success    bool   `json:"success"`
	Value      string `json:"value,omitempty"`
	Message    string `json:"message,omitempty"`
}
