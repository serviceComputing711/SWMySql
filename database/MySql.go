package database

import (
	"database/sql"
    _ "github.com/go-sql-driver/mysql"
	"fmt"
	"encoding/json"
)

type SWSql struct {
	db *sql.DB
}

/*
**	close the database
*/
func (m *SWSql) Close() error {
	return m.db.Close()
}

/*
**	open the database with Path
**	if data not exsist, create a database
**	return a database in SWSql
**	if wrong, return wrong message
*/
func StartDB(name string) (*SWSql, error) {
	//open database
	db, err := sql.Open(name, "root:@tcp(127.0.0.1:3306)/testuser?parseTime=true")
	fmt.Println("Open the database: " + name)
	if err != nil {
		return nil, err
	}
	return &SWSql{db}, nil
}

/*
**	initial database
**	initial with 7 tables, tables "user" record the information of the assign user
**	other tables record the information of films, people, etc
**	tables is empty
*/
func (m *SWSql) InitDB() {
	fmt.Println("Initial DB")
	//get connection
	tx,_ := m.db.Begin()
	//create table
	tx.Exec("CREATE TABLE films (ID char(10), Mes char(1000))")
	tx.Exec("CREATE TABLE people (ID char(10), Mes char(1000))")
	tx.Exec("CREATE TABLE planets (ID char(10), Mes char(1000))")
	tx.Exec("CREATE TABLE species (ID char(10), Mes char(1000))")
	tx.Exec("CREATE TABLE starships (ID char(10), Mes char(1000))")
	tx.Exec("CREATE TABLE vehicles (ID char(10), Mes char(1000))")
	tx.Exec("CREATE TABLE user (Token char(50), Name char(1000))")
    	//fresh the connection of tx
    	//commit connection
    	tx.Commit()

	// _, err := tx.CreatetableIfNotExists([]byte("films"))
	// _, err = tx.CreatetableIfNotExists([]byte("people"))
	// _, err = tx.CreatetableIfNotExists([]byte("planets"))
	// _, err = tx.CreatetableIfNotExists([]byte("species"))
	// _, err = tx.CreatetableIfNotExists([]byte("starships"))
	// _, err = tx.CreatetableIfNotExists([]byte("vehicles"))
	// _, err = tx.CreatetableIfNotExists([]byte("user"))
}

/*
**	add information to table
**	has 3 arguement:
**	table: tables type, as films, people, etc
**	key: object id
**	value: object information as Json
**	return wrong message
*/
func (m *SWSql) AddObj(table string, ID string, Mes string) (error) {
	//construct command
	com := "INSERT INTO " + string(table) + "(ID, Mes) VALUES (?, ?)"
	res, err := m.db.Exec(com, ID, Mes)
	fmt.Println(res)
	return err
}

/*
**	delete a object from database
**	has 2 arguement
**	table: tables type, as films, people, etc
**	key: object id
**	if has error, return wrong message
*/
func (m *SWSql) DeleteObj(table string, key string) (error) {
	//construct command
	com := "DELETE FROM " + string(table) + " WHERE ID = ?"
	res, err := m.db.Exec(com, key)
	fmt.Println(res)
	return err
}

/*
**	search object by ID
**	has 2 arguement
**	table: tables type, as films, people, etc
**	key: object ID
**	return object by Json
*/
func (m *SWSql) SearchByID (table string, ID string) (string) {
	com := "SELECT ID FROM " + table + " WHERE ID = ?"
	rows, _ := m.db.Query(com, ID)
	var cID string
	var Mes string
	for rows.Next() {
		rows.Scan(&cID, &Mes)
	}
	defer rows.Close()
    	return Mes
}

/*
**	search film by Name
**	has 1 arguement
**	key: object Name
**	return object by Json
*/
func (m *SWSql) SearchFilmByName (name string) (string) {
	var p Film
	com := "SELECT * FROM films"
	rows, _ := m.db.Query(com)
	var res string
	for rows.Next() {
		var ID string
		var Mes string
		rows.Scan(&ID, &Mes)
		json.Unmarshal([]byte(Mes), &p)
		if p.getName() == name {
			res = Mes
			break
		}
	}
	defer rows.Close()
    	return res
}

/*
**	search person by Name
**	has 1 arguement
**	key: object Name
**	return object by Json
*/
func (m *SWSql) SearchPersonByName (name string) (string) {
	var p Person
	com := "SELECT * FROM people"
	rows, _ := m.db.Query(com)
	var res string
	for rows.Next() {
		var ID string
		var Mes string
		rows.Scan(&ID, &Mes)
		json.Unmarshal([]byte(Mes), &p)
		if p.getName() == name {
			res = Mes
			break
		}
	}
	defer rows.Close()
    	return res
}

/*
**	search planet by Name
**	has 1 arguement
**	key: object Name
**	return planet by Json
*/
func (m *SWSql) SearchPlanetByName (name string) (string) {
    var p Planet
	com := "SELECT * FROM planets"
	rows, _ := m.db.Query(com)
	var res string
	for rows.Next() {
		var ID string
		var Mes string
		rows.Scan(&ID, &Mes)
		json.Unmarshal([]byte(Mes), &p)
		if p.getName() == name {
			res = Mes
			break
		}
	}
	defer rows.Close()
    	return res
}

/*
**	search species by Name
**	has 1 arguement
**	key: object Name
**	return species by Json
*/
func (m *SWSql) SearchSpeciesByName (name string) (string) {
	var p Species
	com := "SELECT * FROM species"
	rows, _ := m.db.Query(com)
	var res string
	for rows.Next() {
		var ID string
		var Mes string
		rows.Scan(&ID, &Mes)
		json.Unmarshal([]byte(Mes), &p)
		if p.getName() == name {
			res = Mes
			break
		}
	}
	defer rows.Close()
    	return res
}

/*
**	search starship by Name
**	has 1 arguement
**	key: object Name
**	return starship by Json
*/
func (m *SWSql) SearchStarshipByName (name string) (string) {
	var p Starship
	com := "SELECT * FROM starships"
	rows, _ := m.db.Query(com)
	var res string
	for rows.Next() {
		var ID string
		var Mes string
		rows.Scan(&ID, &Mes)
		json.Unmarshal([]byte(Mes), &p)
		if p.getName() == name {
			res = Mes
			break
		}
	}
	defer rows.Close()
    	return res
}

/*
**	search vehicle by Name
**	has 1 arguement
**	key: object Name
**	return Vehicle by Json
*/
func (m *SWSql) SearchVehicleByName (name string) (string) {
	var p Vehicle
	com := "SELECT * FROM vehicles"
	rows, _ := m.db.Query(com)
	var res string
	for rows.Next() {
		var ID string
		var Mes string
		rows.Scan(&ID, &Mes)
		json.Unmarshal([]byte(Mes), &p)
		if p.getName() == name {
			res = Mes
			break
		}
	}
	defer rows.Close()
    	return res
}

/*
**	check  in or not
**	has 1 arguement
**	Token: Token
**	return bool
*/
func  (m *SWSql) IsLogIn (Token string) (bool) {
	flag := false
    	com := "SELECT * FROM user"
	rows, _ := m.db.Query(com)
	for rows.Next() {
		var thisToken string
		var thisName string
		rows.Scan(&thisToken, &thisName)
		if thisToken == Token {
			flag = true
			break
		}
	}
	defer rows.Close()
    	return flag
}

/*
**	user log in
**	add a key-value to table("user")
**	key: Token
**	value: name
**	if has error, return wrong message
*/
func (m *SWSql) LogIn(name string, Token string) (error) {
	//construct command
	com := "INSERT INTO user(Token, Name) VALUES (?, ?)"
	res, err := m.db.Exec(com, Token, name)
	fmt.Println(res)
	return err
}

/*
**	search information by page
**	has 2 arguement
**	table: tables, like people, films, etc
**	page: page number
*/
func (m *SWSql) SearchByPage (table string, page int) (string) {
	endNum := 5 * page
	startNum := endNum - 5
	currentNum := 0
	com := "SELECT * FROM vehicles"
	rows, _ := m.db.Query(com)
	var res string
	for rows.Next() {
		var ID string
		var Mes string
		rows.Scan(&ID, &Mes)
    		if currentNum < endNum && currentNum >= startNum {
    			res = res + Mes
    		}
		currentNum++
		if currentNum == endNum {
			break
		}
    	}
    	return res
}
