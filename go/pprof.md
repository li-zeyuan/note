下载pporf：
cpu：export PPROF_TMPDIR=/home/foo; go tool pprof http://127.0.0.1:31513/debug/pprof/profile?second=30
内存：export PPROF_TMPDIR=/home/foo; go tool pprof --inuse_space http://127.0.0.1:31513/debug/pprof/heap?second=30
分析pporf：
cpu：go tool pprof --http=0.0.0.0:9999 pprof.xxx.cpu.008.pb.gz
内存：go tool pprof --http=0.0.0.0:9999 pprof.xxxx.008.pb.gz
