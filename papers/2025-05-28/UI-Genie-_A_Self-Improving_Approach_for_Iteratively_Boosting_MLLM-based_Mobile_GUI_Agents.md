# UI-Genie: A Self-Improving Approach for Iteratively Boosting MLLM-based Mobile GUI Agents

**URL**: http://arxiv.org/abs/2505.21496v1

## 原始摘要

In this paper, we introduce UI-Genie, a self-improving framework addressing
two key challenges in GUI agents: verification of trajectory outcome is
challenging and high-quality training data are not scalable. These challenges
are addressed by a reward model and a self-improving pipeline, respectively.
The reward model, UI-Genie-RM, features an image-text interleaved architecture
that efficiently pro- cesses historical context and unifies action-level and
task-level rewards. To sup- port the training of UI-Genie-RM, we develop
deliberately-designed data genera- tion strategies including rule-based
verification, controlled trajectory corruption, and hard negative mining. To
address the second challenge, a self-improvement pipeline progressively expands
solvable complex GUI tasks by enhancing both the agent and reward models
through reward-guided exploration and outcome verification in dynamic
environments. For training the model, we generate UI- Genie-RM-517k and
UI-Genie-Agent-16k, establishing the first reward-specific dataset for GUI
agents while demonstrating high-quality synthetic trajectory gen- eration
without manual annotation. Experimental results show that UI-Genie achieves
state-of-the-art performance across multiple GUI agent benchmarks with three
generations of data-model self-improvement. We open-source our complete
framework implementation and generated datasets to facilitate further research
in https://github.com/Euphoria16/UI-Genie.


## AI 摘要

本文提出了UI-Genie框架，通过奖励模型(UI-Genie-RM)和自改进流程解决GUI智能体的两大挑战：轨迹结果验证困难和高质训练数据难以扩展。UI-Genie-RM采用图文交错架构处理历史上下文，统一动作级和任务级奖励，并配合规则验证、轨迹破坏和负样本挖掘等数据生成策略。自改进流程通过奖励引导探索和动态环境验证，逐步扩展可解决的复杂GUI任务。实验表明，经过三代数据-模型自改进后，UI-Genie在多个GUI基准测试中达到最优性能。相关代码和数据集已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-28T22:01:31Z
- **目录日期**: 2025-05-28
