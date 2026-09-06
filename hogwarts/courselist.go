//go:build !solution

package hogwarts

func dfs(head string, prereqs map[string][]string, visited map[string]int) []string {
	result := []string{}
	visited[head] = 1
	for _, node := range prereqs[head] {
		if visited[node] == 0 {
			result = append(result, dfs(node, prereqs, visited)...)
		} else if visited[node] == 1 {
			panic(1)
		}
	}
	visited[head] = 2
	result = append(result, head)
	return result
}

func GetCourseList(prereqs map[string][]string) []string {
	courseList := []string{}
	visited := make(map[string]int)
	for node := range prereqs {
		_, ok := visited[node]
		if !ok {
			courseList = append(courseList, dfs(node, prereqs, visited)...)
		}
	}
	return courseList
}
