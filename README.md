# go-achist
go-achist is crawer for AtCoder user history page.

## REST API

Plese write hook url in sendMessage method.

```
go run main.go
```

## Deploy to heroku

Please change following code

```main.go
- log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
+ log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"), api.MakeHandler()))
```

## CLI

CLIツールはcliブランチに移動して，ビルドして使ってください．

### Usage

```
go-achist <userID>
```

## Dependencies

[PuerkitoBio/goquery](https://github.com/PuerkitoBio/goquery)
[go-json-rest/rest](https://github.com/ant0ine/go-json-rest)
