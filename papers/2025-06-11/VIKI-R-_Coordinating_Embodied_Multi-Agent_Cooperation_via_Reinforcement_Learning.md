# VIKI-R: Coordinating Embodied Multi-Agent Cooperation via Reinforcement Learning

**URL**: http://arxiv.org/abs/2506.09049v1

## 原始摘要

Coordinating multiple embodied agents in dynamic environments remains a core
challenge in artificial intelligence, requiring both perception-driven
reasoning and scalable cooperation strategies. While recent works have
leveraged large language models (LLMs) for multi-agent planning, a few have
begun to explore vision-language models (VLMs) for visual reasoning. However,
these VLM-based approaches remain limited in their support for diverse
embodiment types. In this work, we introduce VIKI-Bench, the first hierarchical
benchmark tailored for embodied multi-agent cooperation, featuring three
structured levels: agent activation, task planning, and trajectory perception.
VIKI-Bench includes diverse robot embodiments, multi-view visual observations,
and structured supervision signals to evaluate reasoning grounded in visual
inputs. To demonstrate the utility of VIKI-Bench, we propose VIKI-R, a
two-stage framework that fine-tunes a pretrained vision-language model (VLM)
using Chain-of-Thought annotated demonstrations, followed by reinforcement
learning under multi-level reward signals. Our extensive experiments show that
VIKI-R significantly outperforms baselines method across all task levels.
Furthermore, we show that reinforcement learning enables the emergence of
compositional cooperation patterns among heterogeneous agents. Together,
VIKI-Bench and VIKI-R offer a unified testbed and method for advancing
multi-agent, visual-driven cooperation in embodied AI systems.


## AI 摘要

VIKI-Bench是首个针对具身多智能体协作的分层基准测试，包含智能体激活、任务规划和轨迹感知三个层级，支持多样化机器人形态和多视角视觉输入。研究者提出VIKI-R框架，通过思维链标注微调预训练的视觉语言模型（VLM），再结合多级奖励信号的强化学习。实验表明VIKI-R在所有任务层级均显著优于基线方法，且强化学习能促进异构智能体形成组合式协作模式。该研究为具身AI系统中的视觉驱动多智能体协作提供了统一测试平台和方法。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-11T07:01:15Z
- **目录日期**: 2025-06-11
