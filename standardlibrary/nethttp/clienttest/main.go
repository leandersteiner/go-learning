package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet,
		"http://localhost:8080/hello", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("X-Secret-Password", "GOPHER")

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("unexpected status: got %v", res.Status))
	}
	if err != nil {
		panic(err)
	}
	bytes, err := ioutil.ReadAll(res.Body)
	fmt.Printf("%+v\n", string(bytes))
}
