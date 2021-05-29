package models

import (
	"database/sql"
	"encoding/json"
)

type NullInt32 struct {
	sql.NullInt32
}

type NullInt64 struct {
	sql.NullInt64
}

type NullBool struct {
	sql.NullBool
}

type NullFloat64 struct {
	sql.NullFloat64
}

type NullString struct {
	sql.NullString
}

type NullTime struct {
	sql.NullTime
}

func (v *NullInt32) UnmarshalJSON(value []byte) error {
	err := json.Unmarshal(value, v.Int32)
	v.Valid = err == nil
	return err
}
func (v *NullInt32) MarshalJSON() ([]byte, error) {
	if !v.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(v.Int32)
}

func (v *NullInt64) UnmarshalJSON(value []byte) error {
	err := json.Unmarshal(value, v.Int64)
	v.Valid = err == nil
	return err
}
func (v *NullInt64) MarshalJSON() ([]byte, error) {
	if !v.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(v.Int64)
}

func (v *NullBool) UnmarshalJSON(value []byte) error {
	err := json.Unmarshal(value, v.Bool)
	v.Valid = err == nil
	return err
}
func (v *NullBool) MarshalJSON() ([]byte, error) {
	if !v.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(v.Bool)
}

func (v *NullFloat64) UnmarshalJSON(value []byte) error {
	err := json.Unmarshal(value, v.Float64)
	v.Valid = err == nil
	return err
}
func (v *NullFloat64) MarshalJSON() ([]byte, error) {
	if !v.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(v.Float64)
}

func (v *NullString) UnmarshalJSON(value []byte) error {
	err := json.Unmarshal(value, v.String)
	v.Valid = err == nil
	return err
}
func (v *NullString) MarshalJSON() ([]byte, error) {
	if !v.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(v.String)
}

func (v *NullTime) UnmarshalJSON(value []byte) error {
	err := json.Unmarshal(value, v.Time)
	v.Valid = err == nil
	return err
}
func (v *NullTime) MarshalJSON() ([]byte, error) {
	if !v.Valid {
		return json.Marshal(nil)
	}
	return json.Marshal(v.Time)
}
