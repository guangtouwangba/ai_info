# Robotic Task Ambiguity Resolution via Natural Language Interaction

**URL**: http://arxiv.org/abs/2504.17748v1

## 原始摘要

Language-conditioned policies have recently gained substantial adoption in
robotics as they allow users to specify tasks using natural language, making
them highly versatile. While much research has focused on improving the action
prediction of language-conditioned policies, reasoning about task descriptions
has been largely overlooked. Ambiguous task descriptions often lead to
downstream policy failures due to misinterpretation by the robotic agent. To
address this challenge, we introduce AmbResVLM, a novel method that grounds
language goals in the observed scene and explicitly reasons about task
ambiguity. We extensively evaluate its effectiveness in both simulated and
real-world domains, demonstrating superior task ambiguity detection and
resolution compared to recent state-of-the-art baselines. Finally, real robot
experiments show that our model improves the performance of downstream robot
policies, increasing the average success rate from 69.6% to 97.1%. We make the
data, code, and trained models publicly available at
https://ambres.cs.uni-freiburg.de.


## AI 摘要

AmbResVLM是一种新型方法，用于解决语言条件策略中的任务描述模糊性问题。该方法通过将语言目标与观察场景关联，并显式推理任务歧义，显著提升了机器人对模糊指令的处理能力。在模拟和真实场景的广泛测试中，AmbResVLM在任务歧义检测和解决方面均优于现有最优基线。实际机器人实验表明，该方法将下游策略的平均成功率从69.6%提升至97.1%。相关数据、代码和训练模型已公开。这项研究填补了语言条件策略中任务描述推理的空白，有效减少了因指令误解导致的策略失败。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-25T09:01:13Z
- **目录日期**: 2025-04-25
