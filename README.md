# greet

A tiny Go library for demonstrating
[gopkg-charmed](https://github.com/canonical/gopkg-charmed), the gopkg.in
service run as a Juju charm, on the host `gopkg-ps7.pfe.staging.canonical.com`.

```go
import greet "gopkg-ps7.pfe.staging.canonical.com/minulo/greet.v1" // branch v1: greet.Hello(name)
import greet "gopkg-ps7.pfe.staging.canonical.com/minulo/greet.v2" // branch v2: greet.Hello(name, lang)
```

The service reads this repository's branches and answers each import path with
the matching one: `.v1` gets branch `v1` and `.v2` gets branch `v2`.
There are no version tags on purpose: on a host other than gopkg.in, Go would
pick a `v1.x.y` tag for `greet.v2` too.

Branch `next-v1` is a fix waiting to be released to `v1` during the demo.
