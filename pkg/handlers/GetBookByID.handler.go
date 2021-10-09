package handlers

import (
	"encoding/json"
	. "net/http"
	"strconv"
	
	"github.com/gorilla/mux"
	
	. "go_rest_api/pkg/mocks"
)
// ⚫️⚫️☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰

func GetBookByID(writer ResponseWriter, request *Request) {
	// ™ ____________
	/*| read our dynamic parameter: {id} |*/
	vars := mux.Vars(request)
	id, _ := strconv.Atoi(vars["id"])
	
	/*| iterate over all the mock book data & compare id's |*/
	for _, book := range BooksMockData {
		// compare the book id with our dynamic: {id}
		if book.Id == id {
			/* if ids are equal send book as response */
			writer.WriteHeader(StatusOK)
			writer.Header().Add("Content-Type", "application/json")
			err := json.NewEncoder(writer).Encode(book)
			if err != nil { return }
			break
		}
	}
}
// ⚫️⚫️☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰
