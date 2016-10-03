FROM scratch
EXPOSE 8080

COPY server /
COPY geolite/GeoLite2-City.mmdb geolite/
CMD ["/server"]
