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

该研究探索了如何通过强化学习训练视觉语言模型(VLM)进行图像推理，而无需思维链监督。研究发现，直接在VLM上应用强化学习会导致模型走捷径，影响泛化能力。为此，作者提出"描述-推理-回答"的输出格式：先生成详细图像描述，再构建推理链。在27.3万无思维链的视觉问答对上训练后，模型Visionary-R1仅用强化学习就在多个视觉推理基准上超越了GPT-4o、Claude3.5-Sonnet和Gemini-1.5-Pro等强大模型。研究表明，先解释图像再进行推理是避免捷径学习的关键。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-22T01:28:39Z
- **目录日期**: 2025-05-22
