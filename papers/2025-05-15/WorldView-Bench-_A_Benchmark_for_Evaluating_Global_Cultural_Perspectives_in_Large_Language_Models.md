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

当前大语言模型（LLMs）的训练和对齐方式主要强化西方中心主义认知，导致文化同质化，难以反映全球文明多样性。为此，研究者提出WorldView-Bench基准，通过自由生成评估（而非传统分类基准）来量化模型的全球文化包容性（GCI）。基于Senturk等人的"多重世界观"理论，研究采用两种干预策略：（1）上下文实现的多重LLMs；（2）多智能体系统（MAS）实现的多重LLMs。实验显示，MAS方法使视角分布熵值从13%提升至94%，积极情感占比达67.7%，显著提升了文化平衡性。这为构建更具包容性的AI系统提供了新思路。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-15T09:01:44Z
- **目录日期**: 2025-05-15
