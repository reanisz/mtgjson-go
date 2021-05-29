// Package models contains the types for schema ''.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"errors"

	"github.com/xo/xoutil"
)

// Ruling represents a row from 'rulings'.
type Ruling struct {
	ID   NullInt64     `json:"id" db:"id"`     // id
	Date xoutil.SqTime `json:"date" db:"date"` // date
	Text NullString    `json:"text" db:"text"` // text
	UUID NullString    `json:"uuid" db:"uuid"` // uuid

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Ruling exists in the database.
func (r *Ruling) Exists() bool {
	return r._exists
}

// Deleted provides information if the Ruling has been deleted from the database.
func (r *Ruling) Deleted() bool {
	return r._deleted
}

// Insert inserts the Ruling to the database.
func (r *Ruling) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if r._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by sequence
	const sqlstr = `INSERT INTO rulings (` +
		`date, text, uuid` +
		`) VALUES (` +
		`?, ?, ?` +
		`) RETURNING id`

	// run query
	XOLog(sqlstr, r.Date, r.Text, r.UUID)
	err = db.QueryRow(sqlstr, r.Date, r.Text, r.UUID).Scan(&r.ID)
	if err != nil {
		return err
	}

	// set existence
	r._exists = true

	return nil
}

// Update updates the Ruling in the database.
func (r *Ruling) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !r._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if r._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE rulings SET (` +
		`date, text, uuid` +
		`) = ( ` +
		`?, ?, ?` +
		`) WHERE id = $4`

	// run query
	XOLog(sqlstr, r.Date, r.Text, r.UUID, r.ID)
	_, err = db.Exec(sqlstr, r.Date, r.Text, r.UUID, r.ID)
	return err
}

// Save saves the Ruling to the database.
func (r *Ruling) Save(db XODB) error {
	if r.Exists() {
		return r.Update(db)
	}

	return r.Insert(db)
}

// Upsert performs an upsert for Ruling.
//
// NOTE: PostgreSQL 9.5+ only
func (r *Ruling) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if r._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO rulings (` +
		`id, date, text, uuid` +
		`) VALUES (` +
		`?, ?, ?, ?` +
		`) ON CONFLICT (id) DO UPDATE SET (` +
		`id, date, text, uuid` +
		`) = (` +
		`EXCLUDED.id, EXCLUDED.date, EXCLUDED.text, EXCLUDED.uuid` +
		`)`

	// run query
	XOLog(sqlstr, r.ID, r.Date, r.Text, r.UUID)
	_, err = db.Exec(sqlstr, r.ID, r.Date, r.Text, r.UUID)
	if err != nil {
		return err
	}

	// set existence
	r._exists = true

	return nil
}

// Delete deletes the Ruling from the database.
func (r *Ruling) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !r._exists {
		return nil
	}

	// if deleted, bail
	if r._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM rulings WHERE id = $1`

	// run query
	XOLog(sqlstr, r.ID)
	_, err = db.Exec(sqlstr, r.ID)
	if err != nil {
		return err
	}

	// set deleted
	r._deleted = true

	return nil
}

// Card returns the Card associated with the Ruling's UUID (uuid).
//
// Generated from foreign key 'rulings_uuid_fkey'.
func (r *Ruling) Card(db XODB) (*Card, error) {
	return CardByUUID(db, r.UUID.String)
}

// RulingByID retrieves a row from 'rulings' as a Ruling.
//
// Generated from index 'rulings_id_pkey'.
func RulingByID(db XODB, id NullInt64) (*Ruling, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, date, text, uuid ` +
		`FROM rulings ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	r := Ruling{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&r.ID, &r.Date, &r.Text, &r.UUID)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
