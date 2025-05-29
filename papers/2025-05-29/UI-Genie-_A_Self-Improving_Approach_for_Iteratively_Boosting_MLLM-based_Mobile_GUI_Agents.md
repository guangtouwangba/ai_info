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

本文介绍了UI-Genie框架，解决GUI智能体的两大挑战：轨迹结果验证困难和高质训练数据难以扩展。该框架包含奖励模型UI-Genie-RM（采用图文交错架构处理历史上下文，统一动作级和任务级奖励）和自改进流程（通过奖励引导探索动态优化任务解决能力）。研究构建了首个GUI奖励专用数据集UI-Genie-RM-517k和16k合成轨迹数据，无需人工标注。实验显示，经过三代数据-模型自改进后，UI-Genie在多个基准测试中达到最优性能。作者开源了框架实现和数据集以促进相关研究。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-29T01:29:11Z
- **目录日期**: 2025-05-29
