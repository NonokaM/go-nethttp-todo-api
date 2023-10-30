# go-nethttp-todo-api
goの標準パッケージnet/httpのみを使用したTODOアプリのAPI

| エンドポイント | リクエストメソッド | 動作 |
|--------------|----------------|-----------------------|
| `/todos`     | GET            | 全てのTODO項目を取得 |
| `/todos/{id}`| GET            | 指定されたIDのTODO項目を取得 |
| `/todos`     | POST           | 新しいTODO項目を作成 |
| `/todos/{id}`| PUT            | 指定されたIDのTODO項目を更新 |
| `/todos/{id}`| DELETE         | 指定されたIDのTODO項目を削除 |
