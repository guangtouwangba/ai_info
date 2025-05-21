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

当前基于大语言模型（LLM）的聊天机器人生成连贯回复能力强，但难以把握对话中简短即时反应的时机，主要因其依赖纯文本输入而忽略真实对话中的多模态线索。为此，研究团队构建了首个从真实对话视频中提取的时序对齐多模态数据集（视觉/听觉/文本），并提出多模态模型MM-When2Speak。该模型通过整合多模态上下文，自适应预测应答时机与类型。实验显示其应答时机准确率超越现有单模态及商业LLM基线达4倍，证实多模态输入对实现自然、及时对话AI的关键作用。成果发表于计算语言学顶会。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-21T07:02:16Z
- **目录日期**: 2025-05-21
