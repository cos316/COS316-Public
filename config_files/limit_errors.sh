#!/usr/bin/env bash
# limit_errors.sh
#
# DESCRIPTION
#   Read in lines from stdin, sending them to stderr if they seem like errors
#   and to stdout otherwise.
#
#   This is a workaround to reduce the number of spurious error messages
#   printed during Vagrant provisioning.
#
# USAGE:
#   ./script_with_error_msgs 2>&1 | ./limit_errors

# list of keywords that cause messages to be flagged as errors
keywords=(
  abnormal
  abort
  "access denied"
  corrupt
  "could not"
  crash
  critical
  danger
  deprecate
  error
  expired
  fail
  fatal
  invalid
  "not found"
  problem
  terminate
  unable
  unexpected
  warn
)

# Collect keywords into a regex
regex=""
for word in "${keywords[@]}"; do
  if [ -z "$regex" ]; then
    regex="$word"
  else
    regex="$regex|$word"
  fi
done

while IFS='$\n' read -r line; do
    if echo "$line" | grep -Eqi "$regex"; then
      echo "$line" 1>&2
    else
      echo "$line"
    fi
done
