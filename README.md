# Short URL Service

A simple, static redirection service that redirects recognised paths to any other URL.

## Configuration

Set environment variables:

- `HOST` and `PORT` configure the HTTP listen server
- `CSV` contains redirection data in CSV format

Environment variables can be set in the environment or a `.env` file.

> :warning: CSV data **is not** loaded from a file!

## Redirection data CSV

The CSV data must have two columns for the path and redirection URL, respectively. The first row is reserved for headings. Follow this example to get started:

```
path,url
/some/path,https://some-url.com
```

## License

See [LICENSE.md](./LICENSE.md)
