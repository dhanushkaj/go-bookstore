package utils

type ApiError struct {
	Err    string
	Status int
}

func (e ApiError) Error() string {
	return e.Err
}
