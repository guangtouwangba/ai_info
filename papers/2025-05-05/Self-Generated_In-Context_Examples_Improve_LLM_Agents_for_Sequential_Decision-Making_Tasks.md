# Self-Generated In-Context Examples Improve LLM Agents for Sequential Decision-Making Tasks

**URL**: http://arxiv.org/abs/2505.00234v2

## 原始摘要

Many methods for improving Large Language Model (LLM) agents for sequential
decision-making tasks depend on task-specific knowledge engineering--such as
prompt tuning, curated in-context examples, or customized observation and
action spaces. Using these approaches, agent performance improves with the
quality or amount of knowledge engineering invested. Instead, we investigate
how LLM agents can automatically improve their performance by learning
in-context from their own successful experiences on similar tasks. Rather than
relying on task-specific knowledge engineering, we focus on constructing and
refining a database of self-generated examples. We demonstrate that even a
naive accumulation of successful trajectories across training tasks boosts test
performance on three benchmarks: ALFWorld (73% to 89%), Wordcraft (55% to 64%),
and InterCode-SQL (75% to 79%)--matching the performance the initial agent
achieves if allowed two to three attempts per task. We then introduce two
extensions: (1) database-level selection through population-based training to
identify high-performing example collections, and (2) exemplar-level selection
that retains individual trajectories based on their empirical utility as
in-context examples. These extensions further enhance performance, achieving
91% on ALFWorld--matching more complex approaches that employ task-specific
components and prompts. Our results demonstrate that automatic trajectory
database construction offers a compelling alternative to labor-intensive
knowledge engineering.


## AI 摘要

该研究探讨了如何通过自动学习自身成功经验来提升大语言模型（LLM）在序列决策任务中的表现，而无需依赖特定任务的知识工程。通过构建和优化自生成案例数据库，即使简单累积训练任务中的成功轨迹，也能显著提升ALFWorld（73%→89%）、Wordcraft（55%→64%）和InterCode-SQL（75%→79%）三个基准测试的表现，效果相当于允许初始模型每个任务尝试2-3次。进一步引入群体训练筛选高性能案例集（91% ALFWorld）和基于效用的轨迹选择方法后，性能甚至媲美复杂的手工优化方案，为替代人工密集型知识工程提供了新思路。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-05T02:31:36Z
- **目录日期**: 2025-05-05
