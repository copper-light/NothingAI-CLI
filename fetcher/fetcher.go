package fetcher

import (
	"fmt"
	"io"
	"net/http"
)

func Request(method string, url string) (string, error) {
	client := http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return "", err
	}
	//defer resp.Body.Close()

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	var result string
	switch resp.StatusCode {
	case http.StatusOK:
		result = string(body)
	default:
		err = fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	//if resp.StatusCode != http.StatusOK {
	//	return nil, fmt.Errorf("bad status: %s", resp.Status)
	//}

	//htmlBytes, err := io.ReadAll(resp)
	//if err != nil {
	//	return nil, err
	//}
	defer resp.Body.Close()
	return result, err
}

//func httpResponseHandler(resp *http.Response) (string, error) {
//	var err error = nil
//	var body string = ""
//	switch response.StatusCode {
//	case http.StatusOK:
//		if
//		body = string()
//	default:
//		err := fmt.Errorf("unexpected status code: %d", response.StatusCode)
//	}
//	return body, err
//}
