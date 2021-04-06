# News App

This repository contains a sample News App implementation capable of creating an aggregator API above multiple RSS/Atom providers.

## Installing

```
go get github.com/petomalina/news-app
```

## Running the API

The server uses environment variables for configuration. Available configs for the `cmd/news` server are:
| env | desc | default |
| ---- | --- | --- |
| PORT | port on which the server should listen | 8080 |
| PROVIDERS | list of providers in a map such as `skynews:skynewsurl/%s.xml` | `[]` |

The `PROVIDERS` variable uses formatted strings to target *categories* within the provider (the `%s`). The server expects
this to be provided by the application. Values of categories must be known by the app and are not provided by the server. For docs on how to
create maps with multiple values, see [github.com/sethvargo/go-envconfig](https://github.com/sethvargo/go-envconfig#maps).

An example of running a server with the Sky News Provider:
```shell
PROVIDERS=SkyNews:http://feeds.skynews.com/feeds/rss/%s.xml go run ./cmd/news/main.go
```

## API

The API has 3 endpoints as described in the table below:
| endpoint | desc/params |
| --- | --- |
| /health | provides a health-check for orchestrators |
| /providers | provides a list of registered providers |
| /fetch | query params: `p` (required) is provider, `c` (required) is category, `sort=` is time sort (only `desc` supported) - otherwise sorted by relevance |
