package main

import "os"

func main() {
	// フォルダを 1 階層だけ作成
	os.Mkdir("setting", 0755)

	// 深いフォルダを 1 回で作成
	os.MkdirAll("setting/myapp/networksettings", 0755)
}
