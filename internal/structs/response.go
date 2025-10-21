package structs

type Response struct {
	StatusCode    int
	Url           string
	Method        string
	Authorization *string
	ContentType   string
	Body          *[]byte
}
