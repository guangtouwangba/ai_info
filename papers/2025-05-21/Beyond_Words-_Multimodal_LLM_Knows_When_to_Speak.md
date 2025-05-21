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

当前大型语言模型(LLM)聊天机器人在生成连贯回复方面表现优异，但在把握对话时机(特别是简短即时反应)上存在局限。为解决这一问题，研究团队构建了包含视觉、听觉和文本的多模态对话数据集，并提出了MM-When2Speak模型。该模型能整合多模态信号，自适应预测最佳回应时机和类型。实验表明，该模型在回应时机准确性上比主流商业LLM提升4倍，显著优于单模态基线。研究证实多模态输入对实现自然、即时的对话AI至关重要。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-21T09:02:27Z
- **目录日期**: 2025-05-21
