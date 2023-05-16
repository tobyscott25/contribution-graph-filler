// should make multiple commits some days

package commits

import "math/rand"

func NumberOfCommits(isWeekday bool) int {

	commits := 0

	chance := 10
	if isWeekday {
		chance = 70
	}

	for i := 0; i < 3; i++ {
		randomNumber := rand.Intn(100)
		if randomNumber < chance {
			commits++
		}
	}

	return commits
}
