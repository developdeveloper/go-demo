package move

// Movable 能移动
type Movable interface {
	Walk(distance int) int // 走了多久
	Swim(distance int) int // 游了多久
}
