package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
    sum := 0
	for _, val := range birdsPerDay {
        sum += val
    }

    return sum
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	sum := 0
    for _, val := range birdsPerDay[(week - 1) * 7 : (week * 7)] {
        sum += val
    }

    return sum
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i, val := range birdsPerDay {
        if i == 0 || i % 2 == 0 {
            birdsPerDay[i] = val + 1
        }
    }

    return birdsPerDay
}
