package schemas

type HomeEvent struct {
	ApiKey        string         `json:"apiKey" validate:"required"`
	Source        string         `json:"source" validate:"required,oneof=desktop mobile app"`
	SecretKey     *string        `json:"secretKey"`
	Name          *string        `json:"name"`
	Url           *string        `json:"url" validate:"url"`
	User          *User          `json:"user"`
	Identity      Identity       `json:"identity"`
	TestGroup     *TestGroup     `json:"testGroup"`
	DeviceId      *string        `json:"deviceId"`
	Info          *Info          `json:"info"`
	AnonymousUser *AnonymousUser `json:"anonymousUser"`
}
