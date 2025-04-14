# DocAgent: A Multi-Agent System for Automated Code Documentation Generation

**URL**: http://arxiv.org/abs/2504.08725v1

## 原始摘要

High-quality code documentation is crucial for software development
especially in the era of AI. However, generating it automatically using Large
Language Models (LLMs) remains challenging, as existing approaches often
produce incomplete, unhelpful, or factually incorrect outputs. We introduce
DocAgent, a novel multi-agent collaborative system using topological code
processing for incremental context building. Specialized agents (Reader,
Searcher, Writer, Verifier, Orchestrator) then collaboratively generate
documentation. We also propose a multi-faceted evaluation framework assessing
Completeness, Helpfulness, and Truthfulness. Comprehensive experiments show
DocAgent significantly outperforms baselines consistently. Our ablation study
confirms the vital role of the topological processing order. DocAgent offers a
robust approach for reliable code documentation generation in complex and
proprietary repositories.


## AI 摘要

DocAgent是一种新型的多智能体协作系统，通过拓扑代码处理逐步构建上下文，并利用专门角色（阅读器、搜索器、编写器、验证器、协调器）协作生成代码文档。针对现有大型语言模型（LLM）自动生成文档存在不完整、无用或事实错误的问题，DocAgent提出了一个评估框架，从完整性、实用性和真实性三方面进行衡量。实验表明，DocAgent显著优于基线方法，且拓扑处理顺序对其性能至关重要。该系统为复杂或专有代码库提供了可靠的自动化文档生成方案。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-14T11:01:10Z
- **目录日期**: 2025-04-14
