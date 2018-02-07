package eadicomcollector

type EaCompRequest struct {
	Day            string
	Month          string
	Year           string
	Pid            string
	OutputDir      string
}

type EaCompResponse struct {
	Guid string
}
