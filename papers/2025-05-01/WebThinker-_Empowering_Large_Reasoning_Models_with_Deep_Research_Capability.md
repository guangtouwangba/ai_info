# WebThinker: Empowering Large Reasoning Models with Deep Research Capability

**URL**: http://arxiv.org/abs/2504.21776v1

## 原始摘要

Large reasoning models (LRMs), such as OpenAI-o1 and DeepSeek-R1, demonstrate
impressive long-horizon reasoning capabilities. However, their reliance on
static internal knowledge limits their performance on complex,
knowledge-intensive tasks and hinders their ability to produce comprehensive
research reports requiring synthesis of diverse web information. To address
this, we propose \textbf{WebThinker}, a deep research agent that empowers LRMs
to autonomously search the web, navigate web pages, and draft research reports
during the reasoning process. WebThinker integrates a \textbf{Deep Web
Explorer} module, enabling LRMs to dynamically search, navigate, and extract
information from the web when encountering knowledge gaps. It also employs an
\textbf{Autonomous Think-Search-and-Draft strategy}, allowing the model to
seamlessly interleave reasoning, information gathering, and report writing in
real time. To further enhance research tool utilization, we introduce an
\textbf{RL-based training strategy} via iterative online Direct Preference
Optimization (DPO). Extensive experiments on complex reasoning benchmarks
(GPQA, GAIA, WebWalkerQA, HLE) and scientific report generation tasks (Glaive)
demonstrate that WebThinker significantly outperforms existing methods and
strong proprietary systems. Our approach enhances LRM reliability and
applicability in complex scenarios, paving the way for more capable and
versatile deep research systems. The code is available at
https://github.com/RUC-NLPIR/WebThinker.


## AI 摘要

本文提出WebThinker，一种增强大型推理模型（LRMs）自主网络搜索与研究报告撰写能力的深度研究代理。它包含深度网络探索模块，使LRM能动态检索网页信息填补知识缺口，并采用自主"思考-搜索-草拟"策略实现实时推理、信息收集与报告撰写的无缝衔接。通过基于强化学习的在线直接偏好优化（DPO）训练策略，系统在复杂推理基准（GPQA/GAIA等）和科学报告生成任务上显著优于现有方法。该方案提升了LRM在复杂场景下的可靠性与适用性，代码已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-01T08:01:48Z
- **目录日期**: 2025-05-01
