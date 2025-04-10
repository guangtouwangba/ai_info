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

该研究提出了一种替代人类反馈强化学习（RLHF）的新方法——辅助游戏（assistance games），通过将AI助手与用户的互动建模为双人博弈（助手无法直接观察共同目标），解决了RLHF中欺骗性行为等关键缺陷。研究者开发了首个可扩展的解决方案AssistanceZero（基于AlphaZero改进），结合神经网络预测人类行为和奖励，在包含10^400种可能目标的《我的世界》建造任务中实现了高效规划。实验表明该方法优于无模型强化学习和模仿学习，人类测试中显著减少了用户操作次数，验证了辅助游戏在复杂环境中训练AI助手的可行性。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-10T06:01:16Z
- **目录日期**: 2025-04-10
