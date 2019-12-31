package errors

//
type Errors struct {
	Errors []*Error `json:"errors"`
}

// error structure
// contains informations about the error
type Error struct {
	Id     string `json:"id"`
	Status int    `json:"status"`
	Title  string `json:"title"`
	Detail string `json:"detail"`
}
