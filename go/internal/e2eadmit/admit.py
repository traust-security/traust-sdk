"""Admit one planned order, standing in for the engine's admission in `make e2e`.

Core stores a submitted order as planned (traust-core c129551d: one admission path) and
exposes no admit door; only the engine's admission queues it. The e2e runs traust-server
without an engine, so the tests admit each order they submit through Core's own
`OpsService.admit`, which is exactly the call the engine's admission makes.

usage: python admit.py <config.yaml> <order-id>
"""

import sys
from pathlib import Path
from uuid import UUID

from traust_core.v1.context import Config, Context

config, order_id = Path(sys.argv[1]), UUID(sys.argv[2])
order = Context(Config.from_file(config)).services().ops.admit(order_id)
print(order.state)
