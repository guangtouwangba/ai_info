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

AmbResVLM是一种新型方法，旨在解决语言条件策略中的任务描述模糊性问题。该方法通过将语言目标与观察场景结合，并显式推理任务歧义，显著提升了机器人对模糊指令的处理能力。实验表明，AmbResVLM在模拟和真实环境中均优于现有基线方法，能有效检测和解决任务歧义。实际机器人测试中，该方法将下游策略的平均成功率从69.6%提升至97.1%。相关数据、代码和训练模型已开源。该研究为提升语言条件策略的可靠性提供了重要突破。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-26T01:28:19Z
- **目录日期**: 2025-04-26
