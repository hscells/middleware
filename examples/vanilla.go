/*
Example for How to use with standard net/http package.
*/
package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/jeevatkm/middleware"
)

var htmlString = `
		<!DOCTYPE html>
		<html>
		<head>

		<style>body {
		    background-color: #d0e4fe;
		}

		h1 {
		    color: orange;
		    text-align: center;
		}

		p {
		    font-family: "Times New Roman";
		    font-size: 20px;
		}
		</style>
		</head>

		<body>

		<h1>Minify #NAME# Example</h1>

		<p>Example #NAME# paragraph.</p>
		<div id="example"></div>

		<script>
			var day;
			switch (new Date().getDay()) {
			    case 0:
			        day = "Sunday";
			        break;
			    case 1:
			        day = "Monday";
			        break;
			    case 2:
			        day = "Tuesday";
			        break;
			    case 3:
			        day = "Wednesday";
			        break;
			    case 4:
			        day = "Thursday";
			        break;
			    case 5:
			        day = "Friday";
			        break;
			    case  6:
			        day = "Saturday";
			        break;
			}
			document.getElementById("example").innerHTML = "Today is " + day;
			</script>

		</body>
		</html>`

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(strings.Replace(htmlString, "#NAME#", "Vanilla", -1)))
}

func main() {
	homeHandler := http.HandlerFunc(home)

	// Adding Minify middleware
	// Note: If you use any Gzip middleware, add Minify middleware after that
	http.Handle("/", middleware.Minify(homeHandler))

	log.Println("Starting server at localhost:8000")
	http.ListenAndServe(":8000", nil)
}
