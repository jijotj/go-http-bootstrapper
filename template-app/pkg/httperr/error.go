package httperr

const (
	StatusUnknown = 520
)

type Error struct {
	HTTPStatus int    `json:"httpStatus"`
	Code       string `json:"code"`
	Message    string `json:"message"`
}

func (err Error) Error() string {
	return err.Message
}

var UnknownError = Error{HTTPStatus: StatusUnknown, Code: "520", Message: "UNKNOWN_ERROR"}
