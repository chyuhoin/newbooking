package entity

import (
	"fmt"
)

type HotelWithRoom struct {
	Hotel
	Rooms []Room
}

// CalcMinCost 使用多重背包计算最小的花费
func (hwr *HotelWithRoom) CalcMinCost(adult int, child int) (float64, *[]Room) {
	rooms := make([]Room, 0)

	//第一步：二进制拆分
	for _, room := range hwr.Rooms {
		tmp := room
		for now := 1 << 10; now > 0; now >>= 1 {
			if room.Remain > now {
				tmp.Remain = now
				rooms = append(rooms, tmp)
				room.Remain -= now
			}
		}
		rooms = append(rooms, room)
	}

	//第二步：倒序dp
	var dp [35][35]float64
	var method [35][35][]Room
	for i := 0; i < 35; i++ {
		for j := 0; j < 35; j++ {
			dp[i][j] = 192608170.0
			method[i][j] = make([]Room, 0)
		}
	}
	dp[0][0] = 0
	for _, room := range rooms {
		var x, y int
		fmt.Sscanf(room.Capacity, "%d,%d", &x, &y)
		x *= room.Remain
		y *= room.Remain
		p := room.Price * float64(room.Remain)
		for i := 30; i >= x; i-- {
			for j := 30; j >= y; j-- {
				newCost := dp[i-x][j-y] + p
				if newCost < dp[i][j] {
					dp[i][j] = newCost
					method[i][j] = append(method[i-x][j-y], room)
				}
			}
		}
	}

	//第三步：枚举小孩变大人的情况，统计答案
	ans := 192608170.0
	var res *[]Room
	for i := 0; i < 35; i++ {
		for j := 0; j < 35; j++ {
			if i >= adult && i+j >= adult+child && ans > dp[i][j] {
				ans = dp[i][j]
				res = &method[i][j]
			}
		}
	}

	//由于需要得出具体方案，所以还要把二进制拆分的加回去（合并同类项）
	//这一步也可以用排序来合并，但是一般来讲res里的元素个数都不会超过3个，所以直接写俩for就行
	tres := make([]Room, 0)
	for _, room := range *res {
		unique := true
		for j, pre := range tres {
			if room.Id == pre.Id {
				tres[j].Remain = room.Remain + pre.Remain //用remain表示租多少个
				unique = false
				break
			}
		}
		if unique {
			tres = append(tres, room)
		}
	}
	return ans, &tres
}
