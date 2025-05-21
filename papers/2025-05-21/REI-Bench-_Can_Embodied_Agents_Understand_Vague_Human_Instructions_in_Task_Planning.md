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

机器人任务规划将人类指令分解为可执行动作序列，但现实中非专家用户（如老人和儿童）的指令常存在模糊指代表达（REs），严重影响基于大语言模型（LLM）的任务规划性能。研究发现REs模糊性可使成功率骤降77.9%，主要源于规划器遗漏对象。为此，研究团队提出首个含模糊REs的机器人任务规划基准REI-Bench，并开发了任务导向的上下文认知方法，通过生成清晰指令显著提升性能。该成果推动了人机交互研究，尤其提升了机器人服务非专业用户的实用性。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-21T01:29:46Z
- **目录日期**: 2025-05-21
