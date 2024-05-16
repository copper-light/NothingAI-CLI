package fetch

import (
	"encoding/json"
	"fmt"
	"git.datacentric.kr/handh/NothingAI-CLI/settings"
	"io"
	"net/http"
	"strings"
)

func Request(method string, url string) (map[string]interface{}, error) {
	var result map[string]interface{} = nil
	client := http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	switch resp.StatusCode {
	case http.StatusOK:
		result = make(map[string]interface{})
		err = json.Unmarshal([]byte(string(body)), &result)
	default:
		err = fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return result, err
}

func GetResources(url string, resource string) ([]any, error) {
	var data []any
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = settings.GetServerHost() + url
	}
	for url != "" {
		body, err := Request("GET", url)
		if err != nil {
			return nil, err
		}
		code, _ := body["code"].(float64)
		if int(code) == 200 {
			results := body["results"].(map[string]any)
			model := results[resource].(map[string]any)
			if model["next"] == nil {
				url = ""
			} else {
				url = model["next"].(string)
			}
			data = append(data, model["data"].([]any)...)
		} else {
			message := ""
			if body["message"] != nil {
				message = body["message"].(string)
			} else {
				message = "Unknown Error"
			}
			err = fmt.Errorf("[%v] %s", code, message)
			return nil, err
		}
	}

	return data, nil
}

func DescribeResource(url string, id string, resource string) (map[string]any, error) {
	var data map[string]any
	if url != "" {
		fmt.Println(url)
		url = fmt.Sprintf(url, id)
		url = settings.GetServerHost() + url
		fmt.Println(url)
		body, err := Request("GET", url)
		if err != nil {
			return nil, err
		}
		code, _ := body["code"].(float64)
		if int(code) == 200 {
			results := body["results"].(map[string]any)
			data = results[resource].(map[string]any)
		} else {
			message := ""
			if body["message"] != nil {
				message = body["message"].(string)
			} else {
				message = "Unknown Error"
			}
			err = fmt.Errorf("[%v] %s", code, message)
			return nil, err
		}
	}
	return data, nil
}
