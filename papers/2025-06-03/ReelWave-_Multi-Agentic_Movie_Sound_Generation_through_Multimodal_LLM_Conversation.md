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

这篇论文提出了一种受电影制作启发的多智能体框架，用于生成与视频内容同步的音频。该框架由自主的声音导演智能体监督，通过多模态大语言模型与其他智能体协作，分别处理"画内音"（需时间对齐）和"画外音"（环境声和背景音乐）。针对画内音，系统检测视频中说话人物后，训练预测模型生成随时间变化的音频控制信号（响度、音高、音色），指导拟音师生成同步声音。拟音师与作曲家和配音演员智能体合作，自动生成画外音。该框架能将文本/视频条件分解为原子级声音生成指令，实现与视觉内容同步的丰富音频生成。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-03T10:02:03Z
- **目录日期**: 2025-06-03
