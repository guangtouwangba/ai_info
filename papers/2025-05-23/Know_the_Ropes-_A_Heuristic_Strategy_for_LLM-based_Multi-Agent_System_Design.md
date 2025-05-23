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

Know-The-Ropes (KtR) 是一个多智能体框架，通过将领域先验转化为层次化算法蓝图来解决单一大模型面临的上下文限制、角色过载和领域迁移问题。KtR递归分解任务为类型化子任务，采用轻量级增强技术（如思维链、微调等）实现高效协作。实验显示：在背包问题（3-8项）中，3个GPT-4o-mini智能体将准确率从零样本的3%提升至95%；在任务分配问题（6-15项）中，6个o3-mini智能体实现小规模100%、大规模84%的准确率，远超零样本11%的表现。该方法证明算法感知分解结合定向增强可使轻量模型达成可靠协作，无需持续扩大模型规模。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-23T03:20:03Z
- **目录日期**: 2025-05-23
