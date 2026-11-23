package exposure

import (
	"time"

	"github.com/traust-security/traust-sdk/go/v1/traust"

	"github.com/traust-security/traust-sdk/go/examples/webapp/internal/web"
)

// severity words the record's declared severity; an undeclared one is unrated.
func severity(record *traust.AdvisoryRecord) string {
	if s := web.Text(record.DeclaredSeverity); s != "" {
		return s
	}
	return "unrated"
}

// published words when the record's source published it.
func published(record *traust.AdvisoryRecord) string {
	if record.PublishedAt == nil {
		return "not yet"
	}
	return record.PublishedAt.Format(time.DateOnly)
}

// reach words whether an affected dependency is itself vulnerable or only leads to one.
func reach(a traust.AffectedDependency) string {
	if a.Vulnerable {
		return "vulnerable"
	}
	return "leads to a vulnerable package"
}
