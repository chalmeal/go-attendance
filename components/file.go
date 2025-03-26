package components

import (
	"fmt"
	"os"
)

// Jsonファイルの読み込み
func LoadJsonFile(fileName string, loadFunc func(*os.File) error) error {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("jsonファイルのオープンに失敗しました。")
		fmt.Println("Error:", err)
		file.Close()
		return err
	}
	defer file.Close()

	return func() error {
		err := loadFunc(file)
		if err != nil || os.IsNotExist(err) {
			fmt.Println("jsonファイルの読み込みに失敗しました。")
			fmt.Println("Error:", err)
			return nil
		}
		return err
	}()
}
