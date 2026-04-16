local fail = 'JSONNUNIT--failed';
local pass = 'JSONNUNIT--passed';

local resultIsPass(result) =
  if std.isArray(result) then
    std.all(std.map(function(r) r == pass, result))
  else if std.isString(result) then
    result == pass
  else
    error 'Invalid result type: ' + std.type(result);

local isPass(testcase) =
  std.foldl(function(acc, tc) acc && isPass(tc), testcase.tests.testcases, true) &&
  std.foldl(function(acc, r) acc && resultIsPass(r.result), testcase.tests.tests, true);

local indent(level) =
  std.repeat('  ', level);

local formatOutput(testcase, level=0) =
  std.foldl(function(acc, tc) acc + '%sDESC %s\n' % [ indent(level), tc.name ] + formatOutput(tc, level=level + 1), testcase.tests.testcases, '') +
  std.foldl(function(acc, r) acc + '%s%s %s\n' % [ indent(level), if resultIsPass(r.result) then 'PASS' else 'FAIL', r.name ], testcase.tests.tests, '');

local run = function(jsonnetunit)
  local output = formatOutput({ tests: jsonnetunit });
  if isPass({ tests: jsonnetunit }) then
    output
  else
    error 'FAIL Some tests failed.\n\n' + output;

{
  run: run,
}
