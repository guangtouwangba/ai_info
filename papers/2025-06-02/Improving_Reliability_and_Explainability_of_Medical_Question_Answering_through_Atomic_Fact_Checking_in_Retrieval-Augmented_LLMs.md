# Improving Reliability and Explainability of Medical Question Answering through Atomic Fact Checking in Retrieval-Augmented LLMs

**URL**: http://arxiv.org/abs/2505.24830v1

## 原始摘要

Large language models (LLMs) exhibit extensive medical knowledge but are
prone to hallucinations and inaccurate citations, which pose a challenge to
their clinical adoption and regulatory compliance. Current methods, such as
Retrieval Augmented Generation, partially address these issues by grounding
answers in source documents, but hallucinations and low fact-level
explainability persist. In this work, we introduce a novel atomic fact-checking
framework designed to enhance the reliability and explainability of LLMs used
in medical long-form question answering. This method decomposes LLM-generated
responses into discrete, verifiable units called atomic facts, each of which is
independently verified against an authoritative knowledge base of medical
guidelines. This approach enables targeted correction of errors and direct
tracing to source literature, thereby improving the factual accuracy and
explainability of medical Q&amp;A. Extensive evaluation using multi-reader
assessments by medical experts and an automated open Q&amp;A benchmark demonstrated
significant improvements in factual accuracy and explainability. Our framework
achieved up to a 40% overall answer improvement and a 50% hallucination
detection rate. The ability to trace each atomic fact back to the most relevant
chunks from the database provides a granular, transparent explanation of the
generated responses, addressing a major gap in current medical AI applications.
This work represents a crucial step towards more trustworthy and reliable
clinical applications of LLMs, addressing key prerequisites for clinical
application and fostering greater confidence in AI-assisted healthcare.


## AI 摘要

大型语言模型(LLMs)在医学领域展现出丰富知识，但存在幻觉和错误引用问题。研究者提出新型原子事实核查框架，将LLM生成的医学问答分解为可独立验证的"原子事实"，对照权威医学指南进行验证。该方法实现了40%的答案改进和50%的幻觉检测率，能追踪每个事实到原始文献，显著提升回答的准确性和可解释性。多专家评估和开放问答基准测试证实，该框架为医学AI应用提供了更透明、可靠的解决方案，是推动LLM临床可信应用的重要进展。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-02T18:01:55Z
- **目录日期**: 2025-06-02
