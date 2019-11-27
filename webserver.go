package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"log"
	"flag"
)


func main() {
	portNr := flag.Int("http", 0, "http port")
	portNrTLS := flag.Int("https", 8443, "https port")
	flag.Parse()	

	http.HandleFunc("/", MainServer)
	http.HandleFunc("/git/", GitServer)

	fmt.Println(*portNr, ",", *portNrTLS)

	if *portNr != 0 {
		fmt.Printf("Listening on localhost:%d\n", *portNr)
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *portNr), nil))
	} else {
		fmt.Println("Listening on secure localhost:%d\n", *portNrTLS)
		log.Fatal(http.ListenAndServeTLS(fmt.Sprintf(":%d", *portNrTLS), "./sslforfree/certificate.crt", "./sslforfree/private.key", nil))
	}
}

func MainServer(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Tried to access ", r.URL.Path, " ...")
	var p string
	r1 := regexp.MustCompile("(?:^/(?P<path>(?:uio))?(?P<subpath>/?.*/)?(?P<site>.+\\.(?P<filetype>html|css|js)?)?$)?")
	match := r1.FindStringSubmatch(r.URL.Path)
	fullmatch, path, subpath, site, filetype := match[0], match[1], match[2], match[3], match[4]
	fmt.Printf("Path: (%s), Subpath: (%s), Site: (%s) with (%s)\n", path, subpath, site, filetype)
	if site == "" {
		site = "index.html"
	}

	if fullmatch == "/" {
		p = "./content/index.html"
	} else if filetype == "css" {
		p = "./assets/" + site
	} else {
		p = fmt.Sprintf("./content/%s%s", subpath, site)
	}
	fmt.Println("Serving:", p)
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
