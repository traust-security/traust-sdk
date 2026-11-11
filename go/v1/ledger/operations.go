package ledger

var (
	countersignOp = submissionOp[CountersignInput]{kind: KindCountersign}
	severityOp    = submissionOp[SeverityInput]{kind: KindSeverity}
)
