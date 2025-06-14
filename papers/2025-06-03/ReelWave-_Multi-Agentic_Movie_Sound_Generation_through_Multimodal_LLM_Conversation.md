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

这篇论文提出了一种多智能体框架ReelWave，用于生成与电影场景匹配的音频内容。该框架由Sound Director智能体监督，通过多模态LLM与其他智能体协作，分别处理"画内音"和"画外音"生成。对于画内音，系统检测视频中说话人物后，训练预测模型生成随时间变化的音频控制信号(响度、音高、音色)，指导Foley Artist智能体生成同步音频。画外音则由Foley Artist、Composer和Voice Actor智能体协同生成。该框架将文本/视频条件分解为原子级音频生成指令，能根据电影片段生成丰富且相关的音频内容。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-03T20:02:08Z
- **目录日期**: 2025-06-03
