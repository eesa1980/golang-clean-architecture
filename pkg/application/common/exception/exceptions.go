package exception

type Exception struct {
	StatusCode  int    `json:"statusCode"`
	Title       string `json:"title"`
	Explanation string `json:"code"`
	Message     string `json:"message"`
}

type Http struct {
	StatusCode int
	Message    string
}

type NotFound struct {
	Message string
}

type BadRequest struct {
	Message string
}

type InternalServer struct {
	Message string
}

func (e *Http) Error() string {
	return e.Message
}

func (e *NotFound) Error() string {
	return e.Message
}

func (e *InternalServer) Error() string {
	return e.Message
}

func (e *BadRequest) Error() string {
	return e.Message
}
