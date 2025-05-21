# Visionary-R1: Mitigating Shortcuts in Visual Reasoning with Reinforcement Learning

**URL**: http://arxiv.org/abs/2505.14677v1

## 原始摘要

Learning general-purpose reasoning capabilities has long been a challenging
problem in AI. Recent research in large language models (LLMs), such as
DeepSeek-R1, has shown that reinforcement learning techniques like GRPO can
enable pre-trained LLMs to develop reasoning capabilities using simple
question-answer pairs. In this paper, we aim to train visual language models
(VLMs) to perform reasoning on image data through reinforcement learning and
visual question-answer pairs, without any explicit chain-of-thought (CoT)
supervision. Our findings indicate that simply applying reinforcement learning
to a VLM -- by prompting the model to produce a reasoning chain before
providing an answer -- can lead the model to develop shortcuts from easy
questions, thereby reducing its ability to generalize across unseen data
distributions. We argue that the key to mitigating shortcut learning is to
encourage the model to interpret images prior to reasoning. Therefore, we train
the model to adhere to a caption-reason-answer output format: initially
generating a detailed caption for an image, followed by constructing an
extensive reasoning chain. When trained on 273K CoT-free visual question-answer
pairs and using only reinforcement learning, our model, named Visionary-R1,
outperforms strong multimodal models, such as GPT-4o, Claude3.5-Sonnet, and
Gemini-1.5-Pro, on multiple visual reasoning benchmarks.


## AI 摘要

近期研究表明，通过强化学习（如GRPO）训练大型语言模型（LLM）可使其基于简单问答对发展推理能力。本研究将该方法扩展至视觉语言模型（VLM），提出无需思维链（CoT）标注、仅通过视觉问答对和强化学习训练VLM进行图像推理。发现直接应用强化学习会导致模型从简单问题中学习捷径，降低泛化能力。为此提出"描述-推理-回答"输出格式：先生成详细图像描述，再构建推理链。在27.3万无CoT标注数据上训练的Visionary-R1模型，在多项视觉推理基准上超越GPT-4o等主流多模态模型。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-21T19:01:02Z
- **目录日期**: 2025-05-21
