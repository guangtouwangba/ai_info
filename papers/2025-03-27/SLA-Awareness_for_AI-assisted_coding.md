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

AI辅助编码工具（如CodeLLMs）能显著减少开发时间，让开发者更专注于创造性工作。这些工具可自动完成代码生成、补全、摘要和翻译等任务，但对实时响应要求严格：代码补全需低首字延迟（TTFT），翻译任务需低端到端延迟（E2E）。现有服务模式无法兼顾多任务并发和资源利用率。为此，研究者提出CATO系统，通过任务编排优化延迟和资源使用。实验显示，CATO使TTFT关键任务吞吐量提升10%，资源利用率提高41.1%，代码摘要E2E延迟降低18%，代码生成TTFT延迟减少14%，优于现有方案。（99字）

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-27T00:01:55Z
- **目录日期**: 2025-03-27
