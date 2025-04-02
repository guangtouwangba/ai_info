# BALROG: Benchmarking Agentic LLM and VLM Reasoning On Games

**URL**: http://arxiv.org/abs/2411.13543v2

## 原始摘要

Large Language Models (LLMs) and Vision Language Models (VLMs) possess
extensive knowledge and exhibit promising reasoning abilities, however, they
still struggle to perform well in complex, dynamic environments. Real-world
tasks require handling intricate interactions, advanced spatial reasoning,
long-term planning, and continuous exploration of new strategies-areas in which
we lack effective methodologies for comprehensively evaluating these
capabilities. To address this gap, we introduce BALROG, a novel benchmark
designed to assess the agentic capabilities of LLMs and VLMs through a diverse
set of challenging games. Our benchmark incorporates a range of existing
reinforcement learning environments with varying levels of difficulty,
including tasks that are solvable by non-expert humans in seconds to extremely
challenging ones that may take years to master (e.g., the NetHack Learning
Environment). We devise fine-grained metrics to measure performance and conduct
an extensive evaluation of several popular open-source and closed-source LLMs
and VLMs. Our findings indicate that while current models achieve partial
success in the easier games, they struggle significantly with more challenging
tasks. Notably, we observe severe deficiencies in vision-based decision-making,
as several models perform worse when visual representations of the environments
are provided. We release BALROG as an open and user-friendly benchmark to
facilitate future research and development in the agentic community. Code and
Leaderboard at balrogai.com.


## AI 摘要

BALROG是一个新型基准测试，旨在评估大语言模型(LLM)和视觉语言模型(VLM)在复杂动态环境中的智能体能力。该基准整合了多种难度级别的强化学习环境，从人类可快速解决到极难任务(如NetHack)。研究发现，当前模型在简单任务中表现尚可，但在复杂任务(尤其是基于视觉的决策)中表现显著不足，部分模型在提供视觉信息时性能反而下降。BALROG提供了细粒度评估指标，并作为开源基准发布，以促进智能体领域的研究发展。代码和排行榜可在balrogai.com获取。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-02T18:01:48Z
- **目录日期**: 2025-04-02
