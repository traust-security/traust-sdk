from hashlib import sha256
from pathlib import Path
from sys import argv

from traust_core.v1.context import Context
from traust_core.v1.models.identity import IdentitySource, Principal
from traust_core.v1.models.operations import ConfigDigest, OrderRequest
from traust_core.v1.models.vocabulary import ActorKind

config_path, pack_path, token_path = map(Path, argv[1:])
context = Context.from_file(config_path)
principal = Principal(
    identity="operator",
    kind=ActorKind.HUMAN,
    display="operator",
    verified_by=IdentitySource.LOCAL_USER,
)
digest = sha256(pack_path.read_bytes()).hexdigest()
service = context.services().ops
# Plan/admit has no REST door; only this test seed calls Core's owning service.
planned = service.plan(
    OrderRequest(asset="command_result", priority=1),
    principal,
    [ConfigDigest(name="pack:smoke", digest=digest)],
)
service.admit(planned.id)
token_path.write_text(
    context.local_issuer.mint("smoke-runner", machine=True).token.get_secret_value()
)
token_path.chmod(0o600)
