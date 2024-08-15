package models

type Report struct {
	QtyTransactions int    `json:"qty_transactions"`
	CustomerName    string `json:"customer_name"`
	CustomerId      string `json:"customer_id"`
	Month           int    `json:"month"`
	Year            int    `json:"year"`
}

func NewReportModel() *Report {
	return &Report{}
}
