# SLA-Awareness for AI-assisted coding

**URL**: http://arxiv.org/abs/2503.19876v1

## 原始摘要

The integration of AI-assisted coding tools within development environments
drastically reduces development time, and allows developers to focus more on
creative and critical aspects of software engineering through the use of Code
Large Language Models (CodeLLMs). These coding assistants automate repetitive
and time-consuming coding tasks such as code generation, code completion, code
summarization, and code translation. Responsiveness is a crucial requirement of
these coding assistants to maintain real-time interactivity, such that their
use does not impede the developers' workflows. Different coding tasks have
unique characteristics and latency requirements: Time-To-First-Token (TTFT)
latency is essential for code completion tasks, while End-To-End (E2E) latency
is crucial for code translation tasks. Managing these varying requirements
simultaneously while optimizing resource usage poses significant challenges.
Existing work adopts the Model-as-a-Service paradigm for serving individual
CodeLLMs, but cannot effectively manage latency requirements of concurrent
coding tasks and sequences of CodeLLM inference calls, due to a lack of
end-to-end latency awareness. Another challenge is keeping resource utilization
high, when the serving system is deployed on a shared cluster environment. To
address these challenges, we propose Coding Assistant Task Orchestrator (CATO),
a runtime system designed to serve a diverse assortment of coding tasks while
meeting latency requirements and maximizing resource utilization. Our
experiments demonstrate that when all types of coding tasks were served
simultaneously, for TTFT-critical tasks, CATO improves overall Goodput rate and
resource utilization by up to 10% and 41.1%, respectively. P95 E2E latency was
also reduced by 18% for code summarization tasks, and P95 TTFT for code
generation tasks were reduced by 14% compared against state-of-the-art systems.


## AI 摘要

AI辅助编码工具（如CodeLLMs）能显著减少开发时间，让开发者更专注于创造性工作。然而，不同编码任务（如代码补全和翻译）对延迟要求各异，现有系统难以同时满足实时性和资源利用率。为此，研究者提出CATO系统，通过端到端延迟感知和任务编排，优化并发任务处理。实验显示，CATO在TTFT关键任务中提升吞吐率10%、资源利用率41.1%，代码摘要的P95端到端延迟降低18%，代码生成的P95首令牌延迟减少14%，优于现有方案。该系统有效平衡了延迟需求与资源效率。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-26T20:01:47Z
- **目录日期**: 2025-03-26
