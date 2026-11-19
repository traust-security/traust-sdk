# The gap log — what the example had to build itself

Version one of `examples/go` runs end to end against today's SDK. This log
records every place it had to build machinery that a plug-and-play SDK would
own. Each entry names the appliance or read-side step that closes it. When the
second version of the example is written after the appliance lands, its line
count against the 155-line baseline is the receipt.

| GAP | Felt as | The gap | Closed by |
|-----|---------|---------|-----------|
| 1 | `onboard` | Onboarding is not idempotent: a consumer invents a unique product name per run instead of asking "what do I already have?", and registering needs the label dance. | read-side S1 (resolution door) + an upsert-shaped onboarding |
| 2 | `runOneAsset` | The submitter must carry the pack's digest to pin the order to it — no start door that resolves "run this subject" for you. | appliance step 4 (`StartRunner` + `SubjectResolver`) |
| 3 | `runOneAsset` | One claim, once. A real worker loops, and the loop is the consumer's: claim cooldown, bounded lanes, fair rotation, panic containment, supervision, drain. The reference consumer in production built exactly this; its integration layer is thousands of lines. | appliance step 4 (the runner lifecycle) |
| 4 | `readOneScreen` | The dashboard read is a report-shaped fold, not a read model: fixed five-int columns, severity a `map[string]int`, no caller-chosen grouping or filtering, and `SummarizeFindings` takes exactly one repository — N repositories means N calls. | read-side S3 + S4 (typed metrics query, typed severity) |
| 5 | `main` | No drain protocol: the example reaches for a blanket timeout because "finish in-flight, then exit" sequencing is the consumer's to invent. | appliance step 4 (drain built into the runner) |
| 6 | `downloadOneReport` | The report is computed on the request path — the call blocks until the whole view is folded and rendered, and returns bytes, not a handle to a job's artifact. | read-side S6 (reports as jobs) |
| 7 | not yet felt | One `BaseURL` for everything: the client config cannot point reads and writes at different deployments, so the query/command/ops serving split is unreachable from the SDK. | appliance step 2 / schema-performance-audit (`Endpoints{Query, Command, Ops}`) |

## What the example did not have to build

The seams that already work, named so the appliance keeps them: `traust.Config`
+ `FileTokenSource`, the `…All` iterators, `ops.Pack` and its pin, the
`Executor` interface, and `RunOnce` for a single claim.
