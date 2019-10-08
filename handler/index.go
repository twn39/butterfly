package handler

import (
	"encoding/json"
	"github.com/twn39/butterfly/cache"
	"github.com/twn39/butterfly/github"
	"github.com/twn39/butterfly/middleware"
	"html/template"
	"net/http"
	"time"
	//"time"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	settings := r.Context().Value("settings").(*middleware.Settings)
	githubClient := r.Context().Value("github").(*github.Client)
	cacheStorage := r.Context().Value("cache").(cache.Cache)

	// get issues
	data := cacheStorage.Get("issues", func(item *cache.Item) string {
		item.SetExpire(600 * time.Second)
		issues := githubClient.GetIssues("created", "desc")
		return issues
	})

	issues := new([]github.SingleIssueResult)
	err := json.Unmarshal([]byte(data), issues)
	if err != nil {
		panic(err)
	}
	// get users
	user := cacheStorage.Get("users", func(item *cache.Item) string {
		item.SetExpire(600 * time.Second)
		result := githubClient.GetUser()
		return result
	})

	userResult := new(github.UserResult)
	err = json.Unmarshal([]byte(user), userResult)
	if err != nil {
		panic(err)
	}

	tmp, err := template.ParseFiles("views/layout.html",
		"views/head.html",
		"views/index.html",
		"views/aside.html",
		"views/header.html")

	if err != nil {
		panic(err)
	}

	err = tmp.Execute(w, map[string]interface{}{
		"Title":  settings.Site.Title,
		"Issues": issues,
		"User":   userResult,
	})
}
