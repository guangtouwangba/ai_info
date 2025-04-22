# FlowReasoner: Reinforcing Query-Level Meta-Agents

**URL**: http://arxiv.org/abs/2504.15257v1

## 原始摘要

This paper proposes a query-level meta-agent named FlowReasoner to automate
the design of query-level multi-agent systems, i.e., one system per user query.
Our core idea is to incentivize a reasoning-based meta-agent via external
execution feedback. Concretely, by distilling DeepSeek R1, we first endow the
basic reasoning ability regarding the generation of multi-agent systems to
FlowReasoner. Then, we further enhance it via reinforcement learning (RL) with
external execution feedback. A multi-purpose reward is designed to guide the RL
training from aspects of performance, complexity, and efficiency. In this
manner, FlowReasoner is enabled to generate a personalized multi-agent system
for each user query via deliberative reasoning. Experiments on both engineering
and competition code benchmarks demonstrate the superiority of FlowReasoner.
Remarkably, it surpasses o1-mini by 10.52% accuracy across three benchmarks.
The code is available at https://github.com/sail-sg/FlowReasoner.


## AI 摘要

这篇论文提出了一个名为FlowReasoner的查询级元代理，用于自动化设计针对每个用户查询的多代理系统。其核心思想是通过外部执行反馈激励基于推理的元代理。研究团队首先通过提炼DeepSeek R1赋予FlowReasoner生成多代理系统的基本推理能力，再通过强化学习结合外部反馈进行增强。设计的多目标奖励函数从性能、复杂度和效率三方面指导训练，使FlowReasoner能针对每个查询生成个性化的多代理系统。实验表明，该方法在工程和竞赛代码基准测试中表现优异，准确率比o1-mini高出10.52%。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-22T07:01:24Z
- **目录日期**: 2025-04-22
