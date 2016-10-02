//go:generate  /home/juno/neonworkspace/gowork/bin/statik -src=/home/juno/git/github.com/remotejob/weathermap_ang2_webpack/dist

package main // import "github.com/remotejob/weathermap_go"

import (
	"log"
	"net/http"

	// "github.com/gorilla/mux"

	"github.com/rakyll/statik/fs"
	"github.com/rs/cors"
	// "github.com/remotejob/goDevice"

	_ "github.com/remotejob/weathermap_go/statik"
)

// func timeoutHandler(h http.Handler) http.Handler {

// 	return http.TimeoutHandler(h, 1*time.Second, "timed out")
// }
//GetWeather return weather
func GetWeather(w http.ResponseWriter, r *http.Request) {

	log.Println("GetWeather")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"result": "lslslslssl"}`))

}

//Middleware to define mobile
// func Middleware(h http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// log.Println("middleware", r.URL)
// 		deviceType := goDevice.GetType(r)
// 		if deviceType == "Mobile" {

// 			context.Set(r, "device", "mobile")
// 			w.Header().Set("Mobile", "true")
// 		} else {
// 			// detectMobile.Detect(w, r)
// 			context.Set(r, "device", "desk")
// 			w.Header().Set("Mobile", "false")
// 		}
// 		// log.Println("remoteAddr", r.RemoteAddr)
// 		h.ServeHTTP(w, r)
// 	})
// }

// func MiddleOne(w http.ResponseWriter, r *http.Request) error {
// 	context.Set(r, "rideTheMiddlewareChain", "like a wave")
// 	return nil
// }

func main() {
	statikFS, err := fs.New()
	if err != nil {
		log.Fatalf(err.Error())
	}
	c := cors.New(cors.Options{
	// AllowedOrigins: []string{"*"},
	})
	// hd := http.HandleFunc("/api/weather", GetWeather)

	hd := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		path := r.RequestURI

		log.Println(path)
		if path == "/api/weather" {
			ip := r.Header.Get("ip")
			log.Println("ip", ip)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte("{\"hello\": \"world\"}"))
		} else {
			http.FileServer(statikFS).ServeHTTP(w, r)
		}

	})
	// router := httprouter.New()
	// r := mux.NewRouter()
	// // fs := http.FileServer(http.Dir("/home/juno/neonworkspace/gowork/src/github.com/remotejob/godocker/assets"))
	// fs := http.FileServer(http.Dir("assets"))

	// http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	// http.Handle("/assets", http.FileServer(http.Dir("/home/juno/neonworkspace/gowork/src/github.com/remotejob/godocker/assets")))

	// http.Handle("/", http.FileServer(statikFS))
	// http.Handle("/", restiful.Handle(MiddleOne, http.FileServer(statikFS)))
	// router.Handler("GET", "/", http.FileServer(statikFS))
	log.Println("Listening at port 8000!!")
	log.Fatal(http.ListenAndServe(":8000", c.Handler(hd)))
}
