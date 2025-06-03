# ReelWave: Multi-Agentic Movie Sound Generation through Multimodal LLM Conversation

**URL**: http://arxiv.org/abs/2503.07217v3

## 原始摘要

Current audio generation conditioned by text or video focuses on aligning
audio with text/video modalities. Despite excellent alignment results, these
multimodal frameworks still cannot be directly applied to compelling movie
storytelling involving multiple scenes, where "on-screen" sounds require
temporally-aligned audio generation, while "off-screen" sounds contribute to
appropriate environment sounds accompanied by background music when applicable.
Inspired by professional movie production, this paper proposes a multi-agentic
framework for audio generation supervised by an autonomous Sound Director
agent, engaging multi-turn conversations with other agents for on-screen and
off-screen sound generation through multimodal LLM. To address on-screen sound
generation, after detecting any talking humans in videos, we capture
semantically and temporally synchronized sound by training a prediction model
that forecasts interpretable, time-varying audio control signals: loudness,
pitch, and timbre, which are used by a Foley Artist agent to condition a
cross-attention module in the sound generation. The Foley Artist works
cooperatively with the Composer and Voice Actor agents, and together they
autonomously generate off-screen sound to complement the overall production.
Each agent takes on specific roles similar to those of a movie production team.
To temporally ground audio language models, in ReelWave, text/video conditions
are decomposed into atomic, specific sound generation instructions synchronized
with visuals when applicable. Consequently, our framework can generate rich and
relevant audio content conditioned on video clips extracted from movies.


## AI 摘要

这篇论文提出了一种多代理框架ReelWave，用于电影音频生成。该框架受专业电影制作启发，通过一个自主的"声音导演"代理协调多个代理（拟音师、作曲家和配音演员）进行多轮对话，分别生成"画内音"（与视频严格同步）和"画外音"（环境声和背景音乐）。对于画内音，系统先检测视频中说话人物，训练预测模型生成随时间变化的音频控制信号（响度、音高和音色），再由拟音师代理生成同步声音。通过将文本/视频条件分解为原子级声音生成指令，该框架能为电影片段生成丰富相关的音频内容。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-03T12:02:01Z
- **目录日期**: 2025-06-03
