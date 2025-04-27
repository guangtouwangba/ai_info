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

本文提出AmbResVLM方法，用于解决机器人语言条件策略中的任务描述模糊性问题。该方法通过将语言目标与观察场景进行关联，并显式推理任务歧义，有效提升了任务模糊性的检测和解决能力。实验表明，在模拟和真实环境中，其性能均优于现有最优基线方法，成功将下游机器人策略的平均成功率从69.6%提升至97.1%。相关数据、代码和训练模型已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-27T16:01:12Z
- **目录日期**: 2025-04-27
