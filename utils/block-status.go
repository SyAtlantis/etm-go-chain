package utils

var (
	totalAmount      int64 = 10000000000000000 //创世快发行量
	distance         int64 = 10112000          // 每个里程碑的长度
	rewardOffset     int   = 1                 // 开始奖励的高度
	lastRewardHeight int64 = 59328000          // 最后一个有奖励的高度
	
	milestones []int64 = []int64{
		600000000, // Initial Reward
		500000000, // Milestone 1
		400000000, // Milestone 2
		400000000, // Milestone 3
		300000000, // Milestone 4
		200000000, // Milestone 5
	}
)

type BlockStatus struct {
}

// 计算某个高度的里程碑
func (bs *BlockStatus) CalcMilestone(height int64) int {
	location := int((height - int64(rewardOffset)) / distance)
	
	if location > len(milestones)-1 {
		return len(milestones) - 1
	} else {
		return location
	}
}

// 计算某个高度的奖励
func (bs *BlockStatus) CalcReward(height int64) int64 {
	if height < int64(rewardOffset) || height <= 1 || height > lastRewardHeight {
		return 0
	} else {
		return milestones[bs.CalcMilestone(height)]
	}
}

//计算某个高度的发行量
func (bs *BlockStatus) CalcSupply(height int64) int64 {
	slots := NewSlots()
	if height > lastRewardHeight { //超出最后奖励高度不计算
		height = lastRewardHeight
	}
	
	m_height := height - height%int64(slots.Delegates)
	milestone := bs.CalcMilestone(m_height)
	supply := totalAmount
	rewards := map[int][]int64{}
	
	if m_height <= 0 {
		return supply
	}
	
	var _blocks, _rewards int64
	m_height = m_height - int64(rewardOffset) + 1
	for i := 0; i < len(milestones); i++ {
		if milestone >= i {
			_rewards = milestones[i]
			
			if m_height <= 0 {
				break
			} else if m_height < distance {
				_blocks = m_height % distance
			} else {
				_blocks = distance
			}
			rewards[i] = []int64{_blocks, _rewards}
			
			m_height -= distance
			
		} else {
			break
		}
	}
	
	for _, reward := range rewards {
		supply += reward[0] * reward[1]
	}
	
	return supply
}
