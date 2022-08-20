package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"database/sql"
)

func readCSVFromUrl(url string) ([][]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	reader := csv.NewReader(resp.Body)
	reader.Comma = ';'
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func connection() {
    const(
        host =  "localhost"
		port = 5432
		user = postgres
		password = "Cuihua2020@"
		dbname = "cbiproject_postgres"
	)

    psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
    
	db. err := sql.Open("postgres", psqlconn)

	defer db.Close()

}

func main() {
    url := "https://data.cityofchicago.org/api/views/xhc6-88s9/rows.csv?accessType=DOWNLOAD"
	data, err := readCSVFromUrl(url)
	if err != nil {
		panic(err)
	}

	for idx, row := range data {
		// skip header
		if idx == 0 {
			continue
		}

		if idx == 6 {
			break
		}

		fmt.Println(row[2])
	}
}