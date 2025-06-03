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

当前知识编辑（KE）研究多关注简单任务，现有评估框架下许多方法得分接近完美，但缺乏真实场景验证。为此，研究者提出基于脚本的评估基准ScEdit，涵盖反事实和时序编辑，整合词级和文本级评估方法，将传统事实型（"是什么"）评估扩展至行动型（"怎么做"）评估。实验表明，所有KE方法在现有指标上均表现下降，文本级评估面临挑战，凸显任务复杂性。该基准为KE研究提供了更贴近实际应用的评估体系，代码已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-03T20:01:21Z
- **目录日期**: 2025-06-03
