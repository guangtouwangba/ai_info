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

这篇论文指出，随着大语言模型(LLMs)的快速发展，存在审稿人可能依赖AI撰写同行评审的风险。研究者创建了一个包含788,984条AI生成与人工撰写评审对比的数据集，涵盖ICLR和NeurIPS两大AI会议8年的投稿论文。通过评估18种现有AI文本检测算法，发现难以准确识别AI生成的评审内容。研究还探索了基于稿件内容的"Anchor"检测方法，并分析了模型对人类文本经AI编辑后的敏感性。结果表明，当前技术难以可靠检测单个评审是否由AI生成，亟需开发新工具应对这一伦理挑战。数据集已公开。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-26T13:10:14Z
- **目录日期**: 2025-05-26
