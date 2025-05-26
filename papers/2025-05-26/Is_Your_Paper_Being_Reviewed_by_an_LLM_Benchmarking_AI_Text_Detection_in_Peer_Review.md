# Is Your Paper Being Reviewed by an LLM? Benchmarking AI Text Detection in Peer Review

**URL**: http://arxiv.org/abs/2502.19614v2

## 原始摘要

Peer review is a critical process for ensuring the integrity of published
scientific research. Confidence in this process is predicated on the assumption
that experts in the relevant domain give careful consideration to the merits of
manuscripts which are submitted for publication. With the recent rapid
advancements in large language models (LLMs), a new risk to the peer review
process is that negligent reviewers will rely on LLMs to perform the often time
consuming process of reviewing a paper. However, there is a lack of existing
resources for benchmarking the detectability of AI text in the domain of peer
review. To address this deficiency, we introduce a comprehensive dataset
containing a total of 788,984 AI-written peer reviews paired with corresponding
human reviews, covering 8 years of papers submitted to each of two leading AI
research conferences (ICLR and NeurIPS). We use this new resource to evaluate
the ability of 18 existing AI text detection algorithms to distinguish between
peer reviews fully written by humans and different state-of-the-art LLMs.
Additionally, we explore a context-aware detection method called Anchor, which
leverages manuscript content to detect AI-generated reviews, and analyze the
sensitivity of detection models to LLM-assisted editing of human-written text.
Our work reveals the difficulty of identifying AI-generated text at the
individual peer review level, highlighting the urgent need for new tools and
methods to detect this unethical use of generative AI. Our dataset is publicly
available at:
https://huggingface.co/datasets/IntelLabs/AI-Peer-Review-Detection-Benchmark.


## AI 摘要

该研究针对同行评审中可能滥用大语言模型(LLMs)的风险，构建了一个包含788,984条AI生成与人工撰写评审对照的数据集，涵盖ICLR和NeurIPS两大AI会议8年间的投稿论文。研究评估了18种现有AI文本检测算法，发现现有方法难以准确识别AI生成的单篇评审，并提出了利用论文内容的"Anchor"上下文感知检测方法。研究揭示了当前检测AI生成评审文本的困难，强调亟需开发新工具应对这一学术伦理问题。相关数据集已公开供研究使用。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-26T02:32:13Z
- **目录日期**: 2025-05-26
