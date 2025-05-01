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

这篇论文提出了WebThinker，一种增强大型推理模型（LRM）自主网络研究能力的深度研究代理。它包含**深度网络探索模块**，使模型能动态搜索、导航和提取网络信息以填补知识空白，并采用**自主思考-搜索-草拟策略**，实时交替进行推理、信息收集和报告撰写。此外，通过基于强化学习的迭代在线直接偏好优化（DPO）训练策略提升工具使用能力。实验表明，WebThinker在复杂推理基准（GPQA等）和科学报告生成任务（Glaive）上显著优于现有方法和商业系统，提高了LRM在复杂场景中的可靠性和适用性。代码已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-01T07:01:43Z
- **目录日期**: 2025-05-01
