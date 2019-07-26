package logic

import (
	"doudizhu/game/card"
	"fmt"
	"log"
	"time"
)

// 斗地主执行
type TDdz struct {
	pHandCardList1 card.TCards // 第1家的牌
	pHandCardList2 card.TCards // 第2家的牌
	pHandCardList3 card.TCards // 第3家的牌
	pUnderCardList card.TCards // 3张底牌
	nDZPosition    int         // 抢地主号位
	chDefineDZ     chan int    // 抢地主的通道
	nMultiples     int         // 斗地主的倍数
}

// NewDDZ 新建一个斗地主类
func NewDDZ() *TDdz {
	p := &TDdz{}
	return p
}

// DealCard 发牌阶段
func (t *TDdz) DealCard() {
	// 需要洗一副牌， 并且发送给三个玩家
	pCardHead := card.NewCards(54)
	fmt.Println(pCardHead)

	// 需要给三家每一家发17张牌
	t.pHandCardList1 = make(card.TCards, 17)
	t.pHandCardList2 = make(card.TCards, 17)
	t.pHandCardList3 = make(card.TCards, 17)
	t.pUnderCardList = make(card.TCards, 3)

	copy(t.pHandCardList1, pCardHead[0:17])
	copy(t.pHandCardList2, pCardHead[17:34])
	copy(t.pHandCardList3, pCardHead[34:51])
	copy(t.pUnderCardList, pCardHead[51:54])

	// 发牌完毕给大家手上排个序好看些
	t.pHandCardList1.Sort()
	t.pHandCardList2.Sort()
	t.pHandCardList3.Sort()
	t.pUnderCardList.Sort()

	t.Print()
}

// 抢地主
func (t *TDdz) GrabLandlord() {
	ch := make(chan int, 1)
	t.chDefineDZ = ch

	go func() {
		log.Println("模拟客户端叫地主")
		t.nDZPosition = 1
		t.nMultiples *= 2 // 倍数要乘以2
		ch <- 1
	}()

	select {
	case <-ch:
		log.Println("叫地主")
	case <-time.After(time.Second * 3):
		log.Println("超时")
	}

	// 给地主分配额外的3张手牌
	switch t.nDZPosition {
	case 0:
		log.Println("无人抢地主，游戏应该结束，未来在做其他的异常处理")
	case 1:
		t.pHandCardList1 = append(t.pHandCardList1, t.pUnderCardList...)
		t.pHandCardList1.Sort()
	case 2:
		t.pHandCardList2 = append(t.pHandCardList2, t.pUnderCardList...)
		t.pHandCardList2.Sort()
	case 3:
		t.pHandCardList3 = append(t.pHandCardList3, t.pUnderCardList...)
		t.pHandCardList3.Sort()
	default:
		log.Panicln("未知的位置")
	}

	t.Print()
}

// 打印
func (t *TDdz) Print() {
	fmt.Println(t.pHandCardList1)
	fmt.Println(t.pHandCardList2)
	fmt.Println(t.pHandCardList3)
	fmt.Println(t.pUnderCardList)
}
