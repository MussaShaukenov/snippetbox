package main

import (
<<<<<<< HEAD
<<<<<<< HEAD
	"errors"
	"fmt"
	"github.com/MussaShaukenov/snippetbox/internal/models"
=======
	"fmt"
	"html/template"
>>>>>>> parent of ab1feef (chapter 5: database-driven responses)
=======
	"database/sql"
	"errors"
	"fmt"
>>>>>>> parent of 7edf881 (chapter 5: dynamic html templates)
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
<<<<<<< HEAD

<<<<<<< HEAD
=======
>>>>>>> parent of 7edf881 (chapter 5: dynamic html templates)
	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

<<<<<<< HEAD
	app.render(w, http.StatusOK, "home.tmpl", data)
=======
	// Initialize a slice containing the paths to the two files. It's important
	// to note that the file containing our base template must be the *first*
	// file in the slice.
	files := []string{
		"./ui/html/base.html",
		"./ui/html/partials/nav.html",
		"./ui/html/pages/home.html",
	}
=======
	for _, snippet := range snippets {
		fmt.Fprintf(w, "%+v\n", snippet)
	}

	// Initialize a slice containing the paths to the two files. It's important
	// to note that the file containing our base template must be the *first*
	// file in the slice.
	//files := []string{
	//	"./ui/html/base.html",
	//	"./ui/html/partials/nav.html",
	//	"./ui/html/pages/home.html",
	//}
>>>>>>> parent of 7edf881 (chapter 5: dynamic html templates)

	// Use the template.ParseFiles() function to read the files and store the
	// templates in a template set. Notice that we can pass the slice of file
	// paths as a variadic parameter?
<<<<<<< HEAD
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err) // use serverError() helper
		return
	}
	// Use the ExecuteTemplate() method to write the content of the "base"
	// template as the response body.
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, err) // use serverError() helper
	}
>>>>>>> parent of ab1feef (chapter 5: database-driven responses)
=======
	//ts, err := template.ParseFiles(files...)
	//if err != nil {
	//	app.serverError(w, err) // use serverError() helper
	//	return
	//}
	// Use the ExecuteTemplate() method to write the content of the "base"
	// template as the response body.
	//err = ts.ExecuteTemplate(w, "base", nil)
	//if err != nil {
	//	app.serverError(w, err) // use serverError() helper
	//}
>>>>>>> parent of 7edf881 (chapter 5: dynamic html templates)
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w) // use notFound() helper
		return
	}
<<<<<<< HEAD
<<<<<<< HEAD

=======
	// Use the SnippetModel object's Get method to retrieve the data for a
	// specific record based on its ID. If no matching record is found,
	// return a 404 Not Found response.
>>>>>>> parent of 7edf881 (chapter 5: dynamic html templates)
	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

<<<<<<< HEAD
	data := app.newTemplateData(r)
	data.Snippet = snippet

	app.render(w, http.StatusOK, "view.tmpl", data)
=======

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
>>>>>>> parent of ab1feef (chapter 5: database-driven responses)
=======
	// Write the snippet data as a plain HTTP response body
	fmt.Fprintf(w, "%+v", snippet)
>>>>>>> parent of 7edf881 (chapter 5: dynamic html templates)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Create a new snippet..."))
}
