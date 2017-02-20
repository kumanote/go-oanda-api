package data

type GetPositions struct {
	Positions []Position `json:"positions"`
}

type Position struct {
	Instrument string  `json:"instrument"`
	Units      int     `json:"units"`
	Side       string  `json:"side"`
	AvgPrice   float64 `json:"avgPrice"`
}

type ClosePositionResult struct {
	Ids        []int64 `json:"ids"`
	Instrument string  `json:"instrument"`
	TotalUnits int     `json:"totalUnits"`
	Price      float64 `json:"price"`
}
