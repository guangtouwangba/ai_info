# Relative Overfitting and Accept-Reject Framework

**URL**: http://arxiv.org/abs/2505.07783v1

## 原始摘要

Currently, the scaling law of Large Language Models (LLMs) faces challenges
and bottlenecks. This paper posits that noise effects, stemming from changes in
the signal-to-noise ratio under diminishing marginal returns, are the root
cause of these issues. To control this noise, we investigated the differences
between models with performance advantages and disadvantages, introducing the
concept of "relative overfitting." Based on their complementary strengths, we
have proposed an application framework, Accept-Reject (AR). In Natural Language
Processing (NLP), we use LLMs and Small Language Models (SLMs) as the medium
for discussion. This framework enables SLMs to exert a universal positive
influence on LLM decision outputs, rather than the intuitively expected
negative influence. We validated our approach using self-built models based on
mainstream architectures and pre-trained mainstream models across multiple
datasets, including basic language modeling, long-context tasks, subject
examination, and question-answering (QA) benchmarks. The results demonstrate
that through our structure, compared to increasing the LLM's parameters, we can
achieve better performance improvements with significantly lower parameter and
computational costs in many scenarios. These improvements are universal,
stable, and effective. Furthermore, we explore the potential of "relative
overfitting" and the AR framework in other machine learning domains, such as
computer vision (CV) and AI for science. We hope the proposed approach can help
scale laws overcome existing bottlenecks.


## AI 摘要

当前大语言模型(LLM)的扩展规律面临瓶颈，本文认为边际收益递减下信噪比变化导致的噪声效应是主因。通过分析优劣模型差异，提出"相对过拟合"概念及"接受-拒绝"(AR)应用框架。在NLP领域，该框架使小语言模型(SLM)能普遍正向影响LLM决策，而非预期的负向影响。实验表明，相比增加LLM参数量，该结构能以更低成本获得更优性能提升，且效果稳定普适。研究还探讨了该框架在CV和AI科研等领域的潜力，有望帮助扩展规律突破现有瓶颈。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-14T00:01:26Z
- **目录日期**: 2025-05-14
