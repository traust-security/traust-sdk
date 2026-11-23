package web

import "github.com/traust-security/traust-sdk/go/v1/traust"

// Items reads a generated optional list; nil is an empty one.
func Items[T any](list *[]T) []T {
	if list == nil {
		return nil
	}
	return *list
}

// Text reads a generated optional string; nil is an empty one.
func Text(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// Value reads a generated optional value; nil is its zero.
func Value[T any](p *T) T {
	var zero T
	if p == nil {
		return zero
	}
	return *p
}

// Severity words a finding's severity; an unmeasured one is unrated.
func Severity(s *traust.Severity) string {
	if s == nil {
		return "unrated"
	}
	return string(*s)
}

// Validity words a disposition's validity; an undecided one says so.
func Validity(d *traust.Disposition) string {
	if d == nil || d.Validity == nil {
		return "undecided"
	}
	return string(*d.Validity)
}

// Short is a digest's first characters, or none without one.
func Short(s *string) string {
	if s == nil {
		return "none"
	}
	return (*s)[:min(len(*s), 12)]
}
