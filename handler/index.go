package handler

import (
	//"github.com/twn39/butterfly/cache"
	"github.com/twn39/butterfly/github"
	"github.com/twn39/butterfly/middleware"
	"html/template"
	"net/http"
	//"time"
)

type ResponseData struct {
	Title  string
	Issues []github.RepoIssueResult
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	settings := r.Context().Value("settings").(*middleware.Settings)
	githubClient := r.Context().Value("github").(*github.Client)
	//cacheStorage := r.Context().Value("cache").(cache.Cache)

	issues := githubClient.GetIssues("created", "desc")

	//data := cacheStorage.Get("issues", func(item *cache.Item) string {
	//	item.SetExpire(600 * time.Second)
	//	return "twn39@163.com"
	//})

	//fmt.Printf("%s", data)

	tmp, err := template.ParseFiles("views/layout.html", "views/head.html", "views/index.html")
	if err != nil {
		panic(err)
	}

	err = tmp.Execute(w, ResponseData{
		Title:  settings.Site.Title,
		Issues: issues,
	})
}
