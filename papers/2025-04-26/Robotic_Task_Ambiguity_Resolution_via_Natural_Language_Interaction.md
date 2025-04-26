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

研究人员提出了一种名为AmbResVLM的新方法，用于解决语言条件策略中任务描述模糊的问题。该方法通过将语言目标与场景结合并显式推理任务歧义，显著提升了机器人对模糊指令的处理能力。实验表明，AmbResVLM在模拟和真实环境中都优于现有技术，能有效检测和解决任务歧义。实际机器人测试中，该方法将下游策略的平均成功率从69.6%提升至97.1%。相关数据、代码和模型已开源。这一进展有望提高语言控制机器人的可靠性和实用性。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-26T14:00:59Z
- **目录日期**: 2025-04-26
