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

当前基于大语言模型（LLM）的聊天机器人在生成连贯回复方面表现优异，但难以把握对话中的适时回应时机，尤其是简短即时的反应。这主要源于其依赖纯文本输入，缺乏真实对话中的多模态线索。为此，研究者构建了首个含视觉、听觉和文本对齐的多模态对话数据集，并提出MM-When2Speak模型，通过整合多模态信号预测最佳回应时机与类型。实验显示该模型在回应时机准确率上较单模态和商业LLM提升达4倍，证实多模态输入对实现自然流畅对话AI的关键作用。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-21T16:02:14Z
- **目录日期**: 2025-05-21
