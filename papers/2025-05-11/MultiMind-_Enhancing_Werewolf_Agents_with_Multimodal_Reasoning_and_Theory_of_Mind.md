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

多模态社交推理AI框架MultiMind在狼人杀游戏中取得突破。该研究首次将面部表情、语音语调等非语言信息与文本内容结合，并引入心理理论模型追踪玩家间的相互怀疑程度。通过蒙特卡洛树搜索优化沟通策略，MultiMind能有效降低自身嫌疑。实验表明，该框架在人机对抗和人类玩家测试中均表现优异，实现了大型语言模型在多模态社交推理领域的重要进展，使AI更接近人类的多模态社交认知能力。研究以《一夜终极狼人杀》为测试平台，解决了现有文本型AI忽视非语言线索和他人视角建模的局限性。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-11T01:29:51Z
- **目录日期**: 2025-05-11
