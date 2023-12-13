package recordServices

import (
	"fmt"

	"github.com/Gibad-brave-monkey/blog-go-backend/internal/storage"
)

func SaveRecord(data storage.Record) error {
	fmt.Println(data, "DATA")
	// Some logic
	return nil
}
