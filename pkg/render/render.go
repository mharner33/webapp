package render

import (
	"bytes"
	//"http/template"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

// Renders the template with name templ
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	//Create a template cache so we don't need to read from disk
	tc, err := createTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal(ok)
	}

	buff := new(bytes.Buffer)
	err = t.Execute(buff, nil)
	if err != nil {
		log.Println(err)
	}

	_, err = buff.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	//get all of the *.tmpl files from templates folder
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	//range through the .tmpl files
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*layout.tmpl")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts

	}
	return myCache, nil
}
