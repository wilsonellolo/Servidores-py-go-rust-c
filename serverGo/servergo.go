package main

import (
	"html/template"
	"net/http"
	"os"
	"github.com/go-redis/redis"
	"path/filepath"
	"fmt"
)

var templatesDir = os.Getenv("static")

type Punto struct {
	Tiempo string
	Uso string
	Ultimo string
}

func pagina(w http.ResponseWriter, r *http.Request) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "52.33.234.92:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	pong, err := rdb.Ping().Result()
	fmt.Println(pong, err)
	
	t, _ := rdb.LRange("tiempo", -10, -1).Result()
	fmt.Println(t)
	u, _ := rdb.LRange("memoria", -10, -1).Result()
	fmt.Println(u)
	
	
	tmplPath := filepath.Join("static", "index.html")
	tmpl := template.Must(template.ParseFiles(tmplPath))
	data := []Punto{{t[0],u[0],""},{t[1],u[1],""},{t[2],u[2],""},{t[3],u[3],""},{t[4],u[4],""},{t[5],u[5],""},{t[6],u[6],""},{t[7],u[7],""},{t[8],u[8],""},{t[9],u[9],u[9]}}

	tmpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/", pagina)
	http.ListenAndServe(":9001", nil)
}


