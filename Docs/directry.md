# ディレクトリ構成
## 構成一覧
```tree
.
├── cmd
│   └── server
│       └── main.go
├── go.mod
├── go.sum
├── infrastructure
│   ├── config
│   ├── db
│   ├── di
│   ├── logger
│   └── repository
└── internal
    ├── application
    │   ├── interface
    │   └── usecase
    └── domain
        ├── output
        │   └── entity
        │       ├── repository
        │       └── services
        └── user
            ├── entity
            │   └── user.go
            ├── repository
            │   └── user_repository.go
            ├── services
            │   └── user_service.go
            └── value_obj
                └── role.go
```

---

## 詳細
### cmd
#### server
プロジェクトのメインディレクトリ
main.goをここに配置し、起動する際は ./cmd/server内で実行コマンドを入力する

--- 

### infrastructure
データベース、設定、外部APIなどの外部リソースとの連携を行う。
#### config
各設定情報を管理する
#### db
データベース情報を管理する
#### di
依存性注入の管理を行う
今回はwireを使用する
#### logger
log設定情報を管理する
#### repository
各リポジトリを管理する
ドメインごとにリポジトリを分けるので、用意

---

## internal
外部に公開しない業務ロジックやサービスをここに記載する
### application
ユーザーの操作に対して、ユースケースの実装を行う
具体的な操作を表現する
### domain
各ドメインに関連するロジックの実装を行う
#### entity
エンティティオブジェクトの定義
各ドメインの型となるオブジェクトインスタンスをここで作成する
#### value_obj
値オブジェクトの定義（不変の値であること）
必要に応じてディレクトリを用意する。
権限の確認やバリデーションなどのビジネスルールをvalue objectで定義する
#### repository
DB操作を行うInterfaceを定義し、実装層への抽象化を実現する
#### service
各ドメインのロジックの補助となる。
ビジネスルールの実装を行う

---
## 疑問点
## application/usecaseとdomain/serviceの違い
### usecase層
システムの操作単位を定義する
### service
業務ルールを定義する
## *参考文献*
[https://qiita.com/mizuko_dev/items/a8a3864e23d82ba2a60d]
[https://architecting.hateblo.jp/entry/2025/05/21/000044]
ドメイン/アプリケーション層[https://qiita.com/kotobuki5991/items/22712c7d761c659a784f]
サービス[https://note.com/happy_avocet7237/n/nfbb69b5345b5]