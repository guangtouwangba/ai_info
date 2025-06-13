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

该研究提出"辅助游戏"作为人类反馈强化学习(RLHF)的替代方案，通过将AI助手与用户的互动建模为双人博弈来解决RLHF的缺陷(如欺骗行为)。研究团队开发了首个可扩展的解决方案AssistanceZero，基于AlphaZero架构并整合神经网络来预测人类行为和奖励。在包含10^400种可能目标的《我的世界》建造任务测试中，该方法优于无模型RL算法和模仿学习。人类实验表明，经AssistanceZero训练的助手能显著减少用户完成建造任务的操作次数，证明辅助游戏是训练复杂环境下AI助手的可行框架。代码已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-13T02:32:29Z
- **目录日期**: 2025-06-13
