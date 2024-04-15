# Ginとは
- Go言語のためのWebフレームワーク
    - ルーティングだけではなく多くの機能を含む
    - REST APIを非常に簡単に構築することが可能
- Go言語で作られおり優れたパフォーマンス
- Go言語のWebフレームワークの中では比較的歴史があり、今でも人気が高い

## Ginの特徴
- 軽量で高速
  - Go言語の効率的な性能を活用し、非常に高いパフォーマンス
- シンプルなAPI設計
  - シンプルで直感的な記述により簡単にAPIが実装可能
- 強力なルーティング機能
  - パラメータ付きURL、グループ化、URLパターンマッチングなど複雑なルーティングに対応
- ミドルウェア
  - ミドルウェアをサポートしており、ロギング、認証、CORS処理などの共通処理を簡単に追加可能
- 拡張性
  - カスタムミドルウェアや外部ライブラリと簡単に統合が可能

## Ginのインストール
https://gin-gonic.com/ja/docs/quickstart/

```shell
go get -u github.com/gin-gonic/gin
```

# 機能
## パスパラメータ
### パスパラメータ定義
`:parameterName`で定義
```go
r.GET("/items/:id", itemController.FindById)
```
### パスパラメータ利用
パスパラメータは常に `string` typeになる
```go
itemId := ctx.Param("id")
```

## クエリパラメータ
### クエリパラメータの定義
特別のキーワードの設定は不要
```go
r.GET("/items", itemController.FIndByName)
```
### クエリパラメータの利用
```go
name := ctx.Query("name")
```
`?name=kakao`の場合 `kakao`が取得できる


## バリデーション
Ginでバリデーションを使う方法
1. リクエストデータの構造体（DTO）を定義
```go
type CreateItemInput struct {
	Name string `json:"name"`
	Price uint `json:"price"`
	Description string `json:"description"`
}
```
2.  DTOにバリデーションのためのタグを追加
```go
type CreateItemInput struct {
    Name string `json:"name" binding:"required,min=2"`
    Price uint `json:"price" binding:"required,min=1,max=999999"`
    Description string `json:"description"`
}
```
3. リクエストデータをDTOにバインド
```go
ctx.ShouldBindJSON(&input)
```

# ライブラリ
## air
Hot Reloadができるようにしてくれるライブラリ
- https://github.com/cosmtrek/air

### 初期化
```shell
air init
```
`command not found`が出る場合
`.zshrc` にパスを指定
```zsh
export PATH="go env GOPATHで確認したパス/bin:$PATH"
```

## godotenv
- `.env`ファイルをgoでも使用可能にする。
- `godotenv.Load()` の引数に何も指定しない場合基本値として `.env` が使われる

```go
func Init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
```
# フレームワーク
## gorm
install
```shell
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite
go get -u gorm.io/driver/postgres
```
`.env` ファイルに接続情報設定

```env
ENV=prod
DB_HOST=localhost
DB_USER=ginuser
DB_PASSWORD=ginpassword
DB_NAME=fleamarket
DB_PORT=5432
```

databaseへの接続設定
```go
func SetupDB() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=$s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	return db
}
```

### 
`gorm.Model`には以下の情報が含む構造体である
```go
type Model struct {
    ID        uint `gorm:"primarykey"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt DeletedAt `gorm:"index"`
}
```
`gorm:""`タグを利用してマッピングすることが可能
```go
type Item struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Price       uint `gorm:"not null"`
	Description string
	SoldOut     bool `gorm:"not null;default:false"`
}
```
migrationを行う

mainパッケージでmain関数を作ってやる理由はアプリケーションの実行とmigrationを分けるためである。
```go
package main

import (
	"go-gin/infra"
	"go-gin/models"
)

func main() {
	infra.Init()
	db := infra.SetupDB()

	if err := db.AutoMigrate(&models.Item{}); err != nil {
		panic("Failed to migrate database")
	}
}
```