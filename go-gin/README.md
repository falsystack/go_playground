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

