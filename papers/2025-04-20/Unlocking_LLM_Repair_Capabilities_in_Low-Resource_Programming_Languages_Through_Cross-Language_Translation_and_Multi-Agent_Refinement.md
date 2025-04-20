# Unlocking LLM Repair Capabilities in Low-Resource Programming Languages Through Cross-Language Translation and Multi-Agent Refinement

**URL**: http://arxiv.org/abs/2503.22512v3

## 原始摘要

Recent advances in leveraging LLMs for APR have demonstrated impressive
capabilities in fixing software defects. However, current LLM-based approaches
predominantly focus on mainstream programming languages like Java and Python,
neglecting less prevalent but emerging languages such as Rust due to expensive
training resources, limited datasets, and insufficient community support. This
narrow focus creates a significant gap in repair capabilities across the
programming language spectrum, where the full potential of LLMs for
comprehensive multilingual program repair remains largely unexplored. To
address this limitation, we introduce a novel cross-language program repair
approach LANTERN that leverages LLMs' differential proficiency across languages
through a multi-agent iterative repair paradigm. Our technique strategically
translates defective code from languages where LLMs exhibit weaker repair
capabilities to languages where they demonstrate stronger performance, without
requiring additional training. A key innovation of our approach is an LLM-based
decision-making system that dynamically selects optimal target languages based
on bug characteristics and continuously incorporates feedback from previous
repair attempts. We evaluate our method on xCodeEval, a comprehensive
multilingual benchmark comprising 5,068 bugs across 11 programming languages.
Results demonstrate significant enhancement in repair effectiveness,
particularly for underrepresented languages, with Rust showing a 22.09%
improvement in Pass@10 metrics. Our research provides the first empirical
evidence that cross-language translation significantly expands the repair
capabilities of LLMs and effectively bridges the performance gap between
programming languages with different levels of popularity, opening new avenues
for truly language-agnostic automated program repair.


## AI 摘要

最新研究提出了一种跨语言程序修复方法LANTERN，利用大语言模型(LLM)在不同编程语言间的能力差异，通过多代理迭代修复范式提升修复效果。该方法将缺陷代码从LLM表现较弱的语言(如Rust)翻译至其擅长的语言进行修复，无需额外训练。实验在11种语言的5068个错误基准测试中显示，该方法显著提高了冷门语言的修复效果，其中Rust的Pass@10指标提升了22.09%。研究首次证明跨语言翻译能有效扩展LLM的修复能力，弥合不同流行度语言间的性能差距，为真正的语言无关自动程序修复开辟了新途径。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-20T00:02:04Z
- **目录日期**: 2025-04-20
