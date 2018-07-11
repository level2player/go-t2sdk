package t2sdk

type IResponse interface {
}

type Response struct {
	FunctionNo     int
	ReturnCode     int
	RequestContent string
}
