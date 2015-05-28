package main

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

const htmlData = "PCFkb2N0eXBlIGh0bWw+Cgo8aHRtbD4KCjxoZWFkPgogICAgPHRpdGxlPkhsNyBVSTwvdGl0bGU+CiAgICA8bWV0YSBuYW1lPSJ2aWV3cG9ydCIgY29udGVudD0id2lkdGg9ZGV2aWNlLXdpZHRoIj4KICAgIDxsaW5rIHJlbD0ic3R5bGVzaGVldCIgaHJlZj0iaHR0cHM6Ly9uZXRkbmEuYm9vdHN0cmFwY2RuLmNvbS9ib290c3dhdGNoLzMuMC4wL3NsYXRlL2Jvb3RzdHJhcC5taW4uY3NzIj4KICAgIDxzY3JpcHQgdHlwZT0idGV4dC9qYXZhc2NyaXB0IiBzcmM9Imh0dHBzOi8vYWpheC5nb29nbGVhcGlzLmNvbS9hamF4L2xpYnMvanF1ZXJ5LzIuMC4zL2pxdWVyeS5taW4uanMiPjwvc2NyaXB0PgogICAgPHNjcmlwdCB0eXBlPSJ0ZXh0L2phdmFzY3JpcHQiIHNyYz0iaHR0cHM6Ly9uZXRkbmEuYm9vdHN0cmFwY2RuLmNvbS9ib290c3RyYXAvMy4xLjEvanMvYm9vdHN0cmFwLm1pbi5qcyI+PC9zY3JpcHQ+CiAgICA8c3R5bGUgdHlwZT0idGV4dC9jc3MiPgogICAgICAgIGJvZHkgewogICAgICAgICAgICBwYWRkaW5nLXRvcDogMjBweDsKICAgICAgICB9CgogICAgICAgIC5mb290ZXIgewogICAgICAgICAgICBib3JkZXItdG9wOiAxcHggc29saWQgI2VlZTsKICAgICAgICAgICAgbWFyZ2luLXRvcDogNDBweDsKICAgICAgICAgICAgcGFkZGluZy10b3A6IDQwcHg7CiAgICAgICAgICAgIHBhZGRpbmctYm90dG9tOiA0MHB4OwogICAgICAgIH0KICAgICAgICAvKiBNYWluIG1hcmtldGluZyBtZXNzYWdlIGFuZCBzaWduIHVwIGJ1dHRvbiAqLwoKICAgICAgICAuanVtYm90cm9uIHsKICAgICAgICAgICAgdGV4dC1hbGlnbjogY2VudGVyOwogICAgICAgICAgICBiYWNrZ3JvdW5kLWNvbG9yOiB0cmFuc3BhcmVudDsKICAgICAgICB9CgogICAgICAgIC5qdW1ib3Ryb24gLmJ0biB7CiAgICAgICAgICAgIGZvbnQtc2l6ZTogMjFweDsKICAgICAgICAgICAgcGFkZGluZzogMTRweCAyNHB4OwogICAgICAgIH0KICAgICAgICAvKiBDdXN0b21pemUgdGhlIG5hdi1qdXN0aWZpZWQgbGlua3MgdG8gYmUgZmlsbCB0aGUgZW50aXJlIHNwYWNlIG9mIHRoZSAubmF2YmFyICovCgogICAgICAgIC5uYXYtanVzdGlmaWVkIHsKICAgICAgICAgICAgYmFja2dyb3VuZC1jb2xvcjogI2VlZTsKICAgICAgICAgICAgYm9yZGVyLXJhZGl1czogNXB4OwogICAgICAgICAgICBib3JkZXI6IDFweCBzb2xpZCAjY2NjOwogICAgICAgIH0KCiAgICAgICAgLm5hdi1qdXN0aWZpZWQgPiBsaSA+IGEgewogICAgICAgICAgICBwYWRkaW5nLXRvcDogMTVweDsKICAgICAgICAgICAgcGFkZGluZy1ib3R0b206IDE1cHg7CiAgICAgICAgICAgIGNvbG9yOiAjNzc3OwogICAgICAgICAgICBmb250LXdlaWdodDogYm9sZDsKICAgICAgICAgICAgdGV4dC1hbGlnbjogY2VudGVyOwogICAgICAgICAgICBib3JkZXItYm90dG9tOiAxcHggc29saWQgI2Q1ZDVkNTsKICAgICAgICAgICAgYmFja2dyb3VuZC1jb2xvcjogI2U1ZTVlNTsKICAgICAgICAgICAgLyogT2xkIGJyb3dzZXJzICovCgogICAgICAgICAgICBiYWNrZ3JvdW5kLXJlcGVhdDogcmVwZWF0LXg7CiAgICAgICAgICAgIC8qIFJlcGVhdCB0aGUgZ3JhZGllbnQgKi8KCiAgICAgICAgICAgIGJhY2tncm91bmQtaW1hZ2U6IC1tb3otbGluZWFyLWdyYWRpZW50KHRvcCwgI2Y1ZjVmNSAwJSwgI2U1ZTVlNSAxMDAlKTsKICAgICAgICAgICAgLyogRkYzLjYrICovCgogICAgICAgICAgICBiYWNrZ3JvdW5kLWltYWdlOiAtd2Via2l0LWdyYWRpZW50KGxpbmVhciwgbGVmdCB0b3AsIGxlZnQgYm90dG9tLCBjb2xvci1zdG9wKDAlLCAjZjVmNWY1KSwgY29sb3Itc3RvcCgxMDAlLCAjZTVlNWU1KSk7CiAgICAgICAgICAgIC8qIENocm9tZSxTYWZhcmk0KyAqLwoKICAgICAgICAgICAgYmFja2dyb3VuZC1pbWFnZTogLXdlYmtpdC1saW5lYXItZ3JhZGllbnQodG9wLCAjZjVmNWY1IDAlLCAjZTVlNWU1IDEwMCUpOwogICAgICAgICAgICAvKiBDaHJvbWUgMTArLFNhZmFyaSA1LjErICovCgogICAgICAgICAgICBiYWNrZ3JvdW5kLWltYWdlOiAtbXMtbGluZWFyLWdyYWRpZW50KHRvcCwgI2Y1ZjVmNSAwJSwgI2U1ZTVlNSAxMDAlKTsKICAgICAgICAgICAgLyogSUUxMCsgKi8KCiAgICAgICAgICAgIGJhY2tncm91bmQtaW1hZ2U6IC1vLWxpbmVhci1ncmFkaWVudCh0b3AsICNmNWY1ZjUgMCUsICNlNWU1ZTUgMTAwJSk7CiAgICAgICAgICAgIC8qIE9wZXJhIDExLjEwKyAqLwoKICAgICAgICAgICAgZmlsdGVyOiBwcm9naWQ6IERYSW1hZ2VUcmFuc2Zvcm0uTWljcm9zb2Z0LmdyYWRpZW50KHN0YXJ0Q29sb3JzdHI9JyNmNWY1ZjUnLCBlbmRDb2xvcnN0cj0nI2U1ZTVlNScsIEdyYWRpZW50VHlwZT0wKTsKICAgICAgICAgICAgLyogSUU2LTkgKi8KCiAgICAgICAgICAgIGJhY2tncm91bmQtaW1hZ2U6IGxpbmVhci1ncmFkaWVudCh0b3AsICNmNWY1ZjUgMCUsICNlNWU1ZTUgMTAwJSk7CiAgICAgICAgICAgIC8qIFczQyAqLwogICAgICAgIH0KCiAgICAgICAgLm5hdi1qdXN0aWZpZWQgPiAuYWN0aXZlID4gYSwKICAgICAgICAubmF2LWp1c3RpZmllZCA+IC5hY3RpdmUgPiBhOmhvdmVyLAogICAgICAgIC5uYXYtanVzdGlmaWVkID4gLmFjdGl2ZSA+IGE6Zm9jdXMgewogICAgICAgICAgICBiYWNrZ3JvdW5kLWNvbG9yOiAjZGRkOwogICAgICAgICAgICBiYWNrZ3JvdW5kLWltYWdlOiBub25lOwogICAgICAgICAgICBib3gtc2hhZG93OiBpbnNldCAwIDNweCA3cHggcmdiYSgwLCAwLCAwLCAuMTUpOwogICAgICAgIH0KCiAgICAgICAgLm5hdi1qdXN0aWZpZWQgPiBsaTpmaXJzdC1jaGlsZCA+IGEgewogICAgICAgICAgICBib3JkZXItcmFkaXVzOiA1cHggNXB4IDAgMDsKICAgICAgICB9CgogICAgICAgIC5uYXYtanVzdGlmaWVkID4gbGk6bGFzdC1jaGlsZCA+IGEgewogICAgICAgICAgICBib3JkZXItYm90dG9tOiAwOwogICAgICAgICAgICBib3JkZXItcmFkaXVzOiAwIDAgNXB4IDVweDsKICAgICAgICB9CgogICAgICAgIEBtZWRpYShtaW4td2lkdGg6IDc2OHB4KSB7CiAgICAgICAgICAgIC5uYXYtanVzdGlmaWVkIHsKICAgICAgICAgICAgICAgIG1heC1oZWlnaHQ6IDUycHg7CiAgICAgICAgICAgIH0KICAgICAgICAgICAgLm5hdi1qdXN0aWZpZWQgPiBsaSA+IGEgewogICAgICAgICAgICAgICAgYm9yZGVyLWxlZnQ6IDFweCBzb2xpZCAjZmZmOwogICAgICAgICAgICAgICAgYm9yZGVyLXJpZ2h0OiAxcHggc29saWQgI2Q1ZDVkNTsKICAgICAgICAgICAgfQogICAgICAgICAgICAubmF2LWp1c3RpZmllZCA+IGxpOmZpcnN0LWNoaWxkID4gYSB7CiAgICAgICAgICAgICAgICBib3JkZXItbGVmdDogMDsKICAgICAgICAgICAgICAgIGJvcmRlci1yYWRpdXM6IDVweCAwIDAgNXB4OwogICAgICAgICAgICB9CiAgICAgICAgICAgIC5uYXYtanVzdGlmaWVkID4gbGk6bGFzdC1jaGlsZCA+IGEgewogICAgICAgICAgICAgICAgYm9yZGVyLXJhZGl1czogMCA1cHggNXB4IDA7CiAgICAgICAgICAgICAgICBib3JkZXItcmlnaHQ6IDA7CiAgICAgICAgICAgIH0KICAgICAgICB9CiAgICAgICAgLyogUmVzcG9uc2l2ZTogUG9ydHJhaXQgdGFibGV0cyBhbmQgdXAgKi8KCiAgICAgICAgQG1lZGlhIHNjcmVlbiBhbmQobWluLXdpZHRoOiA3NjhweCkgewogICAgICAgICAgICAvKiBSZW1vdmUgdGhlIHBhZGRpbmcgd2Ugc2V0IGVhcmxpZXIgKi8KCiAgICAgICAgICAgIC5tYXN0aGVhZCwgLm1hcmtldGluZywgLmZvb3RlciB7CiAgICAgICAgICAgICAgICBwYWRkaW5nLWxlZnQ6IDA7CiAgICAgICAgICAgICAgICBwYWRkaW5nLXJpZ2h0OiAwOwogICAgICAgICAgICB9CiAgICAgICAgfQogICAgPC9zdHlsZT4KICAgIDxzY3JpcHQgdHlwZT0idGV4dC9qYXZhc2NyaXB0Ij4KICAgICAgICBmdW5jdGlvbiBzd2l0Y2hlbmNvZGUoKSB7CiAgICAgICAgICAgIGlmIChkb2N1bWVudC5nZXRFbGVtZW50QnlJZCgiY2hlY2tib3gtZW5jIikKICAgICAgICAgICAgICAgIC5jaGVja2VkKSB7CiAgICAgICAgICAgICAgICAkKCIjdGV4dC1hcmVhIikKICAgICAgICAgICAgICAgICAgICAudmFsKGJ0b2EoJCgiI3RleHQtYXJlYSIpCiAgICAgICAgICAgICAgICAgICAgICAgIC52YWwoKSkpCiAgICAgICAgICAgIH0gZWxzZSB7CgogICAgICAgICAgICAgICAgJCgiI3RleHQtYXJlYSIpCiAgICAgICAgICAgICAgICAgICAgLnZhbChhdG9iKCQoIiN0ZXh0LWFyZWEiKQogICAgICAgICAgICAgICAgICAgICAgICAudmFsKCkpKQogICAgICAgICAgICB9CgogICAgICAgIH0KCiAgICAgICAgZnVuY3Rpb24gc2VuZGhsN2VuY2RhdGEoKSB7CiAgICAgICAgICAgIHZhciByZXEgPSB7CiAgICAgICAgICAgICAgICBIbDdjOiB7CiAgICAgICAgICAgICAgICAgICAgU2VySXA6ICQoIiNhZGRyZXNzLWlkIikKICAgICAgICAgICAgICAgICAgICAgICAgLnZhbCgpLAogICAgICAgICAgICAgICAgICAgIFBvcnQ6ICQoIiNwb3J0LWlkIikKICAgICAgICAgICAgICAgICAgICAgICAgLnZhbCgpLAogICAgICAgICAgICAgICAgfSwKICAgICAgICAgICAgICAgIGRhdGE6ICQoIiN0ZXh0LWFyZWEiKQogICAgICAgICAgICAgICAgICAgIC52YWwoKQogICAgICAgICAgICB9OwogICAgICAgICAgICAkLmFqYXgoewogICAgICAgICAgICAgICAgICAgIHVybDogIi9zZW5kbXNnIiwKICAgICAgICAgICAgICAgICAgICB0eXBlOiAiUE9TVCIsCiAgICAgICAgICAgICAgICAgICAgZGF0YTogSlNPTi5zdHJpbmdpZnkocmVxKSwKICAgICAgICAgICAgICAgICAgICBkYXRhVHlwZTogInRleHQiCiAgICAgICAgICAgICAgICB9KQogICAgICAgICAgICAgICAgLnN1Y2Nlc3MoZnVuY3Rpb24oanNvbkRhdGEpIHsKICAgICAgICAgICAgICAgICAgICB2YXIgaW5lckh0bWwgPSAiIgogICAgICAgICAgICAgICAgICAgIGluZXJIdG1sID0gaW5lckh0bWwuY29uY2F0KCc8ZGl2IGNsYXNzPSJhbGVydCBhbGVydC13YXJuaW5nIGFsZXJ0LWRpc21pc3NhYmxlIj48YnV0dG9uIHR5cGU9ImJ1dHRvbiIgY2xhc3M9ImNsb3NlIiBkYXRhLWRpc21pc3M9ImFsZXJ0Ij4mdGltZXM7PC9idXR0b24+PGI+SGw3IHNlcnZlciByZXNwb25zZTwvYj4gJykKICAgICAgICAgICAgICAgICAgICBpbmVySHRtbCA9IGluZXJIdG1sLmNvbmNhdChqc29uRGF0YSkKICAgICAgICAgICAgICAgICAgICBpbmVySHRtbCA9IGluZXJIdG1sLmNvbmNhdCgnIDwvZGl2PicpCiAgICAgICAgICAgICAgICAgICAgJCgiI3dlbGwtaWQiKQogICAgICAgICAgICAgICAgICAgICAgICAuYXBwZW5kKGluZXJIdG1sKQogICAgICAgICAgICAgICAgfSkKICAgICAgICAgICAgICAgIC5lcnJvcihmdW5jdGlvbihqc29uRGF0YSkgewogICAgICAgICAgICAgICAgICAgIHZhciBpbmVySHRtbCA9ICIiCiAgICAgICAgICAgICAgICAgICAgaW5lckh0bWwgPSBpbmVySHRtbC5jb25jYXQoJzxkaXYgY2xhc3M9ImFsZXJ0IGFsZXJ0LXdhcm5pbmcgYWxlcnQtZGlzbWlzc2FibGUiPjxidXR0b24gdHlwZT0iYnV0dG9uIiBjbGFzcz0iY2xvc2UiIGRhdGEtZGlzbWlzcz0iYWxlcnQiPiZ0aW1lczs8L2J1dHRvbj48cD5Db25uZXRpb24gZXJyb3I8L3A+ICcpCiAgICAgICAgICAgICAgICAgICAgaW5lckh0bWwgPSBpbmVySHRtbC5jb25jYXQoSlNPTi5zdHJpbmdpZnkoanNvbkRhdGEpKQogICAgICAgICAgICAgICAgICAgIGluZXJIdG1sID0gaW5lckh0bWwuY29uY2F0KCcgPC9kaXY+JykKICAgICAgICAgICAgICAgICAgICAkKCIjd2VsbC1pZCIpCiAgICAgICAgICAgICAgICAgICAgICAgIC5hcHBlbmQoaW5lckh0bWwpCiAgICAgICAgICAgICAgICB9KQogICAgICAgIH0KICAgIDwvc2NyaXB0Pgo8L2hlYWQ+Cgo8Ym9keT4KICAgIDxkaXYgY2xhc3M9ImNvbnRhaW5lciI+CiAgICAgICAgPGRpdiBjbGFzcz0id2VsbCIgaWQ9IndlbGwtaWQiPgogICAgICAgICAgICA8ZGl2IGNsYXNzPSJwYW5lbC1mb290ZXIiPkhMNyBTZXJ2ZXIgc2V0dGluZ3MKICAgICAgICAgICAgPC9kaXY+CiAgICAgICAgICAgIDx0YWJsZSBjbGFzcz0idGFibGUgdGFibGUtYm9yZGVyZWQgdGFibGUtY29uZGVuc2VkIHRhYmxlLWhvdmVyIHRhYmxlLXN0cmlwZWQiPgogICAgICAgICAgICAgICAgPHRib2R5PgogICAgICAgICAgICAgICAgICAgIDx0cj4KICAgICAgICAgICAgICAgICAgICAgICAgPHRkPgogICAgICAgICAgICAgICAgICAgICAgICAgICAgPGRpdiBjbGFzcz0iZm9ybS1ncm91cCI+CiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgPGxhYmVsIGNsYXNzPSJjb250cm9sLWxhYmVsIj5ITDcgc2VydmVyIGFkZHJlc3M8L2xhYmVsPgogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIDxkaXYgY2xhc3M9ImNvbnRyb2xzIj4KICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgPGlucHV0IGNsYXNzPSJmb3JtLWNvbnRyb2wgaW5wdXQtc20iIGlkPSJhZGRyZXNzLWlkIiB2YWx1ZT0iMTkyLjE2OC4xMjMuMTA3IiB0eXBlPSJ0ZXh0Ij4KICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICA8L2Rpdj4KICAgICAgICAgICAgICAgICAgICAgICAgICAgIDwvZGl2PgogICAgICAgICAgICAgICAgICAgICAgICA8L3RkPgogICAgICAgICAgICAgICAgICAgICAgICA8dGQ+CiAgICAgICAgICAgICAgICAgICAgICAgICAgICA8ZGl2IGNsYXNzPSJmb3JtLWdyb3VwIj4KICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICA8bGFiZWwgY2xhc3M9ImNvbnRyb2wtbGFiZWwiPlBvcnQgbnVtYmVyPC9sYWJlbD4KICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICA8ZGl2IGNsYXNzPSJjb250cm9scyI+CiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgIDxpbnB1dCBjbGFzcz0iZm9ybS1jb250cm9sIGlucHV0LXNtIiBpZD0icG9ydC1pZCIgdmFsdWU9IjE5MDE5IiB0eXBlPSJ0ZXh0Ij4KICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICA8L2Rpdj4KICAgICAgICAgICAgICAgICAgICAgICAgICAgIDwvZGl2PgogICAgICAgICAgICAgICAgICAgICAgICA8L3RkPgogICAgICAgICAgICAgICAgICAgICAgICA8dGQ+CiAgICAgICAgICAgICAgICAgICAgICAgICAgICA8ZGl2IGNsYXNzPSJmb3JtLWdyb3VwIj4KICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICA8bGFiZWwgY2xhc3M9ImNvbnRyb2wtbGFiZWwiPjwvbGFiZWw+CiAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgPGRpdiBjbGFzcz0iY29udHJvbHMiPgogICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICA8YSBvbmNsaWNrPSJzZW5kaGw3ZW5jZGF0YSgpIiBjbGFzcz0iYnRuIHB1bGwtbGVmdCBidG4taW5mbyI+UyBFIE4gRDwvYT4KICAgICAgICAgICAgICAgICAgICAgICAgICAgICAgICA8L2Rpdj4KICAgICAgICAgICAgICAgICAgICAgICAgICAgIDwvZGl2PgogICAgICAgICAgICAgICAgICAgICAgICA8L3RkPgogICAgICAgICAgICAgICAgICAgIDwvdHI+CiAgICAgICAgICAgICAgICA8L3Rib2R5PgogICAgICAgICAgICA8L3RhYmxlPgogICAgICAgICAgICA8ZGl2IGNsYXNzPSJwYW5lbC1mb290ZXIiIGlkPSJzZWFyY2gtZm9vdGVyIj5FY25vZGVkIERhdGEKICAgICAgICAgICAgPC9kaXY+CiAgICAgICAgICAgIDx0ZXh0YXJlYSBjbGFzcz0iZm9ybS1jb250cm9sIGlucHV0LWJsb2NrLWxldmVsIiBzdHlsZT0ibWFyZ2luLXRvcDogMHB4OyBtYXJnaW4tYm90dG9tOiAwcHg7IGhlaWdodDogMzYwcHg7IiBpZD0idGV4dC1hcmVhIj48L3RleHRhcmVhPgogICAgICAgICAgICA8ZGl2IGNsYXNzPSJjaGVja2JveCI+CiAgICAgICAgICAgICAgICA8bGFiZWw+CiAgICAgICAgICAgICAgICAgICAgPGlucHV0IHR5cGU9ImNoZWNrYm94IiBpZD0iY2hlY2tib3gtZW5jIiBvbmNsaWNrPSJzd2l0Y2hlbmNvZGUoKSIgdmFsdWU9InRydWUiPkVuY29kZSB0byBCYXNlNjQ8L2xhYmVsPgogICAgICAgICAgICA8L2Rpdj4KICAgICAgICA8L2Rpdj4KICAgIDwvZGl2Pgo8L2JvZHk+Cgo8L2h0bWw+Cg=="

//main srv class
type EaFolderCompressorSrv struct {
	fcom FolderCompressor
}

//start and init srv
func (srv *EaFolderCompressorSrv) Start(listenPort int) error {
	http.HandleFunc("/compress", srv.Compress)
	http.HandleFunc("/", srv.Redirect)
	http.HandleFunc("/index.html", srv.index)
	if err := http.ListenAndServe(":"+strconv.Itoa(listenPort), nil); err != nil {
		return errors.New("error: can't start listen http server")
	}
	return nil
}

//serve main page request
func (srv *EaFolderCompressorSrv) index(rwr http.ResponseWriter, req *http.Request) {
	rwr.Header().Set("Content-Type: text/html", "*")

	content, err := ioutil.ReadFile("index.html")
	if err != nil {
		log.Println("warning: start page not found, return included page")
		val, _ := base64.StdEncoding.DecodeString(htmlData)
		rwr.Write(val)
		return
	}
	rwr.Write(content)
}

func (service *EaFolderCompressorSrv) Redirect(responseWriter http.ResponseWriter, request *http.Request) {
	http.Redirect(responseWriter, request, "/index.html", 301)
}

//serve cEcho responce
func (srv *EaFolderCompressorSrv) Compress(rwr http.ResponseWriter, httpreq *http.Request) {
	defer httpreq.Body.Close()
	bodyData, err := ioutil.ReadAll(httpreq.Body)
	if err != nil {
		strErr := "error: Can't read http body data"
		http.Error(rwr, err.Error(), http.StatusInternalServerError)
		log.Println(strErr)
		return
	}
	//parse http request
	var req EaCompRequest
	if err := json.Unmarshal(bodyData, &req); err != nil {
		strErr := "error: can't parse request data"
		http.Error(rwr, err.Error(), http.StatusInternalServerError)
		log.Println(strErr)
		return
	}
	TgF := req.ArchPathPrefix + "\\" + req.Year + "\\" + req.Month + "\\" + req.Day
	NHash := req.ArchPathPrefix + req.Year + req.Month + req.Day + req.Pid
	ResF := req.OutputDir + "\\" + NHash + ".zip"

	err = srv.fcom.CompressFolder("7z.exe", "a", "-tzip", TgF, ResF)
	if err != nil {
		log.Println(err)
		http.Error(rwr, err.Error(), http.StatusInternalServerError)
		return

	}
	res := EaCompResponse{Guid: NHash}
	js, err := json.Marshal(res)
	if err != nil {
		http.Error(rwr, err.Error(), http.StatusInternalServerError)
		return
	}
	rwr.Header().Set("Content-Type", "application/json")
	rwr.Write(js)
}
