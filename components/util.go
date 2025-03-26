package components

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/chalmeal/delight/config"
)

// jsonファイルパス変換
func ConvertJsonFilePath(fileName string) string {
	if fileName == "" {
		fmt.Println("ファイル名が空です。")
		return ""
	}
	filepath := config.JSON_ROOT_PATH + fileName + config.JSON_EXTENSION
	return filepath
}

// map[string]map[string]string{}型をjson形式に変換する
func ConvertToJson(data map[string]map[int]string) ([]byte, error) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("json形式への変換に失敗しました。")
		fmt.Println("Error:", err)
		return nil, err
	}
	return jsonData, nil
}

// json形式のデータをファイルに保存する
func SaveJsonFile(filePath string, jsonData []byte) error {
	// 保存先ディレクトリが存在しない場合は作成する
	dirpath := config.JSON_ROOT_PATH + config.OUTPUT_JSON_PATH
	if _, err := os.Stat(dirpath); os.IsNotExist(err) {
		err := os.MkdirAll(dirpath, os.ModePerm)
		if err != nil {
			fmt.Println("jsonファイルの保存先ディレクトリの作成に失敗しました。")
			fmt.Println("Error:", err)
			return err
		}
	}

	// ファイルを作成する
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("jsonファイルの作成に失敗しました。")
		fmt.Println("Error:", err)
		return err
	}
	defer file.Close()

	// ファイルにjsonデータを書き込む
	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Println("jsonファイルの書き込みに失敗しました。")
		fmt.Println("Error:", err)
		return err
	}

	fmt.Println("jsonファイルを保存しました。")
	fmt.Printf("保存先: %s\n", filePath)

	return nil
}
