package handlers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	. "net/http"
	"strconv"
	
	"github.com/gorilla/mux"
	
	. "go_rest_api/pkg/mocks"
	. "go_rest_api/pkg/models"
)
// ⚫️⚫️☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰

func PutUpdateBook(writer ResponseWriter, request *Request) {
	// ™ ____________
	var updatedBook Book
	
	/*| read dynamic id variable: {id} parameter |*/
	vars := mux.Vars(request)
	id, _ := strconv.Atoi(vars["id"])
	
	/*| read request body |*/
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil { log.Fatalln(err) }
	}(request.Body)
	
	body, err := ioutil.ReadAll(request.Body)
	if err != nil { log.Fatalln(err) }
	
	err = json.Unmarshal(body, &updatedBook)
	if err != nil { return }
	
	/*| iterate through all the mock book data |*/
	for index, book := range BooksMockData {
		// ™ ____________
		if book.Id == id {
			/*| update & send response when book id matches dynamics id |*/
			book.Title = updatedBook.Title
			book.Author = updatedBook.Author
			book.Desc = updatedBook.Desc
			
			// once the id is correct grab the book at that index & update it
			BooksMockData[index] = book
			// response
			writer.WriteHeader(StatusOK)
			writer.Header().Add("Content-Type", "application/json")
			err := json.NewEncoder(writer).Encode(book)
			if err != nil { return }
			break
		}
	}
}
// ⚫️⚫️☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰
