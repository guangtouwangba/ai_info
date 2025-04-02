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

BALROG是一个评估大语言模型(LLMs)和视觉语言模型(VLMs)在复杂环境中代理能力的新基准。它整合了多种难度级别的强化学习环境，从简单任务到极难任务(如NetHack)。研究发现，当前模型在简单游戏中表现尚可，但在复杂任务中表现不佳，尤其是基于视觉的决策能力存在严重缺陷——某些模型在提供视觉环境时表现反而更差。该基准采用细粒度指标进行评估，并开源以促进未来研究。代码和排行榜可在balrogai.com获取。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-02T07:01:51Z
- **目录日期**: 2025-04-02
