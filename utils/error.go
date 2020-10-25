package utils

type ErrorBag struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

func (e ErrorBag) Error() string {
	panic("implement me")
}