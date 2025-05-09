# www-dl

To build:

```bash
go build main.go
```

To run:

```bash
./main <URL>
```

Example:

```bash
go build main.go
./main https://go.dev/
```

![](./assets/build.gif)

## Issues

```bash
./main https://go.dev/
```

Returns: Response Status 200 OK

```bash
./main https://stackoverflow.com/questions/34530237/find-files-in-a-directory-containing-desired-string-in-python
```

Returns: Response Status [403 Forbidden](https://developer.mozilla.org/en-US/docs/Web/HTTP/Reference/Status/403)

## Acknowledgements

- [Go by Example](https://gobyexample.com/)
- [VHS](https://github.com/charmbracelet/vhs)
