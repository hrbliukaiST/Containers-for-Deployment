// websockets.go
package main

import (
	"fmt"
	"net/http"

	"strings"

	"math/rand"

	"github.com/gorilla/websocket"

	"regexp"
	// submsgstr "strings";
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity

		for {
			// Read message from browser
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}
			var radomtext [5]string
			radomtext[0] = "Logic will get you from A to B. Imagination will take you everywhere."
			radomtext[1] = "There are 10 kinds of people. Those who know binary and those who don't.  \n \n Repository  https://github.com/fred/myrepo/ "
			//radomtext[2] = "There are two ways of constructing a software design. One way is to make it so simple that there are obviously no deficiencies and the other is to make it so complicated that there are no obvious deficiencies. \n \n
			//Repository  <a href=\"https://github.com/fred/myrepo/\">Visit https://github.com/fred/myrepo</a> "
			radomtext[2] = "There are two ways of constructing a software design. One way is to make it so simple that there are obviously no deficiencies and the other is to make it so complicated that there are no obvious deficiencies."
			radomtext[3] = "It's not that I'm so smart, it's just that I stay with problems longer."
			radomtext[4] = "It is pitch dark. You are likely to be eaten by a grue."

			var id = rand.Intn(5)

			//text := string(msg)
			text := radomtext[id]

			// Define a regular expression to match URLs in the text
			urlRegex := regexp.MustCompile(`(http[s]?|ftp):\/?\/?([^:\/\s]+)(:[0-9]+)?(\/\S*)?`)

			// Replace URLs with clickable links
			newText := urlRegex.ReplaceAllStringFunc(text, func(matchedURL string) string {
				// Check if the URL has a scheme, if not, add "http://" by default
				if !strings.HasPrefix(matchedURL, "http://") && !strings.HasPrefix(matchedURL, "https://") {
					matchedURL = "http://" + matchedURL
				}
				return fmt.Sprintf(`<a href="%s">%s</a>`, matchedURL, matchedURL)
			})

			// Print the modified text
			fmt.Println(newText, msgType, msg)

			//msg = append(msg, []byte(newText.Bytes()))
			fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(newText))

			// Write message back to browser
			// if err = conn.WriteMessage(msgType, msg); err != nil {
			// 	return
			// }
			if err = conn.WriteMessage(msgType, []byte(newText)); err != nil {
				return
			}

			// if err = conn.WriteMessage(websocket.TextMessage, []byte(newText)); err != nil {
			// 	return
			// }
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.ListenAndServe(":8888", nil)
}
