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

该研究提出了一种无需任务特定知识工程（如提示调整或人工筛选示例）的LLM智能体自学习方法。通过构建和优化自生成的成功轨迹数据库，智能体能在类似任务中通过上下文学习自动提升性能。实验表明，在ALFWorld、Wordcraft和InterCode-SQL三个基准上，仅简单累积成功轨迹即可显著提升测试表现（如ALFWorld从73%升至89%）。进一步引入基于群体训练的数据库级筛选和基于效用的轨迹级筛选后，ALFWorld准确率达91%，媲美复杂的人工优化方法。这为替代费力的知识工程提供了高效自动化方案。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-05T06:02:06Z
- **目录日期**: 2025-05-05
