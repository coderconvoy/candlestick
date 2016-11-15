package main

import (
	"candlestick/types"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"text/template"
	"time"
)

//Webby Section
//
//

var temps *template.Template
var game *types.Game

func handle(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/")
	a, err := strconv.Atoi(path)
	if err == nil {
		//Play round of game
		if game.HumanTurn(a) {
			for {
				a, _ := game.TryTurn()
				if a == types.TURN_HUMAN {
					break
				}
			}
		}

	} else {
		game.Message = fmt.Sprintf("Error:%d", err)

	}

	err = temps.ExecuteTemplate(w, "main.html", game)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintln(w, err)
	}
}

func main() {
	rand.Seed(time.Now().Unix())

	var err error

	fmap := template.FuncMap{
		"add": func(i ...int) int {
			res := 0
			for _, v := range i {
				res += v
			}
			return res
		},
	}

	temps, err = template.New("").Funcs(fmap).ParseGlob("templates/*.html")
	if err != nil {
		fmt.Println(err)
		return
	}

	game = types.NewGame(5)

	http.HandleFunc("/", handle)

	fmt.Printf("Server Starting\n")
	err = http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println(err)
	}
}
