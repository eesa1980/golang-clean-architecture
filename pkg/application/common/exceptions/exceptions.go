package application

type Exception struct {
	StatusCode  int    `json:"statusCode"`
	Title       string `json:"title"`
	Explanation string `json:"code"`
	Message     string `json:"message"`
}

type HttpException struct {
	StatusCode int
	Message    string
}

type NotFoundException struct {
	Message string
}

type BadRequestException struct {
	Message string
}

type InternalServerException struct {
	Message string
}

type CustomException struct {
	Exception
}

func (e *HttpException) Error() string {
	return e.Message
}

func (e *NotFoundException) Error() string {
	return e.Message
}

func (e *InternalServerException) Error() string {
	return e.Message
}

func (e *BadRequestException) Error() string {
	return e.Message
}

func (e *CustomException) Error() string {
	return e.Message
}
