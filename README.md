# TODO API一覧
- 一覧取得 /todo
- 登録 /todo
- 更新 /todo/:id
- 完了 /todo/:id/done
- 削除 /todo/:id
- 完了一覧 /todo/done

**todosテーブル**
todo_id int
content varchar(1000)
is_done boolean
created_at datetime
updated_at datetime
deleted_flg boolean

## Golang, Echo疑問
- 例外時NewHttpErrorで返すべき？
→ NewHttpErrorで返すとHttpErrorHandlerにエラーが渡される。HttpErrorHandlerに独自のエラーハンドリングを設定しておくとエラー処理を統一することができるので便利
- ポインタの概念がよくわからない
- 例外処理修正

## 課題
- 更新もバリデーションする
- エラーハンドリング
- logging実装
