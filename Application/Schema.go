package Application

type Schema struct {
	Redirects []Redirect `json:"redirects"`
}

type Redirect struct {
	From string `json:"from"`
	To   string `json:"to"`
}
