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

DocAgent是一个创新的多智能体协作系统，通过拓扑代码处理和增量上下文构建自动生成高质量的代码文档。它包含五个专门智能体（阅读器、搜索器、编写器、验证器和协调器）协同工作，解决了现有大型语言模型生成文档不完整、无用或不准确的问题。研究提出了评估框架，从完整性、实用性和真实性三个维度验证效果。实验表明DocAgent显著优于基线方法，消融研究证实了拓扑处理顺序的关键作用。该系统为复杂和专有代码库提供了可靠的文档生成解决方案。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-15T01:28:32Z
- **目录日期**: 2025-04-15
