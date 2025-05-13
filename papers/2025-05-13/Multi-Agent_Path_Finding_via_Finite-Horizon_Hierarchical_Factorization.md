# Multi-Agent Path Finding via Finite-Horizon Hierarchical Factorization

**URL**: http://arxiv.org/abs/2505.07779v1

## 原始摘要

We present a novel algorithm for large-scale Multi-Agent Path Finding (MAPF)
that enables fast, scalable planning in dynamic environments such as automated
warehouses. Our approach introduces finite-horizon hierarchical factorization,
a framework that plans one step at a time in a receding-horizon fashion. Robots
first compute individual plans in parallel, and then dynamically group based on
spatio-temporal conflicts and reachability. The framework accounts for conflict
resolution, and for immediate execution and concurrent planning, significantly
reducing response time compared to offline algorithms. Experimental results on
benchmark maps demonstrate that our method achieves up to 60% reduction in
time-to-first-action while consistently delivering high-quality solutions,
outperforming state-of-the-art offline baselines across a range of problem
sizes and planning horizons.


## AI 摘要

本文提出了一种新型大规模多智能体路径规划(MAPF)算法，适用于自动化仓库等动态环境。该算法采用有限时域分层分解框架，以滚动时域方式逐步规划：机器人先并行计算个体路径，再根据时空冲突和可达性动态分组。该方法支持冲突解决、即时执行和并行规划，相比离线算法显著降低响应时间。基准测试显示，该算法能减少60%的首动作时间，同时保持高质量解决方案，在各种问题规模和规划时域下均优于现有最优离线基准算法。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-13T18:01:44Z
- **目录日期**: 2025-05-13
