package mian

import (
	"etm-go-chain/utils"
	"fmt"
	"time"
)

func main() {
	slots :=utils.NewSlots()
	
	t0 := slots.BeginEpochTime()
	fmt.Println("BeginEpochTime",t0,t0.Unix()*1000)
	
	t1 := slots.GetTime()
	fmt.Println("GetTime",t1)
	
	t2 := slots.GetRealTime()
	fmt.Println("GetRealTime",t2,time.Now().Unix()*1000)
	
	t3 := slots.GetSlotNumber()
	fmt.Println("GetSlotNumber",t3)
	
	t6 := slots.GetSlotTime(t3)
	fmt.Println("GetSlotTime",t6)
	
	t4 := slots.GetNextSlot()
	fmt.Println("GetNextSlot",t4)
	
	t5 := slots.GetHeightPerDay()
	fmt.Println("GetHeightPerDay",t5)
}