package schemas

type User struct {
	Id                 string  `json:"id" validate:"required"`
	Name               *string `json:"name"`
	Email              *string `json:"email"`
	AllowMailMarketing *bool   `json:"allowMailMarketing"`
}

type AnonymousUser struct {
	Name               *string `json:"name"`
	Email              *string `json:"email"`
	AllowMailMarketing *bool   `json:"allowMailMarketing"`
}

type Info = map[string]string

type Identity struct {
	AnonymousUserId string `json:"anonymousUserId" validate:"required"`
	BrowserId       string `json:"browserId" validate:"required"`
	Session         string `json:"session" validate:"required"`
}

type TestGroup struct {
	Code       *string `json:"code"`
	Experiment *string `json:"experiment"`
	Group      *string `json:"group"`
	Session    *string `json:"session"`
	TestCode   *string `json:"testCode"`
}
