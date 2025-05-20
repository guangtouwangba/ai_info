# REI-Bench: Can Embodied Agents Understand Vague Human Instructions in Task Planning?

**URL**: http://arxiv.org/abs/2505.10872v2

## 原始摘要

Robot task planning decomposes human instructions into executable action
sequences that enable robots to complete a series of complex tasks. Although
recent large language model (LLM)-based task planners achieve amazing
performance, they assume that human instructions are clear and straightforward.
However, real-world users are not experts, and their instructions to robots
often contain significant vagueness. Linguists suggest that such vagueness
frequently arises from referring expressions (REs), whose meanings depend
heavily on dialogue context and environment. This vagueness is even more
prevalent among the elderly and children, who robots should serve more. This
paper studies how such vagueness in REs within human instructions affects
LLM-based robot task planning and how to overcome this issue. To this end, we
propose the first robot task planning benchmark with vague REs (REI-Bench),
where we discover that the vagueness of REs can severely degrade robot planning
performance, leading to success rate drops of up to 77.9%. We also observe that
most failure cases stem from missing objects in planners. To mitigate the REs
issue, we propose a simple yet effective approach: task-oriented context
cognition, which generates clear instructions for robots, achieving
state-of-the-art performance compared to aware prompt and chains of thought.
This work contributes to the research community of human-robot interaction
(HRI) by making robot task planning more practical, particularly for non-expert
users, e.g., the elderly and children.


## AI 摘要

这篇论文研究了基于大语言模型(LLM)的机器人任务规划中模糊指代表达(REs)的影响。作者发现REs导致的模糊性会使规划成功率下降高达77.9%，主要问题是对象缺失。为此，他们创建了首个含模糊REs的机器人规划基准REI-Bench，并提出任务导向的上下文认知方法，能生成清晰指令，显著提升性能。该研究使机器人任务规划对非专业用户(如老人和儿童)更具实用性，推动了人机交互领域发展。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-20T16:02:09Z
- **目录日期**: 2025-05-20
