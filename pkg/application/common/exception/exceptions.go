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
	StatusCode int
	Message    string
}

type BadRequest struct {
	StatusCode int
	Message    string
}

type InternalServer struct {
	StatusCode int
	Message    string
}

type Custom struct {
	StatusCode int
	Title      string
	Message    string
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

func (e *Custom) Error() string {
	return e.Message
}

func (e *Exception) Error() string {
	return e.Message
}

// NewException is a generic function that returns a pointer to an exception
func NewException[E NotFound | InternalServer | BadRequest](message string) *E {
	return &E{
		Message: message,
	}
}

// NewCustomException is a generic function that returns a pointer to a custom exception
func NewCustomException(statusCode int, message string) *Custom {
	return &Custom{
		Message:    message,
		StatusCode: statusCode,
	}
}
