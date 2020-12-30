package tag

// type record struct {
// 	Name  string
// 	Age   int
// 	Score float32
// }

// type record struct {
// 	Name  string  `json:"nickname"`
// 	Age   int     `json:"age"`
// 	Score float32 `json:"score"`
// }

type record struct {
	Name  string  `json:"nickname" xml:"nickname,attr"`
	Age   int     `json:"age" xml:"age,attr"`
	Score float32 `json:"score" xml:"score,attr"`
}
