package server

import (
	"fmt"
	"net/http"
)

func countHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		w.WriteHeader(http.StatusAccepted)
		fmt.Fprint(w, "fetching data from source")
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(htmlErrorBody))
	}
}

func Run() error {
	server := http.NewServeMux()

	server.HandleFunc("/count", countHandler)
	return http.ListenAndServe(":5000", server)
}

var htmlErrorBody = `
<!doctype html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <title>Your Friendly ERROR</title>
  </head>
  <body>
      <h3 style="color: red;">
      Wrong Method <br/>
      <code>/count</code> accepts only <code>POST</code>
    </h3>
      <br/> 
      <h2>Example:</h2>
      <br/> 
      <code>
            $ curl -XPOST -H "Content-type: application/json" -d '{"start_date": "2024-01-01", "end_date": "2024-03-30"}' 'localhost:5000/count'
      </code>
      <br />
      <code>
            $ echo '{ "start_date": "2024-01-01", "end_date": "2024-03-20" }' | restish post :5000/count
      </code>
  </body>
</html>
      `
