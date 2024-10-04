package usecase

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"tech-shelf/model"
	"os"
	"net/url"
)

type IBookUsecase interface {
	FetchBooksFromRakutenAPI(query string) (*model.RakutenBookAPIResponse, error)
}

type bookUsecase struct{}

func NewBookUsecase() IBookUsecase {
	return &bookUsecase{}
}

func (bu *bookUsecase)FetchBooksFromRakutenAPI(query string) (*model.RakutenBookAPIResponse, error) {

	encodedQuery := url.QueryEscape(query)
	apiUrl := fmt.Sprintf("https://app.rakuten.co.jp/services/api/BooksTotal/Search/20170404?format=json&keyword=%s&booksGenreId=000&applicationId=%s", encodedQuery, os.Getenv("RAKUTEN_APP_ID"))

	resp, err := http.Get(apiUrl)
	if err != nil {
			return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
			return nil, err
	}

	var result model.RakutenBookAPIResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
			return nil, err
	}

	return &result, nil
}