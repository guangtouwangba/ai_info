# MEDAL: A Framework for Benchmarking LLMs as Multilingual Open-Domain Chatbots and Dialogue Evaluators

**URL**: http://arxiv.org/abs/2505.22777v2

## 原始摘要

As the capabilities of chatbots and their underlying LLMs continue to
dramatically improve, evaluating their performance has increasingly become a
major blocker to their further development. A major challenge is the available
benchmarking datasets, which are largely static, outdated, and lacking in
multilingual coverage, limiting their ability to capture subtle linguistic and
cultural variations. This paper introduces MEDAL, an automated multi-agent
framework for generating, evaluating, and curating more representative and
diverse open-domain dialogue evaluation benchmarks. Our approach leverages
several state-of-the-art LLMs to generate user-chatbot multilingual dialogues,
conditioned on varied seed contexts. A strong LLM (GPT-4.1) is then used for a
multidimensional analysis of the performance of the chatbots, uncovering
noticeable cross-lingual performance differences. Guided by this large-scale
evaluation, we curate a new meta-evaluation multilingual benchmark and
human-annotate samples with nuanced quality judgments. This benchmark is then
used to assess the ability of several reasoning and non-reasoning LLMs to act
as evaluators of open-domain dialogues. We find that current LLMs struggle to
detect nuanced issues, particularly those involving empathy and reasoning.


## AI 摘要

本文提出MEDAL框架，利用多智能体系统自动生成、评估和优化多语言开放域对话评测基准。该框架采用先进大语言模型生成多语言对话，并通过GPT-4.1进行多维度分析，发现显著的跨语言性能差异。基于大规模评估结果，研究团队构建了一个包含细致人工标注的新元评测基准，用于测试不同大语言模型的对话评估能力。研究发现当前模型在识别共情和推理等细微问题方面仍存在困难。该工作解决了现有评测数据集静态、过时且缺乏多语言覆盖的局限性。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-10T00:01:48Z
- **目录日期**: 2025-06-10
