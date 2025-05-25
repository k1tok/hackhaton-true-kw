package v1

type UseCases struct {
	CheckUsecase CheckUseCase
}

type HTTPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
