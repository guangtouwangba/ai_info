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

该研究提出了WebThinker，一种增强大型推理模型（LRMs）能力的深度研究代理。它通过动态网络探索模块和自主的“思考-搜索-撰写”策略，使模型能实时获取网络信息并生成研究报告。为解决静态知识限制，WebThinker采用基于强化学习的迭代优化训练方法。实验表明，在复杂推理基准（如GPQA、GAIA）和科学报告生成任务中，WebThinker显著优于现有方法和专有系统，提升了LRMs在知识密集型任务中的可靠性和适用性。代码已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-01T06:01:44Z
- **目录日期**: 2025-05-01
