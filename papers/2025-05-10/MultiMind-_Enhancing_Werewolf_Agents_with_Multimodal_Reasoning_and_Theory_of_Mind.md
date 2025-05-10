# MultiMind: Enhancing Werewolf Agents with Multimodal Reasoning and Theory of Mind

**URL**: http://arxiv.org/abs/2504.18039v2

## 原始摘要

Large Language Model (LLM) agents have demonstrated impressive capabilities
in social deduction games (SDGs) like Werewolf, where strategic reasoning and
social deception are essential. However, current approaches remain limited to
textual information, ignoring crucial multimodal cues such as facial
expressions and tone of voice that humans naturally use to communicate.
Moreover, existing SDG agents primarily focus on inferring other players'
identities without modeling how others perceive themselves or fellow players.
To address these limitations, we use One Night Ultimate Werewolf (ONUW) as a
testbed and present MultiMind, the first framework integrating multimodal
information into SDG agents. MultiMind processes facial expressions and vocal
tones alongside verbal content, while employing a Theory of Mind (ToM) model to
represent each player's suspicion levels toward others. By combining this ToM
model with Monte Carlo Tree Search (MCTS), our agent identifies communication
strategies that minimize suspicion directed at itself. Through comprehensive
evaluation in both agent-versus-agent simulations and studies with human
players, we demonstrate MultiMind's superior performance in gameplay. Our work
presents a significant advancement toward LLM agents capable of human-like
social reasoning across multimodal domains.


## AI 摘要

本文介绍了MultiMind框架，这是首个将多模态信息整合到社交推理游戏（如《一夜终极狼人杀》）中的AI代理系统。它通过分析面部表情、语音语调和文本内容，结合心理理论（ToM）模型来追踪玩家间的相互怀疑程度，并利用蒙特卡洛树搜索（MCTS）优化沟通策略以降低自身嫌疑。实验证明，MultiMind在AI对战和人类玩家测试中均表现出色。该研究推动了大型语言模型在多模态社交推理方面的发展，使其更接近人类自然交流方式。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-10T10:02:26Z
- **目录日期**: 2025-05-10
