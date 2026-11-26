package exposure

import (
	"time"

	"github.com/google/uuid"

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

// impactRefs is the refs an advisory's affected dependencies name, each once, in core's order;
// core prepares an advisory's impacts one ref at a time.
func impactRefs(affected []traust.AffectedDependency) []uuid.UUID {
	refs := make([]uuid.UUID, 0, len(affected))
	seen := make(map[uuid.UUID]bool, len(affected))
	for _, a := range affected {
		if a.RepositoryRefId != nil && !seen[*a.RepositoryRefId] {
			seen[*a.RepositoryRefId] = true
			refs = append(refs, *a.RepositoryRefId)
		}
	}
	return refs
}
