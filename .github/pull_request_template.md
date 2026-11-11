## What

<!-- Which extension point, lifecycle step or core endpoint does this add or change? -->

## Checklist

- [ ] Targets `development`
- [ ] It's an extension point the SDK calls, a lifecycle step in `ops`, or a regenerated `v1/traust`; a missing endpoint went to traust-core instead
- [ ] No decision core owns (authorization, ordering, identity, IDs, time) derived, cached or defaulted here
- [ ] Generated code regenerated, not edited; no spec or hand-written type/enum committed
- [ ] Any non-answer stops with a typed error; no fallbacks
- [ ] Nothing added to frozen packages (`storage`, `ledger`, `skills`, `types`, `enums`, `validate`) except bug fixes
- [ ] Comments only for a non-obvious why; doc comments one or two lines; no status prose in READMEs
- [ ] Table-driven tests cover the edges (nil deps, cancellation, each problem status, malformed body)
- [ ] `make lint test security guardrails` green

## Changed tests

<!-- test · old → new · why, for each edited or deleted existing test. Leave empty if you only added tests. -->

- [ ] No existing test changed
- [ ] A person reviewed each changed test and the behaviour change is intended
