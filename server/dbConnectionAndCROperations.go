package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

var db pgx.Conn

func dbConnection() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("postgresql://localhost:5432/Rethenya"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

}

func createTable() {
	createSql :=
		` 
	Create table if not exists ports(
		id SERIAL PRIMARY KEY,
		name VARCHAR,
		code VARCHAR,
		city VARCHAR,
		state VARCHAR,
		country VARCHAR
	);
	`

	_, error := db.Exec(context.Background(), createSql)
	if error != nil {
		fmt.Fprintf(os.Stderr, "Table creation failed: %v\n", error)
		os.Exit(1)
	}
}
func Createnewport(id, name, code, city, state, country string) {
	_, newerr := db.Exec(context.Background(), "Insert into ports(name,code,city,state,country) values ($1,$2,$3,$4,$5)", id, name, code, city, state, country)
	if newerr != nil {
		fmt.Fprintf(os.Stderr, "insertion failed: %v\n", newerr)
		os.Exit(1)
	}
}

func Getportdetails(id string) (Id, name, code, city, state, country string) {
	results, err := db.Query(context.Background(), "select * from ports where id=$1", id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "insertion failed: %v\n", err)
		os.Exit(1)
	}
	error := results.Scan(&id, &name, &code, &city, &state, &country)
	if error != nil {
		fmt.Fprintf(os.Stderr, "Fetching the port detail failed:%v", error)
	}
	return id, name, code, city, state, country
}

func UpdatePortDetails(id, name, code, city, state, country string) {
	_, err := db.Exec(context.Background(), "Update ports set name=$1 where id=$2", id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Updating the port failed:%v", err)
	}
}

func checkPortId(id string) (isexists bool) {
	result, err := db.Query(context.Background(), "select exists(select 1 from ports where id=$1)", id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Updating the port failed:%v", err)
	}
	result.Scan(&isexists)
	return isexists
}

func DeletePortDetails(id string) {
	_, err := db.Exec(context.Background(), "Delete from ports where id=$2", id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "deleting the port failed:%v", err)
	}
}
