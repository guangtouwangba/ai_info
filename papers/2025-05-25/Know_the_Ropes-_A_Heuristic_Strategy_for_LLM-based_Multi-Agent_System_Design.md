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

大型语言模型（LLM）作为单一智能体存在上下文限制、角色过载和领域迁移脆弱性等问题。传统多智能体方案虽能缓解，但带来了任务分解模糊、合约不明确和验证开销等新问题。为此，研究者提出Know-The-Ropes（KtR）框架，将领域先验转化为层次化算法蓝图，通过类型化子任务分解和轻量级增强（如思维链、微调等）提升性能。实验显示：在背包问题（3-8项）中，3个GPT-4o-mini智能体将准确率从3%提升至95%；在任务分配问题（6-15项）中，6个o3-mini智能体实现最高100%准确率。该方法证明：算法感知分解+定向增强可使中小模型达成可靠协作，无需依赖巨型单体模型。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-25T09:02:16Z
- **目录日期**: 2025-05-25
