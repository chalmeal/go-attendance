package domain

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/chalmeal/delight/components"
	"github.com/chalmeal/delight/config"
)

const (
	// 対象のjsonファイル名
	JSON_STUDENT_INFO = "student_info"
	STUDENT_SEAT      = "student_seat"
)

type Student struct {
	// 出席番号
	AttendanceNumber int `json:"attendance_number"`
	// 名前
	Name string `json:"name"`
	// 性別
	Gender string `json:"gender"`
}

// student_info.jsonを読み込む
func LoadStudentInfo() ([]Student, error) {
	var students []Student
	// JSONファイルのパスを取得
	filePath := components.ConvertJsonFilePath(JSON_STUDENT_INFO)
	// JSONファイルを読み込む
	err := components.LoadJsonFile(filePath,
		func(file *os.File) error {
			decoder := json.NewDecoder(file)
			e := decoder.Decode(&students)
			if e != nil {
				fmt.Println("JSONファイルの読み込みに失敗しました。")
				fmt.Println("Error:", e)
				return e
			}
			return nil
		})
	if err != nil {
		return nil, err
	}

	return students, nil
}

// 席順をランダムに自動確定
func CreateAutoSeat(students []Student) Seat {
	// それぞれのスライスを作成
	var maleStudents []Student
	var femaleStudents []Student

	// 男性と女性のスライスを分ける
	for _, student := range students {
		// MALEならmaleStudentsに追加
		if student.Gender == config.MALE {
			maleStudents = append(maleStudents, student)
		} else if student.Gender == config.FEMALE {
			// FEMALEならfemaleStudentsに追加
			femaleStudents = append(femaleStudents, student)
		} else {
			// 性別が入力されていない場合はエラーを返す
			fmt.Println("性別が入力されていない学生がいます。")
			fmt.Println("出席番号:", student.AttendanceNumber, "名前:", student.Name)
			return Seat{}
		}
	}

	// 男性の配列をランダムに並び替える
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(maleStudents), func(i, j int) {
		maleStudents[i], maleStudents[j] = maleStudents[j], maleStudents[i]
	})

	// 女性の配列をランダムに並び替える
	rand.Shuffle(len(femaleStudents), func(i, j int) {
		femaleStudents[i], femaleStudents[j] = femaleStudents[j], femaleStudents[i]
	})

	var male MaleSeat
	var female FemaleSeat
	var seat Seat

	// 男性の席を作成
	for idx, student := range maleStudents {
		male.MaleSeatNumber = idx + 1
		male.MaleName = student.Name
		seat.MaleSeat = append(seat.MaleSeat, male)
	}

	// 女性の席を作成
	for idx, student := range femaleStudents {
		female.FemaleSeatNumber = idx + 1
		female.FemaleName = student.Name
		seat.FemaleSeat = append(seat.FemaleSeat, female)
	}

	return seat
}

// 席順をjson形式で保存する
func SaveSeatInfo(seat Seat) error {
	// maleSeatをmap[]string]string{}型に変換
	maleSeatMap := make(map[int]string)
	for _, student := range seat.MaleSeat {
		maleSeatMap[student.MaleSeatNumber] = student.MaleName
	}

	// femaleSeatMapをmap[]string]string{}型に変換
	femaleSeatMap := make(map[int]string)
	for _, student := range seat.FemaleSeat {
		femaleSeatMap[student.FemaleSeatNumber] = student.FemaleName
	}

	// maleSeatとfemaleSeatをjson形式に変換
	seatMap := make(map[string]map[int]string)
	seatMap[config.MALE] = maleSeatMap
	seatMap[config.FEMALE] = femaleSeatMap

	// seatMapをjson形式に変換
	jsonData, err := components.ConvertToJson(seatMap)
	if err != nil {
		return err
	}

	// json形式のデータをファイルに保存する
	currentTime := time.Now().Format(config.JSON_DATE_FORMAT)
	filepath := components.ConvertJsonFilePath(config.OUTPUT_JSON_PATH + STUDENT_SEAT + currentTime)
	err = components.SaveJsonFile(filepath, jsonData)
	if err != nil {
		return err
	}

	return nil
}
