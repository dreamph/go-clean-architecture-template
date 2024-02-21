package template

type HtmlTemplate interface {
	Execute(fileName string, data map[string]interface{}) (string, error)
}
