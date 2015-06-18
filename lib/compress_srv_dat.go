package eadicomcollector

type EaCompRequest struct {
	ArchPathPrefix string
	Day            string
	Month          string
	Year           string
	Pid            string
	OutputDir      string
}

type EaCompResponse struct {
	Guid string
}
