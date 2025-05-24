# Know the Ropes: A Heuristic Strategy for LLM-based Multi-Agent System Design

**URL**: http://arxiv.org/abs/2505.16979v1

## 原始摘要

Single-agent LLMs hit hard limits--finite context, role overload, and brittle
domain transfer. Conventional multi-agent fixes soften those edges yet expose
fresh pains: ill-posed decompositions, fuzzy contracts, and verification
overhead that blunts the gains. We therefore present Know-The-Ropes (KtR), a
framework that converts domain priors into an algorithmic blueprint hierarchy,
in which tasks are recursively split into typed, controller-mediated subtasks,
each solved zero-shot or with the lightest viable boost (e.g.,
chain-of-thought, micro-tune, self-check). Grounded in the No-Free-Lunch
theorem, KtR trades the chase for a universal prompt for disciplined
decomposition. On the Knapsack problem (3-8 items), three GPT-4o-mini agents
raise accuracy from 3% zero-shot to 95% on size-5 instances after patching a
single bottleneck agent. On the tougher Task-Assignment problem (6-15 jobs), a
six-agent o3-mini blueprint hits 100% up to size 10 and 84% on sizes 13-15,
versus 11% zero-shot. Algorithm-aware decomposition plus targeted augmentation
thus turns modest models into reliable collaborators--no ever-larger monoliths
required.


## AI 摘要

Know-The-Ropes (KtR) 提出了一种解决单一大模型局限性的框架，通过将领域先验转化为分层算法蓝图，递归分解任务为控制器协调的子任务，并采用轻量级增强技术（如思维链、微调等）。基于"没有免费午餐"定理，KtR放弃通用提示，转向结构化分解。实验显示：在背包问题中，3个GPT-4o-mini智能体将准确率从3%提升至95%；在更复杂的任务分配问题中，6个o3-mini智能体在10项任务上达到100%准确率，而零样本方法仅11%。该框架证明算法感知分解结合针对性增强可使小型模型成为可靠协作系统。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-24T18:02:12Z
- **目录日期**: 2025-05-24
