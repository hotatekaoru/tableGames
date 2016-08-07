package tic

var cell [2]int

// 
func CallAI(b Board, turn string) [2]int {

	nextTurn := NextTurnMap[turn]

	// 次のターンで勝利できるか
	judge := false
	if judge = b.reach(turn); judge { return cell }

	// 次のターンで敗北を阻止できるか
	if judge = b.reach(nextTurn); judge { return cell }

	return b.searchEmpty(turn)
}

func (b *Board)reach (color string) bool  {

	for _, w := range winBoardIndexArray {
		slice := []string{b[w[0].row][w[0].col],b[w[1].row][w[1].col],b[w[2].row][w[2].col]}

		count1, count2 := 0, 0
		isEnemy := true
		empCell := [2]int{0,0}
		for i, s := range slice {
			switch s {
			case color:
				count1 += 1
			case "":
				count2 += 1
				empCell[0] = w[i].row
				empCell[1] = w[i].col
			default:
				isEnemy = false
			}
		}

		if count1 == 2 && isEnemy {
			cell = empCell
			return true
		}
	}

	return false
}

func (b *Board)searchEmpty (color string) [2]int {
	for row, rows := range b {
		for col, _ := range rows {
			if b[row][col] == "" {
				b[row][col] = color
				return [2]int{row, col}
			}
		}
	}
	return [2]int{0,0}
}

