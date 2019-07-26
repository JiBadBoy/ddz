package card

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func init() {
	// 以时间作为初始化种子
	rand.Seed(time.Now().UnixNano())
}

type TCards []*TCard

// NewCards 新建一副牌
func NewCards(nLen int) TCards {
	cards := make(TCards, nLen)
	// 给牌堆里面的牌赋初始值
	for i := 0; i < nLen; i++ {
		pCard := NewCard(i + 1) //这里从1开始
		cards[i] = pCard
	}

	cards.randomize()

	return cards
}

// randomize 洗牌
func (t TCards) randomize() {
	for i := 0; i < len(t); i++ {
		nRand := rand.Intn(len(t) - 1) // 从 0 - 53 开始
		// 两两指针进行交换
		t.Swap(i, nRand)
	}
}

// 长度
func (t TCards) Len() int {
	return len(t)
}

// 比较大小
func (t TCards) Less(i, j int) bool {
	// 先比较点数
	if t[i].nPoint == t[j].nPoint {
		// 点数相同比较花色
		return t[i].nFlower > t[j].nFlower
	}
	return t[i].nPoint < t[j].nPoint
}

// 交换两个牌的牌值
func (t TCards) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t TCards) String() string {
	// 顺带打印看结果
	strResult := ""
	for i := range t {
		strResult = fmt.Sprintf("%s%s ", strResult, t[i].String())
	}
	return strResult
}

// 排序
func (t TCards) Sort() {
	sort.Sort(t)
}
