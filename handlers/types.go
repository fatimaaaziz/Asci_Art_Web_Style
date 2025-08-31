package ourcode

type PageData struct {
	Input  string
	Banner string
	Result string
}

type errorData struct {
	ErNUM   int
	ErrMess string
}

func Initialiseerr(ErrNumb int, ErrMes string) errorData {
	return errorData{
		ErNUM:   ErrNumb,
		ErrMess: ErrMes,
	}
}
