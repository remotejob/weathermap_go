package geocity

import (
	"log"
	"net"

	geoip2 "github.com/oschwald/geoip2-golang"
)

//ByIP get location by Ip address
func ByIP(db *geoip2.Reader, ipstr string) (latitude float64, longitude float64) {

	ip := net.ParseIP(ipstr)
	record, err := db.City(ip)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("Engish city name: %v\n", record.City.Names["en"])
	// fmt.Println("Country", record.Country.Names["en"])
	// fmt.Println("coordinats", record.Location.Latitude, record.Location.Longitude)

	return record.Location.Latitude, record.Location.Longitude
}
