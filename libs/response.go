package libs

type Response struct {
	Code int64       `json:"code"`
	Msg  interface{} `json:"message"`
	Data interface{} `json:"data"`
}

type Lists struct {
	Data  interface{} `json:"data"`
	Total int         `json:"total"`
}

var CodeSuc int64 = 1
var CodeErr int64 = 0    // エラー
var CodeErr2 int64 = 2   // 検索結果なし
var CodeErr3 int64 = 3   // JSON読み込み失敗
var CodeErr4 int64 = 4   // トークンなし
var CodeErr5 int64 = 5   // ユーザーIDが見つからない
var CodeErr6 int64 = 6   // 同期操作失敗
var CodeErr7 int64 = 7   // お気に入り同期操作失敗
var CodeErr8 int64 = 8   // レポート同期操作失敗
var CodeErr9 int64 = 9   // docKeyが取得できない
var CodeErr10 int64 = 10 // lungCodeが取得できない
var CodeErr11 int64 = 11 // ドキュメントの取得失敗
var CodeErr12 int64 = 12 // ドキュメントの保存失敗
var CodeErr13 int64 = 13 // リクエストエラー
var CodeErr14 int64 = 14 // ParentOrgnaizationId未入力
var CodeErr15 int64 = 15 // kgpの所属一覧:操作失敗
var CodeErr16 int64 = 16 // 保存失敗
var CodeErr17 int64 = 17 // ユーザー登録失敗
var CodeErr18 int64 = 18 // ユーザー更新失敗
var CodeErr19 int64 = 19 // ユーザー削除失敗
var CodeErr20 int64 = 20 // 32MB超過
var CodeErr21 int64 = 21 // CSV開けない
var CodeErr22 int64 = 22 // CSV読み込み失敗
var CodeErr23 int64 = 23 // CSV作成失敗
var CodeErr24 int64 = 24 // CSVファイル圧縮失敗
var CodeErr25 int64 = 25 // zipファイル名取得失敗
var CodeErr26 int64 = 26 // zipファイル取得失敗
var CodeErr27 int64 = 27 // 検索履歴削除失敗
var CodeErr28 int64 = 28 // お気に入り削除失敗
var CodeErr29 int64 = 29 // 検索履歴同期失敗
var CodeErr30 int64 = 30 // 検索履歴追加失敗
var CodeErr31 int64 = 31 // flash point保存失敗
var CodeErr32 int64 = 32 // flash point取得失敗
var CodeErr33 int64 = 33 // お気に入りコンテンツ追加失敗
var CodeErr34 int64 = 34 // 取得に失敗
var CodeErr35 int64 = 35 // ログ登録失敗
var CodeErr36 int64 = 36 // content登録失敗
var CodeErr37 int64 = 37 // jsonファイル取得失敗
var CodeErr38 int64 = 38 // jsonのデコード失敗
var MegSuc string = "操作が成功しました。"
var MegErr string = "操作が失敗しました。"
var MegReadJSONErr string = "データの読み込みに失敗しました。"

func ApiResource(code int64, objects interface{}, msg string) (r *Response) {
	r = &Response{Code: code, Data: objects, Msg: msg}
	return
}
