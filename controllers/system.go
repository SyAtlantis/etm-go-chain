package controllers

import "github.com/astaxie/beego"

// Operations about system
type SystemController struct {
	beego.Controller
}


// @Title getHeight
// @Description The latest block height
// @Success 200 {int} height
// @Failure 403 height not exist
// @router /getHeight [get]
func (s *SystemController) GetHeight() {
	s.Data["json"] = "height"
	s.ServeJSON()
}

// @Title getMilestone
// @Description The chain milestone
// @Success 200 {int} milestone
// @Failure 403 milestone not exist
// @router /getMilestone [get]
func (s *SystemController) GetMilestone() {
	s.Data["json"] = "milestone"
	s.ServeJSON()
}

// @Title getReward
// @Description The chain reward
// @Success 200 {int} reward
// @Failure 403 reward not exist
// @router /getReward [get]
func (s *SystemController) GetReward() {
	s.Data["json"] = "Reward"
	s.ServeJSON()
}
// @Title getStatus
// @Description The chain status
// @Success 200 {object} status
// @Failure 403 status not exist
// @router /getStatus [get]
func (s *SystemController) GetStatus() {
	s.Data["json"] = "status"
	s.ServeJSON()
}