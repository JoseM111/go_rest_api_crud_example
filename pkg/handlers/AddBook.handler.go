package handlers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	. "net/http"
	
	"go_rest_api/pkg/mocks"
	. "go_rest_api/pkg/models"
)
// ⚫️⚫️☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰

func AddBooks(writer ResponseWriter, request *Request) {
	// ™ ____________
	/* read the request body */
	var book Book
	
	// so the body closes when we are done our request
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil { return }
	}(request.Body)
	
	body, err := ioutil.ReadAll(request.Body)
	if err != nil { log.Fatalln() }
	
	err = json.Unmarshal(body, &book)
	if err != nil { return }
	
	/* append to the mock book data */
	book.Id = rand.Intn(100)
	mocks.BooksMockData = append(mocks.BooksMockData, book)
	
	/* send a 201 created response */
	writer.WriteHeader(StatusCreated)
	writer.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(book)
	if err != nil { return }
}
// ⚫️⚫️☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰
















