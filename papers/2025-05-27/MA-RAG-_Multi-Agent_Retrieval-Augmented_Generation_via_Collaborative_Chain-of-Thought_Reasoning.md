# MA-RAG: Multi-Agent Retrieval-Augmented Generation via Collaborative Chain-of-Thought Reasoning

**URL**: http://arxiv.org/abs/2505.20096v1

## 原始摘要

We present MA-RAG, a Multi-Agent framework for Retrieval-Augmented Generation
(RAG) that addresses the inherent ambiguities and reasoning challenges in
complex information-seeking tasks. Unlike conventional RAG methods that rely on
either end-to-end fine-tuning or isolated component enhancements, MA-RAG
orchestrates a collaborative set of specialized AI agents: Planner, Step
Definer, Extractor, and QA Agents, to tackle each stage of the RAG pipeline
with task-aware reasoning. Ambiguities may arise from underspecified queries,
sparse or indirect evidence in retrieved documents, or the need to integrate
information scattered across multiple sources. MA-RAG mitigates these
challenges by decomposing the problem into subtasks, such as query
disambiguation, evidence extraction, and answer synthesis, and dispatching them
to dedicated agents equipped with chain-of-thought prompting. These agents
communicate intermediate reasoning and progressively refine the retrieval and
synthesis process. Our design allows fine-grained control over information flow
without any model fine-tuning. Crucially, agents are invoked on demand,
enabling a dynamic and efficient workflow that avoids unnecessary computation.
This modular and reasoning-driven architecture enables MA-RAG to deliver
robust, interpretable results. Experiments on multi-hop and ambiguous QA
benchmarks demonstrate that MA-RAG outperforms state-of-the-art training-free
baselines and rivals fine-tuned systems, validating the effectiveness of
collaborative agent-based reasoning in RAG.


## AI 摘要

MA-RAG是一个多智能体框架，用于改进检索增强生成（RAG）任务中的模糊查询和复杂推理问题。不同于传统RAG方法，MA-RAG通过协同多个专用智能体（规划器、步骤定义器、提取器和问答代理）分解任务，逐步处理查询消歧、证据提取和答案合成。这些代理通过思维链提示进行动态推理，优化检索与合成流程，无需微调模型。实验表明，MA-RAG在复杂问答任务中优于无训练的基线方法，媲美微调系统，验证了多智能体协作在RAG中的有效性。其模块化设计实现了高效、可解释的信息处理。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-27T05:02:15Z
- **目录日期**: 2025-05-27
