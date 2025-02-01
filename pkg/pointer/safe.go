// This file contains functionalities to safely dereference a pointer to a primitive type.
// If the pointer is nil, it will return the zero or default value of the primitive type.
package pointer

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

// SafeString ...
func SafeString(s *string) string {
	if s != nil {
		return strings.TrimSpace(*s)
	}
	return ""
}

// SafeEnumToString safely convert pointer of string alias into string
func SafeEnumToString[t ~string](s *t) string {
	if s != nil {
		return strings.TrimSpace(string(*s))
	}
	return ""
}

func SafeUuid(s *uuid.UUID) uuid.UUID {
	if s != nil {
		return *s
	}
	return uuid.Nil
}

// SafeInt ...
func SafeInt(i *int) int {
	if i != nil {
		return *i
	}
	return 0
}

// SafeInt32 ...
func SafeInt32(i *int32) int32 {
	if i != nil {
		return *i
	}
	return 0
}

// SafeInt64 ...
func SafeInt64(i *int64) int64 {
	if i != nil {
		return *i
	}
	return 0
}

// DEPRECATED: use SafeBool instead
func SafeBoolean(b *bool) bool {
	if b != nil {
		return *b
	}
	return false
}

func SafeBool(b *bool) bool {
	if b != nil {
		return *b
	}
	return false
}

func SafePointerUUIDAsString(i *uuid.UUID) string {
	if i != nil && *i != uuid.Nil {
		return i.String()
	}

	return ""
}

func SafePointerUUIDAsPointerString(i *uuid.UUID) *string {
	if i != nil && *i != uuid.Nil {
		return String(i.String())
	}

	return nil
}

// / NewArrString
func NewArrString(s []string) *[]string {
	return &s
}

// SafeArrString convert pointer array of string into array of string
func SafeArrString(arr *[]string) []string {
	out := make([]string, 0)
	if arr != nil {
		for _, s := range *arr {
			so := strings.TrimSpace(s)
			if so != "" {
				out = append(out, so)
			}
		}
	}
	return out
}

// SafeArrEnumToArrString convert pointer array of enum into array of string
func SafeArrEnumToArrString[t ~string](arr *[]t) []string {
	out := make([]string, 0)
	if arr != nil {
		for _, s := range *arr {
			so := strings.TrimSpace(string(s))
			if so != "" {
				out = append(out, so)
			}
		}
	}
	return out
}

// SafeConvertArrOfStringAlias safely convert a pointer array of string alias into another array of string alias of different type
func SafeConvertArrOfStringAlias[T1, T2 ~string](arr *[]T1) []T2 {
	out := make([]T2, 0)
	if arr != nil {
		for _, s := range *arr {
			so := strings.TrimSpace(string(s))
			if so != "" {
				out = append(out, T2(so))
			}
		}
	}
	return out
}

// SafeFloat32 ...
func SafeFloat32(f *float32) float32 {
	if f != nil {
		return *f
	}
	return 0
}

// SafeFloat64 ...
func SafeFloat64(f *float64) float64 {
	if f != nil {
		return *f
	}
	return 0
}

// SafeOf this is function for safely dereference a pointer
func SafeOf[Value any](v *Value) (output Value) {
	if v != nil {
		return *v
	}
	return
}

// Safe unix if nil set to now
func SafeTimeUnixNilToNow(date *int64) int64 {
	resp := time.Now().UnixMilli()
	if date != nil && *date > 0 {
		resp = *date
	}
	return resp
}

func SafeStringWithDefault(s *string, defaultVal string) string {
	if s != nil {
		return *s
	}
	return defaultVal
}

func SafeInt64WithDefault(i *int64, defaultVal int64) int64 {
	if i != nil {
		return *i
	}
	return defaultVal
}

func SafeUuidWithDefault(i *uuid.UUID, defaultVal uuid.UUID) uuid.UUID {
	if i != nil {
		return *i
	}
	return defaultVal
}

func SafeArrUuidWithDefault(u *[]uuid.UUID, defaultVal []uuid.UUID) []uuid.UUID {
	if u != nil {
		return *u
	}
	return defaultVal
}

// SafeOf this is function for safely dereference a pointer
func SafeOfWithDefault[Value any](v *Value, defaultValue Value) (output Value) {
	if v != nil {
		return *v
	}
	return defaultValue
}
