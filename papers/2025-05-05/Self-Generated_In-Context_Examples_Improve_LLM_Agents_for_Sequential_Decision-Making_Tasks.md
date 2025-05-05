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

这篇论文研究了如何通过自动学习大语言模型（LLM）代理在顺序决策任务中的成功经验来提升性能，而无需依赖特定任务的知识工程。作者构建了一个自生成示例数据库，发现仅通过积累成功轨迹就能显著提高ALFWorld（73%到89%）、Wordcraft（55%到64%）和InterCode-SQL（75%到79%）三个基准测试的表现。进一步引入基于群体的训练和示例级选择方法后，ALFWorld性能提升至91%，媲美复杂的手工优化方法。结果表明，自动轨迹数据库构建可有效替代费时的人工知识工程。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-05T23:02:07Z
- **目录日期**: 2025-05-05
