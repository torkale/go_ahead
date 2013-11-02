package main
import "fmt"

func slices() {
  primes := []int{2, 3, 5, 7, 11, 13}

  fmt.Println("primes[1:4] ==", primes[1:4])

  zeroes := make([]int, 5)
  fmt.Println("zeroes ==", zeroes)

  for i, v := range primes {
     fmt.Printf("(%d) = %d\n", i, v)
  }
  for _, v := range primes {
     fmt.Printf("%d\n", v)
  }
}
