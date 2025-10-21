package http

import (
	"net/http"

	"github.com/JoaoPedr0Maciel/charm/internal/structs"
	"github.com/JoaoPedr0Maciel/charm/internal/utils"
)

func MakeRequest(opts structs.RequestOptions) (*http.Response, error) {
	return utils.DoRequest(opts)
}

func Get(url, bearer, basic, contentType string) (*http.Response, error) {
	return MakeRequest(structs.RequestOptions{
		Method:      "GET",
		URL:         url,
		Bearer:      bearer,
		Basic:       basic,
		ContentType: contentType,
	})
}

func Post(url, bearer, basic, contentType, data string) (*http.Response, error) {
	return MakeRequest(structs.RequestOptions{
		Method:      "POST",
		URL:         url,
		Bearer:      bearer,
		Basic:       basic,
		ContentType: contentType,
		Data:        data,
	})
}

func Put(url, bearer, basic, contentType, data string) (*http.Response, error) {
	return MakeRequest(structs.RequestOptions{
		Method:      "PUT",
		URL:         url,
		Bearer:      bearer,
		Basic:       basic,
		ContentType: contentType,
		Data:        data,
	})
}

func Patch(url, bearer, basic, contentType, data string) (*http.Response, error) {
	return MakeRequest(structs.RequestOptions{
		Method:      "PATCH",
		URL:         url,
		Bearer:      bearer,
		Basic:       basic,
		ContentType: contentType,
		Data:        data,
	})
}

func Delete(url, bearer, basic, contentType, data string) (*http.Response, error) {
	return MakeRequest(structs.RequestOptions{
		Method:      "DELETE",
		URL:         url,
		Bearer:      bearer,
		Basic:       basic,
		ContentType: contentType,
		Data:        data,
	})
}
