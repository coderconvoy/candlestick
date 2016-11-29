package main

import (
	"fmt"
	"github.com/coderconvoy/candlestick/types"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"text/template"
	"time"
)

//Webby Section
//
//

type usersession struct {
	g        *types.Game
	lastUsed time.Time
}

var temps *template.Template

var sessions map[string]*usersession
var seshMutex sync.Mutex

func handle(w http.ResponseWriter, r *http.Request) {

	c, err := r.Cookie("uid")
	if err != nil {
		c = &http.Cookie{Name: "uid", Value: fmt.Sprintf("%d:%d", time.Now().Unix(), rand.Intn(10000)), Expires: time.Now().Add(time.Hour * 24)}
	}

	seshMutex.Lock()
	userS, ok := sessions[c.Value]
	if !ok {
		//Here we need to make a live session to work with.
		userS = &usersession{types.NewGame(5), time.Now()}
		sessions[c.Value] = userS
	}
	seshMutex.Unlock()

	game := userS.g
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

	http.SetCookie(w, c)

	err = temps.ExecuteTemplate(w, "main.html", game)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintln(w, err)
	}
}

func killCookies() {
	for {
		time.Sleep(time.Hour)
		seshMutex.Lock()
		for k, v := range sessions {
			if time.Now().After(v.lastUsed.Add(time.Hour * 24)) {
				delete(sessions, k)
			}
		}
		seshMutex.Unlock()

	}
}

func main() {
	rand.Seed(time.Now().Unix())

	var err error

	//session builder
	sessions = make(map[string]*usersession)
	go killCookies()

	//Template builder
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

	http.HandleFunc("/", handle)

	fmt.Printf("Server Starting\n")
	err = http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println(err)
	}
}
