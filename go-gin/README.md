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
## ミドルウェアとは
- リクエスト処理の前後に任意の処理を挿入出来る機能
  - ex) ロギング、認証、リクエストやレスポンスの加工
### ミドルウェアを使用するメリット
- コードの再利用性
  - ロギング、認証など一般的な機能をモジュール化し、アプリケーションの異なる部分で再利用可能
- 責務の分離
  - ビジネスロジックから共通機能を分離することで、各コード部分が特定の目的に集中することができ、コードの可読性と保守性が向上
### Ginにおけるミドルウェアの使い方
- 前からの引数から順次実行される、順番に注意が必要
- 全エンドポイントに適用する方法
  - ルーターの`Use`メソッドを使用
```go
r := gin.Default()
r.Use(middleware)
```
- 個別のエンドポイントに適用する方法
  - エンドポイントに直接ミドルウェアを渡す
```go
itemRouter.GET("", Middleware1, itemController.Findall)
itemRouter.POST("", Middleware1, Middleware2, itemController.Create)
```
- ルートグループに対して適用する方法
  - グループ作成時にミドルウェアを渡す
```go
itemRouter := r.Group("/items", Middleware1)
authRouter := r.Group("/auth", Middleware1, Middleware2)
```
### カスタムミドルウェアの作成
- 独自のミドルウェアを作成することが可能
  - gin.HandleFuncを返却する関数を定義
```go
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 前処理
		t := time.Now()
		c.Set("example", "12345")
		
		// 次のミドルウェア or 目的の処理
		c.Next()
		
		// 後処理
		latency := time.Since(t)
		log.Print(latency)
    }   
}
```
---

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

## bcrypt
hash化
```go
hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
```
hash化されたパスワードと平文のパスワードの比較
```go
err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(password))
```

## jwt
install
```shell
go get -u github.com/golang-jwt/jwt/v5
```
### Tokenの生成とサイン
```go
func createToken(userId uint, email string) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   userId, // subject : userの識別子、ここではuserId
		"email": email,
		"exp":   time.Now().Add(time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return nil, err
	}
	return &tokenString, nil
}
```
### トークンのDecode
```go
token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
        return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
    }
    return []byte(os.Getenv("SECRET_KEY")), nil
})
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

### テーブルにマイグレーション
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

## testify
### install
```shell
go get github.com/stretchr/testify
```

`.env.test`を作成

```.env
ENV=test
```


# ETC
randomな文字列生成
```shell
openssl rand -hex 32
```