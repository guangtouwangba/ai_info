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

该研究从实际效率角度重新审视了测试阶段的扩展规律，发现小模型的效果被严重高估。传统基于计算最优性的方法忽视了推理策略（如Best-of-N、长链思维）引入的内存访问瓶颈。通过分析0.6B到32B参数的模型，研究者提出了包含计算和内存成本的"动力学扩展定律"，表明资源应优先分配给超过阈值的大模型，因为注意力机制（而非参数量）成为主要成本因素。基于此，研究提出了稀疏注意力的新扩展范式，实验显示在相同资源下，稀疏模型在AIME问题求解准确率上最高可提升60分，证明其是实现测试阶段扩展潜力的关键。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-06T23:01:14Z
- **目录日期**: 2025-06-06
