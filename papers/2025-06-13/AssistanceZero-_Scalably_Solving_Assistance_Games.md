# AssistanceZero: Scalably Solving Assistance Games

**URL**: http://arxiv.org/abs/2504.07091v2

## 原始摘要

Assistance games are a promising alternative to reinforcement learning from
human feedback (RLHF) for training AI assistants. Assistance games resolve key
drawbacks of RLHF, such as incentives for deceptive behavior, by explicitly
modeling the interaction between assistant and user as a two-player game where
the assistant cannot observe their shared goal. Despite their potential,
assistance games have only been explored in simple settings. Scaling them to
more complex environments is difficult because it requires both solving
intractable decision-making problems under uncertainty and accurately modeling
human users' behavior. We present the first scalable approach to solving
assistance games and apply it to a new, challenging Minecraft-based assistance
game with over $10^{400}$ possible goals. Our approach, AssistanceZero, extends
AlphaZero with a neural network that predicts human actions and rewards,
enabling it to plan under uncertainty. We show that AssistanceZero outperforms
model-free RL algorithms and imitation learning in the Minecraft-based
assistance game. In a human study, our AssistanceZero-trained assistant
significantly reduces the number of actions participants take to complete
building tasks in Minecraft. Our results suggest that assistance games are a
tractable framework for training effective AI assistants in complex
environments. Our code and models are available at
https://github.com/cassidylaidlaw/minecraft-building-assistance-game.


## AI 摘要

本文提出了一种基于辅助游戏（assistance games）的新方法AssistanceZero，用于训练AI助手，以解决传统人类反馈强化学习（RLHF）中的欺骗行为等问题。该方法将人机交互建模为双人博弈，通过扩展AlphaZero并结合神经网络预测人类行为和奖励，在不确定性下进行规划。研究在一个包含超过10^400种可能目标的Minecraft建造任务中验证了该方法的有效性，结果显示AssistanceZero优于无模型强化学习和模仿学习。人类实验表明，该方法能显著减少玩家完成任务所需的操作步骤，证明了辅助游戏在复杂环境中训练AI助手的可行性。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-13T11:01:55Z
- **目录日期**: 2025-06-13
