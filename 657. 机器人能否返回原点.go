package leetcode

func judgeCircle(moves string) bool {
	x := 0
	y := 0
	b:=[]byte(moves)
	for i:=0;i<len(b);i++{
		if b[i]=='U'{
			y++
		}else if b[i]=='D'{
			y--
		}else if b[i]=='L'{
			x--
		}else {
			x++
		}
	}


	return x == 0 && y == 0
}
