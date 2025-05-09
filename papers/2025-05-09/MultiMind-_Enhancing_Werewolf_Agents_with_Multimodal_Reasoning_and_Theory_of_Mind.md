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

该研究提出了首个整合多模态信息的社交推理游戏（SDG）智能体框架MultiMind，以《一夜终极狼人杀》为测试平台。该框架突破现有纯文本方法的局限，通过分析面部表情、语音语调等多模态线索，结合心智理论（ToM）模型追踪玩家间的相互怀疑程度，并利用蒙特卡洛树搜索（MCTS）优化沟通策略以降低自身嫌疑。实验表明，MultiMind在AI对战和人类玩家测试中均表现优异，标志着大语言模型在跨模态人类社交推理能力上的重要突破。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-09T16:02:32Z
- **目录日期**: 2025-05-09
