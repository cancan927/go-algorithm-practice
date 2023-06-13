package Table

type Table interface {
	Put(k, v int)
	Get(k int) int
	Del(k int) bool
	Exist(k int) bool
	First(k int) int
	Last(k int) int
	Floor(k int)
	Ceiling(k int)
}
