package base

type ValidateRequest struct {
	UserId      string
	Title       string
	Description string
}

func Validate(req *ValidateRequest) []string {
	res := make([]string, 0)

	if req == nil {
		res = append(res, ReqValidateMessage)
	}

	if req != nil && req.UserId == emptyField {
		res = append(res, UserIdValidateMessage)
	}

	if req != nil && req.Title == emptyField {
		res = append(res, TitleValidateMessage)
	}

	if req != nil && req.Description == emptyField {
		res = append(res, DescriptionValidateMessag)
	}

	return res
}

const (
	ReqValidateMessage        = "req is nil"
	UserIdValidateMessage     = "UserId is empty"
	DescriptionValidateMessag = "Desс is empty"
	TitleValidateMessage      = "Title is empty"
	emptyField                = ""
)
