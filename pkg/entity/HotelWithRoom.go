package entity

import (
	"fmt"
	"math"
)

type HotelWithRoom struct {
	Hotel
	Rooms []Room
}

// CalcMinCost 使用多重背包计算最小的花费
func (hwr *HotelWithRoom) CalcMinCost(adult int, child int) float64 {
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
	for i := 0; i < 35; i++ {
		for j := 0; j < 35; j++ {
			dp[i][j] = 192608170.0
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
				dp[i][j] = math.Min(dp[i-x][j-y]+p, dp[i][j])
			}
		}
	}

	//第三步：枚举小孩变大人的情况，统计答案
	ans := 192608170.0
	for i := 0; i < 35; i++ {
		for j := 0; j < 35; j++ {
			if i >= adult && i+j >= adult+child {
				ans = math.Min(ans, dp[i][j])
			}
		}
	}
	return ans
}
