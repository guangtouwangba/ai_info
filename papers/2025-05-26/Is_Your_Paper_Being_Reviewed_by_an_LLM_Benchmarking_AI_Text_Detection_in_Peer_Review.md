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

这篇论文探讨了同行评审中AI生成文本的检测问题。研究者创建了一个包含788,984条AI撰写与人工撰写的同行评审数据集，覆盖了ICLR和NeurIPS两个顶级AI会议8年间的投稿论文。通过评估18种现有AI文本检测算法，研究发现目前方法难以可靠识别AI生成的评审内容。论文还提出了一种基于稿件内容的检测方法Anchor，并分析了检测模型对AI修改人工文本的敏感性。结果表明，在个体评审层面检测AI文本极具挑战性，亟需开发新工具来应对生成式AI在学术评审中的不当使用。数据集已公开。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-26T07:02:30Z
- **目录日期**: 2025-05-26
