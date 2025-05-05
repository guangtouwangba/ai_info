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

这篇论文研究了如何通过自动学习提升大语言模型(LLM)在序列决策任务中的表现，而无需依赖特定任务的知识工程。研究者构建了一个自生成的成功轨迹数据库，通过积累训练任务中的成功经验来提升测试性能。在ALFWorld、Wordcraft和InterCode-SQL三个基准测试中，性能分别从73%、55%、75%提升至89%、64%、79%。进一步引入的群体训练和基于效用的轨迹选择方法，使ALFWorld性能达到91%，媲美需要复杂任务特定组件的方法。结果表明，自动轨迹数据库构建可以替代费时的人工知识工程。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-05T03:20:19Z
- **目录日期**: 2025-05-05
