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

MA-RAG是一个多智能体框架，用于增强检索生成（RAG）任务，解决复杂信息查询中的模糊性和推理挑战。不同于传统RAG方法，MA-RAG通过协作的专用智能体（规划器、步骤定义器、提取器和问答代理）分解任务，逐步处理查询消歧、证据提取和答案合成。代理通过思维链提示进行推理，动态优化检索和生成流程，无需微调模型即可实现细粒度控制。实验表明，MA-RAG在模糊和多跳问答任务中优于无训练基线，媲美微调系统，验证了多代理协作推理的有效性和可解释性。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-27T22:02:15Z
- **目录日期**: 2025-05-27
