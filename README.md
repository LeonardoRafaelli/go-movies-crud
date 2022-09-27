This api is just a crud template of movies, where I can register new movies and then manage them.
There is no HTML ou CSS on this application. Just a running server on localhost:8080.

Has been used on this GO api:

* `encoding/json` to encode and decode data that comes and go from front-end
* `math/rand` to give each movie a random ID
* `net/http` to create a connection with the web browser
* `strconv` to convert the integer ID of math/rand into a string
* `github.com/gorilla/mux` to create routes and get the params from the url
