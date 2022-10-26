package main

import "fmt"

type Event struct {
	notifier Notifier //被知晓的通知者
	One      Listener //打人的人
	Another  Listener //被打的
	Msg      string
}

// 广播

type Listener interface {
	OnFriendBeFight(event *Event)
	Title() string
	GetName() string
	GetParty() string
}

// 通知

type Notifier interface {
	AddListener(listener Hero)
	RemoveListener(listener Hero)
	Notifier(event *Event)
}

// 观察者

type Hero struct {
	Name  string
	Party string
}

// 朋友被攻击

func (hero *Hero) OnFriendBeFight(event *Event) {
	if hero.Name == event.One.GetName() || hero.Name == event.Another.GetName() {
		return
	}
	if hero.Party == event.One.GetParty() {
		fmt.Println(hero.Name + "拍手叫好")
	}
	if hero.Party == event.Another.GetParty() {
		fmt.Println(hero.Name + "回击" + event.One.GetName())
		//hero.Fight(event.One, event.notifier) 不能再嵌套 避免无线循环
		return
	}
}
func (hero *Hero) GetName() string {
	return hero.Name
}
func (hero *Hero) GetParty() string {
	return hero.Party
}
func (hero *Hero) Title() string {
	return fmt.Sprintf("[%s]%s", hero.Party, hero.Name)
}

// 主动出击

func (hero *Hero) Fight(another Listener, baixiaosheng Notifier) {
	fmt.Println(hero.Name + "打了" + another.GetName())
	event := new(Event)
	event.One = hero
	event.Another = another
	event.notifier = baixiaosheng
	event.Msg = fmt.Sprintf("%s 打了 %s", hero.Name, hero.Title())
	baixiaosheng.Notifier(event)
}

// 通知人

type BaiXiaoSheng struct {
	heroList []Hero
}

func (s *BaiXiaoSheng) AddListener(hero Hero) {
	s.heroList = append(s.heroList, hero)
}

func (s *BaiXiaoSheng) RemoveListener(hero Hero) {

	for i, l := range s.heroList {
		if l == hero {
			//拼接slice
			s.heroList = append(s.heroList[:i], s.heroList[i+1:]...)
		}
	}
}

func (s *BaiXiaoSheng) Notifier(event *Event) {
	for _, hero := range s.heroList {
		hero.OnFriendBeFight(event)
	}
}

func main() {
	hero1 :=
		Hero{
			Name:  "1",
			Party: "1",
		}
	hero2 :=
		Hero{
			Name:  "2",
			Party: "1",
		}
	hero3 :=
		Hero{
			Name:  "3",
			Party: "1",
		}
	hero4 :=
		Hero{
			Name:  "4",
			Party: "2",
		}
	hero5 :=
		Hero{
			Name:  "5",
			Party: "2",
		}
	hero6 :=
		Hero{
			Name:  "6",
			Party: "2",
		}
	bai := BaiXiaoSheng{}
	bai.AddListener(hero1)
	bai.AddListener(hero2)
	bai.AddListener(hero3)
	bai.AddListener(hero4)
	bai.AddListener(hero5)
	bai.AddListener(hero6)
	hero1.Fight(&hero4, &bai)
}
