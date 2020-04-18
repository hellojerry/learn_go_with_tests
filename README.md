## assorted stuff

1. race condition detector: go test -race
2. benchmarker: go test -bench=.

: tests need to be in a file labeled "xx_test.go"
- test function must start with Test
- test function only takes in one arg: t \*test.ing.T (it's a framework hook)
