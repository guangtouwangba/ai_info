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

机器人任务规划将人类指令分解为可执行动作序列。然而，现实指令常存在模糊指代表达(REs)，严重影响基于大语言模型(LLM)的规划性能，成功率可下降77.9%。研究团队创建首个含模糊REs的基准测试REI-Bench，发现多数失败源于对象缺失。为此提出任务导向情境认知方法，能生成清晰指令，性能优于现有技术。该研究提升了人机交互实用性，尤其有助于服务老人和儿童等非专业用户。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-21T02:31:26Z
- **目录日期**: 2025-05-21
