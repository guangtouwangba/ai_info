# Beyond Words: Multimodal LLM Knows When to Speak

**URL**: http://arxiv.org/abs/2505.14654v1

## 原始摘要

While large language model (LLM)-based chatbots have demonstrated strong
capabilities in generating coherent and contextually relevant responses, they
often struggle with understanding when to speak, particularly in delivering
brief, timely reactions during ongoing conversations. This limitation arises
largely from their reliance on text input, lacking the rich contextual cues in
real-world human dialogue. In this work, we focus on real-time prediction of
response types, with an emphasis on short, reactive utterances that depend on
subtle, multimodal signals across vision, audio, and text. To support this, we
introduce a new multimodal dataset constructed from real-world conversational
videos, containing temporally aligned visual, auditory, and textual streams.
This dataset enables fine-grained modeling of response timing in dyadic
interactions. Building on this dataset, we propose MM-When2Speak, a multimodal
LLM-based model that adaptively integrates visual, auditory, and textual
context to predict when a response should occur, and what type of response is
appropriate. Experiments show that MM-When2Speak significantly outperforms
state-of-the-art unimodal and LLM-based baselines, achieving up to a 4x
improvement in response timing accuracy over leading commercial LLMs. These
results underscore the importance of multimodal inputs for producing timely,
natural, and engaging conversational AI.


## AI 摘要

当前基于大语言模型(LLM)的聊天机器人在生成连贯回复方面表现优异，但在把握对话时机(尤其是简短即时应答)方面存在不足。研究团队提出了MM-When2Speak多模态模型，通过整合视觉、听觉和文本信息来预测最佳应答时机和类型。基于真实对话视频构建的新型多模态数据集支持了这一研究。实验表明，该模型在应答时机预测准确率上比主流商业LLM提升高达4倍，显著优于单模态基线系统。这些发现证明多模态输入对于实现自然流畅的人机对话至关重要。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-21T15:02:30Z
- **目录日期**: 2025-05-21
