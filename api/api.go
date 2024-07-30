package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Handler to load index page
//
// Parameters:
//     w - http.ResponseWriter
//     r - *http.Request
//
// Returns:
//     Nothing


type Reply struct{

	Percentage float32
	Segments []float32
}



func (button ButtonService) buttonHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":

		var reply = Reply{
			Percentage: button.Repository.Value(),
			Segments: button.Repository.Segments(),
		}
		log.Println("GET")
		var jsonData, err = json.Marshal(reply)
		if err != nil {
			log.Println(err)
		}
		fmt.Fprintf(w, "%s",jsonData)
	case "POST":
		var percentage = button.Repository.Value()
		var seg = button.Repository.Click()
		var reply = map[string]string{"percentage": fmt.Sprintf("%.3f",percentage), "segment": strconv.Itoa(int(seg))}
		log.Println("POST")
		var jsonData, err = json.Marshal(reply)
		if err != nil {
			log.Println(err)
		}
		fmt.Fprintf(w, "%s",jsonData)
	default:
		log.Println("Method not supported")
	}
}



func AddRoutes(root string, buttonService ButtonService) {
	http.HandleFunc(root+ "/button", buttonService.buttonHandler)
}