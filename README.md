












## modelsフォルダについて

### modelとdb操作用のテーブルを分けること

#### modelについて
- name_model.go // 単数系命名すること
- logを扱う場合はlog_name_model.goとファイル名を命名すること
- gettingとsettingを定義するfile
- modelを定義、キャピタルから始まるモデルを定義しないこと

#### entitiesについて 
- entitiesフォルダの中のファイルdb処理用のTBLを定義すること
- file名に関してはtbl_name.goとlog_name.goと切り分けて命名すること
- table定義に関しては先頭にTBLとLogをつけること // Logに関しては最初だけ大文字の方が見やすいから

#### バリデーション関数
汚いもっとうまく期待
#### テストコード
汚いもっと綺麗に書い直す。



