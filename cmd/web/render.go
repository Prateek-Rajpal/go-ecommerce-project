package main

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

type templateData struct {
	StringMap       map[string]string 
	IntMap          map[string]int
	FloatMap        map[string]float32
	Data            map[string]interface{} // any kind of data 
	CSRFToken       string
	Flash           string // message that displayed once and then goes away
	Warning         string 
	Error           string
	IsAuthenticated int // is user authenticated
	Api             string // route to our api
	CSSVersion      string
}

var functions = template.FuncMap{}

//go:embed templates
var templateFs embed.FS // go embed is a Go directive that tells the Go toolchain to embed the files and directories within the "templates" directory into the variable "templateFs." This allows us to compile my application, including all of its associated templates into a single binary.

func (app *application) addDefaultData(td *templateData, r *http.Request) *templateData {
	return td
}

func (app *application) renderTemplate(w http.ResponseWriter, r *http.Request, page string, td *templateData, partials ...string) error {
	var t *template.Template
	var err error
	templateToRender := fmt.Sprintf("templates/%s.page.gohtml", page)
	
	// checking whether template is available in cache
	_, templateInMap := app.templateCache[templateToRender]

	if app.config.env == "prod" && templateInMap {
		t = app.templateCache[templateToRender]
	} else { // in other environments don't use templateCache
		t, err = app.parseTemplate(partials, page, templateToRender)
		if err != nil {
			app.errorLog.Println(err)
			return err
		}
	}

	if td == nil {
		td = &templateData{}
	}
	td = app.addDefaultData(td, r)

	err = t.Execute(w, td)
	if err != nil {
		app.errorLog.Println(err)
		return err
	}
	return nil
}

func (app *application) parseTemplate(partials []string, page, templateToRender string) (*template.Template, error) {

	var t *template.Template
	var err error

	//build partials
	if len(partials) > 0 {
		for i, x := range partials {
			partials[i] = fmt.Sprintf("templates/%s.partial.gohtml", x)

		}
	}
	if len(partials) > 0 {
		t, err = template.New(fmt.Sprintf("%s.page.gohtml", page)).Funcs(functions).ParseFS(templateFs, "templates/base.layout.gohtml", strings.Join(partials, ","), templateToRender)
	} else {
		t, err = template.New(fmt.Sprintf("%s.page.gohtml", page)).Funcs(functions).ParseFS(templateFs, "templates/base.layout.gohtml", templateToRender)
	}
	if err != nil {
		app.errorLog.Println(err)
		return nil, err
	}
	app.templateCache[templateToRender] = t
	return t, nil

}
