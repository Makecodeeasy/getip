package pkg

import (
	"io/ioutil"
	"net/http"
)

func GetIp() (string, error) {
	client := &http.Client{}
	url := "http://ip.sb"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("User-Agent", "curl/7.29.0")
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(b))
	return string(b), nil
}
