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

这篇论文提出了一种基于多智能体框架的电影音频生成方法，受专业电影制作启发。系统由自主的"声音导演"智能体监督，通过多模态大语言模型与其他智能体协作，分别处理"画内音"（需时间对齐）和"画外音"（环境声与背景音乐）。针对画内音，系统检测视频中说话人物后，训练预测模型生成随时间变化的音频控制信号（响度、音高、音色），再由"拟音师"智能体生成同步声音。"拟音师"与"作曲家"、"配音演员"智能体协作生成画外音。通过将文本/视频条件分解为原子化指令，该方法能生成与电影片段匹配的丰富音频内容。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-03T13:12:10Z
- **目录日期**: 2025-06-03
