package recordsHandlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	recordServices "github.com/Gibad-brave-monkey/blog-go-backend/internal/services/records-services"
	"github.com/Gibad-brave-monkey/blog-go-backend/internal/storage"
	"github.com/go-chi/chi/v5"
)

type RecordsHandlers struct {
}

func (rh *RecordsHandlers) SaveRecord(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var record storage.Record

	err := decoder.Decode(&record)
	if err != nil {
		http.Error(w, "error decode JSON", http.StatusBadRequest)
		return
	}

	if record.Title == "" {
		http.Error(w, `{"error": "field 'Title' is required!"}`, http.StatusBadRequest)
		return
	}

	recordServices.SaveRecord(record)
}

func (rh *RecordsHandlers) GetRecords(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetRecords")
}

func (rh *RecordsHandlers) GetRecordById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	fmt.Println("GetRecordById", id)
}
