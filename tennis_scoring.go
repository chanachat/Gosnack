package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type game struct {
	scoreA int
	scoreB int
	winA   bool
	winB   bool
}

var point = []int{
	0: 0,
	1: 15,
	2: 30,
	3: 40,
}

var wordPoint = map[int]string{
	0:  "love",
	15: "Fifteen",
	30: "Thirty",
	40: "Forty",
}

func (g *game) playerAGetPoint() {
	g.scoreA++
}

func (g *game) playerBGetPoint() {
	g.scoreB++
}

func getPoint(score int) int {
	if score <= 3 {
		return point[score]
	}
	return 40
}

func (g *game) currentScore() {
	// Get player A score
	pa := getPoint(g.scoreA)
	wa := wordPoint[pa]
	// Get player B score
	pb := getPoint(g.scoreB)
	wb := wordPoint[pb]

	const defultSentence = "(Score A: ?1 - B: ?2 Input A: ?3 - B: ?4 => Output :"

	// Validate score
	if g.scoreA > 3 && g.scoreA > g.scoreB && (g.scoreA-g.scoreB) == 1 {
		fmt.Println(generateSentence(defultSentence, g, pa, pb), "A -", wa)
	} else if g.scoreB > 3 && g.scoreB > g.scoreA && (g.scoreB-g.scoreA) == 1 {
		fmt.Println(generateSentence(defultSentence, g, pa, pb), wb, "- A")
	} else if g.scoreA > 3 && g.scoreA > g.scoreB && (g.scoreA-g.scoreB) >= 2 {
		fmt.Println(generateSentence(defultSentence, g, pa, pb), "A Wins")
		g.winA = true
	} else if g.scoreB > 3 && g.scoreB > g.scoreA && (g.scoreB-g.scoreA) >= 2 {
		fmt.Println(generateSentence(defultSentence, g, pa, pb), "B Wins")
		g.winB = true
	} else if g.scoreB >= 3 && g.scoreB == g.scoreA {
		fmt.Println(generateSentence(defultSentence, g, pa, pb), "Deuce")
	} else {
		fmt.Println(generateSentence(defultSentence, g, pa, pb), wa, "-", wb)
	}
}

func generateSentence(sen string, g *game, pa int, pb int) string {
	return (strings.NewReplacer("?1", strconv.Itoa(g.scoreA),
		"?2", strconv.Itoa(g.scoreB),
		"?3", strconv.Itoa(pa),
		"?4", strconv.Itoa(pb))).Replace(sen)
}

func isFinishMatch(setA int, setB int) bool {

	const wordAwinner = "Finished tennis results is \"A\" are winner and \"B\" are Losser"
	const wordBwinner = "Finished tennis results is \"B\" are winner and \"A\" are Losser"

	if setA >= 6 && setA > setB && (setA-setB) >= 2 {
		fmt.Println(wordAwinner)
		return true
	} else if setB >= 6 && setB > setA && (setB-setA) >= 2 {
		fmt.Println(wordBwinner)
		return true
	} else if setA == 7 && setB == 5 {
		fmt.Println(wordAwinner)
		return true
	} else if setB == 7 && setA == 5 {
		fmt.Println(wordBwinner)
		return true
	} else if setA == 7 && setB == 6 {
		fmt.Println(wordAwinner)
		return true
	} else if setB == 6 && setA == 7 {
		fmt.Println(wordBwinner)
		return true
	}
	return false
}

func isFinishSet(winA bool, winB bool, setA *int, setB *int) bool {
	if winA || winB {
		if winA {
			*setA++
		} else {
			*setB++
		}
		return true
	}
	return false
}

func main() {
	sn := 1
	var setA int
	var setB int
	for {
		if isFinishMatch(setA, setB) {
			break
		} else {
			fmt.Println("Set : ", sn, "A Vs B")
			g := game{}
			for {
				if isFinishSet(g.winA, g.winB, &setA, &setB) {
					break
				} else {
					rand.Seed(time.Now().UTC().UnixNano())
					if (rand.Int() % 2) == 0 {
						g.playerAGetPoint()
					} else {
						g.playerBGetPoint()
					}
					g.currentScore()
				}
			}
			sn++
		}
	}
}
