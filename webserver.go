package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"log"
)

var legalPage = map[string]bool {
	"/index.html": true,
	//"/stylesheet.css": true,
	"/": true,
}

func main() {
	fmt.Println("Listening on localhost:80")
	http.HandleFunc("/", MainServer)
	http.HandleFunc("/git/", GitServer)
//	log.Fatal(http.ListenAndServe(":8080", nil))
//	log.Fatal(http.ListenAndServe(":80", nil))
	log.Fatal(http.ListenAndServeTLS(":8443", "./certs/leaf.pem", "./certs/leaf.key", nil))
//	log.Fatal(http.ListenAndServeTLS(":443", "./certs/cert.pub", "priv.pem", nil))
}

func MainServer(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Tried to access ", r.URL.Path, " ...")
	var p string
	r1 := regexp.MustCompile("(?:^/(?P<path>(?:uio))?(?P<subpath>/?.*/)?(?P<site>.+\\.(?P<filetype>html|css|js)?)?$)?")
	match := r1.FindStringSubmatch(r.URL.Path)
	fullmatch, path, subpath, site, filetype := match[0], match[1], match[2], match[3], match[4]
	orderedPath := strings.Split(subpath, "/")
	fmt.Printf("Path: (%s), Subpath: (%s), Site: (%s) with (%s)\n", path, subpath, site, filetype)

	if fullmatch == "/" {
		p = "./content/index.html"
	} else if filetype == "css" {
		p = "./assets/" + site
	} else if legalPage[subpath + site] {
		p = "./content" + strings.Join(orderedPath[2:],"/")
	} else {
		p = "./noncontent/404.html"
		p = "." + r.URL.Path
		fmt.Println("Forbidden:", r.RemoteAddr)
	}
	fmt.Println(p)
	http.ServeFile(w, r, p)
}

func GitServer(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Git tried to access ", r.URL.Path, " ...")
	var p string
	orderedPath := strings.Split(r.URL.Path[1:], "/") // [git, f, <repo>, ...<path>]
	fmt.Println(orderedPath)

	if orderedPath[1] == "f" {
		p = "https://github.com/mazunki/" + orderedPath[2] + "/blob/master/" + strings.Join(orderedPath[3:],"/") + "?raw=True"
	} else if len(orderedPath)>2 {
		p = "https://github.com/mazunki/" + orderedPath[1] + "/blob/master/" + strings.Join(orderedPath[2:],"/")
	} else if len(orderedPath) == 2  {
		p = "https://github.com/mazunki/" + orderedPath[1]
	} else {
		p = "https://github.com/mazunki/"
	}
	fmt.Println("Redirecting to", p)
	http.Redirect(w, r, p, http.StatusSeeOther)
}
