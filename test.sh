#!/bin/bash

set -u

for test in tests/*.jsonnet; do
  echo "Running test: $test"
  jsonnet -S "$test" > /dev/null 2>&1 || {
    if grep -q '_fail.jsonnet$' <<< "$test"; then
      echo "Failed as expected: $test"
    else
      echo "Test failed unexpectedly: $test"
      exit 1
    fi
  }
done
