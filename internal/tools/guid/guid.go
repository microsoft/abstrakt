package guid

////////////////////////////////////////////////////////////
// util -- misc. helpers
//
// 1. GUID - a type alias for strings used as IDs
// 2. tolerateMiscasedKey - controls whether lookup methods
//    in this package will compensate for names which differ
//    only in casing -- e.g. "abc" is equivalent to "Abc".
//
////////////////////////////////////////////////////////////

import (
	"strings"
)

// If true, tolerate Find requests where the case is incorrect.
// e.g. if asked for Name="abc", then okay to return the object
// named "Abc", if found.  If false, use only exact name matching.

//TolerateMiscasedKey -- how strict the yaml parsing should be on case sensitivity
const TolerateMiscasedKey = true

// For ID fields, we currently use strings containing GUID-formatted
// information (e.g., "d6e4a5e9-696a-4626-ba7a-534d6ff450a5").
// For now, it's enough to treat it like a string, but might as well
// introduce the type to identify what's just a string and what is
// an ID.

// GUID -- an alias for an ID
type GUID string

// EmptyGUID -- equivalent to an uninitialized GUID.
const EmptyGUID = GUID("")

// IsEmpty - true if GUID represents empty value.
func (LHS GUID) IsEmpty() bool {
	return EmptyGUID.Equals(LHS)
}

// Equals -- true if two GUIDs compare equal.  Case differences are tolerated.
func (LHS GUID) Equals(RHS GUID) bool {
	if LHS == EmptyGUID && RHS == EmptyGUID {
		return true
	}
	if LHS == EmptyGUID || RHS == EmptyGUID {
		return false
	}

	if LHS == RHS { // exact match first...
		return true
	}

	// tolerate case mismatches...
	if strings.EqualFold(string(LHS), string(RHS)) {
		return true
	}

	return false
}
