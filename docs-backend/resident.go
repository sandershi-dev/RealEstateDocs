package main

import (
	"fmt"
	"database/sql"
	"github.com/google/uuid"
    "encoding/json"
    "strings"
)
type Resident struct{
	ResidentId uuid.UUID
	FirstName string
	LastName string
	Address string
	PhoneNumber int
	Email string
	ResidentStatus string
}
func ResidentJsonToString(jsonString string)(map[string]interface{}){
    var data map[string]interface{}

	// Unmarshal the JSON string (converted to a byte slice) into the map
	err := json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return nil
	}
    return data
}
func getAllResidents(db *sql.DB)([]Resident, error){
	// An albums slice to hold data from returned rows.
    var residents []Resident

    rows, err := db.Query("SELECT * FROM Resident")
    if err != nil {
        return nil, fmt.Errorf("getAllResident %v", err)
    }
    defer rows.Close()
    // Loop through rows, using Scan to assign column data to struct fields.
    for rows.Next() {
        var res Resident
        if err := rows.Scan(&res.ResidentId, &res.FirstName, &res.LastName, &res.Address,&res.PhoneNumber,&res.Email,&res.ResidentStatus); err != nil {
            return nil, fmt.Errorf("getAllResidents: %v", err)
        }
        residents = append(residents, res)
    }
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("getAllResidents: %v", err)
    }
    return residents, nil
}
func getResidentByName(db *sql.DB, Name string)([]Resident,error){
    // An albums slice to hold data from returned rows.
    var residents []Resident

    rows,err := db.Query("SELECT * FROM Resident WHERE FirstName = ? OR LastName =?",Name, Name)
	if err != nil {
        return nil, fmt.Errorf("getAllResident %v", err)
    }
	defer rows.Close()

    for rows.Next() {
        var res Resident
        if err := rows.Scan(&res.ResidentId, &res.FirstName, &res.LastName, &res.Address,&res.PhoneNumber,&res.Email,&res.ResidentStatus); err != nil {
            return nil, fmt.Errorf("getAllResidents: %v", err)
        }
        residents = append(residents, res)
    }
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("getAllResidents: %v", err)
    }
    return residents, nil
}
func getResidentByID(db *sql.DB, id string )(Resident,error){
    // An albums slice to hold data from returned rows.
    var resident Resident
	resident_id := uuid.MustParse(id)
    row := db.QueryRow("SELECT * FROM Resident WHERE ResidentID = ?",resident_id)
    
    if err := row.Scan(&resident.ResidentId, &resident.FirstName, &resident.LastName, &resident.Address,&resident.PhoneNumber,&resident.Email,&resident.ResidentStatus); err != nil {
        if err == sql.ErrNoRows {
            return resident, fmt.Errorf("albumsById %d: no such resident", resident_id)
        }
        return resident, fmt.Errorf("residentByName %d: %v", resident_id, err)
    }

    return resident, nil
}
func addResident(db *sql.DB,resident Resident)(int64,error){
	result, err := db.Exec("INSERT INTO Resident (first_name,last_name,address,phone_number,email,resident_status) VALUES (?,?,?,?,?,?) RETURNING resident_id",resident.FirstName,resident.LastName,resident.Address,resident.PhoneNumber,resident.Email,resident.ResidentStatus)
    if err != nil {
        return 0, fmt.Errorf("addResident: %v", err)
    }
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return 0, fmt.Errorf("addResident: %v", err)
    }
    return rowsAffected, nil
}

func deleteResidentbyId(db *sql.DB,id string)(int64,error){
    resident_id := uuid.MustParse(id)
    result, err := db.Exec("DELETE FROM Resident WHERE resident_id = ?", resident_id)
    if err != nil {
        return 0, fmt.Errorf("deleteResident: %v", err)
    }
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return 0, fmt.Errorf("deleteResident: %v", err)
    }
    return rowsAffected, nil
    
}
func updateResident(db *sql.DB, id string, fields map[string]string)(int64,error){
    resident_id := uuid.MustParse(id)
    var sb strings.Builder

    for col, value := range fields{
        sb.WriteString(fmt.Sprintf("%s = '%s',",col,value)) 
    }

    updateFields := sb.String()

    if len(updateFields) > 0 {
		updateFields = updateFields[:len(updateFields)-1] // Remove trailing comma
	}
    result, err := db.Exec("UPDATE Resident SET "+updateFields+" WHERE resident_id = ?",resident_id)
    if err != nil {
        return 0, fmt.Errorf("UpdateResident: %v", err)
    }
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return 0, fmt.Errorf("UpdateResident: %v", err)
    }
    return rowsAffected, nil
}