package main

import (
    "fmt"
    "github.com/julienschmidt/httprouter"
    "net/http"
)


func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func main() {
    router := httprouter.New()
    // router.GET("/", Index)
    router.GET("/:name", Hello)

    http.ListenAndServe(":8083", router)
}
