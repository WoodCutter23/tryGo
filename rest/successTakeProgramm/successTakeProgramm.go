package takeProgramm

import (
	"database/sql"
	"fmt"
	"net/http"
	V "netControllers/validators/getSchedule"
)

func SuccessGetProgramm(dbConnect string)  {
	http.HandleFunc("/successGetProgramm", func(w http.ResponseWriter, r *http.Request) {

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
		_, err = db.Exec("UPDATE basic SET  actual_programm = 1 WHERE controller_id = ?",
			id)
	})
}