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

这篇论文提出了一种名为"协助游戏"(assistance games)的新方法替代传统的人类反馈强化学习(RLHF)来训练AI助手。研究团队开发了AssistanceZero算法，通过扩展AlphaZero架构并整合预测人类行为和奖励的神经网络，成功解决了复杂环境下的决策不确定性难题。该方法在包含超过10^400种可能目标的《我的世界》建造任务中表现优异，超越了无模型强化学习和模仿学习。人类参与实验表明，经AssistanceZero训练的AI助手能显著减少用户完成任务所需的操作步骤。这证实了协助游戏框架在复杂环境中训练高效AI助手的可行性。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-13T05:02:28Z
- **目录日期**: 2025-06-13
