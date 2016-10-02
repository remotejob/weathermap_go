package geocity

import (
	"log"
	"testing"

	"github.com/oschwald/geoip2-golang"
)

func TestByIp(t *testing.T) {
	db, err := geoip2.Open("../geolite/GeoLite2-City.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	latitude, longitude := ByIP(db, "188.238.62.67")

	log.Println(latitude, longitude)

}
