package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"
)

const sampleJson string = `
    {
        "user": {
        "name": "山田 一郎",
        "age": 25,
        "gender": "男"
        }
    }
`

func main() {
	api := rest.NewApi()
    api.Use(rest.DefaultDevStack...)  
    router, err := rest.MakeRouter(
        rest.Get("/api/v1/user", func(w rest.ResponseWriter, r *rest.Request) {
            w.WriteJson(sampleJson)
        }),
    )
    if err != nil {
        log.Fatal(err)
    }
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":5000", api.MakeHandler()))
}