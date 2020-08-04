package main

import (
	crand "crypto/rand"
	"errors"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"time"

	_ "github.com/1k-ct/sampleGoWeb/db"
	_ "github.com/1k-ct/sampleGoWeb/server"
	_ "github.com/gin-gonic/gin"
)

/*
func main() {
	r := gin.Default()
	var s user.Service
	todos, _ := s.GetAll()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world go gin")
		c.String(http.StatusOK, "", todos)
	})
	db.Init()
	server.Init()
	defer db.Close()
}
*/
func myShuffle(nums ...interface{}) interface{} {
	rand.Shuffle(len(nums), func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})
	return nums
}

//
func gacha() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100)
}

// pureGacha 純粋らしい、おそらく
func pureGacha(n int) int {
	seed, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	rand.Seed(seed.Int64())
	return rand.Intn(n)
}
func GachaProbabilityMachine() (int, error) {
	n := gacha()
	if n < 10 {
		return n, nil
	} else if n < 20 {
		return n, nil
	} else if n < 100 {
		return n, nil
	}
	return 0, errors.New("Unexpected numbers")
}
func main() {
	nums := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	rand.Shuffle(len(nums), func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})
	fmt.Println(nums)
	fmt.Println(myShuffle(nums))
	randN := gacha()
	fmt.Println(randN)
	randN = gacha()
	fmt.Println(randN)
	monster := []Monster{
		{
			rarity: 5, // ★★★★★
			prob:   0.01,
			ids: []int{
				5001,
				5002,
				5003,
			},
		},
		{
			rarity: 6, // ★★★★★★
			prob:   0.1,
			ids: []int{
				6001,
				6002,
				6003,
			},
		},
	}
	fmt.Println(monster[0].ids)
	ans, _ := GachaProbabilityMachine()
	fmt.Println(ans)
}

type id struct{ id []int }
type Monster struct {
	rarity int     // ★★★★★
	prob   float64 // 確立
	ids    []int
}
