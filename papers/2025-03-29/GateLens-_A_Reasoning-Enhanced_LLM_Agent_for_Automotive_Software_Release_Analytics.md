# GateLens: A Reasoning-Enhanced LLM Agent for Automotive Software Release Analytics

**URL**: http://arxiv.org/abs/2503.21735v1

## 原始摘要

Ensuring the reliability and effectiveness of software release decisions is
critical, particularly in safety-critical domains like automotive systems.
Precise analysis of release validation data, often presented in tabular form,
plays a pivotal role in this process. However, traditional methods that rely on
manual analysis of extensive test datasets and validation metrics are prone to
delays and high costs. Large Language Models (LLMs) offer a promising
alternative but face challenges in analytical reasoning, contextual
understanding, handling out-of-scope queries, and processing structured test
data consistently; limitations that hinder their direct application in
safety-critical scenarios. This paper introduces GateLens, an LLM-based tool
for analyzing tabular data in the automotive domain. GateLens translates
natural language queries into Relational Algebra (RA) expressions and then
generates optimized Python code. It outperforms the baseline system on
benchmarking datasets, achieving higher F1 scores and handling complex and
ambiguous queries with greater robustness. Ablation studies confirm the
critical role of the RA module, with performance dropping sharply when omitted.
Industrial evaluations reveal that GateLens reduces analysis time by over 80%
while maintaining high accuracy and reliability. As demonstrated by presented
results, GateLens achieved high performance without relying on few-shot
examples, showcasing strong generalization across various query types from
diverse company roles. Insights from deploying GateLens with a partner
automotive company offer practical guidance for integrating AI into critical
workflows such as release validation. Results show that by automating test
result analysis, GateLens enables faster, more informed, and dependable release
decisions, and can thus advance software scalability and reliability in
automotive systems.


## AI 摘要

本文介绍了GateLens，一种基于大语言模型（LLM）的工具，用于分析汽车领域的表格数据。GateLens将自然语言查询转换为关系代数表达式，并生成优化的Python代码。它在基准数据集上表现优于基线系统，F1分数更高，处理复杂和模糊查询时更稳健。消融研究表明关系代数模块至关重要。工业评估显示，GateLens减少80%以上的分析时间，同时保持高准确性和可靠性。该工具无需依赖少量示例即可实现高性能，在不同公司角色的查询类型中展现出强泛化能力，有助于提升汽车系统软件的发布决策效率和可靠性。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-29T23:01:22Z
- **目录日期**: 2025-03-29
