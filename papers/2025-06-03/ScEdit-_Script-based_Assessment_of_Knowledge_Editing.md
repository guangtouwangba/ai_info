# ScEdit: Script-based Assessment of Knowledge Editing

**URL**: http://arxiv.org/abs/2505.23291v2

## 原始摘要

Knowledge Editing (KE) has gained increasing attention, yet current KE tasks
remain relatively simple. Under current evaluation frameworks, many editing
methods achieve exceptionally high scores, sometimes nearing perfection.
However, few studies integrate KE into real-world application scenarios (e.g.,
recent interest in LLM-as-agent). To support our analysis, we introduce a novel
script-based benchmark -- ScEdit (Script-based Knowledge Editing Benchmark) --
which encompasses both counterfactual and temporal edits. We integrate
token-level and text-level evaluation methods, comprehensively analyzing
existing KE techniques. The benchmark extends traditional fact-based
("What"-type question) evaluation to action-based ("How"-type question)
evaluation. We observe that all KE methods exhibit a drop in performance on
established metrics and face challenges on text-level metrics, indicating a
challenging task. Our benchmark is available at
https://github.com/asdfo123/ScEdit.


## AI 摘要

当前知识编辑（KE）研究虽受关注，但任务设计简单，现有评估框架下多数方法得分接近满分，缺乏真实场景验证。为此，研究者提出基于脚本的评估基准ScEdit，涵盖反事实和时序编辑，结合词级和文本级评估，将传统事实型（"What"类问题）评测扩展到行动型（"How"类问题）。实验表明，所有KE方法在现有指标上均表现下降，文本级评估面临挑战，凸显任务难度。该基准旨在推动KE技术面向实际应用（如LLM智能体），代码已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-03T13:11:05Z
- **目录日期**: 2025-06-03
