package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tableGames/logic/tic"
	"os/exec"
)

func T01G01(c *gin.Context) {

	// 外部ファイル起動の練習
	out, err := exec.Command("./normal", "").CombinedOutput()
	if err != nil {
		println(err)
	}
	println(string(out))

	c.HTML(http.StatusOK, "t01.html", gin.H{})

}

func T02P01(c *gin.Context) {

	tic.GameStart(c)

	// 手番を初期化する
	nextTurn := tic.FirstTurn()
	// 盤面の入力値を取得する
	board := tic.BoardFormValue(c)

	c.HTML(http.StatusOK, "t02.html", gin.H{
		"turn":   nextTurn,
		"board":  board,
		"win":    false,
		"draw":   false,
		"winner": "",
	})

}

func T02P02(c *gin.Context) {

	// 手番の入力値を取得する
	turn, nextTurn := tic.TurnFormValue(c)
	// 盤面の入力値を取得する
	board := tic.BoardFormValue(c)

	// 勝敗、引き分け、勝者の変数宣言
	win, draw, winner := tic.JudgeWinner(board, turn)

	c.HTML(http.StatusOK, "t02.html", gin.H{
		"turn":   nextTurn,
		"board":  board,
		"win":    win,
		"draw":   draw,
		"winner": winner,
	})

}
