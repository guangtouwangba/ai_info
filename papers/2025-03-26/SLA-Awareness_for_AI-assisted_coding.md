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

AI辅助编码工具通过代码大语言模型(CodeLLMs)显著提升开发效率，自动处理重复性任务(如代码生成、补全等)。不同任务对延迟有独特要求：代码补全关注首词延迟(TTFT)，代码翻译则重视端到端延迟(E2E)。现有方案因缺乏端到端延迟感知，难以同时满足并发任务的延迟要求和资源优化。为此，研究者提出CATO系统，能协调多样化编码任务，在满足延迟要求的同时最大化资源利用率。实验表明，CATO使TTFT关键任务吞吐量提升10%，资源利用率提高41.1%，代码摘要E2E延迟降低18%，代码生成TTFT延迟减少14%。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-26T08:01:39Z
- **目录日期**: 2025-03-26
