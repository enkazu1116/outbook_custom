# クリーンアーキテクチャ × DDD × モジュラーモノリス
## フォルダ構成
```tree
.
├── cmd
│   └── server
│       └── main.go
├── Docs
│   └── directry.md
├── go.mod
├── go.sum
├── infrastructure
│   ├── config
│   ├── db
│   ├── di
│   │   └── wire.go
│   ├── logger
│   ├── repository
│   │   └── user_repository.go
│   └── security
│       └── bcrypt.go
├── internal
│   ├── application
│   │   ├── dto
│   │   │   └── user
│   │   │       └── dto.go
│   │   ├── interface
│   │   ├── port
│   │   │   └── PasswordHasher.go
│   │   └── usecase
│   │       └── user
│   │           └── create_user_usecase.go
│   └── domain
│       ├── output
│       │   └── entity
│       │       ├── repository
│       │       └── services
│       └── user
│           ├── entity
│           │   └── user.go
│           ├── repository
│           │   └── user_repository.go
│           ├── services
│           │   └── user_service.go
│           └── value_obj
│               └── role.go
└── server
```
---
## 詳細
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
リポジトリの実装層
internal/domain/repositoryで定義したinterfaceを用いて実装を行う。
#### security
パスワードのハッシュ化の実装層（Adapter）
internal/application/portで定義したinterfaceを用いて実装を行う。

---
## internal
外部に公開しない業務ロジックやサービスをここに記載する
### application
ユーザーの操作に対して、ユースケースの実装を行う
具体的な操作を表現する
#### dto
ユーザーの入力を受け付けるオブジェクト
ドメインとは切り離し、ユースケースの実装を行うapplication内に配置する
#### port
アプリが外部とやり取りするための接続口を定義
パスワードハッシュ化アダプターと現在はやり取りをしている。
#### usecase
ユースケース層
アプリケーションのロジックをここで定義する
今回モジュラーモノリス構成も採用しているため、ドメインごとに分けて管理する

---

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

### port層について
外部の技術実装とやりとりをする役割
Q. 外部はinfrastructureはもそれにあたるのか？
A. 「外部」とはビジネスロジック層（internal）からみた外の世界を指す。
	つまり、infrastructureはinternalから見れば外の世界とみなす。

---
## *参考文献*
[https://qiita.com/mizuko_dev/items/a8a3864e23d82ba2a60d]
[https://architecting.hateblo.jp/entry/2025/05/21/000044]
ドメイン/アプリケーション層[https://qiita.com/kotobuki5991/items/22712c7d761c659a784f]
サービス[https://note.com/happy_avocet7237/n/nfbb69b5345b5]
パスワードハッシュ化[https://qiita.com/osamu_0/items/e747921cce4ca9e48720]
DTO1[https://zenn.dev/motojouya/articles/extend_data_transfer_object]
DTO2[https://note.com/happy_avocet7237/n/nf1b3dd7dbac9]
PORT1[https://note.com/happy_avocet7237/n/n27413b6d18af]
PORT2[https://qiita.com/Sicut_study/items/d59a362176a8412458fb]