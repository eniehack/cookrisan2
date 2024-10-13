# cookrisan2

## 開発環境の構築

### 必要なもの

- git
- [go](https://go.dev)
- [sql-migrate](https://github.com/rubenv/sql-migrate)
- GNU make

### 構築

1. git clone && cd
2. `sql-migrate up -config db/migrate.yml -env=dev`

- バイナリのビルド: `make`
  - `cmd/crawler`にあるものをbuildした場合、`bin/crawler`にバイナリが置かれる