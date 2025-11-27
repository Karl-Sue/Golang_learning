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
    start := 7*(week-1)
    end := start + 7
    result := 0
	for i:=start;i<end;i++{
        result += birdsPerDay[i]
    }
    return result
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
    birds := birdsPerDay
	for i:=0; i<len(birdsPerDay); i+=2 {
        birds[i]++
    }
    return birds
}
