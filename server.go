//go:generate  /home/juno/neonworkspace/gowork/bin/statik -src=/home/juno/git/github.com/remotejob/weathermap_ang2_webpack/dist

package main // import "github.com/remotejob/weathermap_go"

import (
	"log"
	"net/http"

	"github.com/rakyll/statik/fs"
	_ "github.com/remotejob/weathermap_go/statik"
)

func main() {
	statikFS, err := fs.New()
	if err != nil {
		log.Fatalf(err.Error())
	}

	// // fs := http.FileServer(http.Dir("/home/juno/neonworkspace/gowork/src/github.com/remotejob/godocker/assets"))
	// fs := http.FileServer(http.Dir("assets"))

	// http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	// http.Handle("/assets", http.FileServer(http.Dir("/home/juno/neonworkspace/gowork/src/github.com/remotejob/godocker/assets")))
	http.Handle("/", http.FileServer(statikFS))
	log.Println("Listening at port 8080!!")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
