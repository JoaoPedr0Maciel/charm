package structs

type Request struct {
	Url           string
	Method        string
	Authorization *string
	ContentType   string
	Body          *[]byte
}
