package tic

var cell [2]int

func CallAI(b Board, turn string) {

	println("CPU起動")
	nextTurn := NextTurnMap[turn]

	// 次のターンで勝利できるか
	if judge := b.reach(turn); judge == true { b.setBoard(turn); return }

	// 次のターンで敗北を阻止できるか
	if judge := b.reach(nextTurn); judge == true { b.setBoard(turn); return }

	b.searchEmpty(turn)
	return
}

func (b *Board)reach (color string) bool  {

	for _, w := range winBoardIndexArray {
		slice := []string{b[w[0].row][w[0].col],b[w[1].row][w[1].col],b[w[2].row][w[2].col]}

		count1, count2 := 0, 0
		for i, s := range slice {
			switch s {
			case color:
				count1 += 1
			case "":
				count2 += 1
				cell[0] = w[i].row
				cell[1] = w[i].col
			default:
				return false
			}
		}

		if count1 == 2 {return true}
	}

	return false
}

func (b *Board)setBoard (color string) { b[cell[0]][cell[1]] = color }

func (b *Board)searchEmpty (color string) {
	for row, rows := range b {
		for col, _ := range rows {
			if b[row][col] == "" {
				b[row][col] = color
			}
		}
	}
}

