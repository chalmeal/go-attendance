// entorypoint v1
package main

import (
	"github.com/chalmeal/delight/domain"
)

func main() {
	// 学生情報を読み込む
	students, err := domain.LoadStudentInfo()
	if err != nil {
		return
	}

	// ランダムに席順を決定する
	seat := domain.CreateAutoSeat(students)
	if seat.IsEmpty() {
		return
	}

	// 席順をjson形式で保存する
	err = domain.SaveSeatInfo(seat)
	if err != nil {
		return
	}
}
