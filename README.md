# Shorty

A simple redirection service that routes any path to any URL - or in brief, a short URL service, depending on the length of your domain name.

## System requirements

- [Go v1.21](https://go.dev/dl/)

## Quickstart

Build the app first:

```sh
go build -o shorty
```

You can use `--help` or `-h` with any command to learn more about Shorty usage.

```sh
./shorty --help
```

Start the app:

```sh
./shorty start
```

## Configuration

Shorty can be configured via CLI flags or environment variables. Refer to help text for more detail.

The most important flag is `--token|-t|SHORTY_TOKEN` which is unset by default, meaning that authentication is not enabled and anyone can modify your redirects. Set a token secretly and add it to your client to prevent unauthenticated access.

## Usage

You can use the following commands to manage your redirects:

- `shorty add <path> <destination>` adds a redirect
- `shorty rm <path>` removes a redirect
- `shorty get <path>` gets the destination of a redirect

These commands connect to the running Shorty instance, so you need to make sure you have started it first. Make sure to set `--token|-t|SHORTY_TOKEN` as well if you have enabled token authentication.

If you provide `--url|-u|SHORTY_URL` then you can use Shorty as a remote CLI. By default it will connect to the local Shorty instance.

Since all Shorty interaction goes through a REST API, it's easy to use another client (such as [Postman](https://www.postman.com/)) or develop your own client to manage Shorty. This is left as an exercise for the reader.

## Docker usage

This project is packaged as a Docker image in [GitHub Container Registry](https://github.com/annybs/shorty/pkgs/container/shorty) for ease of use.

```sh
docker run --rm -ti -v './your-data-dir:/shorty/data' -p '3000:3000' ghcr.io/annybs/shorty:develop
```

## License

See [LICENSE.md](./LICENSE.md)
