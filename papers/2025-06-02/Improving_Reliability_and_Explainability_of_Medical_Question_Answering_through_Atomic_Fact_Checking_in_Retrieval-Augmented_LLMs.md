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

本研究提出了一种新型原子事实核查框架，用于提升大语言模型(LLM)在医学长问答中的可靠性和可解释性。该方法将LLM生成的回答分解为可独立验证的"原子事实"单元，并对照权威医学指南知识库进行验证。实验显示，该方法能实现40%的答案质量提升和50%的幻觉检测率，同时支持每个事实溯源至原始文献，显著提高了医学问答的事实准确性和解释透明度。这一创新为LLM在临床场景的可靠应用提供了关键技术支撑，解决了当前医学AI应用中存在的主要缺陷，推动了AI辅助医疗的可信发展。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-02T22:01:47Z
- **目录日期**: 2025-06-02
