package models

import "time"

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

type ReportBigOperation struct {
	City         string    `json:"city"`
	OriginalCity string    `json:"original_city"`
	Amount       float64   `json:"amount"`
	CustomerName string    `json:"customer_name"`
	CustomerId   string    `json:"customer_id"`
	CreatedAt    time.Time `json:"created_at"`
}

func NewReportBigOperation() *ReportBigOperation {
	return &ReportBigOperation{}
}
