# SWE-PolyBench: A multi-language benchmark for repository level evaluation of coding agents

**URL**: http://arxiv.org/abs/2504.08703v1

## 原始摘要

Coding agents powered by large language models have shown impressive
capabilities in software engineering tasks, but evaluating their performance
across diverse programming languages and real-world scenarios remains
challenging. We introduce SWE-PolyBench, a new multi-language benchmark for
repository-level, execution-based evaluation of coding agents. SWE-PolyBench
contains 2110 instances from 21 repositories and includes tasks in Java (165),
JavaScript (1017), TypeScript (729) and Python (199), covering bug fixes,
feature additions, and code refactoring. We provide a task and
repository-stratified subsample (SWE-PolyBench500) and release an evaluation
harness allowing for fully automated evaluation. To enable a more comprehensive
comparison of coding agents, this work also presents a novel set of metrics
rooted in syntax tree analysis. We evaluate leading open source coding agents
on SWE-PolyBench, revealing their strengths and limitations across languages,
task types, and complexity classes. Our experiments show that current agents
exhibit uneven performances across languages and struggle with complex problems
while showing higher performance on simpler tasks. SWE-PolyBench aims to drive
progress in developing more versatile and robust AI coding assistants for
real-world software engineering. Our datasets and code are available at:
https://github.com/amazon-science/SWE-PolyBench


## AI 摘要

研究人员开发了名为SWE-PolyBench的多语言基准测试工具，用于评估基于大语言模型的编码代理在真实场景中的表现。该工具包含2110个测试实例，涵盖Java、JavaScript、TypeScript和Python四种语言，涉及bug修复、功能添加和代码重构等任务。研究还提出了基于语法树分析的新评估指标，并测试了主流开源编码代理，发现它们在处理复杂问题时表现不佳，且不同语言间性能差异明显。SWE-PolyBench旨在推动开发更强大、通用的AI编程助手。相关数据和代码已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-14T14:01:16Z
- **目录日期**: 2025-04-14
