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

本文提出了一种新型原子事实核查框架，用于提高医疗长文本问答中大型语言模型(LLMs)的可靠性和可解释性。该方法将LLM生成回答分解为可独立验证的"原子事实"，并对照权威医学指南进行验证。实验显示，该方法能显著提升回答准确性(最高40%)和可解释性，幻觉检测率达50%，且每个原子事实都可溯源至原始文献。这一框架解决了当前医疗AI应用中准确性不足和解释性差的关键问题，为LLMs在临床中的可信应用迈出了重要一步，有助于增强AI辅助医疗的可信度。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-02T07:02:29Z
- **目录日期**: 2025-06-02
