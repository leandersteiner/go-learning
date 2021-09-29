package main

import (
	"context"
	"io/ioutil"
	"learning/context/values/tracker"
	"net/http"
)

type Logger interface {
	Log(context.Context, string)
}

type RequestDecorator func(*http.Request) *http.Request

type BusinessLogic struct {
	RequestDecorator RequestDecorator
	Logger           Logger
	Remote           string
}

func (bl BusinessLogic) businessLogic(ctx context.Context, user string, data string) (string, error) {
	bl.Logger.Log(ctx, "starting businessLogic for "+user+" with "+data)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, bl.Remote+"?query="+data, nil)
	if err != nil {
		bl.Logger.Log(ctx, "error building remote request:"+err.Error())
		return "", err
	}
	req = bl.RequestDecorator(req)
	resp, err := http.DefaultClient.Do(req)
	// processing continues
	defer resp.Body.Close()
	answer, err := ioutil.ReadAll(resp.Body)
	return string(answer), err
}

func main() {
	bl := BusinessLogic{
		RequestDecorator: tracker.Request,
		Logger:           tracker.Logger{},
		Remote:           "http://www.example.com/query",
	}
	bl.businessLogic(context.Background(), "ls", "data")
}
