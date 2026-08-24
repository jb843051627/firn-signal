基于 Go 实现的冰川冰芯与积雪层钻孔热信号采集 Web 项目，一款现场操作后端服务，完成钻孔、探头、温度读数、校准、深度剖面重建、质量评估与发布管理。

# Firn signal baseline

The service stores borehole thermal scans in a file-backed SQLite database. It is a
single Go process and serves the operations page and JSON API from the same binary.

Build with `./build_benzhi_docker.sh firn-signal-run linux/amd64` or
`./build_benzhi_docker.sh firn-signal-run linux/arm64`. The container keeps the
complete Go toolchain so the repository can be inspected, built, and tested offline.

Set `FIRN_DB` to choose the database file and `FIRN_ADDR` to choose the HTTP address.
