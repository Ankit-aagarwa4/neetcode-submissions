func carFleet(target int, position []int, speed []int) int {
	m := make(map[int]int)
    // arr := make([]float64, len(position))
    j := 0
	for i := 0; i < len(position); i++ {
		m[position[i]] = speed[i]
	}

	sort.Ints(position)
	time1 := float64(target - position[len(position) - 1]) / float64(m[position[len(position) - 1]])
	for i:= len(position) - 1; i > 0; i-- {
        time2 := float64(target - position[i - 1]) / float64(m[position[i - 1]])
        if time1 < time2 {
            time1 = time2
            j++;
        } 
	}
	return j + 1
}
