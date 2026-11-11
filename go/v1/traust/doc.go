// Package traust is the Go SDK for traust-core's REST API: every operation is a `(T, error)` call
// on Client, every paged listing an `…All` iterator, and a refusal a *ProblemError. Start with
// NewClient; traustest is a fake core for testing code that uses it, and ops runs work.
package traust
