local fail = 'JSONNUNIT--failed';
local pass = 'JSONNUNIT--passed';

local resultIsPass(result) =
  if std.isArray(result) then
    std.all(std.map(function(r) r == pass, result))
  else if std.isString(result) then
    result == pass
  else
    error 'Invalid result type: ' + std.type(result);

local indent(level) = std.repeat('  ', level);

local collect(testcase, level=0) =
  local result = { pass: true, output: '', failedTests: [], successfulTests: [] };
  local tcr = std.foldl(
    function(acc, tc)
      local c = collect(tc, level=level + 1);
      {
        pass: acc.pass && c.pass,
        output: acc.output + '%s%s\n' % [ indent(level), tc.name ] + c.output,
        successfulTests: acc.successfulTests + c.successfulTests,
        failedTests: acc.failedTests + c.failedTests,
      },
    testcase.tests.testcases,
    result
  );
  local tr = std.foldl(
    function(acc, r)
      {
        local pass = resultIsPass(r.result),
        pass: acc.pass && pass,
        output: acc.output + '%s%s %s\n' % [ indent(level), if pass then '✔' else '✗', r.name ],
        successfulTests: if pass then acc.successfulTests + [r.name] else acc.successfulTests,
        failedTests: if pass then acc.failedTests else acc.failedTests + [r.name],
      },
    testcase.tests.tests,
    result
  );
  {
    pass: tcr.pass && tr.pass,
    output: tcr.output + tr.output,
    failedTests: tcr.failedTests + tr.failedTests,
    successfulTests: tcr.successfulTests + tr.successfulTests,
  };

local run = function(jsonnetunit)
  local output = collect({ tests: jsonnetunit });
  local summary = '\n%d ✔ passing\n%d ✗ failing\n' % [ std.length(output.successfulTests), std.length(output.failedTests) ];
  if output.pass then
    output.output + summary
  else
    error 'FAIL Some tests failed.\n\n' + output.output + summary;

{
  run: run,
}
