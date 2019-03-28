package main

import (
    "fmt"
    "github.com/julienschmidt/httprouter"
    "net/http"
    // "log"
		"encoding/json"
		"github.com/gocolly/colly"
)

type Result struct {
	Nb_comments string
	Rating string
}


func getResultFromAddress(url_name string) Result {

	var nb_comments = "-1"
	var rating ="-1"

	c := colly.NewCollector(
			// colly.AllowedDomains("https://www.tripadvisor.fr"),
	)

	c.OnError(func(_ *colly.Response, err error) {
	    fmt.Println("Something went wrong:", err)
	})

	// Find number of comments
	c.OnHTML("span.reviews_header_count", func(e *colly.HTMLElement) {
		nb_comments = e.Text
		fmt.Printf("reviews_header_count found : ")
		fmt.Printf("%+v\n", e.Text)
	})

	// Find restaurants rating
	 c.OnHTML("span.restaurants-detail-overview-cards-RatingsOverviewCard__overallRating--nohTl", func(e *colly.HTMLElement) {
	 	fmt.Println("restaurants-detail found ",e.Text)
		rating = e.Text
	})

	// Find rating
	c.OnHTML("span.ui_bubble_rating bubble_45::after", func(e *colly.HTMLElement) {
	 fmt.Println("ui buble rating found")
 })

	// Log
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

// URL test:
// https://www.tripadvisor.fr/Restaurant_Review-g60763-d1236281-Reviews-Club_A_Steakhouse-New_York_City_New_York.html
	c.Visit(url_name)
	// c.Visit("https://www.tripadvisor.fr/Restaurant_Review-g60763-d1236281-Reviews-Club_A_Steakhouse-New_York_City_New_York.html")

  fmt.Printf("address : ")
  fmt.Println(url_name)

	// Build res
	var res = Result{nb_comments,rating}
  return res
}


func Generate(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
    url_name := r.URL.Query().Get("url_name")
    var res =getResultFromAddress(url_name)
    data, _ := json.Marshal(res)

    // Disp res
    // Write content-type, statuscode, payload
    // w.Header().Set("Content-Type", "text/plain")
    if origin := r.Header.Get("Origin"); origin != "" {
             w.Header().Set("Access-Control-Allow-Origin", origin)
         }
     w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
     w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
     w.Header().Set("Access-Control-Allow-Credentials", "true")

    w.Header().Set("text.plain", "application/json; charset=UTF-8")
    w.WriteHeader(201)
    // w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8083")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Add("Access-Control-Allow-Methods", "PUT")
    w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
    w.Header().Add("Access-Control-Allow-Methods", "PUT")
    // w.Header().Set("Access-Control-Allow-Headers","Content-Type,access-control-allow-origin, access-control-allow-headers")
    fmt.Fprintf(w, "%s", data)
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    fmt.Fprintf(w, "Server is up !\n")
}


func main() {
    router := httprouter.New()
    // router.GET("/", )
    router.GET("/ping", Hello)
		router.GET("/getTA", Generate)

    http.ListenAndServe(":8083", router)
}
