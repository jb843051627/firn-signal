package api

import "net/http"

func (h *Handler) dashboard(response http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		http.NotFound(response, request)
		return
	}
	response.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = response.Write([]byte(`<!doctype html><html><head><meta charset="utf-8"><title>Firn Signal</title><style>body{font:16px system-ui;background:#e8f0f5;color:#163447;margin:2rem}main{max-width:920px;margin:auto;background:#fff;padding:2rem;border-radius:18px;box-shadow:0 8px 30px #183b4d22}h1{margin-top:0;color:#0c5d78}.grid{display:grid;grid-template-columns:repeat(auto-fit,minmax(180px,1fr));gap:1rem}.card{padding:1rem;background:#edf7fa;border-left:4px solid #1d8ca8;border-radius:8px}button{background:#0c5d78;color:white;border:0;padding:.7rem 1rem;border-radius:6px}</style></head><body><main><h1>Firn signal operations</h1><p>钻孔热信号现场校准与剖面验收</p><section class="grid"><div class="card"><strong id="boreholes">-</strong><br>active boreholes</div><div class="card"><strong id="scans">-</strong><br>thermal scans</div><div class="card"><strong id="status">loading</strong><br>database</div></section><p><button onclick="refresh()">Refresh field view</button></p><pre id="details"></pre></main><script>async function refresh(){const [b,s,h]=await Promise.all([fetch('/api/boreholes').then(r=>r.json()),fetch('/api/scans').catch(()=>({})),fetch('/healthz').then(r=>r.json())]);document.querySelector('#boreholes').textContent=Array.isArray(b)?b.length:'?';document.querySelector('#scans').textContent=Array.isArray(s)?s.length:'API';document.querySelector('#status').textContent=h.path||'ready';document.querySelector('#details').textContent=JSON.stringify(h,null,2)}refresh()</script></body></html>`))
}
