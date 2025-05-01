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

WebThinker是一种深度研究智能体，旨在增强大型推理模型（LRM）的动态知识获取能力。它通过"深度网络探索器"模块实现自主网页搜索和导航，并采用"自主思考-搜索-撰写"策略，实时整合推理、信息收集和报告撰写。结合基于强化学习的在线直接偏好优化（DPO）训练方法，WebThinker在复杂推理基准（GPQA、GAIA等）和科研报告生成任务中显著优于现有系统。该方案提高了LRM在知识密集型任务中的可靠性，为构建更强大的深度研究系统开辟了新途径。代码已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-01T05:01:38Z
- **目录日期**: 2025-05-01
