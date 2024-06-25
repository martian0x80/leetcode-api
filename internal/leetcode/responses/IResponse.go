package responses

type IResponse interface {
	ToByte() ([]byte, error)
	ToString() (string, error)
}
