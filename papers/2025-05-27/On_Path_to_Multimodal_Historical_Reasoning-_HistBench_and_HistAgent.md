# On Path to Multimodal Historical Reasoning: HistBench and HistAgent

**URL**: http://arxiv.org/abs/2505.20246v1

## 原始摘要

Recent advances in large language models (LLMs) have led to remarkable
progress across domains, yet their capabilities in the humanities, particularly
history, remain underexplored. Historical reasoning poses unique challenges for
AI, involving multimodal source interpretation, temporal inference, and
cross-linguistic analysis. While general-purpose agents perform well on many
existing benchmarks, they lack the domain-specific expertise required to engage
with historical materials and questions. To address this gap, we introduce
HistBench, a new benchmark of 414 high-quality questions designed to evaluate
AI's capacity for historical reasoning and authored by more than 40 expert
contributors. The tasks span a wide range of historical problems-from factual
retrieval based on primary sources to interpretive analysis of manuscripts and
images, to interdisciplinary challenges involving archaeology, linguistics, or
cultural history. Furthermore, the benchmark dataset spans 29 ancient and
modern languages and covers a wide range of historical periods and world
regions. Finding the poor performance of LLMs and other agents on HistBench, we
further present HistAgent, a history-specific agent equipped with carefully
designed tools for OCR, translation, archival search, and image understanding
in History. On HistBench, HistAgent based on GPT-4o achieves an accuracy of
27.54% pass@1 and 36.47% pass@2, significantly outperforming LLMs with online
search and generalist agents, including GPT-4o (18.60%), DeepSeek-R1(14.49%)
and Open Deep Research-smolagents(20.29% pass@1 and 25.12% pass@2). These
results highlight the limitations of existing LLMs and generalist agents and
demonstrate the advantages of HistAgent for historical reasoning.


## AI 摘要

近期大语言模型（LLM）在多个领域取得显著进展，但在人文学科（尤其是历史领域）的应用仍待探索。历史推理涉及多模态源解释、时间推理和跨语言分析，对AI提出独特挑战。为此，研究者推出HistBench基准测试（含414个专家设计的历史问题），覆盖29种语言和多种历史时期，评估AI的历史推理能力。测试发现现有LLM表现不佳，于是开发了专用历史智能体HistAgent（集成OCR、翻译、档案搜索等工具）。基于GPT-4o的HistAgent在测试中表现显著优于通用模型（如GPT-4o和DeepSeek-R1），准确率达27.54%（pass@1），凸显领域专用方案的优势。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-27T20:01:00Z
- **目录日期**: 2025-05-27
