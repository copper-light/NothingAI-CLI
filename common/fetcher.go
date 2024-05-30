package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"git.datacentric.kr/handh/NothingAI-CLI/common/utils"
	"git.datacentric.kr/handh/NothingAI-CLI/settings"
	"github.com/iancoleman/orderedmap"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func Request(method string, requestURL string, params map[string]any) (*orderedmap.OrderedMap, error) {
	var bodyStream *bytes.Buffer = nil
	//var result map[string]interface{} = nil
	client := http.Client{}

	if params != nil {
		if strings.ToUpper(method) == "GET" {
			query := url.Values{}
			for key, item := range params {
				value, ok := item.(string)
				if ok {
					query.Set(key, value)
				}
			}
			if len(query) > 0 {
				requestURL += "?" + query.Encode()
			}
		} else { // strings.ToUpper(method) == "POST"
			jsonData, err := json.Marshal(params)
			if err != nil {
				return nil, err
			}
			bodyStream = bytes.NewBuffer(jsonData)
		}
	}
	var req *http.Request
	var err error
	if bodyStream == nil {
		req, err = http.NewRequest(method, requestURL, nil)
	} else {
		req, err = http.NewRequest(method, requestURL, bodyStream)
	}

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	bodyData, err := responseParser(resp)
	return bodyData, err
}

func PostFiles(requestURL string, filePathList []string) (*orderedmap.OrderedMap, int, error) {
	result := 0
	client := http.Client{}
	var response *orderedmap.OrderedMap
	for i, filePath := range filePathList {
		fmt.Printf("[%d/%d] Send file : %s\n", i+1, len(filePathList), filePath)
		file, err := os.Open(filePath)
		if err != nil {
			err = fmt.Errorf("Error opening file %s: %s\n", filePath, err)
			return nil, result, err
		}
		defer file.Close()

		var requestBody bytes.Buffer
		writer := multipart.NewWriter(&requestBody)

		dirPath := filepath.Dir(filePath)
		if dirPath == "." {
			dirPath = "/"
		}
		part, err := writer.CreateFormFile(dirPath, filepath.Base(file.Name()))
		if err != nil {
			return nil, result, err
		}

		_, err = io.Copy(part, file)
		if err != nil {
			return nil, result, err
		}
		writer.Close()

		req, err := http.NewRequest("POST", requestURL, &requestBody)
		if err != nil {
			return nil, result, err
		}

		req.Header.Set("Content-Type", writer.FormDataContentType())

		resp, err := client.Do(req)
		if err != nil {
			return nil, result, err
		}
		bodyData, err := responseParser(resp)
		if err != nil {
			return bodyData, result, err
		}
		result++
	}
	return response, result, nil
}

func responseParser(response *http.Response) (*orderedmap.OrderedMap, error) {
	byteBody, err := io.ReadAll(response.Body)
	defer response.Body.Close()
	body := orderedmap.New()
	var bodyData orderedmap.OrderedMap
	err = json.Unmarshal([]byte(string(byteBody)), &body)
	if err != nil {
		switch response.StatusCode {
		case http.StatusNotFound:
			err = fmt.Errorf("error: resource not found")
		default:
			err = fmt.Errorf("error: unexpected status code: %d", response.StatusCode)
		}
		return nil, err
	}
	value, ok := body.Get("code")
	if !ok {
		return nil, fmt.Errorf("returned data is not json format")
	}

	code := int(value.(float64))
	if code == 200 {
		value, ok = body.Get("data")
		if ok {
			bodyData = value.(orderedmap.OrderedMap)
			err = nil
		}
	} else if code == 400 {
		value, ok = body.Get("detail")
		message := value.(string)
		err = fmt.Errorf(fmt.Sprintf("%s", message))
	} else if code == 404 {
		err = fmt.Errorf("resource not found")
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

	return &bodyData, err
}

func SendFiles(resourceType string, id string, path string) (int, error) {
	files := utils.GetFileList(path)
	if len(files) == 0 {
		return 0, fmt.Errorf("error: no files found in path: %s", path)
	}
	_, result, err := PostFiles(fmt.Sprintf(RESOURCE_FILES_URL, settings.GetServerHost(), resourceType, id), files)
	return result, err
}

func GetResources(resourceType string) ([]any, error) {
	var results []any
	var items []any
	requestURL := fmt.Sprintf(RESOURCE_URL, settings.GetServerHost(), resourceType)
	bodyData, err := Request("GET", requestURL, nil)
	if err != nil {
		return nil, err
	}
	if bodyData != nil {
		value, _ := bodyData.Get("items")
		items = value.([]any)
		results = append(results, items...)
		value, _ = bodyData.Get("next")
		if value != nil {
			requestURL = value.(string)
			nextResult, err := GetResources(requestURL)
			if err != nil {
				return nil, err
			} else {
				results = append(results, nextResult...)
			}
		}
	}
	return results, nil
}

func DescribeResource(resourceType string, id string) (*orderedmap.OrderedMap, error) {
	requestURL := fmt.Sprintf(RESOURCE_DETAIL_URL, settings.GetServerHost(), resourceType, id)
	bodyData, err := Request("GET", requestURL, nil)
	if err != nil {
		return nil, err
	}
	return bodyData, nil
}

func CreateResource(resourceType string, data map[string]any) (int, error) {
	var result = -1
	requestURL := fmt.Sprintf(RESOURCE_URL, settings.GetServerHost(), resourceType)
	bodyData, err := Request("POST", requestURL, data)
	if err != nil {
		return result, err
	}
	if bodyData != nil {
		id, _ := bodyData.Get("id")
		result = int(id.(float64))
	}
	return result, err
}

func DeleteResource(resourceType string, id string) (bool, error) {
	requestURL := fmt.Sprintf(RESOURCE_DETAIL_URL, settings.GetServerHost(), resourceType, id)
	_, err := Request("DELETE", requestURL, nil)
	if err != nil {
		return false, err
	}
	return true, nil
}

func ExecExperiment(id string) (bool, error) {
	requestURL := fmt.Sprintf(EXEC_PERIEMENT_URL, settings.GetServerHost(), id)
	_, err := Request("GET", requestURL, nil)
	if err != nil {
		return false, err
	}
	return true, nil
}

func LogTask(id string) ([]any, error) {
	requestURL := fmt.Sprintf(LOG_URL, settings.GetServerHost(), id)
	bodyData, err := Request("GET", requestURL, nil)
	var logs []any
	if err != nil {
		return nil, err
	} else {
		value, ok := bodyData.Get("items")
		if ok {
			logs = value.([]any)
		} else {
			logs = []any{}
		}
	}
	return logs, nil
}
