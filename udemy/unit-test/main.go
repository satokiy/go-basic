package add

// import "fmt"
func Add(x, y int) int {
	return x + y
}

func Divide(x, y int) float64 {
	if y == 0 {
		return 0.
	}
	return float64(x) / float64(y)
}

func main() {
	// x,y := 3,5
	// fmt.Printf("add: %v\n", Add(x,y))
	// fmt.Printf("divide: %v\n", Divide(x,y))
}

/*
  単体テストの実行
  $ go test -v ./udemy/unit-test/
  $ go test -v -cover -coverprofile=./udemy/unit-test/coverage.out ./udemy/unit-test/
  $ go tool cover -html=./udemy/unit-test/coverage.out
*/
