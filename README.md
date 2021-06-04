# iTunes Ranking Service

Go API Server REST HTTP service to get land specific rank of top entries.

## Usage

### Build and run the server

```shell
bin/build-as-docker.sh
```

> Please install [docker-compose](https://docs.docker.com/compose/) to build and manage the service containerized in docker.

### Running the server

To run the server, follow these simple steps:

```shell
go run main.go
```

## TODO:

* [x] Erstelle einen [HTTP Server](main.go) mit den folgenden beiden Endpunkten in Golang:
* [x] Der Endpunkt [/ENV ](api/routers.go#L45) 
    * [x] [lädt beim Aufruf die Liste der 100 populärsten Podcasts](itunes/service.go#L35)  
    * [x] für [5 verschiedene Länder](itunes/country.go#L5)
    * [x] und [schreibt diese in geeigneter Form in eine Datenbank](itunes/service.go#L95).
    * [x] Die 100 populärsten Podcasts lassen sich von iTunes über die folgende URL als JSON
      abrufen: `https://itunes.apple.com/{cc}/rss/topaudiopodcasts/limit=100/json` (`{cc}` = Country Code)
    * [x] Die iTunesID ist im JSON unter dem Pfad [feed](itunes/feed.go) > [entry](itunes/entry.go) > [id > attribute > im:id](itunes/id.go)` zu finden
    * [x] Die 5 Länder mit den dazugehörigen [Country-Codes](itunes/country.go#L5) sind:
        * [x] USA (us)
        * [x] Deutschland (de)
        * [x] Frankreich (fr)
        * [x] Italien (it)
        * [x] und Spanien (es)
    * [x] Die Datenbank sollte entweder
      * [x] MySQL, 
      * [x] Postgres
      * [x] oder [SQLite](itunes/service.go#L24) sein
* [x] Der Endpunkt [/rank](api/routers.go#L38) 
  * [x] [Erfordert](api/api.go#L54) die [iTunesID als Eingabe-Parameter](api/routers.go#L38) 
  * [x] und [liefert eine Liste von Land/Position-Tupeln](api/api.go#L91) zurück
  * [x] [In dem Tupel ist das Land](itunes/service.go#L124) der `Country Code` 
  * [x] und die `Position` ist die [Position](itunes/rank-result.go) innerhalb der Top-100
  * [x] Die Server-Antwort [soll mithilfe der Daten aus der Datenbank](itunes/service.go#L116) erstellt werden
  * [x] Ist die `iTunesID` unbekannt, so [soll eine leere Liste zurückgegeben werden](api/api.go#L54)
  * [x] Die Antwort soll [als JSON](api/api.go#L46) zurückgegeben werden
* [x] Der Server soll als [Docker-Image](bin/build-as-docker.sh) ausgeliefert werden, d.h. ein entsprechendes [Dockerfile](Dockerfile) soll vorhanden sein. 
* [x] Die Integration in [eine CI/CD](.github/workflows/go.yml) ist nicht notwendig.
