package recordsHandlers

import (
	"fmt"
	"net/http"
)

type RecordsHandlers struct {
}

func (rh *RecordsHandlers) SaveRecord(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Need to Implement")
	r.Header.Add("status", "300")
}

func (rh *RecordsHandlers) GetRecords(w http.ResponseWriter, t *http.Request) {
	fmt.Println("Need to implement")
}
