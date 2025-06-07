# Kinetics: Rethinking Test-Time Scaling Laws

**URL**: http://arxiv.org/abs/2506.05333v1

## 原始摘要

We rethink test-time scaling laws from a practical efficiency perspective,
revealing that the effectiveness of smaller models is significantly
overestimated. Prior work, grounded in compute-optimality, overlooks critical
memory access bottlenecks introduced by inference-time strategies (e.g.,
Best-of-$N$, long CoTs). Our holistic analysis, spanning models from 0.6B to
32B parameters, reveals a new Kinetics Scaling Law that better guides resource
allocation by incorporating both computation and memory access costs. Kinetics
Scaling Law suggests that test-time compute is more effective when used on
models above a threshold than smaller ones. A key reason is that in TTS,
attention, rather than parameter count, emerges as the dominant cost factor.
Motivated by this, we propose a new scaling paradigm centered on sparse
attention, which lowers per-token cost and enables longer generations and more
parallel samples within the same resource budget. Empirically, we show that
sparse attention models consistently outperform dense counterparts, achieving
over 60 points gains in low-cost regimes and over 5 points gains in high-cost
regimes for problem-solving accuracy on AIME, encompassing evaluations on
state-of-the-art MoEs. These results suggest that sparse attention is essential
for realizing the full potential of test-time scaling because, unlike training,
where parameter scaling saturates, test-time accuracy continues to improve
through increased generation. The code is available at
https://github.com/Infini-AI-Lab/Kinetics.


## AI 摘要

该研究从实际效率角度重新审视测试阶段的扩展规律，发现小模型的有效性被显著高估。传统基于计算最优性的方法忽视了推理策略（如Best-of-N、长链思维）引入的内存访问瓶颈。通过分析0.6B到32B参数模型，提出了同时考虑计算和内存成本的"动力学扩展定律"，表明资源应优先分配给超过阈值的大模型，因为注意力机制（而非参数量）成为主要成本因素。研究提出以稀疏注意力为核心的新扩展范式，实证显示稀疏模型在相同资源下性能显著优于稠密模型，在AIME基准上最高可获得60分提升。稀疏注意力能充分发挥测试阶段扩展潜力，因为生成量的增加会持续提升准确率。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-07T10:01:18Z
- **目录日期**: 2025-06-07
