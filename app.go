package main

import (
    "fmt"
    "net/http"
    "text/template"
    "strings"
    "io/ioutil"
    "io"
    "encoding/json"
    "net/url"
    "log"
    "strconv"
    "crypto/md5"
)

type Person struct {
    Name string
    Points int
    Slug string
}

func handler(w http.ResponseWriter, r *http.Request) {

    // TODO: Replace _'s with err's
    // TODO: Have index serve a pre-generated HTML page

    peopleJson, _ := ioutil.ReadFile("people.json")

    var people []Person
    _ = json.Unmarshal(peopleJson, &people)

    query_params, _ := url.ParseQuery(r.URL.RawQuery)
    key := query_params["key"]
    var digest string
    if key != nil {
        h := md5.New()
        io.WriteString(h, key[0])
        digest = fmt.Sprintf("%x", h.Sum(nil))
    } else {
        digest = "" 
    }

    // If the supplied key matches the hash of a secret hash
    if digest == "a7360762a2bfffd1af5b6b30affe22e9" {
        t, err := template.ParseFiles("admin.html")
        if err != nil {
            log.Fatal(err)
        }
        t.Execute(w, people)
    } else {
        t, err := template.ParseFiles("index.html")
        if err != nil {
            log.Fatal(err)
        }
        t.Execute(w, people)
    }

}

func pointsHandler(w http.ResponseWriter, r *http.Request) {

    params, _ := url.ParseQuery(r.URL.RawQuery)
    param_key    := params["key"]
    param_slug  := params["slug"]
    param_points := params["points"]

    points := param_points[0]
    slug := param_slug[0]
    key := param_key[0]

    // Check key
    h := md5.New()
    io.WriteString(h, key)
    digest := fmt.Sprintf("%x", h.Sum(nil))

    if digest == "a7360762a2bfffd1af5b6b30affe22e9" && points != "" && slug != "" {

        // Open JSON
        peopleJson, _ := ioutil.ReadFile("people.json")

        // Unmarshal JSON
        var people []Person
        _ = json.Unmarshal(peopleJson, &people)

        for k, p := range people {
            if p.Slug == slug {
                points_int64, _ := strconv.ParseInt(points, 0, 32)
                var points_int int
                points_int = int(points_int64)
                people[k].Points += points_int

                // Marshal JSON
                newPeopleJson, err := json.Marshal(people)
                log.Println(err)

                // Save JSON
                err = ioutil.WriteFile("people.json", newPeopleJson, 777)
                log.Println(err)

                // Redirect to /
                s := []string{"/?key=", key}
                http.Redirect(w, r, strings.Join(s, ""), 302)
            }
        }
    }
}

func main() {
    http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
    http.HandleFunc("/points", pointsHandler)
    http.HandleFunc("/", handler)
    http.ListenAndServe(":9000", nil)
}
