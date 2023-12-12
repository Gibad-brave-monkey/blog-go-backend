package recordsHandlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type RecordsHandlers struct {
}

func (rh *RecordsHandlers) SaveRecord(w http.ResponseWriter, r *http.Request) {
	fmt.Println("SaveRecord")
}

func (rh *RecordsHandlers) GetRecords(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetRecords")
}

func (rh *RecordsHandlers) GetRecordById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	fmt.Println("GetRecordById", id)
}
