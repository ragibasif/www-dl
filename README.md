# www-dl

![](./assets/demo.gif)

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

## Issues

> [!update] This seems to be a HTTP header issue. Will fix soon!

```bash
./main https://stackoverflow.com/questions/34530237/find-files-in-a-directory-containing-desired-string-in-python
```

Returns: Response Status [403 Forbidden](https://developer.mozilla.org/en-US/docs/Web/HTTP/Reference/Status/403)

## Acknowledgements

- [Go by Example](https://gobyexample.com/)
- [VHS](https://github.com/charmbracelet/vhs)
