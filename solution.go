package square

import "math"

type Num int

const (
	SidesTriangle Num = 3
	SidesSquare   Num = 4
	SidesCircle   Num = 0
	Pi                = 3.14
)

func CalcSquare(sideLen float64, sidesNum Num) float64 {
	if sidesNum == SidesTriangle {
		squareTriangle := (math.Sqrt(3) * (math.Pow(sideLen, 2))) / 4
		return squareTriangle
	}
	if sidesNum == SidesSquare {
		squareSquare := math.Pow(sideLen, 2)
		return squareSquare
	}
	if sidesNum == SidesCircle {
		squareCircle := Pi * math.Pow(sideLen, 2)
		return squareCircle
	}
	return 0
}
