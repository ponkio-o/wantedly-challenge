package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
)

type Data struct {
	Message  string `json:"message"`
}

func main(){

	// ルーティング
	router := mux.NewRouter()

	// URL設定
	router.HandleFunc("/", func(w http.ResponseWriter, r  *http.Request){

		// 構造体定義
		var data = Data{}
		data.Message = "Hello World!!"

		// jsonエンコード
		outputJson, err := json.Marshal(&data)
		if err != nil {
			panic(err)
		}

		// jsonヘッダーを出力
		w.Header().Set("Content-Type", "application/json")

		// jsonデータを出力
		fmt.Fprint(w, string(outputJson))

	})

	// ハンドル割当
	http.Handle("/", router)

	// ポート
	http.ListenAndServe(":9000", nil)

}
