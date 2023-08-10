package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"golem/component/go_golem_demo"
	"golem/component/roundtrip"

	"net/http"
)

type RequestBody struct {
	CurrentTotal uint64
}

type ResponseBody struct {
	Message string
}

func init() {
	a := GoGolemDemoImpl{}
	go_golem_demo.SetExportsGolemComponentApi(a)
}

// total State can be stored in global variables
var total uint64

type GoGolemDemoImpl struct {
	total uint64
}

// Implementation of the exported interface


func (e GoGolemDemoImpl) Add(value uint64) {
	total += value
}

func (e GoGolemDemoImpl) Get() uint64 {
	return total
}

func (e GoGolemDemoImpl) Publish() go_golem_demo.Result[struct{}, string] {
	http.DefaultClient.Transport = roundtrip.WasiHttpTransport{}
	var result go_golem_demo.Result[struct{}, string]

	postBody, _ := json.Marshal(RequestBody{
		CurrentTotal: total,
	})
	resp, err := http.Post("http://localhost:9999/post-example", "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		result.SetErr(fmt.Sprintln(err))
		return result
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		result.SetErr(fmt.Sprintln(err))
		return result
	}

	var response ResponseBody
	err = json.Unmarshal(body, &response)
	if err != nil {
		result.SetErr(fmt.Sprintln(err))
		return result
	}

	fmt.Println(response.Message)

	result.Set(struct{}{})
	return result
}

func (e GoGolemDemoImpl) Pause() {
	promise := go_golem_demo.GolemApiHostGolemCreatePromise()
	go_golem_demo.GolemApiHostGolemAwaitPromise(promise)
}

func main() {
}
