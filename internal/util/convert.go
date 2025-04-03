package util

import "github.com/jackc/pgx/v5/pgtype"

// PgTextToString converts a pgtype.Text to a string.
// Returns empty string if the pgtype.Text is null.
func PgTextToString(pgText pgtype.Text) string {
	return pgText.String
}

// PgIntToInt32 converts a pgtype.Int4 to int32.
// Returns 0 if the pgtype.Int4 is null.
func PgIntToInt32(pgInt pgtype.Int4) int32 {
	return pgInt.Int32
}

// StringToPgText converts a string to pgtype.Text.
func StringToPgText(s string) pgtype.Text {
	var pgText pgtype.Text
	pgText.String = s
	pgText.Valid = s != ""
	return pgText
}

// Int32ToPgInt converts an int32 to pgtype.Int4.
func Int32ToPgInt(i int32) pgtype.Int4 {
	var pgInt pgtype.Int4
	pgInt.Int32 = i
	pgInt.Valid = i != 0
	return pgInt
}
