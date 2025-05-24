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

这篇论文探讨了基于异构大语言模型（LLM）的多智能体系统（X-MAS），通过让不同LLM驱动不同智能体，突破单一模型限制，提升系统整体智能。研究者开发了X-MAS-Bench测试平台，对27个LLM在5个领域（21个测试集）和5种功能上进行了超170万次评估，确定了最优模型组合。实验表明，从同构转向异构LLM驱动的多智能体系统能显著提升性能：在纯聊天机器人场景中，MATH数据集上性能提升8.4%；在混合聊天-推理场景中，AIME数据集上性能提升达47%。这为构建可扩展的协作AI系统提供了新方向。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-24T13:08:14Z
- **目录日期**: 2025-05-24
