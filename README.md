# Repro for charmbracelet/bubbletea#964

## Steps to reproduce

1. Checkout this code
1. Run test: `go test -count=1 ./spin2sec_test.go`
1. You should see `TestSpin2SecFailing` test failing, with error:
   ```
   error creating cancelreader: add reader with descriptor 0 to epoll interest list: operation not permitted
   ```