package handler

import (
	"github.com/gorilla/mux"
	"github.com/twn39/butterfly/github"
	"github.com/twn39/butterfly/middleware"
	"html/template"
	"net/http"
)

func DetailHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	githubClient := r.Context().Value("github").(*github.Client)
	settings := r.Context().Value("settings").(*middleware.Settings)

	issue, err := githubClient.GetSingleIssue(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
	} else {
		tmp, err := template.ParseFiles("views/layout.html",
			"views/head.html",
			"views/header.html",
			"views/aside.html",
			"views/detail.html")
		if err != nil {
			panic(err)
		}

		err = tmp.Execute(w, map[string]interface{}{
			"Title":   settings.Site.Title,
			"Issue":   issue,
			"Content": template.HTML(issue.Body),
		})
	}
}
