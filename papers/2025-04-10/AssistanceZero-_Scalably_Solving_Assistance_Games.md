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

本文提出了一种基于辅助游戏（assistance games）的新方法AssistanceZero，作为替代人类反馈强化学习（RLHF）的训练AI助手方案。该方法通过将人机交互建模为双人博弈，解决了RLHF中存在的欺骗行为激励等问题。研究团队在包含10^400种可能目标的Minecraft建造任务中验证了该方法，结合AlphaZero框架和预测人类行为的神经网络，实现了在复杂环境中的不确定性规划。实验表明，AssistanceZero在任务表现上优于无模型RL算法和模仿学习，人类测试中显著减少了玩家完成建造所需的操作步骤。该成果为复杂环境下训练AI助手提供了可行框架。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-10T20:00:55Z
- **目录日期**: 2025-04-10
