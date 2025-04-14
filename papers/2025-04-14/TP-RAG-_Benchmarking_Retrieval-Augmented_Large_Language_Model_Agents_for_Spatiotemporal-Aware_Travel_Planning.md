# TP-RAG: Benchmarking Retrieval-Augmented Large Language Model Agents for Spatiotemporal-Aware Travel Planning

**URL**: http://arxiv.org/abs/2504.08694v1

## 原始摘要

Large language models (LLMs) have shown promise in automating travel
planning, yet they often fall short in addressing nuanced spatiotemporal
rationality. While existing benchmarks focus on basic plan validity, they
neglect critical aspects such as route efficiency, POI appeal, and real-time
adaptability. This paper introduces TP-RAG, the first benchmark tailored for
retrieval-augmented, spatiotemporal-aware travel planning. Our dataset includes
2,348 real-world travel queries, 85,575 fine-grain annotated POIs, and 18,784
high-quality travel trajectory references sourced from online tourist
documents, enabling dynamic and context-aware planning. Through extensive
experiments, we reveal that integrating reference trajectories significantly
improves spatial efficiency and POI rationality of the travel plan, while
challenges persist in universality and robustness due to conflicting references
and noisy data. To address these issues, we propose EvoRAG, an evolutionary
framework that potently synergizes diverse retrieved trajectories with LLMs'
intrinsic reasoning. EvoRAG achieves state-of-the-art performance, improving
spatiotemporal compliance and reducing commonsense violation compared to
ground-up and retrieval-augmented baselines. Our work underscores the potential
of hybridizing Web knowledge with LLM-driven optimization, paving the way for
more reliable and adaptive travel planning agents.


## AI 摘要

本文提出TP-RAG，首个针对时空感知旅游规划的检索增强基准，包含2,348个真实旅游查询、85,575个细粒度POI标注和18,784条高质量旅游轨迹参考。研究发现参考轨迹能显著提升规划的空间效率和POI合理性，但普遍性和鲁棒性仍受冲突参考和噪声数据影响。为此，作者提出EvoRAG进化框架，有效结合检索轨迹与LLM推理能力，相比基线方法显著提升了时空合规性并减少常识错误。该工作展示了网络知识与LLM优化结合的潜力，为开发更可靠、自适应的旅游规划智能体奠定了基础。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-14T10:01:44Z
- **目录日期**: 2025-04-14
