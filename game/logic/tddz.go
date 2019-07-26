package logic

import (
	"doudizhu/game/card"
	"fmt"
)

// 斗地主执行
type TDdz struct {
	pHandCardList1 card.TCards // 第1家的牌
	pHandCardList2 card.TCards // 第2家的牌
	pHandCardList3 card.TCards // 第3家的牌
	pUnderCardList card.TCards // 3张底牌
}

// NewDDZ 新建一个斗地主类
func NewDDZ() *TDdz {
	p := &TDdz{}
	return p
}

// dealCard 发牌阶段
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

	fmt.Println(t.pHandCardList1)
	fmt.Println(t.pHandCardList2)
	fmt.Println(t.pHandCardList3)
	fmt.Println(t.pUnderCardList)
}
