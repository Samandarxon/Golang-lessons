package modules

import "sort"

type Money struct {
	Id        int    `json:"id"`
	Code      string `json:"Code"`
	Ccy       string `json:"Ccy"`
	CcyNm_RU  string `json:"CcyNm_RU"`
	CcyNm_UZ  string `json:"CcyNm_UZ"`
	CcyNm_UZC string `json:"CcyNm_UZC"`
	CcyNm_EN  string `json:"CcyNm_EN"`
	Nominal   string `json:"Nominal"`
	Rate      string `json:"Rate"`
	Diff      string `json:"Diff"`
	Date      string `json:"Date"`
}

type Monies struct {
	Objs []Money
}

func (m Monies) GetAll() *Monies {
	return &Monies{Objs: m.Objs}
}

func (m Monies) Sorted() *Monies {
	sort.Slice(m.Objs, func(i, j int) bool {
		return m.Objs[i].Id < m.Objs[j].Id
	})
	return &Monies{Objs: m.Objs}
}
