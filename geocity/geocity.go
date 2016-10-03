package geocity

import (
	"log"
	"net"
	"strings"

	geoip2 "github.com/oschwald/geoip2-golang"
)

//ByIP get location by Ip address
func ByIP(db *geoip2.Reader, ipstr string) (latitude float64, longitude float64) {

	ip := net.ParseIP(strings.TrimSpace(ipstr))

	record, err := db.City(ip)
	if err != nil {
		log.Fatal(err)
	}

	return record.Location.Latitude, record.Location.Longitude
}
