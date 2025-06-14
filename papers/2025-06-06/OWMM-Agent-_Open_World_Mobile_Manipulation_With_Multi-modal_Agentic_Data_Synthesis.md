# OWMM-Agent: Open World Mobile Manipulation With Multi-modal Agentic Data Synthesis

**URL**: http://arxiv.org/abs/2506.04217v1

## 原始摘要

The rapid progress of navigation, manipulation, and vision models has made
mobile manipulators capable in many specialized tasks. However, the open-world
mobile manipulation (OWMM) task remains a challenge due to the need for
generalization to open-ended instructions and environments, as well as the
systematic complexity to integrate high-level decision making with low-level
robot control based on both global scene understanding and current agent state.
To address this complexity, we propose a novel multi-modal agent architecture
that maintains multi-view scene frames and agent states for decision-making and
controls the robot by function calling. A second challenge is the hallucination
from domain shift. To enhance the agent performance, we further introduce an
agentic data synthesis pipeline for the OWMM task to adapt the VLM model to our
task domain with instruction fine-tuning. We highlight our fine-tuned OWMM-VLM
as the first dedicated foundation model for mobile manipulators with global
scene understanding, robot state tracking, and multi-modal action generation in
a unified model. Through experiments, we demonstrate that our model achieves
SOTA performance compared to other foundation models including GPT-4o and
strong zero-shot generalization in real world. The project page is at
https://github.com/HHYHRHY/OWMM-Agent


## AI 摘要

这篇论文提出了一种新型多模态智能体架构（OWMM-VLM），旨在解决开放世界移动操作（OWMM）任务的挑战。该架构通过维护多视角场景帧和智能体状态进行决策，并采用函数调用控制机器人。针对领域转移导致的幻觉问题，研究者开发了数据合成流程对视觉语言模型（VLM）进行指令微调。实验表明，该模型在全局场景理解、机器人状态跟踪和多模态动作生成方面表现优异，性能超越GPT-4o等基础模型，并展现出强大的零样本泛化能力。项目页面已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-06T05:01:06Z
- **目录日期**: 2025-06-06
