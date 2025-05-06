# AutoLibra: Agent Metric Induction from Open-Ended Feedback

**URL**: http://arxiv.org/abs/2505.02820v1

## 原始摘要

Agents are predominantly evaluated and optimized via task success metrics,
which are coarse, rely on manual design from experts, and fail to reward
intermediate emergent behaviors. We propose AutoLibra, a framework for agent
evaluation, that transforms open-ended human feedback, e.g., "If you find that
the button is disabled, don't click it again", or "This agent has too much
autonomy to decide what to do on its own", into metrics for evaluating
fine-grained behaviors in agent trajectories. AutoLibra accomplishes this by
grounding feedback to an agent's behavior, clustering similar positive and
negative behaviors, and creating concrete metrics with clear definitions and
concrete examples, which can be used for prompting LLM-as-a-Judge as
evaluators. We further propose two meta-metrics to evaluate the alignment of a
set of (induced) metrics with open feedback: "coverage" and "redundancy".
Through optimizing these meta-metrics, we experimentally demonstrate
AutoLibra's ability to induce more concrete agent evaluation metrics than the
ones proposed in previous agent evaluation benchmarks and discover new metrics
to analyze agents. We also present two applications of AutoLibra in agent
improvement: First, we show that AutoLibra-induced metrics serve as better
prompt-engineering targets than the task success rate on a wide range of text
game tasks, improving agent performance over baseline by a mean of 20%. Second,
we show that AutoLibra can iteratively select high-quality fine-tuning data for
web navigation agents. Our results suggest that AutoLibra is a powerful
task-agnostic tool for evaluating and improving language agents.


## AI 摘要

本文提出AutoLibra框架，通过将开放式人类反馈转化为细粒度行为评估指标，解决传统任务成功率评估方法的局限性。该框架能将反馈与智能体行为关联，聚类正负行为，并生成可量化指标，用于大语言模型评估。实验表明，AutoLibra能产生比现有基准更具体的评估指标，发现新分析维度。应用方面：1)作为提示工程目标，在文本游戏中使智能体性能平均提升20%；2)迭代筛选高质量微调数据改进网页导航智能体。研究表明，AutoLibra是评估和改进语言智能体的通用有效工具。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-06T22:01:13Z
- **目录日期**: 2025-05-06
