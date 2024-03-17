# tenntenn/panics

[![pkg.go.dev][gopkg-badge]][gopkg]

`panics` recovers panicking and convert to an error with the panic value.

```go
if err := panics.Recover(f); panics.IsRecovered(err) {
	fmt.Println(panics.Value(err))
}
```

<!-- links -->
[gopkg]: https://pkg.go.dev/github.com/tenntenn/panics
[gopkg-badge]: https://pkg.go.dev/badge/github.com/tenntenn/panics?status.svg
