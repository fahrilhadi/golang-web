package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

type Pages struct {
	Title string
	Name string
}

func TemplateActionIf(writer http.ResponseWriter, request *http.Request)  {
	t := template.Must(template.ParseFiles("./templates/if.gohtml"))
	t.ExecuteTemplate(writer, "if.gohtml", Pages{
		Title: "Template Data Struct",
		Name: "Fahril",
	})
}

func TestTemplateActionIf(t *testing.T)  {
	request := httptest.NewRequest(http.MethodGet, "http://localhost/8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionIf(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateActionOperator(writer http.ResponseWriter, request *http.Request)  {
	t := template.Must(template.ParseFiles("./templates/comparator.gohtml"))
	t.ExecuteTemplate(writer, "comparator.gohtml", map[string]interface{}{
		"Title" : "Template Action Operator",
		"FinalValue" : 90,
	})
}

func TestTemplateActionOperator(t *testing.T)  {
	request := httptest.NewRequest(http.MethodGet, "http://localhost/8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionOperator(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateActionRange(writer http.ResponseWriter, request *http.Request)  {
	t := template.Must(template.ParseFiles("./templates/range.gohtml"))
	t.ExecuteTemplate(writer, "range.gohtml", map[string]interface{}{
		"Title" : "Template Action Range",
		"Hobbies" : []string{
			"Game", "Read", "Code",
		},
	})
}

func TestTemplateActionRange(t *testing.T)  {
	request := httptest.NewRequest(http.MethodGet, "http://localhost/8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionRange(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateActionWith(writer http.ResponseWriter, request *http.Request)  {
	t := template.Must(template.ParseFiles("./templates/address.gohtml"))
	t.ExecuteTemplate(writer, "address.gohtml", map[string]interface{}{
		"Title" : "Template Action With",
		"Name" : "Fahril",
		"Address" : map[string]interface{}{
			"Street" : "Jalan Belum Ada",
			"City" : "Pekanbaru",
		},
	})
}

func TestTemplateActionWith(t *testing.T)  {
	request := httptest.NewRequest(http.MethodGet, "http://localhost/8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionWith(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}