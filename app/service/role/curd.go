package role

import (
	"Gf-Vben/app/model/entity"
	"Gf-Vben/app/service/internal/dao"
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

type Req struct {
	Page     int `p:"page"`
	PageSize int `p:"page_size"`
	Ctx      context.Context
	Query
}

type Query struct {
	Id         int    `p:"id"`
	Uid        int    `p:"uid"`
	Name       string `p:"name"`
	Value      string `p:"value"`
	Desc       string `p:"desc"`
	Permission string `p:"permission"`
}

type Role struct {
	Name string `json:"name"`
}

func (r *Req) List() (g.Map, error) {
	var res []entity.Role
	if err := dao.Role.Ctx(r.Ctx).Scan(&res); err != nil {
		return nil, err
	}
	return g.Map{
		"items":    res,
		"total":    len(res),
		"page":     r.Page,
		"pageSize": r.PageSize,
	}, nil
}
func (r *Req) Add() error {
	panic("implement me")
}

func (r *Req) Edit() error {
	panic("implement me")
}

func (r *Req) Del() error {
	panic("implement me")
}

func (r *Req) Tree() (g.Map, error) {
	panic("implement me")
}

func (r *Req) Options() (g.Map, error) {
	panic("implement me")
}
