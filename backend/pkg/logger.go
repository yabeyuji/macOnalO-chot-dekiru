package pkg

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// MyErr ...
type MyErr struct {
	Layer string
	Part  string
	Value string
	Err   error
}

// NewMyErr ...
func NewMyErr(layer string, part string) *MyErr {
	me := &MyErr{}
	me.Layer = layer
	me.Part = part

	return me
}

// Logging ...
func (me MyErr) Logging(passErr error, value ...interface{}) {
	// テスト中であればロギングしない
	if flag.Lookup("test.v") != nil {
		log.Println("run under go test")
		return
	}

	fileName := time.Now().Format("2006-01-02") + ".log"
	LogName := filepath.Join(ErrorLogPath, fileName)

	// ファイルが存在しなければ作成
	_, err := os.Stat(LogName)
	if err != nil {
		_, err := os.Create(LogName)
		if err != nil {
			log.Fatal(err)
		}
	}

	// エラー以外の情報をjson化
	jsonData, err := json.Marshal(&value)
	if err != nil {
		log.Fatal(err)
	}

	// 実行ファイルと行数取得
	_, fullPath, line, _ := runtime.Caller(1)
	file := strings.TrimLeft(fullPath, currentPath)
	if err != nil {
		log.Fatal(err)
	}

	// エラー情報作成
	var rowData []string
	rowData = append(rowData, time.Now().Format("2006-01-02_15:04:05")) // dateTime
	rowData = append(rowData, me.Layer)                                 // レイヤ名
	rowData = append(rowData, me.Part)                                  // アプリ名
	rowData = append(rowData, fmt.Sprint(file, ":", line))              // ファイル名:行番号
	rowData = append(rowData, passErr.Error())                          // エラー内容
	rowData = append(rowData, string(jsonData))                         // エラー原因
	row := strings.Join(rowData, "\t")                                  // タブ区切り

	// ファイル書き込み
	f, err := os.OpenFile(filepath.Clean(LogName), os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	fmt.Fprintln(f, row)

	return
}
