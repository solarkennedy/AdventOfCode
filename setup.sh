#!/usr/bin/env bash
set -euvx

function fetchInput() {
curl "https://adventofcode.com/$1/day/$2/input" \
  -H "Cookie: session=$AOC_SESSION_TOKEN" > input
  sleep 5s
}

for year in $(seq 2015 2020); do
  for day in $(seq 1 25); do
    echo "Working on Year $year day $day..."
    mkdir -p $year/$day
    cd $year/$day
    ln -sf ../../Makefile.template Makefile 
    [[ -f input ]] || fetchInput $year $day
    cd ../../
  done
done
