## What

<!-- Which domain package, runtime piece (config, transport, errors, Poller) or example-app flow does this add or change? -->

## Checklist

- [ ] Targets `development`
- [ ] It's a regenerated `v1/traust`, a runtime piece every consumer needs, or an example-app flow; a missing endpoint went to traust-core instead
- [ ] No decision core owns (authorization, ordering, identity, IDs, time) derived, cached or defaulted here
- [ ] Generated code regenerated, not edited; no spec or hand-written type/enum committed
- [ ] Any non-answer stops with a typed error; no fallbacks
- [ ] Types come only from core's OpenAPI; nothing consumer-specific
- [ ] Comments only for a non-obvious why; doc comments one or two lines; no status prose in READMEs
- [ ] Table-driven tests cover the edges (nil deps, cancellation, each problem status, malformed body)
- [ ] `make lint test guardrails lean`, `make -C go check-drift e2e` green

## Changed tests

<!-- test · old → new · why, for each edited or deleted existing test. Leave empty if you only added tests. -->

- [ ] No existing test changed
- [ ] A person reviewed each changed test and the behaviour change is intended
