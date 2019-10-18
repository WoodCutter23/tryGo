package setStatistic

import (
	"database/sql"
	_ "database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	V "netControllers/validators/setSchedule"
)

func SetStatistic(dbConnect string)  {
	http.HandleFunc("/setStatistic", func(w http.ResponseWriter, r *http.Request) {

		if err := r.ParseForm(); err != nil {
			_, _ = fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		id := r.FormValue("id")

		data :=r.FormValue("data")

		var validateMethod  = V.ValidateMethod(r.Method)

		if !validateMethod {
			w.WriteHeader(http.StatusMethodNotAllowed)
			_, _ = fmt.Fprintf(w, "This Method Not Allowed")
			return
		}

		validId := V.ValidateId(id)

		if !validId {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = fmt.Fprintf(w, "id must be set & integer")
			return
		}

		db, err := sql.Open("mysql", dbConnect)

		if err != nil {
			panic(err)
		}

		_, err = db.Exec("INSERT into statistic (statistic, controller_id) VALUES(?,?)",
			data, id)

		if err != nil {
			panic(err)
		}

	})
}
