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

这篇论文研究了如何通过自动学习自身成功经验来提升大语言模型(LLM)在序列决策任务中的表现，而无需依赖任务特定的知识工程。研究发现，简单积累训练任务中的成功轨迹就能显著提升在ALFWorld(73%→89%)、Wordcraft(55%→64%)和InterCode-SQL(75%→79%)三个基准测试中的表现。进一步提出的两种扩展方法——基于群体训练的数据库级选择和基于经验效用的样本级选择——将ALFWorld上的表现提升至91%，与需要复杂任务特定组件的方法相当。这表明自动构建轨迹数据库是替代人工知识工程的有效方案。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-05T17:02:08Z
- **目录日期**: 2025-05-05
