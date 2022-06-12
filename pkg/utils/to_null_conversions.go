package utils

import "database/sql"

// StringToNullString ...
func StringToNullString(s string) (ns sql.NullString) {
	if s != "" {
		ns.String = s
		ns.Valid = true
	}
	return ns
}

//Float64ToNullFloat64 ...
func Float64ToNullFloat64(f float64) (nf sql.NullFloat64) {
	if f != 0 {
		nf.Float64 = f
		nf.Valid = true
	}
	return nf
}

// Int64ToNullInt64 ...
func Int64ToNullInt64(i int64) (ni sql.NullInt64) {
	if i != 0 {
		ni.Int64 = i
		ni.Valid = true
	}
	return ni
}
