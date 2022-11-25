/*
Écrire un module exécutable servant une API sur le port 4567 avec comme routes:
- Une route GET “/” affichant l’heure qu’il est.
- Une route GET “/dice” affichant le résultat d’un dé à 1000 faces (D1000)
- Une route GET “/dices” affichant quinze dés aux nombres de faces aléatoires parmi:
une pièce (D2), un D4, un D6, un D8, un D10, un D12, un D20 et un D100
- Ajouter un paramètre optionnel pour préciser le type de tous les dés à l’appel de la
route GET précédente: “/dices?type=d6”
- Une route POST “/randomize-words” prenant un payload sous forme de
application/x-www-form-urlencoded avec comme paramètre:
- words: la phrase contenant des mots à renvoyer dans le désordre
- Une route POST “/semi-capitalize-sentence” prenant un payload sous forme de
application/x-www-form-urlencoded avec comme paramètre:
- sentence: la phrase à renvoyer avec une lettre en majuscule sur deux
Le binaire généré devra s’appeler “miniapi” (ou miniapi.exe), un appel go build dans le
dossier de travail devrait être suffisant pour le générer.*/

package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func main() {

	var choice string

	fmt.Printf("1 : hourHandler -> / | 2 : diceHandler -> /dice | 3 : dicesHandler -> /dices | 4 : randomizeHandler -> /randomize-words\n")
	fmt.Scan(&choice)

	//http.HandleFunc("/hello", helloHandler)

	if choice == "/" {
		http.HandleFunc("/", hourHandler) // affiche l'heure qu'il est           OK
	}

	if choice == "/dice" {
		http.HandleFunc("/dice", diceHandler) //dice” affichant le résultat d’un dé à 1000 faces (D1000)   OK
	}

	if choice == "/dices" {
		http.HandleFunc("/dices", dicesHandler) //Une route GET “/dices” affichant quinze dés aux nombres de faces aléatoires parmi:
		//une pièce (D2), un D4, un D6, un D8, un D10, un D12, un D20 et un D100
	}

	if choice == "randomize-words" {
		//http.HandleFunc("/randomize-words", randomizeHandler)
		/*Une route POST “/randomize-words” prenant un payload sous forme de
		application/x-www-form-urlencoded avec comme paramètre:
		- words: la phrase contenant des mots à renvoyer dans le désordre*/
	}
	http.ListenAndServe(":4567", nil)
}

func randomizeHandler(w http.ResponseWriter, req *http.Request) {

	switch req.Method {
	case http.MethodPost:
		if err := req.ParseForm(); err != nil {
			fmt.Println("Something went bad")
			fmt.Fprintln(w, "Something went bad")
			return
		}

		for key, value := range req.PostForm {
			fmt.Println(key, "=>", value)
		}
		fmt.Fprintf(w, "Information received: %v\n", req.PostForm)
	}

	if _, ok := req.PostForm["sentence"]; !ok {
		fmt.Println("something went wrong")
		fmt.Println(w, "something went wrong")
		return
	}

	fmt.Fprintf(w, "<p>%s</p>", req.PostForm["sentence"][0])

}

// ///////////////////////////////////////////////////////////////////////////////////////////////
func dicesHandler(w http.ResponseWriter, req *http.Request) {

	switch req.Method {
	case http.MethodGet:
		rand.Seed(time.Now().UnixNano())

		if req.URL.Query().Get("type") != "" {
			if req.URL.Query().Get("type")[0] == 'd' {

				faces, _ := strconv.Atoi(req.URL.Query().Get("type")[1:])

				for i := 0; i < 15; i++ {
					fmt.Fprintf(w, "%d ", rand.Intn(faces)+1)
				}
				return

			} else {
				fmt.Fprintf(w, "Bad request")
				return
			}
		} else {

			dices := []int{2, 4, 6, 8, 10, 12, 20, 100}

			for i := 0; i < 15; i++ {
				faces := dices[rand.Intn(len(dices))]
				fmt.Fprintf(w, "%d ", rand.Intn(faces)+1)
			}
			return
		}

	default:
		fmt.Fprintf(w, "Method not allowed")
		return
	}
}

/////////////////////////////////////////////////////////////////////////////////////////////////

func hourHandler(w http.ResponseWriter, req *http.Request) {

	switch req.Method {
	case http.MethodGet:
		day := time.Now()
		fmt.Fprintf(w, "Il est %dh%s", day.Hour(), addZeroIfNecessary(day.Minute()))

	}

}

/////////////////////////////////////////////////////////////////////////////////////////////////

func diceHandler(w http.ResponseWriter, req *http.Request) {

	switch req.Method {
	case http.MethodGet:
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)

		fmt.Fprintf(w, "resultat du dé : |%d|\n", r1.Intn(1000)+1)
	}
}

func addZeroIfNecessary(number int) string {
	if number <= 9 {
		return "0" + fmt.Sprintf("%d", number)
	}
	return fmt.Sprintf("%d", number)
}
