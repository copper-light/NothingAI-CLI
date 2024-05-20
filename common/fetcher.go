package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"git.datacentric.kr/handh/NothingAI-CLI/settings"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func Request(method string, requestURL string, requestData map[string]any) (map[string]interface{}, error) {
	var bodyData *bytes.Buffer = nil
	var result map[string]interface{} = nil
	client := http.Client{}

	if requestData != nil {
		if strings.ToUpper(method) == "GET" {
			params := url.Values{}
			for key, item := range requestData {
				value, ok := item.(string)
				if ok {
					params.Set(key, value)
				}
			}
			if len(params) > 0 {
				requestURL += "?" + params.Encode()
			}
		} else { // strings.ToUpper(method) == "POST"
			jsonData, err := json.Marshal(requestData)
			if err != nil {
				return nil, err
			}
			bodyData = bytes.NewBuffer(jsonData)
		}
	}
	var req *http.Request
	var err error
	if bodyData == nil {
		req, err = http.NewRequest(method, requestURL, nil)
	} else {
		req, err = http.NewRequest(method, requestURL, bodyData)
	}

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	result = make(map[string]interface{})
	err = json.Unmarshal([]byte(string(body)), &result)
	if err != nil {
		switch resp.StatusCode {
		case http.StatusNotFound:
			err = fmt.Errorf("error: not found resource")
		default:
			err = fmt.Errorf("error: unexpected status code: %d", resp.StatusCode)
		}
	}

	return result, err
}

func GetResources(resourceType string) ([]any, error) {
	var results []any
	requestURL := fmt.Sprintf(RESOURCE_URL, settings.GetServerHost(), resourceType)
	body, err := Request("GET", requestURL, nil)
	if err != nil {
		return nil, err
	}
	code, _ := body["code"].(float64)
	if int(code) == 200 {
		data := body["data"].(map[string]any)
		items := data["items"].([]any)
		if data["next"] == nil {
			requestURL = ""
		} else {
			requestURL = data["next"].(string)
		}
		results = append(results, items...)
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

	return results, nil
}

func DescribeResource(resourceType string, id string) (map[string]any, error) {
	var results map[string]any
	requestURL := fmt.Sprintf(RESOURCE_DETAIL_URL, settings.GetServerHost(), resourceType, id)
	body, err := Request("GET", requestURL, nil)
	if err != nil {
		return nil, err
	}
	code, _ := body["code"].(float64)
	if int(code) == 200 {
		results = body["data"].(map[string]any)
	} else {
		message := ""
		if body["message"] != nil {
			message = body["message"].(string)
		} else {
			message = "Unknown Error"
		}
		err = fmt.Errorf("Error: (%v) %s", code, message)
		return nil, err
	}
	return results, nil
}

func CreateResource(resourceType string, body map[string]any) (int, error) {
	var data map[string]any
	var result = -1
	requestURL := fmt.Sprintf(RESOURCE_URL, settings.GetServerHost(), resourceType)
	body, err := Request("POST", requestURL, body)

	if err != nil {
		return result, err
	}
	code, _ := body["code"].(float64)
	if int(code) == 200 {
		data = body["data"].(map[string]any)
		result = int(data["id"].(float64))
	} else {
		message := ""
		if body["message"] != nil {
			message = body["message"].(string)
		} else {
			message = "Unknown Error"
		}
		err = fmt.Errorf("Error: (%v) %s", code, message)
		return result, err
	}
	return result, err
}
