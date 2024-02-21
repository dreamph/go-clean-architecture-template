package jet

import (
	"backend/internal/core/template"
	"bytes"
	"log"

	"fmt"

	"github.com/CloudyKit/jet/v6"
)

type jetHtmlTemplate struct {
	views *jet.Set
}

func NewJetHtmlTemplate(templateDir string, prd bool) template.HtmlTemplate {
	if prd {
		return &jetHtmlTemplate{
			views: jet.NewSet(
				jet.NewOSFileSystemLoader(templateDir),
			),
		}
	}
	return &jetHtmlTemplate{
		views: jet.NewSet(
			jet.NewOSFileSystemLoader(templateDir),
			jet.InDevelopmentMode(),
		),
	}
}

func (h *jetHtmlTemplate) Execute(fileName string, data map[string]interface{}) (string, error) {
	view, err := h.views.GetTemplate(fileName)
	if err != nil {
		return "", err
	}

	vars := make(jet.VarMap)
	for key, element := range data {
		vars.Set(key, element)
	}

	var resp bytes.Buffer
	err = view.Execute(&resp, vars, nil)
	if err != nil {
		return "", err
	}
	return resp.String(), nil
}

func (h *jetHtmlTemplate) Test(fileName string, data map[string]string) {
	view, err := h.views.GetTemplate("example-mail.jet")
	if err != nil {
		log.Println("Unexpected templates err:", err.Error())
	}

	for i := 1; i < 5000; i++ {
		var resp bytes.Buffer
		vars := make(jet.VarMap)
		vars.Set("title", fmt.Sprintf("title_%d", i))
		vars.Set("name", fmt.Sprintf("name_%d", i))
		err := view.Execute(&resp, vars, nil)
		if err != nil {
			return
		}
		fmt.Println(resp.String())
	}
}

/*
func (h *HtmlTemplate) EmailRegister( data map[string]string) {
	view, err := h.views.GetTemplate("confirm-email.jet")
	if err != nil {
		log.Println("Unexpected template err:", err.Error())
	}

	var resp bytes.Buffer
	vars := make(jet.VarMap)
	vars.Set("title", data["title"])
	vars.Set("email", data["email"])
	vars.Set("password",data["password"] )
	view.Execute(&resp, vars, nil)
	fmt.Println(resp.String())
}

*/
