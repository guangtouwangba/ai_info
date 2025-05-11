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

本文介绍了MultiMind框架，这是首个将多模态信息整合到社交推理游戏（如《一夜终极狼人杀》）中的AI代理。传统方法仅依赖文本信息，而MultiMind结合了面部表情、语音语调等关键多模态线索，并采用心理理论（ToM）模型追踪玩家间的相互怀疑程度。通过结合ToM与蒙特卡洛树搜索（MCTS），该代理能制定最小化自身嫌疑的沟通策略。实验表明，MultiMind在AI对战和人类玩家测试中均表现优异，标志着大语言模型在多模态社交推理领域的重要进展，使其更接近人类社交推理能力。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-11T15:02:19Z
- **目录日期**: 2025-05-11
