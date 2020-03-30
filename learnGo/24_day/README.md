go火焰图学习
https://www.jianshu.com/p/a22174de24c7
https://www.cnblogs.com/yjf512/archive/2012/12/27/2835331.html




go-torch -u http://127.0.0.1:6060 -t60 -p > profile-local.svg

cpu：
go-torch -u http://127.0.0.1:6060  --seconds 60 -f cpu.svg
内存：
go-torch  http://127.0.0.1:6060/debug/pprof/heap  mem  -f mem.svg


1.运行
go tool pprof httpdemo http://127.0.0.1:6060/debug/pprof/heap
2.输入web

1.运行
go tool pprof -alloc_objects httpdemo http://127.0.0.1:6060/debug/pprof/heap
2.输入web