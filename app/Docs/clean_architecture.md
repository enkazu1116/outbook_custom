# 原則
#CleanArchitecture 

## SOLID原則
Clean ArchitectureはSOLID原則に基づいています。
SOLID原則の簡単な概要は以下となります。
この原則を守るだけでも、コード品質は向上しそうに思います。

- S: 単一責任の原則
  クラスは単一の責任を持つこと
- O: オープン・クローズドの原則
  拡張に対してはオープンで、変更にはクローズドであるべき
  構造の変更を受け入れることには、厳しく。構造に新たに追加することには、柔軟であることが望ましいとされる。
- L: リスコムの置換原則
  子クラスのオブジェクトを親クラスに置き換えても、プログラムは同じ結果を得られる
  親クラスの概念から外れた事象は許されない
- I: インターフェース分離の原則
  メソッドへの依存を強制しない
  クラスが最低限必要なもののみを実行できる状態にしておく
- D: 依存性逆転の原則
  上位モジュールは、下位モジュールに依存してはならない。
  どちらも抽象化に依存するべきである。
  抽象化は詳細に依存してはならない。
  詳細が抽象化に依存するべきである。
参考[https://qiita.com/baby-degu/items/d058a62f145235a0f007]

## コンポーネント原則
モジュール化・コンポーネント化する際の指針となる。
凝集性は、どのように関心事をまとめるべきか？
結合性は、どのように依存関係を構築するか？

下記の本では、「凝集度・結合度」という言葉で、
この原則を守るための考え方が記載されている。
良いコード/悪いコードで学ぶ設計入門[https://gihyo.jp/book/2025/978-4-297-14622-1]
### 凝集性
- REP（再利用/リリース等価の原則）
  再利用するなら、それぞれが独立してリリース・バージョン管理ができる
- CCP（共通閉包の原則）
  一緒に変更されるクラスは同じコンポーネントにまとめる
- CRP（共通再利用の原則） 
  一緒に使われるクラスは同じコンポーネントにまとめる
  使わないものへの依存は避ける。

### 凝集度
- 高い = 責務・関心事が1箇所に集約されている
- 低い = 責務・関心事が複数箇所に点在している / 1箇所に複数の責務・関心事が混在している
**理想** : 凝集度が高い状態。1箇所に1つの責務・関心事が集約されて、
	 コードでの実現がされていると、把握がしやすい。
	 低い状態を想像すると、イメージがしやすい。複数箇所を見て、どう実現させているのかを把握するのは中々骨を折られる。

### 結合性
- ADP（非循環依存の原則）
  コンポーネント依存に循環を作らない。
- SDP（安定依存の原則）
  変化しやすいものは、変化しにくいものに依存すべき。
- SAP（安定抽象の原則）
  安定したコンポーネントは抽象的であるべき
  
###　結合度
- 密結合 = 関係の度合いが高い
- 疎結合 = 関係の度合いが低い
理想: 疎結合であること

コンポーネント原則参考[https://qiita.com/nozomi2025/items/e71bff92e17ca6e421df]
結合度・凝集度[https://qiita.com/dsudo/items/ee3fee1f558c7f1b359f]
凝集度[https://zenn.dev/miya_tech/articles/0dde1228045af6]
結合度[https://zenn.dev/taiga533/articles/e08ad4f4af5577079b5b]

## 思想
### 基本思想
目的: 関心の分離
- 依存性ルール
  基本的に外側から内側に依存関係は向かう。
- フレームワーク独立性 / UI独立性  / DB独立性
  必要に応じて容易に変更・交換ができる
- テスト容易性
  ビジネスロジック単体でのテストが容易に可能

参考[https://logmi.jp/main/technology/323451]

# 4つの階層
## 重要原則
1. 内側の層は、外側の層を知らない
2. 依存関係は内側に向かって流れる
3. インターフェースによる契約
   契約とは: 実装側と呼び出し側でI/O, Errorのルールを決めること

原則の参考[https://zenn.dev/collabostyle/articles/1089b482fd59fe]
依存性逆転の大切さについて[https://speakerdeck.com/shimabox/kurinakitekutiyakarajian-ruyi-cun-noxiang-kinoda-qie-sa?slide=61]
契約の参考[https://zenn.dev/convers39/articles/83dd5898d4d798]

## 各階層(内側から)
1. Enterprise Business Rules
   Entitiesが唯一定義されている概念。
   ビジネスルールをカプセル化したもの。中身は様々である。
2. Application Business Rules
   1をクライアントとしてエンティティを操作する。
   アプリケーション固有のビジネスルールが含まれる。
   全てのユースケースがカプセル化・実装されている。
   **🌟 ユースケースとは**
   エンティティに入出力するデータの流れを調整し、ユースケースの目標を達成できるように
   エンティティに最重要ビジネスネールを使用するように指示を出す
3. Interface Adapter
   2と4との間でそれぞれのレイヤーで利用できるような型への相互変換を行う。
   **⚠️ 注意**
   Interface Adapterという名前の中にInterfaceが含まれていて、プログラミングのInterfaceと混同してしまうが、
   クリーンアーキテクチャでのInterfaceの定義は異なるよう。
   2にとって便利な形式から4にとって便利な形式に変換する役割を持つ。
4. Frameworks & Drivers
   フレームワークやDBなどの技術について置くためのレイヤー
階層参考[https://gist.github.com/mpppk/609d592f25cab9312654b39f1b357c60]

## Enterprise Business Rules
DDD上のEntity・Value ObjectとClean Architecture上のEntityは概念が少し異なる。
システムを設計・実装する際の私なりの考え方としては
1.  Entityはある関心事や事象をカプセル化したものとして分ける
2. Value ObjectはEntityに関連する周辺の不変的な値
3. Clean Architecture上のEntityは1, 2を合わせたドメインモデル
といった具合で、住み分けを行っている。
Entity関連について参考[https://qiita.com/takasek/items/70ab5a61756ee620aee6]

## Application Business Rules
システムが何をするかを表現する。
業務のロジックやプロセスの大枠がここで定義・実現がされる。

| 用語                   | 意味                         | 使用用途                             |
| -------------------- | -------------------------- | -------------------------------- |
| Use Case             | アプリの主要機能                   | ユーザーの操作をまとめる                     |
| Interactore          | Use Caseの実装クラス             | Use Case内で実装することが多い              |
| Use Case Input Port  | Use Caseを呼び出すためのインターフェース   | Input Boundaryの実体                |
| Input Data           | ユースケースが受け取るデータ             | DTOとしてユーザー入力などアプリへ入ってくるデータの変換受け口 |
| Input Boundary       | Use Caseの境界（Input Portのこと） | 外側の層からUsecaseを呼び出す窓口の概念          |
| Use Case Output Port | ユースケースの結果を外側に返すためのインターフェース | Output Boundaryの実体               |
| Output Data          | 出力データ（Presenterへ渡す内容）      | DTOとしてアプリからUIに出力するデータの変換受け口      |
| Output Boundary      | Presenterインターフェース（出力境界）    | 外側の層へUsecaseの処理結果を返す出口の概念        |
| Repository           | 永続化の抽象インターフェース             | DB操作の実装クラス                       |

####  実際のフォルダ構成
UseCase → internal/domain/usecase
UseCase Input Port, Input Boundary → internal/application/port
UseCase Output Port, Output Boundary → internal/application/port
Repository → internal/domain/repository
Input Data, Output Data → internal/application /dto
自分の場合は、以上のようにフォルダ分けを行った

#### この階層の役割
3層目から渡された情報を変換して、ビジネスロジック・ルールに従って処理を実現し、1層目に渡してやること。
3層目で定義した契約に従って、処理を実現する。
1層目で定義したデータ構造に従って、データを出力する。

参考1[https://zenn.dev/sui_water/articles/88af41dc6d64bc]
テーブル内容参考[https://nrslib.com/clean-architecture-with-java/#outline__6_3]
全体的な参考[https://gist.github.com/mpppk/609d592f25cab9312654b39f1b357c60]

## Interface Adapters
2層と4層それぞれで利用できる型への相互変換を行う。
つまり、外側の層・内側の層それぞれをつなぐ変換層。

| 用語         | 意味                                          | 使用用途                     | データ方向 |
| ---------- | ------------------------------------------- | ------------------------ | ----- |
| Gateway    | 4層からのデータを抽象化する                              | Repositoryや外部APIクライアントなど | 外→内   |
| Presenter  | 2層からOutputを受け取り、UIに適した形にして返す                | DTO変換、Webレスポンンスの生成など     | 内→外   |
| Controller | Webサーバ等からデータを受け取り、Inputに適した形に変換してUsecaseへ渡す | HTTP Handlerなど           | 内←→外  |

#### 実際のフォルダ構成
Gateway → infrastructure/repository・security...など
Presenter → internal/application/interface/handler/presenter
Controller →  internal/application/interface/handler/controller
とすると、全てを実現できる

**⚠️注意**
Presenterは省略も可能。
シンプルな実装内容や小規模な場合は、かえってオーバーとなることがあるため、Controllerに返り値を与えるパターンを使用し、簡潔に実現できるメリットを享受できる
参考[https://izumisy.work/entry/2019/12/12/000521]

## Frameworks & Drivers
フレームワークやDBなど、詳細な技術を置くためのレイヤー
アプリケーションが実際に動くための技術的な土台を担い、実際の処理を実行する責務を持つ。

#### 実際のフォルダ構成
infrastructure/di・db・config・logger...など