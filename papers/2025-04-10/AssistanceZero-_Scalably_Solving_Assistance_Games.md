# AssistanceZero: Scalably Solving Assistance Games

**URL**: http://arxiv.org/abs/2504.07091v1

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

本文提出了一种替代人类反馈强化学习(RLHF)的新方法——辅助游戏(assistance games)，通过将人机交互建模为双人博弈来避免RLHF的欺骗行为等问题。研究者开发了首个可扩展的解决方案AssistanceZero，结合AlphaZero算法和预测人类行为的神经网络，能在大规模不确定环境中进行规划。该方法在包含10^400种可能目标的《我的世界》建造任务中表现优异，超越了无模型强化学习和模仿学习。人类实验表明，经AssistanceZero训练的AI助手能显著减少玩家完成建造任务的操作次数，证实了辅助游戏框架在复杂环境中训练AI助手的可行性。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-10T08:01:01Z
- **目录日期**: 2025-04-10
