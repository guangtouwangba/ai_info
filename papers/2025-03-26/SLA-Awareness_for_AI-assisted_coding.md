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

AI辅助编码工具通过代码大语言模型(CodeLLMs)显著提升开发效率，自动化代码生成、补全等重复任务。不同任务对延迟有独特要求：代码补全关注首词延迟(TTFT)，代码翻译则需优化端到端(E2E)延迟。现有方案因缺乏端到端延迟感知，难以兼顾并发任务需求与资源利用率。为此，研究者提出CATO系统，实验显示其能同时提升TTFT关键任务10%吞吐率和41.1%资源利用率，代码摘要E2E延迟降低18%，代码生成TTFT延迟减少14%，优于现有方案。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-26T23:08:45Z
- **目录日期**: 2025-03-26
