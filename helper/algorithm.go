package helper

import "math/rand"

func NumberOfCommits(isWeekday bool) int {

	commits := 0

	chance := 8
	if isWeekday {
		chance = 50
	}

	for i := 0; i < 3; i++ {
		randomNumber := rand.Intn(100)
		if randomNumber < chance {
			commits++
		}
	}

	return commits
}
