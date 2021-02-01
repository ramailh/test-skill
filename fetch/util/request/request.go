package request

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type result struct {
	data []byte
}

func (res result) Bytes() []byte {
	return res.data
}

func (res result) Map() (map[string]interface{}, error) {
	data := make(map[string]interface{})
	if err := json.Unmarshal(res.data, &data); err != nil {
		log.Println(err)
		return nil, err
	}

	return data, nil
}

func (res result) ArrayOfMap() ([]map[string]interface{}, error) {
	var datas []map[string]interface{}
	if err := json.Unmarshal(res.data, &datas); err != nil {
		log.Println(err)
		return nil, err
	}

	return datas, nil
}

func (res result) ArrayOfInterface() ([]interface{}, error) {
	var datas []interface{}
	if err := json.Unmarshal(res.data, &datas); err != nil {
		log.Println(err)
		return nil, err
	}

	return datas, nil
}

func Post(url string, postData interface{}) (res result, err error) {
	dataJSON, err := json.Marshal(postData)
	if err != nil {
		log.Println(err)
		return res, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(dataJSON))
	if err != nil {
		log.Println(err)
		return res, err
	}
	defer resp.Body.Close()

	res.data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return res, err
	}

	return res, nil
}

func Get(url string) (res result, err error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return res, err
	}

	defer resp.Body.Close()

	res.data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return res, err
	}

	return res, nil
}
