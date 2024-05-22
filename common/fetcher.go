package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"git.datacentric.kr/handh/NothingAI-CLI/settings"
	"github.com/iancoleman/orderedmap"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func Request(method string, requestURL string, requestData map[string]any) (*orderedmap.OrderedMap, error) {
	var bodyData *bytes.Buffer = nil
	//var result map[string]interface{} = nil
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

	//result = make(map[string]interface{})
	result := orderedmap.New()

	err = json.Unmarshal([]byte(string(body)), &result)
	if err != nil {
		switch resp.StatusCode {
		case http.StatusNotFound:
			err = fmt.Errorf("error: resource not found")
		default:
			err = fmt.Errorf("error: unexpected status code: %d", resp.StatusCode)
		}
	}

	return result, err
}

func GetResources(resourceType string) ([]any, error) {
	var results []any
	var items []any
	var objectValue orderedmap.OrderedMap
	var value any
	var ok bool
	requestURL := fmt.Sprintf(RESOURCE_URL, settings.GetServerHost(), resourceType)
	body, err := Request("GET", requestURL, nil)
	if err != nil {
		return nil, err
	}
	value, ok = body.Get("code")
	if !ok {
		return nil, fmt.Errorf("returned data is not json format")
	}
	code := int(value.(float64))
	if code == 200 {
		value, ok = body.Get("data")
		objectValue = value.(orderedmap.OrderedMap)

		value, ok = objectValue.Get("items")
		items = value.([]any)

		value, ok = objectValue.Get("next")
		if value == nil {
			requestURL = ""
		} else {
			requestURL = value.(string)
		}
		results = append(results, items...)
	} else {
		message := ""
		value, ok = body.Get("message")
		if ok {
			message = value.(string)
		} else {
			message = "Unknown Error"
		}
		err = fmt.Errorf("[%v] %s", code, message)
		return nil, err
	}
	return results, nil
}

func DescribeResource(resourceType string, id string) (*orderedmap.OrderedMap, error) {
	var results orderedmap.OrderedMap
	var value any
	var ok bool
	requestURL := fmt.Sprintf(RESOURCE_DETAIL_URL, settings.GetServerHost(), resourceType, id)
	body, err := Request("GET", requestURL, nil)
	if err != nil {
		return nil, err
	}
	value, ok = body.Get("code")
	if !ok {
		return nil, fmt.Errorf("returned data is not json format")
	}

	code := int(value.(float64))
	if code == 200 {
		value, ok = body.Get("data")
		results = value.(orderedmap.OrderedMap)
	} else if code == 404 {
		err = fmt.Errorf("error: %s \"%v\" not found", resourceType, id)
		return nil, err
	} else {
		message := ""
		value, ok = body.Get("message")
		if ok {
			message = value.(string)
		} else {
			message = "Unknown Error"
		}
		err = fmt.Errorf("error: [%v] %s", code, message)
		return nil, err
	}
	return &results, nil
}

func CreateResource(resourceType string, inputData map[string]any) (int, error) {
	//var data map[string]any
	var result = -1
	var value any
	var ok bool
	requestURL := fmt.Sprintf(RESOURCE_URL, settings.GetServerHost(), resourceType)
	body, err := Request("POST", requestURL, inputData)
	if err != nil {
		return result, err
	}
	value, ok = body.Get("code")
	if !ok {
		return result, fmt.Errorf("returned data is not json format")
	}

	code := int(value.(float64))
	if code == 200 {
		value, ok = body.Get("data")
		id, _ := value.(*orderedmap.OrderedMap).Get("id")
		result = int(id.(float64))
	} else {
		message := ""
		value, ok = body.Get("message")
		if ok {
			message = value.(string)
		} else {
			message = "Unknown Error"
		}
		err = fmt.Errorf("[%v] %s", code, message)
		return result, err
	}
	return result, err
}

func DeleteReousrce(resourceType string, id string) (bool, error) {
	result := false
	requestURL := fmt.Sprintf(RESOURCE_DETAIL_URL, settings.GetServerHost(), resourceType, id)
	body, err := Request("DELETE", requestURL, nil)
	if err != nil {
		return result, err
	}

	value, ok := body.Get("code")
	if !ok {
		return result, fmt.Errorf("returned data is not json format")
	}

	code := int(value.(float64))
	if code == 200 {
		result = true
	} else if code == 404 {
		err = fmt.Errorf("error: %s \"%v\" not found", resourceType, id)
		return result, err
	} else {
		message := ""
		value, ok = body.Get("message")
		if ok {
			message = value.(string)
		} else {
			message = "Unknown Error"
		}
		err = fmt.Errorf("[%v] %s", code, message)
		return result, err
	}

	return result, nil
}
