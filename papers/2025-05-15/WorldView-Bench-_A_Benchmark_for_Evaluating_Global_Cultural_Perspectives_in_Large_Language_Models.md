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

当前大语言模型(LLMs)主要基于西方中心主义进行训练和对齐，导致文化同质化问题。为评估全球文化包容性(GCI)，研究者开发了WorldView-Bench基准测试，采用自由生成式评估(而非传统分类测试)来测量文化极化现象。通过两种干预策略：(1)上下文实现的多路复用LLMs；(2)多智能体系统(MAS)实现的多路复用LLMs(多个代表不同文化视角的LLM智能体协作生成回答)。实验显示，MAS方法将视角分布得分(PDS)熵从基线13%提升至94%，情感极性转为67.7%正向，显著提升了文化平衡性。这为开发更具文化包容性的AI系统提供了新思路。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-15T02:30:10Z
- **目录日期**: 2025-05-15
