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

DocAgent是一种新型多智能体协作系统，通过拓扑代码处理和增量上下文构建来自动生成高质量代码文档。该系统包含阅读器、搜索器、编写器、验证器和协调器等专门化智能体，协同工作以提高文档的完整性、实用性和准确性。研究表明，DocAgent在综合实验中显著优于基线方法，其拓扑处理顺序对性能至关重要。该框架为复杂和专有代码库提供了可靠的自动化文档生成方案，解决了现有大型语言模型在生成代码文档时常见的不完整、不实用或事实错误等问题。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-15T02:29:09Z
- **目录日期**: 2025-04-15
