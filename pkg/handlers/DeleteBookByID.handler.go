package handlers

import (
	"encoding/json"
	. "net/http"
	"strconv"
	
	"github.com/gorilla/mux"
	
	. "go_rest_api/pkg/mocks"
	. "go_rest_api/pkg/models"
)
// ⚫️⚫️☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰

func DeleteBookByID(writer ResponseWriter, request *Request) {
	// ™ ____________
	/*| read the dynamic id variable: {id} parameter |*/
	vars := mux.Vars(request)
	id, _ := strconv.Atoi(vars["id"])
	
	/*| iterate over all mock books data |*/
	for index, book := range BooksMockData {
		// #..........
		if book.Id == id {
			/*| delete book & send response to if the book id matches the dynamic id |*/
			BooksMockData = deleteFromIndex(index)
			
			// response
			writer.WriteHeader(StatusOK)
			err := json.NewEncoder(writer).Encode(book)
			if err != nil { return }
			break
		}
	}
}
// ⚫️⚫️☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰

/*| helper function |*/
func deleteFromIndex(index int) []Book {
	return append(BooksMockData[:index], BooksMockData[index+1:]...)
}
// ⚫️⚫️☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰
