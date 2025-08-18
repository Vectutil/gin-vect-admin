package system

import (
	"gin-vect-admin/internal/app/logic/system"
	"gin-vect-admin/internal/app/response"
	"gin-vect-admin/internal/app/types/common"
	systype "gin-vect-admin/internal/app/types/system"
	"gin-vect-admin/pkg/mysql"
	"strconv"

	"github.com/gin-gonic/gin"
)

// RoleHandler 角色处理器
type RoleHandler struct {
}

// NewRoleHandler 创建角色Handler实例
func NewRoleHandler() *RoleHandler {
	return &RoleHandler{}
}

// Create 创建角色
// @title 创建角色
// @Summary 创建新角色
// @Description 创建一个新的角色
// @Tags 角色管理
// @Accept json
// @Produce json
// @Param request body systype.RoleCreateReq true "角色创建请求参数"
// @Success 200 {object} systype.RoleCreateResp "成功返回"
// @Failure 500 {object} response.Response "内部错误"
// @Router /role [post]
func (h *RoleHandler) Create(c *gin.Context) {
	var (
		err       error
		db        = mysql.GetDB()
		req       systype.RoleCreateReq
		res       = &systype.RoleCreateResp{}
		roleLogic = system.NewRoleLogic(db)
	)

	defer func() {
		response.HandleDefault(c, res)(&err)
	}()

	if err = c.ShouldBindJSON(&req); err != nil {
		return
	}

	if err = roleLogic.Create(c.Request.Context(), &req); err != nil {
		return
	}
}

// Update 更新角色
// @title 更新角色
// @Summary 更新角色信息
// @Description 根据角色Id更新角色信息
// @Tags 角色管理
// @Accept json
// @Produce json
// @Param id path int true "角色Id"
// @Param request body systype.RoleUpdateReq true "角色更新请求参数"
// @Success 200 {object} systype.RoleUpdateResp "成功返回"
// @Failure 500 {object} response.Response "内部错误"
// @Router /role [put]
func (h *RoleHandler) Update(c *gin.Context) {
	var (
		err       error
		db        = mysql.GetDB()
		req       systype.RoleUpdateReq
		res       = &systype.RoleUpdateResp{}
		roleLogic = system.NewRoleLogic(db)
	)

	defer func() {
		response.HandleDefault(c, res)(&err)
	}()

	if err = c.ShouldBindJSON(&req); err != nil {
		return
	}

	if err = roleLogic.Update(c.Request.Context(), &req); err != nil {
		return
	}
}

// Delete 删除角色
// @title 删除角色
// @Summary 删除指定角色
// @Description 根据角色Id删除角色
// @Tags 角色管理
// @Accept json
// @Produce json
// @Param id path int true "角色Id"
// @Param request body common.IdReq true "角色更新请求参数"
// @Success 200 {object} systype.RoleDeleteResp "成功返回"
// @Failure 500 {object} response.Response "内部错误"
// @Router /role [delete]
func (h *RoleHandler) Delete(c *gin.Context) {
	var (
		err       error
		db        = mysql.GetDB()
		req       = &common.IdReq{}
		res       = &systype.RoleDeleteResp{}
		roleLogic = system.NewRoleLogic(db)
	)

	defer func() {
		response.HandleDefault(c, res)(&err)
	}()

	if err = c.ShouldBindJSON(&req); err != nil {
		return
	}

	if err = roleLogic.Delete(c.Request.Context(), req.Id); err != nil {
		return
	}
}

// GetById 根据Id获取角色
// @title 获取角色详情
// @Summary 获取指定角色详情
// @Description 根据角色Id获取角色详细信息
// @Tags 角色管理
// @Accept json
// @Produce json
// @Param id path int true "角色Id"
// @Success 200 {object} systype.RoleDataResp "成功返回"
// @Failure 500 {object} response.Response "内部错误"
// @Router /role/{id} [get]
func (h *RoleHandler) GetById(c *gin.Context) {
	var (
		err       error
		db        = mysql.GetDB()
		res       = &systype.RoleDataResp{}
		roleLogic = system.NewRoleLogic(db)
	)

	defer func() {
		response.HandleDefault(c, res)(&err)
	}()

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return
	}
	res, err = roleLogic.GetById(c.Request.Context(), id)
	if err != nil {
		return
	}
}

// List 查询角色列表
// @title 获取角色列表
// @Summary 获取角色列表
// @Description 分页获取角色列表信息
// @Tags 角色管理
// @Accept json
// @Produce json
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(10)
// @Success 200 {object} systype.RoleDataListResp "成功返回"
// @Failure 500 {object} response.Response "内部错误"
// @Router /role [get]
func (h *RoleHandler) List(c *gin.Context) {
	var (
		err       error
		db        = mysql.GetDB()
		req       systype.RoleQueryReq
		res       = &systype.RoleDataListResp{}
		roleLogic = system.NewRoleLogic(db)
	)

	defer func() {
		response.HandleListDefault(c, res)(&err)
	}()

	err = response.ShouldBindForList(c, &req)
	if err != nil {
		return
	}

	res, err = roleLogic.GetList(c.Request.Context(), &req)
	if err != nil {
		return
	}
}
