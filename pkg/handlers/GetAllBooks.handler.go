package handlers

import (
	"encoding/json"
	. "net/http"
	
	. "go_rest_api/pkg/mocks"
)
// ⚫️⚫️☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰

func GetAllBooks(writer ResponseWriter, _ *Request) {
	// ™ ____________
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(StatusOK)
	
	err := json.NewEncoder(writer).Encode(BooksMockData)
	if err != nil { return }
}
// ⚫️⚫️☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰☰

