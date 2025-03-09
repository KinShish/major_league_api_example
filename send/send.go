package send

import (
	"fmt"
	"net/http"
)

func RequestGoogle(method string, url string, body string, count int) {
	client := &http.Client{}
	req, err := http.NewRequest(method, site+url, nil) //"GET", "https://google.com", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", "Bearer ...")
	resp, err := client.Do(req)
	if err != nil {

		fmt.Println(err)
		return
	}
	//resp, err := http.Get("https://google.com")
	//http.Post("http://example.com/upload", "image/jpeg", &buf)
	//http.PostForm("http://example.com/form",	url.Values{"key": {"Value"}, "id": {"123"}})
	defer resp.Body.Close()
	for true {

		bs := make([]byte, 1014)
		n, err := resp.Body.Read(bs)
		fmt.Println(string(bs[:n]))

		if n == 0 || err != nil {
			break
		}
	}
}
