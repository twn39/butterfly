package handler

import (
	"encoding/json"
	"fmt"
	"github.com/twn39/butterfly/cache"
	"github.com/twn39/butterfly/github"
	"github.com/twn39/butterfly/middleware"
	"html/template"
	"net/http"
	"strconv"
	"time"
	//"time"
)

func slice(data string) string {
	return data
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	settings := r.Context().Value("settings").(*middleware.Settings)
	githubClient := r.Context().Value("github").(*github.Client)
	cacheStorage := r.Context().Value("cache").(cache.Cache)
	qs := r.URL.Query()
	var page int
	if qs["page"] == nil {
		page = 1
	} else {
		page, _ = strconv.Atoi(qs["page"][0])
	}

	// get issues
	data := cacheStorage.Get("issues.page."+string(page), func(item *cache.Item) string {
		item.SetExpire(600 * time.Second)
		issues := githubClient.GetIssues(page, "created", "desc")
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

	funcMap := template.FuncMap{
		"slice": func(data string) string {
			strLen := len([]rune(data))
			if strLen > 320 {
				return string([]rune(data)[0:320])
			}

			return data
		},
		"div": func(num int) int {
			return num - 1
		},
		"add": func(num int) int {
			return num + 1
		},
	}
	tmpl, err := template.New("layout.html").Funcs(funcMap).ParseFiles("views/layout.html",
		"views/head.html",
		"views/aside.html",
		"views/index.html",
		"views/header.html")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d", page)

	err = tmpl.ExecuteTemplate(w, "layout.html", map[string]interface{}{
		"Title":  settings.Site.Title,
		"Issues": issues,
		"User":   userResult,
		"Page":   page,
	})
}
