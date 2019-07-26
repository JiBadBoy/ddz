package card

// TCard 扑克牌结构体
type TCard struct {
	nValue  int // 牌值
	nFlower int // 花色
	nPoint  int // 点数
}

// NewCard 新建卡牌
func NewCard(nValue int) *TCard {
	p := &TCard{}
	p.nValue = nValue
	p.nFlower = toFlower(nValue)
	p.nPoint = toCardValue(nValue)
	return p
}

// 从牌值获取花色
func toFlower(nValue int) int {
	if nValue <= 0 || nValue > 54 {
		return flowerNil
	}
	if nValue == 53 {
		return flowerHeiTao
	}
	if nValue == 54 {
		return flowerHongTao
	}
	return ((nValue - 1) / 13) + 1
}

// 从牌值获取点数
func toCardValue(nValue int) int {
	if nValue <= 0 || nValue > 54 {
		return cardPointNil
	}
	if nValue == 53 {
		return cardPointX // 小王
	}
	if nValue == 54 {
		return cardPointY // 大王
	}
	return ((nValue - 1) % 13) + 1
}

// 转化成字符串
func (t *TCard) String() string {
	strResult := ""
	// 花色
	switch t.nFlower {
	case flowerHeiTao:
		{
			strResult = "♠"
		}
	case flowerHongTao:
		{
			strResult = "♥"
		}
	case flowerMeiHua:
		{
			strResult = "♣"
		}
	case flowerFangKuai:
		{
			strResult = "♦"
		}
	}

	// 点数
	switch t.nPoint {
	case cardPoint3:
		{
			strResult = strResult + "3"
		}
	case cardPoint4:
		{
			strResult = strResult + "4"
		}
	case cardPoint5:
		{
			strResult = strResult + "5"
		}
	case cardPoint6:
		{
			strResult = strResult + "6"
		}
	case cardPoint7:
		{
			strResult = strResult + "7"
		}
	case cardPoint8:
		{
			strResult = strResult + "8"
		}
	case cardPoint9:
		{
			strResult = strResult + "9"
		}
	case cardPoint10:
		{
			strResult = strResult + "10"
		}
	case cardPointJ:
		{
			strResult = strResult + "J"
		}
	case cardPointQ:
		{
			strResult = strResult + "Q"
		}
	case cardPointK:
		{
			strResult = strResult + "K"
		}
	case cardPointA:
		{
			strResult = strResult + "A"
		}
	case cardPoint2:
		{
			strResult = strResult + "2"
		}
	case cardPointX:
		{
			strResult = strResult + "小王"
		}
	case cardPointY:
		{
			strResult = strResult + "大王"
		}
	}
	return strResult
}
