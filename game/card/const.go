package card

//花色
const (
	flowerNil      int = iota //留空
	flowerHeiTao              // 黑桃（小王）
	flowerHongTao             // 红桃（大王）
	flowerMeiHua              // 梅花
	flowerFangKuai            // 方块
)

//点数
const (
	cardPointNil int = iota // 留空
	cardPoint3              // 1
	cardPoint4              // 2
	cardPoint5              // 3
	cardPoint6              // 4
	cardPoint7              // 5
	cardPoint8              // 6
	cardPoint9              // 7
	cardPoint10             // 8
	cardPointJ              // 9
	cardPointQ              // 10
	cardPointK              // 11
	cardPointA              // 12
	cardPoint2              // 13
	cardPointX              // 14 (小王)
	cardPointY              // 15 （大王）
)
