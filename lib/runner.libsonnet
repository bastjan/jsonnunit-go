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
  local result = { pass: true, output: '' };
  local tcr = std.foldl(
    function(acc, tc)
      local c = collect(tc, level=level + 1);
      {
        pass: acc.pass && c.pass,
        output: acc.output + '%sDESC %s\n' % [ indent(level), tc.name ] + c.output,
      },
    testcase.tests.testcases,
    result
  );
  local tr = std.foldl(
    function(acc, r)
      {
        pass: acc.pass && resultIsPass(r.result),
        output: acc.output + '%s%s %s\n' % [ indent(level), if resultIsPass(r.result) then 'PASS' else 'FAIL', r.name ],
      },
    testcase.tests.tests,
    result
  );
  {
    pass: tcr.pass && tr.pass,
    output: tcr.output + tr.output,
  };

local run = function(jsonnetunit)
  local output = collect({ tests: jsonnetunit });
  if output.pass then
    output.output
  else
    error 'FAIL Some tests failed.\n\n' + output.output;

{
  run: run,
}
