下载pporf：
cpu：export PPROF_TMPDIR=/home/foo; go tool pprof http://127.0.0.1:31777/debug/pprof/profile?second=30
内存：export PPROF_TMPDIR=/home/foo; go tool pprof --inuse_space http://127.0.0.1:31777/debug/pprof/heap?second=30
trace: curl -s http://localhost:1777/debug/pprof/trace?seconds=5 > trace-collector-trace.pprof

分析pporf：
cpu：go tool pprof --http=0.0.0.0:9998 pprof.xxx.cpu.008.pb.gz
内存：go tool pprof --http=0.0.0.0:9999 pprof.xxxx.008.pb.gz
trace: go tool trace trace-trace.pprof

cpu：export PPROF_TMPDIR=./; go tool pprof http://localhost:1777/debug/pprof/profile?second=30
内存：export PPROF_TMPDIR=/tracing/collector; go tool pprof --inuse_space http://localhost:1777/debug/pprof/heap?second=30