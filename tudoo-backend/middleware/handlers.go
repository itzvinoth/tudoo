package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"tudoo-backend/models"
	"log"
	"net/http"
	"os"
	"strconv"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// response format
type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

// create connection with postgres db
func createConnection() *sql.DB {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Open the connection
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		panic(err)
	}

	// check the connection
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	// return the connection
	return db
}

// CreateBoard create a board in the postgres db "tudoo"
func CreateBoard(w http.ResponseWriter, r *http.Request) {
	// set the header to content type x-www-form-urlencoded
	// Allow all origin to handle cors issue
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// create an empty board of type models.Board
	var board models.Board

	// decode the json request to board
	err := json.NewDecoder(r.Body).Decode(&board)

	if err != nil {
		log.Fatalf("Unable to decode the request body. %v", err)
	}

	// call insert board function and pass the board
	insertID := insertBoard(board)

	// format a response object
	res := response{
		ID:          insertID,
		Message:     "Board created successfully",
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

// GetBoard will retuan a single board by its ID
func GetBoard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	
	// get the boardid from the request params
	params := mux.Vars(r)

	// convert the id type from string to int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int. %v", err)
	}

	// call the getBoard function with board id to retrieve a single board
	board, err := getBoard(int64(id))

	if err != nil {
		log.Fatalf("Unable to get board. %v", err)
	}

	// send the response
	json.NewEncoder(w).Encode(board)	
}

// GetAllBoard will return all the boards
func GetAllBoard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// get all the boards in the db
	boards, err := getAllBoards()

	if err != nil {
		log.Fatalf("Unable to get all boards. %v", err)
	}

	// send all the boars as response
	json.NewEncoder(w).Encode(boards)
}

// UpdateBoard update board's detail in the postgres db
func UpdateBoard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// get the boardid from the request params, key is "id"
	params := mux.Vars(r)

	// convert the id type from string to int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert string into int. %v", err)
	}

	// create an empty board of type models.Board
	var board models.Board

	// decode the json request to board
	err = json.NewDecoder(r.Body).Decode(&board)

	if err != nil {
		log.Fatalf("Unable to decode the request body. %v", err)
	}

	// call updateBoard to update the board
	updateRows := updateBoard(int64(id), board)

	// format the message string
	msg := fmt.Sprintf("Board updated successfully. Total rows/record affected %v", updateRows)

	// format the response message
	res := response{
		ID:       int64(id),
		Message:  msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

// DeleteBoard delete board's detail in the postgres db
func DeleteBoard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// get the boardid from the request params, key is "id"
	params := mux.Vars(r)

	// convert the id in string to int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	// call the deleteBoard, convert the int to int64
	deletedRows := deleteBoard(int64(id))

	// format the message string
	msg := fmt.Sprintf("Board updated successfully. Total rows/record affected %v", deletedRows)

	// format the response message
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

//------------------------- handler functions ----------------
// insert one board in the DB
func insertBoard(board models.Board) int64 {
	
	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create the insert sql query
	// returning boardid will return the id of the inserted board
	sqlStatement := `INSERT INTO boards (name) VALUES ($1) RETURNING boardid`
	
	// the inserted id will store in this id
	var id int64

	// execute the sql statement
	// Scan function will save the insert id in the id
	err := db.QueryRow(sqlStatement, board.Name).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	fmt.Printf("Inserted a single record %v", id)

	// return the inserted id
	return id
}

// get one board from the DB by its boardid
func getBoard(id int64) (models.Board, error) {
	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create a board of models.Board type
	var board models.Board

	// create the select sql query
	sqlStatement := `SELECT * FROM boards WHERE boardid=$1`

	// execute the sql statement
	row := db.QueryRow(sqlStatement, id)

	// unmarshal the row object to board
	err := row.Scan(&board.ID, &board.Name)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return board, nil
	case nil:
		return board, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	// return empty board on error
	return board, err
}

// get one board from the DB by its boardid
func getAllBoards() ([]models.Board, error) {
	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	var boards []models.Board

	// create the select sql query
	sqlStatement := `SELECT * FROM boards`

	// execute the sql statement
	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// close the statement
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var board models.Board

		// unmarshal the row object to board
		err = rows.Scan(&board.ID, &board.Name)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		// append the board in the boards slice
		boards = append(boards, board)

	}

	// return empty board on error
	return boards, err
}

// update board in the DB
func updateBoard(id int64, board models.Board) int64 {

	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create the update sql query
	sqlStatement := `UPDATE boards SET name=$2 WHERE boardid=$1`

	// execute the sql statement
	res, err := db.Exec(sqlStatement, id, board.Name)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}

// delete board in the DB
func deleteBoard(id int64) int64 {

	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create the delete sql query
	sqlStatement := `DELETE FROM boards WHERE boardid=$1`

	// execute the sql statement
	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}