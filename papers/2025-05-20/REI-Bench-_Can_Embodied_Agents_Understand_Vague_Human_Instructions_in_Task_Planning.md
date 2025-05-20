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

这篇论文研究了基于大语言模型(LLM)的机器人任务规划在面临模糊指代表达(RE)时的挑战。作者创建了首个包含模糊RE的机器人任务规划基准测试REI-Bench，发现模糊RE会导致规划成功率下降高达77.9%。主要问题源于规划过程中遗漏对象。为此，作者提出了一种简单有效的"任务导向情境认知"方法，能为机器人生成清晰指令，在性能上优于现有提示和思维链方法。该研究使机器人任务规划对非专业用户(如老人和儿童)更具实用性，推动了人机交互领域的发展。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-20T09:02:22Z
- **目录日期**: 2025-05-20
