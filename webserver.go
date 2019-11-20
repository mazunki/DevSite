package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

var legalPage = map[string]bool {
	"/index.html": true,
	"/stylesheet.css": true,
	"/": true,
}

func main() {
	fmt.Println("Listening on localhost:80")
	http.HandleFunc("/", MainServer)
	http.ListenAndServe(":8080", nil)
}

func MainServer(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Tried to access ", r.URL.Path, " ...")
	var p string
	r1, _ := regexp.Compile("(?:^/(?P<path>(?:uio|git))?(?P<subpath>/?.*/)?(?P<site>.+(?P<filetype>\\.html|\\.css|\\.js)?)?$)")
	match := r1.FindStringSubmatch(r.URL.Path)
	fullmatch, path, subpath, site, filetype := match[0], match[1], match[2], match[3], match[4]
	orderedPath := strings.Split(subpath, "/")
	fmt.Printf("Path: (%s), Subpath: (%s), Site: (%s)\n", path, subpath, site)

	if fullmatch == "/" {
		p = "./content/index.html"
		http.ServeFile(w, r, p)
	} else if filetype == ".css" {
		p = "./content/" + site
		http.ServeFile(w, r, p)
	} else if path == "git" {
		fmt.Println(orderedPath)
		fmt.Printf("(%s)", orderedPath[1])
		if orderedPath[1] == "f" {
			p = "https://github.com/mazunki/" + orderedPath[2] + "/blob/master/" + strings.Join(orderedPath[3:],"/" + "?raw=True")
		} else {
			p = "https://github.com/mazunki/"
		}
		http.Redirect(w, r, p, http.StatusSeeOther)
	} else if legalPage[subpath + site] {
		p = "./content" + site
		http.ServeFile(w, r, p)
	} else {
		p = "./noncontent/404.html"
		fmt.Println("Forbidden:", r.RemoteAddr)
		http.ServeFile(w, r, p)
	}
}
