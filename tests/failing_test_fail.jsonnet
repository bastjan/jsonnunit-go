local JSONNUNIT = import '../lib/jsonnunit.libsonnet';

local runner = import '../lib/runner.libsonnet';

runner.run(
  JSONNUNIT
  .describe(
    'Test "be" functionality',
    JSONNUNIT
    .it('tests JSONNUNIT.to.be.empty', [
      JSONNUNIT.expect([]).to.be.empty,
      JSONNUNIT.expect([ 1 ]).to.be.empty,
    ])
  )
)
