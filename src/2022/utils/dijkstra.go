package utils

import "math"

var intfinity = int(math.Float64bits(math.Inf(0)))

func hasUnvisited[T comparable](visited map[T]bool, distances map[T]int) bool {
	for p, v := range visited {
		if !v && distances[p] == intfinity {
			return true
		}
	}
	return false
}

// GetMinimumDistance finds T with minimum distance that is not yet visited
func getMinimumUnvisited[T comparable](visited map[T]bool, distances map[T]int) T {
	min := intfinity
	var out T
	for k := range distances {
		if yes := visited[k]; !yes && distances[k] < min {
			min = distances[k]
			out = k
		}
	}
	return out
}

func FindShortestPath[T comparable](
	start, end T,
	all []T,
	getAdjacent func(curr T) []T,
	getWeight func(curr, adjacent T) int,
) []T {
	visited := make(map[T]bool, len(all))
	distances := make(map[T]int, len(all))
	prev := make(map[T]T)

	// Init visited to false and distances to infinity
	for _, v := range all {
		visited[v] = false
		distances[v] = intfinity
	}

	distances[start] = 0

	for hasUnvisited(visited, distances) {
		curr := getMinimumUnvisited(visited, distances)

		if curr == end {
			break
		}

		visited[curr] = true

		for _, n := range getAdjacent(curr) {
			if yes := visited[n]; yes {
				continue
			}

			dist := distances[curr] + getWeight(curr, n)
			if dist < distances[n] {
				distances[n] = dist
				prev[n] = curr
			}
		}
	}

	path := make([]T, 0)
	curr := end
	for {
		path = append(path, curr)

		p, found := prev[curr]
		if !found {
			break
		}
		curr = p
	}

	// Reverse
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path
}
