# WorldView-Bench: A Benchmark for Evaluating Global Cultural Perspectives in Large Language Models

**URL**: http://arxiv.org/abs/2505.09595v1

## 原始摘要

Large Language Models (LLMs) are predominantly trained and aligned in ways
that reinforce Western-centric epistemologies and socio-cultural norms, leading
to cultural homogenization and limiting their ability to reflect global
civilizational plurality. Existing benchmarking frameworks fail to adequately
capture this bias, as they rely on rigid, closed-form assessments that overlook
the complexity of cultural inclusivity. To address this, we introduce
WorldView-Bench, a benchmark designed to evaluate Global Cultural Inclusivity
(GCI) in LLMs by analyzing their ability to accommodate diverse worldviews. Our
approach is grounded in the Multiplex Worldview proposed by Senturk et al.,
which distinguishes between Uniplex models, reinforcing cultural
homogenization, and Multiplex models, which integrate diverse perspectives.
WorldView-Bench measures Cultural Polarization, the exclusion of alternative
perspectives, through free-form generative evaluation rather than conventional
categorical benchmarks. We implement applied multiplexity through two
intervention strategies: (1) Contextually-Implemented Multiplex LLMs, where
system prompts embed multiplexity principles, and (2) Multi-Agent System
(MAS)-Implemented Multiplex LLMs, where multiple LLM agents representing
distinct cultural perspectives collaboratively generate responses. Our results
demonstrate a significant increase in Perspectives Distribution Score (PDS)
entropy from 13% at baseline to 94% with MAS-Implemented Multiplex LLMs,
alongside a shift toward positive sentiment (67.7%) and enhanced cultural
balance. These findings highlight the potential of multiplex-aware AI
evaluation in mitigating cultural bias in LLMs, paving the way for more
inclusive and ethically aligned AI systems.


## AI 摘要

当前大语言模型（LLMs）的训练和校准主要基于西方中心主义认知体系，导致文化同质化并削弱全球文明多样性的体现。为评估这一问题，研究者提出WorldView-Bench基准，通过自由生成式评估（而非传统分类测试）来量化模型的文化包容性（GCI）。基于Senturk的"多元世界观"理论，研究采用两种干预策略：1）上下文实现的多元LLMs（系统提示嵌入多元原则）；2）多智能体系统（MAS）实现的多元LLMs（不同文化视角的智能体协作生成回答）。实验显示，MAS方法使观点分布熵值从13%提升至94%，积极情感占比达67.7%，显著提升了文化平衡性。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-16T01:29:06Z
- **目录日期**: 2025-05-16
