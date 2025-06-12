package main

import (
	"fmt"
	"database/sql"
	"github.com/google/uuid"
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
	residentId := uuid.MustParse(id)
    row := db.QueryRow("SELECT * FROM Resident WHERE ResidentID = ?",residentId)
    
    if err := row.Scan(&resident.ResidentId, &resident.FirstName, &resident.LastName, &resident.Address,&resident.PhoneNumber,&resident.Email,&resident.ResidentStatus); err != nil {
        if err == sql.ErrNoRows {
            return resident, fmt.Errorf("albumsById %d: no such resident", residentId)
        }
        return resident, fmt.Errorf("residentByName %d: %v", residentId, err)
    }

    return resident, nil
}
// func addResident(resident Resident)(uuid.UUID,error){
// 	result, err := db.Exec("INSERT INTO Resident ()")
// }