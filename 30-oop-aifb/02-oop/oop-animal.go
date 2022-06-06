package main

type Animal interface {
  eat(food string)  // 吃点东西
  sleep(hours uint) // 睡一会儿
  playAsPet()       // 激发宠物的属性
}

// 猫
type Cat struct {}

func (cat *Cat) eat(food string) {}
func (cat *Cat) sleep(hours uint) {}
func (cat *Cat) playAsPet() {}
// 猫捉老鼠
func (cat *Cat) catchMouse() {}

// 狗
type Dog struct {}

func (dog *Dog) eat(food string) {}
func (dog *Dog) sleep(hours uint) {}
func (dog *Dog) playAsPet() {}
// 狗来看家
func (dog *Dog) defendHouse() {}

// 能帮警察搞事情的动物
type PoliceAnimal interface {
  Animal         // 作为动物来讲，肯定有 Animal 的行为，直接拿过来用
  arrestBadMan() // 抓捕犯罪嫌疑人
}

// 牧羊犬
type HuntawayDog struct {
  Dog // 首先是一直狗
}

func (huntawayDog *HuntawayDog) arrestBadMan() {}

// 昆明犬
type KunmingDog struct {
  Dog // 首先是一直狗
}

func (kunmingDog *KunmingDog) arrestBadMan() {}

// 参数面向接口对象
func atPeace(animal Animal) {
  animal.eat("meat")
  animal.sleep(10)
  animal.playAsPet()
}

// 参数面向接口对象
func doTask(animal PoliceAnimal) {
  animal.arrestBadMan()
}

func main() {
  huntawayDog := &HuntawayDog{}
  atPeace(huntawayDog) // 没有问题，HuntawayDog 是 PoliceAnimal，而 PoliceAnimal 具备 Animal的 一切行为
  doTask(huntawayDog)

  normalDog := &Dog{}
  atPeace(normalDog)   // 没问题，Dog 是 Animal 
  // doTask(normalDog) // 有问题，普通的狗没有抓捕能力 
}