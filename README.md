# weathermap_go
--

    	Package weathermap_go simple web site

        Live DEMO http://104.155.23.78:8080
        Angular2 code https://github.com/remotejob/weathermap_ang2_webpack

    	Use:

    	General:
    		small Docker container image
    			can be demployed on
                    Amazon EC2 Container Service
                    Google Container Engine
                    Azure Container Service etc..

    	Geo Location:
            coordinate by
                free GeoLite2-City DB

    	Weathermap_go:
    		by simple https://openweathermap.org/  API service

    	Site created by Angular 2 framework:
    		code splited and organized by Webpack

    	Probably mostly intresting part:
    		all web static web file incorporate is single file server.go
                how it possible --> https://github.com/rakyll/statik/

    	Size:
    		Doker image after deployment on Google Container Registry
                size without GeoLite2-City.mmdn 3 MB!!
                total size 35 MB included GeoLite2-City.mmdn

    	Docker and Golang:
            Golang: definitely winning in relation to Docker
