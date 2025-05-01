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

WebThinker是一种深度研究代理，旨在增强大型推理模型（LRM）在复杂知识密集型任务中的表现。它通过动态网络搜索和实时信息整合，解决了传统LRM依赖静态知识的局限性。WebThinker包含一个深度网络探索模块，可在推理过程中自主搜索、导航和提取网络信息，并采用“思考-搜索-撰写”策略，无缝结合推理、信息收集和报告撰写。通过基于强化学习的在线优化训练，进一步提升了研究工具的使用效率。实验表明，WebThinker在复杂推理基准和科学报告生成任务中显著优于现有方法，增强了LRM在复杂场景下的可靠性和适用性。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-01T04:04:42Z
- **目录日期**: 2025-05-01
