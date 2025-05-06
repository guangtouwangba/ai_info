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

AutoLibra是一个基于开放式人类反馈的智能体评估框架，将模糊反馈（如"按钮禁用时不应重复点击"）转化为可量化的细粒度行为指标。它通过关联反馈与行为、聚类正负案例，生成包含明确定义和示例的评估指标，并利用LLM作为裁判。该框架提出"覆盖率"和"冗余度"两个元指标优化评估体系。实验表明，AutoLibra生成的指标比传统任务成功率更能提升智能体性能（文本游戏任务平均提升20%），并能筛选高质量微调数据。该框架为语言智能体的评估与优化提供了通用工具。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-06T05:01:14Z
- **目录日期**: 2025-05-06
