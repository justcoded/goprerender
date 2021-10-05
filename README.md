GO Prerender
===========================

> Inspired by https://github.com/goprerender/prerender

GO Prerender is a dockerized server app written in Golang that uses Headless Chrome to render HTML and JS files out of any web page. 
The Prerender server listens for a http request, takes the URL and loads it in Headless Chrome, waits for the page to finish loading by waiting for the network to be idle, and then returns your content.

#### In memory cache

Caches pages in memory using prerender `storage` container. 

### Installation 

* clone repository
* run `make init`
* configure `.env` and `docker-compose.yml` to fit your needs or use the default.
* run `make up` - this will build and start docker containers


### Development

Run `make build-server` to build server Go application.  
Run `make build-storage` to build storage Go application.  
Run `make build` to build both Go applications server and storage.   

#### Usage

To run docker compose stack just run `make up`.

To prerender some URL you need to specify it as a query parameter, like this:

```
http://localhost:3000/render?url=https://www.example.com/
```
