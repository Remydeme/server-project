package response

// structure that contains the return value of a typical response
// Body contains the useful information
// StatusText contains a status message that describe the state after the request
//StatusId contains the status code
type Response struct {
	StatusText string
	StatusId   int
	Body       string
}
