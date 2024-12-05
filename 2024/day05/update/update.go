package update

import "slices"

type Update struct {
	pages []int
	rules map[int][]int
}

func (a Update) Len() int      { return len(a.pages) }
func (a Update) Swap(i, j int) { a.pages[i], a.pages[j] = a.pages[j], a.pages[i] }
func (a Update) Less(i, j int) bool {
	mustComeAfter := a.rules[a.pages[i]]
	return slices.Contains(mustComeAfter, a.pages[j])
}

func New(rules map[int][]int, pages []int) Update {
	return Update{
		pages: pages,
		rules: rules,
	}
}

func (u Update) Mid() int {
	i := (len(u.pages) - 1) / 2
	return u.pages[i]
}
