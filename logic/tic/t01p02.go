package tic

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

// 盤面などで使用する文字列
const maru, batsu = "○", "×"

// Board型の宣言
type Board [3][3]string

// boardFormValue関数の宣言（盤面の値を取得）
func BoardFormValue(c *gin.Context) *Board {
	var board Board
	for row, rows := range board {
		for col, _ := range rows {
			// 盤面のname属性「c00」～「c22」を作成
			name := "c" + strconv.Itoa(row) + strconv.Itoa(col)
			// 盤面の各項目を取得
			print(c.PostForm(strconv.Itoa(row) + strconv.Itoa(col) + "-> " + name))
			board[row][col] = c.PostForm(name)
		}
	}
	return &board
}

// 変数nextTurnMapに、次の手番を取得するマップを設定
var NextTurnMap = map[string]string{
	maru:  batsu,
	batsu: maru,
	"":    maru, // 「""」の場合、ゲーム開始時として「"○"」を取得
}

func FirstTurn() string{
	return NextTurnMap[""]
}

// turnFormValue関数の宣言（手番の値を取得）
func TurnFormValue(c *gin.Context) (string, string) {
	turn := c.PostForm("turn")   // 現在の手番を取得
	nextTurn := NextTurnMap[turn] // マップを使用して次の手番を取得
	return turn, nextTurn
}

// 変数winBoardIndexArrayに、勝敗判定用の構造体の2次元配列を設定 … ③
var winBoardIndexArray = [...][3]struct{ row, col int }{
	// 横（行）の判定
	{{0, 0}, {0, 1}, {0, 2}},
	{{1, 0}, {1, 1}, {1, 2}},
	{{2, 0}, {2, 1}, {2, 2}},
	// 縦（列）の判定
	{{0, 0}, {1, 0}, {2, 0}},
	{{0, 1}, {1, 1}, {2, 1}},
	{{0, 2}, {1, 2}, {2, 2}},
	// 斜めの判定
	{{0, 0}, {1, 1}, {2, 2}},
	{{0, 2}, {1, 1}, {2, 0}},
}

// winメソッドの宣言（勝敗の判定）… ④
func (b *Board) Win(turn string) bool {
	for _, w := range winBoardIndexArray {
		// 3個すべてがそろった場合、勝利と判定
		if (b[w[0].row][w[0].col] == turn) &&
		(b[w[1].row][w[1].col] == turn) &&
		(b[w[2].row][w[2].col] == turn) {
			return true
		}
	}
	return false
}

// drawメソッドの宣言（引き分けの判定） … ⑤
func (b *Board) Draw() bool {
	for row, rows := range b {
		for col, _ := range rows {
			if b[row][col] == "" {
				return false // 未入力の項目がある場合、ゲームを続行
			}
		}
	}
	return true // 未入力の項目がない場合、引き分け
}

// setBarメソッドの宣言（「"-"」の設定） … ⑥
func (b *Board) SetBar() {
	for row, rows := range b {
		for col, _ := range rows {
			if b[row][col] == "" {
				b[row][col] = "-" // 未入力の項目は「"-"」を設定
			}
		}
	}
}
