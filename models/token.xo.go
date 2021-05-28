// Package models contains the types for schema ''.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"
)

// Token represents a row from 'tokens'.
type Token struct {
	ID                     sql.NullInt64  `json:"id" db:"id"`                                         // id
	Artist                 sql.NullString `json:"artist" db:"artist"`                                 // artist
	Asciiname              sql.NullString `json:"asciiName" db:"asciiName"`                           // asciiName
	Availability           sql.NullString `json:"availability" db:"availability"`                     // availability
	Bordercolor            sql.NullString `json:"borderColor" db:"borderColor"`                       // borderColor
	Coloridentity          sql.NullString `json:"colorIdentity" db:"colorIdentity"`                   // colorIdentity
	Colors                 sql.NullString `json:"colors" db:"colors"`                                 // colors
	Edhrecrank             sql.NullInt64  `json:"edhrecRank" db:"edhrecRank"`                         // edhrecRank
	Facename               sql.NullString `json:"faceName" db:"faceName"`                             // faceName
	Flavortext             sql.NullString `json:"flavorText" db:"flavorText"`                         // flavorText
	Frameeffects           sql.NullString `json:"frameEffects" db:"frameEffects"`                     // frameEffects
	Frameversion           sql.NullString `json:"frameVersion" db:"frameVersion"`                     // frameVersion
	Hasfoil                int            `json:"hasFoil" db:"hasFoil"`                               // hasFoil
	Hasnonfoil             int            `json:"hasNonFoil" db:"hasNonFoil"`                         // hasNonFoil
	Isfullart              int            `json:"isFullArt" db:"isFullArt"`                           // isFullArt
	Ispromo                int            `json:"isPromo" db:"isPromo"`                               // isPromo
	Isreprint              int            `json:"isReprint" db:"isReprint"`                           // isReprint
	Keywords               sql.NullString `json:"keywords" db:"keywords"`                             // keywords
	Layout                 sql.NullString `json:"layout" db:"layout"`                                 // layout
	Mcmid                  sql.NullString `json:"mcmId" db:"mcmId"`                                   // mcmId
	Mtgarenaid             sql.NullString `json:"mtgArenaId" db:"mtgArenaId"`                         // mtgArenaId
	Mtgjsonv4id            sql.NullString `json:"mtgjsonV4Id" db:"mtgjsonV4Id"`                       // mtgjsonV4Id
	Multiverseid           sql.NullString `json:"multiverseId" db:"multiverseId"`                     // multiverseId
	Name                   sql.NullString `json:"name" db:"name"`                                     // name
	Number                 sql.NullString `json:"number" db:"number"`                                 // number
	Originaltext           sql.NullString `json:"originalText" db:"originalText"`                     // originalText
	Originaltype           sql.NullString `json:"originalType" db:"originalType"`                     // originalType
	Power                  sql.NullString `json:"power" db:"power"`                                   // power
	Promotypes             sql.NullString `json:"promoTypes" db:"promoTypes"`                         // promoTypes
	Reverserelated         sql.NullString `json:"reverseRelated" db:"reverseRelated"`                 // reverseRelated
	Scryfallid             sql.NullString `json:"scryfallId" db:"scryfallId"`                         // scryfallId
	Scryfallillustrationid sql.NullString `json:"scryfallIllustrationId" db:"scryfallIllustrationId"` // scryfallIllustrationId
	Scryfalloracleid       sql.NullString `json:"scryfallOracleId" db:"scryfallOracleId"`             // scryfallOracleId
	Setcode                sql.NullString `json:"setCode" db:"setCode"`                               // setCode
	Side                   sql.NullString `json:"side" db:"side"`                                     // side
	Subtypes               sql.NullString `json:"subtypes" db:"subtypes"`                             // subtypes
	Supertypes             sql.NullString `json:"supertypes" db:"supertypes"`                         // supertypes
	Tcgplayerproductid     sql.NullString `json:"tcgplayerProductId" db:"tcgplayerProductId"`         // tcgplayerProductId
	Text                   sql.NullString `json:"text" db:"text"`                                     // text
	Toughness              sql.NullString `json:"toughness" db:"toughness"`                           // toughness
	Type                   sql.NullString `json:"type" db:"type"`                                     // type
	Types                  sql.NullString `json:"types" db:"types"`                                   // types
	UUID                   string         `json:"uuid" db:"uuid"`                                     // uuid
	Watermark              sql.NullString `json:"watermark" db:"watermark"`                           // watermark

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Token exists in the database.
func (t *Token) Exists() bool {
	return t._exists
}

// Deleted provides information if the Token has been deleted from the database.
func (t *Token) Deleted() bool {
	return t._deleted
}

// Insert inserts the Token to the database.
func (t *Token) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if t._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by sequence
	const sqlstr = `INSERT INTO tokens (` +
		`artist, asciiName, availability, borderColor, colorIdentity, colors, edhrecRank, faceName, flavorText, frameEffects, frameVersion, hasFoil, hasNonFoil, isFullArt, isPromo, isReprint, keywords, layout, mcmId, mtgArenaId, mtgjsonV4Id, multiverseId, name, number, originalText, originalType, power, promoTypes, reverseRelated, scryfallId, scryfallIllustrationId, scryfallOracleId, setCode, side, subtypes, supertypes, tcgplayerProductId, text, toughness, type, types, uuid, watermark` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`) RETURNING id`

	// run query
	XOLog(sqlstr, t.Artist, t.Asciiname, t.Availability, t.Bordercolor, t.Coloridentity, t.Colors, t.Edhrecrank, t.Facename, t.Flavortext, t.Frameeffects, t.Frameversion, t.Hasfoil, t.Hasnonfoil, t.Isfullart, t.Ispromo, t.Isreprint, t.Keywords, t.Layout, t.Mcmid, t.Mtgarenaid, t.Mtgjsonv4id, t.Multiverseid, t.Name, t.Number, t.Originaltext, t.Originaltype, t.Power, t.Promotypes, t.Reverserelated, t.Scryfallid, t.Scryfallillustrationid, t.Scryfalloracleid, t.Setcode, t.Side, t.Subtypes, t.Supertypes, t.Tcgplayerproductid, t.Text, t.Toughness, t.Type, t.Types, t.UUID, t.Watermark)
	err = db.QueryRow(sqlstr, t.Artist, t.Asciiname, t.Availability, t.Bordercolor, t.Coloridentity, t.Colors, t.Edhrecrank, t.Facename, t.Flavortext, t.Frameeffects, t.Frameversion, t.Hasfoil, t.Hasnonfoil, t.Isfullart, t.Ispromo, t.Isreprint, t.Keywords, t.Layout, t.Mcmid, t.Mtgarenaid, t.Mtgjsonv4id, t.Multiverseid, t.Name, t.Number, t.Originaltext, t.Originaltype, t.Power, t.Promotypes, t.Reverserelated, t.Scryfallid, t.Scryfallillustrationid, t.Scryfalloracleid, t.Setcode, t.Side, t.Subtypes, t.Supertypes, t.Tcgplayerproductid, t.Text, t.Toughness, t.Type, t.Types, t.UUID, t.Watermark).Scan(&t.ID)
	if err != nil {
		return err
	}

	// set existence
	t._exists = true

	return nil
}

// Update updates the Token in the database.
func (t *Token) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !t._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if t._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE tokens SET (` +
		`artist, asciiName, availability, borderColor, colorIdentity, colors, edhrecRank, faceName, flavorText, frameEffects, frameVersion, hasFoil, hasNonFoil, isFullArt, isPromo, isReprint, keywords, layout, mcmId, mtgArenaId, mtgjsonV4Id, multiverseId, name, number, originalText, originalType, power, promoTypes, reverseRelated, scryfallId, scryfallIllustrationId, scryfallOracleId, setCode, side, subtypes, supertypes, tcgplayerProductId, text, toughness, type, types, uuid, watermark` +
		`) = ( ` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`) WHERE id = $44`

	// run query
	XOLog(sqlstr, t.Artist, t.Asciiname, t.Availability, t.Bordercolor, t.Coloridentity, t.Colors, t.Edhrecrank, t.Facename, t.Flavortext, t.Frameeffects, t.Frameversion, t.Hasfoil, t.Hasnonfoil, t.Isfullart, t.Ispromo, t.Isreprint, t.Keywords, t.Layout, t.Mcmid, t.Mtgarenaid, t.Mtgjsonv4id, t.Multiverseid, t.Name, t.Number, t.Originaltext, t.Originaltype, t.Power, t.Promotypes, t.Reverserelated, t.Scryfallid, t.Scryfallillustrationid, t.Scryfalloracleid, t.Setcode, t.Side, t.Subtypes, t.Supertypes, t.Tcgplayerproductid, t.Text, t.Toughness, t.Type, t.Types, t.UUID, t.Watermark, t.ID)
	_, err = db.Exec(sqlstr, t.Artist, t.Asciiname, t.Availability, t.Bordercolor, t.Coloridentity, t.Colors, t.Edhrecrank, t.Facename, t.Flavortext, t.Frameeffects, t.Frameversion, t.Hasfoil, t.Hasnonfoil, t.Isfullart, t.Ispromo, t.Isreprint, t.Keywords, t.Layout, t.Mcmid, t.Mtgarenaid, t.Mtgjsonv4id, t.Multiverseid, t.Name, t.Number, t.Originaltext, t.Originaltype, t.Power, t.Promotypes, t.Reverserelated, t.Scryfallid, t.Scryfallillustrationid, t.Scryfalloracleid, t.Setcode, t.Side, t.Subtypes, t.Supertypes, t.Tcgplayerproductid, t.Text, t.Toughness, t.Type, t.Types, t.UUID, t.Watermark, t.ID)
	return err
}

// Save saves the Token to the database.
func (t *Token) Save(db XODB) error {
	if t.Exists() {
		return t.Update(db)
	}

	return t.Insert(db)
}

// Upsert performs an upsert for Token.
//
// NOTE: PostgreSQL 9.5+ only
func (t *Token) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if t._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO tokens (` +
		`id, artist, asciiName, availability, borderColor, colorIdentity, colors, edhrecRank, faceName, flavorText, frameEffects, frameVersion, hasFoil, hasNonFoil, isFullArt, isPromo, isReprint, keywords, layout, mcmId, mtgArenaId, mtgjsonV4Id, multiverseId, name, number, originalText, originalType, power, promoTypes, reverseRelated, scryfallId, scryfallIllustrationId, scryfallOracleId, setCode, side, subtypes, supertypes, tcgplayerProductId, text, toughness, type, types, uuid, watermark` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`) ON CONFLICT (id) DO UPDATE SET (` +
		`id, artist, asciiName, availability, borderColor, colorIdentity, colors, edhrecRank, faceName, flavorText, frameEffects, frameVersion, hasFoil, hasNonFoil, isFullArt, isPromo, isReprint, keywords, layout, mcmId, mtgArenaId, mtgjsonV4Id, multiverseId, name, number, originalText, originalType, power, promoTypes, reverseRelated, scryfallId, scryfallIllustrationId, scryfallOracleId, setCode, side, subtypes, supertypes, tcgplayerProductId, text, toughness, type, types, uuid, watermark` +
		`) = (` +
		`EXCLUDED.id, EXCLUDED.artist, EXCLUDED.asciiName, EXCLUDED.availability, EXCLUDED.borderColor, EXCLUDED.colorIdentity, EXCLUDED.colors, EXCLUDED.edhrecRank, EXCLUDED.faceName, EXCLUDED.flavorText, EXCLUDED.frameEffects, EXCLUDED.frameVersion, EXCLUDED.hasFoil, EXCLUDED.hasNonFoil, EXCLUDED.isFullArt, EXCLUDED.isPromo, EXCLUDED.isReprint, EXCLUDED.keywords, EXCLUDED.layout, EXCLUDED.mcmId, EXCLUDED.mtgArenaId, EXCLUDED.mtgjsonV4Id, EXCLUDED.multiverseId, EXCLUDED.name, EXCLUDED.number, EXCLUDED.originalText, EXCLUDED.originalType, EXCLUDED.power, EXCLUDED.promoTypes, EXCLUDED.reverseRelated, EXCLUDED.scryfallId, EXCLUDED.scryfallIllustrationId, EXCLUDED.scryfallOracleId, EXCLUDED.setCode, EXCLUDED.side, EXCLUDED.subtypes, EXCLUDED.supertypes, EXCLUDED.tcgplayerProductId, EXCLUDED.text, EXCLUDED.toughness, EXCLUDED.type, EXCLUDED.types, EXCLUDED.uuid, EXCLUDED.watermark` +
		`)`

	// run query
	XOLog(sqlstr, t.ID, t.Artist, t.Asciiname, t.Availability, t.Bordercolor, t.Coloridentity, t.Colors, t.Edhrecrank, t.Facename, t.Flavortext, t.Frameeffects, t.Frameversion, t.Hasfoil, t.Hasnonfoil, t.Isfullart, t.Ispromo, t.Isreprint, t.Keywords, t.Layout, t.Mcmid, t.Mtgarenaid, t.Mtgjsonv4id, t.Multiverseid, t.Name, t.Number, t.Originaltext, t.Originaltype, t.Power, t.Promotypes, t.Reverserelated, t.Scryfallid, t.Scryfallillustrationid, t.Scryfalloracleid, t.Setcode, t.Side, t.Subtypes, t.Supertypes, t.Tcgplayerproductid, t.Text, t.Toughness, t.Type, t.Types, t.UUID, t.Watermark)
	_, err = db.Exec(sqlstr, t.ID, t.Artist, t.Asciiname, t.Availability, t.Bordercolor, t.Coloridentity, t.Colors, t.Edhrecrank, t.Facename, t.Flavortext, t.Frameeffects, t.Frameversion, t.Hasfoil, t.Hasnonfoil, t.Isfullart, t.Ispromo, t.Isreprint, t.Keywords, t.Layout, t.Mcmid, t.Mtgarenaid, t.Mtgjsonv4id, t.Multiverseid, t.Name, t.Number, t.Originaltext, t.Originaltype, t.Power, t.Promotypes, t.Reverserelated, t.Scryfallid, t.Scryfallillustrationid, t.Scryfalloracleid, t.Setcode, t.Side, t.Subtypes, t.Supertypes, t.Tcgplayerproductid, t.Text, t.Toughness, t.Type, t.Types, t.UUID, t.Watermark)
	if err != nil {
		return err
	}

	// set existence
	t._exists = true

	return nil
}

// Delete deletes the Token from the database.
func (t *Token) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !t._exists {
		return nil
	}

	// if deleted, bail
	if t._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM tokens WHERE id = $1`

	// run query
	XOLog(sqlstr, t.ID)
	_, err = db.Exec(sqlstr, t.ID)
	if err != nil {
		return err
	}

	// set deleted
	t._deleted = true

	return nil
}

// TokenByID retrieves a row from 'tokens' as a Token.
//
// Generated from index 'tokens_id_pkey'.
func TokenByID(db XODB, id sql.NullInt64) (*Token, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, artist, asciiName, availability, borderColor, colorIdentity, colors, edhrecRank, faceName, flavorText, frameEffects, frameVersion, hasFoil, hasNonFoil, isFullArt, isPromo, isReprint, keywords, layout, mcmId, mtgArenaId, mtgjsonV4Id, multiverseId, name, number, originalText, originalType, power, promoTypes, reverseRelated, scryfallId, scryfallIllustrationId, scryfallOracleId, setCode, side, subtypes, supertypes, tcgplayerProductId, text, toughness, type, types, uuid, watermark ` +
		`FROM tokens ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	t := Token{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&t.ID, &t.Artist, &t.Asciiname, &t.Availability, &t.Bordercolor, &t.Coloridentity, &t.Colors, &t.Edhrecrank, &t.Facename, &t.Flavortext, &t.Frameeffects, &t.Frameversion, &t.Hasfoil, &t.Hasnonfoil, &t.Isfullart, &t.Ispromo, &t.Isreprint, &t.Keywords, &t.Layout, &t.Mcmid, &t.Mtgarenaid, &t.Mtgjsonv4id, &t.Multiverseid, &t.Name, &t.Number, &t.Originaltext, &t.Originaltype, &t.Power, &t.Promotypes, &t.Reverserelated, &t.Scryfallid, &t.Scryfallillustrationid, &t.Scryfalloracleid, &t.Setcode, &t.Side, &t.Subtypes, &t.Supertypes, &t.Tcgplayerproductid, &t.Text, &t.Toughness, &t.Type, &t.Types, &t.UUID, &t.Watermark)
	if err != nil {
		return nil, err
	}

	return &t, nil
}
