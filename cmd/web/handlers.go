package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"golangify.com/snippetbox/pkg/models"
)

// home - Loads the home page and displays the last 10 notes.
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	s, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}
	app.render(w, r, "home.page.tmpl", &templateData{
		Snippets: s,
	})
}

// showSnippet - Displays data about the selected note
func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	s, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}
	app.render(w, r, "show.page.tmpl", &templateData{
		Snippet: s,
	})
}

// createSnippet - Displays the form to add and saves the entered data to the database
func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		app.render(w, r, "create.page.tmpl", &templateData{})
		return
	}

	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	title := r.FormValue("title")
	content := r.FormValue("content")
	expires := r.FormValue("expires")
	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
}

// deleteSnippet - Displays the deletion form and deletes the note by the desired id.
func (app *application) deleteSnippet(w http.ResponseWriter, r *http.Request) {

	app.render(w, r, "delete.page.tmpl", &templateData{})

	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	rowsAffected, err := app.snippets.Delete(id)
	if err != nil {
		app.serverError(w, err)
		return
	}
	if rowsAffected == 0 {
		app.notFound(w)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
