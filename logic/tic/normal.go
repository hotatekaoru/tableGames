package tic

func CallAI(b Board, turn string) string{

	println("CPU起動")

	// 次のターンで勝利できるか
	b.logic1(turn)

	// 次のターンで敗北を阻止できるか
	return ""
}

func (b *Board)logic1 (turn string) {

	for _, w := range winBoardIndexArray {
		// 3個すべてがそろった場合、勝利と判定
		if (b[w[0].row][w[0].col] == turn) &&
			(b[w[1].row][w[1].col] == turn) &&
			(b[w[2].row][w[2].col] == turn) {
		}
	}

}