package chatstore

var UserMsg map[int32][]string
var AdminMsg map[int32][]string

func init() {
	UserMsg = make(map[int32][]string)
	AdminMsg = make(map[int32][]string)
}