# sampleGoWeb
## メモ
- 新規レコードを作成する時はCreate
- 既存レコードを更新する時はできるだけUpdate
- 更新時に空値を含めてStructで更新したい場合はSave
- 更新時に空値を含めてMapで更新したい時はUpdate

  Saveは強力である一方で、「Updateのつもりで使ったら、該当IDが存在しない場合は意図せず新規レコードが作成されてしまう」など、不測の結果を招きやすいのが玉にキズです。まずはCreateやUpdateで代用できないかを考えましょう。  
  [qiita](https://qiita.com/ttiger55/items/3606b8dd570637c12387)
---
