# jsonnunit

https://www.jsonnunit.io/ but in pure Jsonnet and probably some extra features when I get around to it.

## Usage

```sh
jb init
jb install github.com/bastjan/jsonnunit@main

cat <<EOF > test.jsonnet
local JSONNETUNIT  = import "github.com/bastjan/jsonnunit/lib/jsonnunit.libsonnet";

JSONNETUNIT
  .describe('Test "be" functionality',
    JSONNETUNIT
      .it('tests JSONNUNIT.to.be.empty', [
        JSONNETUNIT.expect([]).to.be.empty,
        JSONNETUNIT.expect({}).to.be.empty
      ])
  )
EOF

jsonnet -S -J vendor -e '(import "github.com/bastjan/jsonnunit/lib/runner.libsonnet").run(std.extVar("t"))' --ext-code-file t=test.jsonnet
```
