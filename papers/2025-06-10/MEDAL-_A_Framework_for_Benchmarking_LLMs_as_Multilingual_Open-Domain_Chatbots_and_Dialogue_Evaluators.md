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

本文介绍了MEDAL框架，一个自动化的多智能体系统，用于生成、评估和筛选更具代表性和多样性的开放域对话评测基准。该框架利用多种先进大语言模型（LLM）生成多语言用户-聊天机器人对话，并通过GPT-4.1进行多维度性能分析，发现显著的跨语言表现差异。基于大规模评估结果，研究团队构建了一个新的多语言元评测基准，并进行了细致的人工标注。测试发现，当前LLM在检测细微问题（如同理心和推理）方面仍有不足。该研究为改进对话系统评估提供了新方法和基准。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-10T02:32:43Z
- **目录日期**: 2025-06-10
