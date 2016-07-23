package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tic-tac-toe/logic/tic"
)

func T01G01(c *gin.Context) {

	// 手番の入力値を取得する
	nextTurn := tic.FirstTurn()
	// 盤面の入力値を取得する
	board := tic.BoardFormValue(c)

	c.HTML(http.StatusOK, "t01.html", gin.H{
		"turn"	: nextTurn,
		"board"		: board,
		"win"		: false,
		"draw"		: false,
		"winner"	: "",
	})

}

func T01P02(c *gin.Context) {

	println("aaaaaaaa")
	// 手番の入力値を取得する
	turn, nextTurn := tic.TurnFormValue(c)
	// 盤面の入力値を取得する
	board := tic.BoardFormValue(c)

	// 勝敗、引き分け、勝者の変数宣言
	win, draw, winner := false, false, ""

	// turnが「""」の場合、ゲーム開始時とする
	if turn != "" { // ゲーム開始時以外
		win = board.Win(turn) // 勝敗の判定

		// 勝敗が付いた場合
		if win {
			// 勝者を設定
			winner = turn
			// 未入力の項目に「"-"」を設定
			board.SetBar()
			// 勝敗が付かない場合
		} else {
			// 引き分けの判定
			draw = board.Draw()
		}
	}

	c.HTML(http.StatusOK, "t01.html", gin.H{
		"turn"	: nextTurn,
		"board"		: board,
		"win"		: win,
		"draw"		: draw,
		"winner"	: winner,
	})

}

