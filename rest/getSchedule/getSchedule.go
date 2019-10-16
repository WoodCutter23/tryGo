package getSchedule

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	V "netControllers/validators/getSchedule"
)

func GetSchedule(dbConnect string)  {
	http.HandleFunc("/getSchedule", func(w http.ResponseWriter, r *http.Request) {

		var validateMethod = V.ValidateMethod(r.Method)

		id := r.URL.Query().Get("controllerId")

		var validateData = V.ValidateData(id)

		if !validateMethod {
			w.WriteHeader(http.StatusMethodNotAllowed)
			_, _ = fmt.Fprintf(w, "This Method Not Allowed")
			return
		}

		if !validateData {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = fmt.Fprintf(w, "Controller id must be set & must be integer")
			return
		}

		db, err := sql.Open("mysql", dbConnect)

		if err != nil {
			panic(err)
		}

		stmtOut, err := db.Prepare("SELECT schedule FROM basic WHERE controller_id = ?")

		defer stmtOut.Close()

		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		var schedule string

		err = stmtOut.QueryRow(id).Scan(&schedule)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = fmt.Fprintf(w, "have not got information aboutthis controller")
			return
		}
		_, _ = fmt.Fprintf(w, schedule)
	})
}
