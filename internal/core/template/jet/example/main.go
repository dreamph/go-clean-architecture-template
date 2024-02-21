package main

import (
	"backend/internal/core/template/jet"
)

type TestObj struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func main() {
	htmlTemplate := jet.NewJetHtmlTemplate("./config/email-template", false)

	var list []TestObj
	list = append(list, TestObj{ID: "1", Name: "name1"})
	list = append(list, TestObj{ID: "2", Name: "name2"})

	data := map[string]interface{}{
		"name": "phol",
		"list": list,
		"num":  1,
	}

	result, err := htmlTemplate.Execute("/example-list.jet", data)
	if err != nil {
		panic(err)
	}
	println(result)
}
