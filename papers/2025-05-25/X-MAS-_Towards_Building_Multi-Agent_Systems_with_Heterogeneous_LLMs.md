# X-MAS: Towards Building Multi-Agent Systems with Heterogeneous LLMs

**URL**: http://arxiv.org/abs/2505.16997v1

## 原始摘要

LLM-based multi-agent systems (MAS) extend the capabilities of single LLMs by
enabling cooperation among multiple specialized agents. However, most existing
MAS frameworks rely on a single LLM to drive all agents, constraining the
system's intelligence to the limit of that model. This paper explores the
paradigm of heterogeneous LLM-driven MAS (X-MAS), where agents are powered by
diverse LLMs, elevating the system's potential to the collective intelligence
of diverse LLMs. We introduce X-MAS-Bench, a comprehensive testbed designed to
evaluate the performance of various LLMs across different domains and
MAS-related functions. As an extensive empirical study, we assess 27 LLMs
across 5 domains (encompassing 21 test sets) and 5 functions, conducting over
1.7 million evaluations to identify optimal model selections for each
domain-function combination. Building on these findings, we demonstrate that
transitioning from homogeneous to heterogeneous LLM-driven MAS can
significantly enhance system performance without requiring structural redesign.
Specifically, in a chatbot-only MAS scenario, the heterogeneous configuration
yields up to 8.4\% performance improvement on the MATH dataset. In a mixed
chatbot-reasoner scenario, the heterogeneous MAS could achieve a remarkable
47\% performance boost on the AIME dataset. Our results underscore the
transformative potential of heterogeneous LLMs in MAS, highlighting a promising
avenue for advancing scalable, collaborative AI systems.


## AI 摘要

本文提出异构LLM驱动的多智能体系统(X-MAS)，突破传统单LLM框架限制，通过整合不同LLM实现集体智能提升。研究团队开发X-MAS-Bench测试平台，对27个LLM在5大领域21个测试集进行170万次评估，发现异构配置能显著提升系统性能：在纯聊天机器人场景中，MATH数据集性能提升8.4%；在混合推理场景中，AIME数据集性能提升达47%。研究表明，无需重构系统架构，仅通过采用异构LLM组合即可有效增强多智能体系统的整体表现，为可扩展协作AI系统发展提供了新方向。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-25T23:01:26Z
- **目录日期**: 2025-05-25
