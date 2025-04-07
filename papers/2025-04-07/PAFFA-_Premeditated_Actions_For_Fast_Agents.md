# PAFFA: Premeditated Actions For Fast Agents

**URL**: http://arxiv.org/abs/2412.07958v2

## 原始摘要

Modern AI assistants have made significant progress in natural language
understanding and tool-use, with emerging efforts to interact with Web
interfaces. However, current approaches that heavily rely on repeated
LLM-driven HTML parsing are computationally expensive and error-prone,
particularly when handling dynamic web interfaces and multi-step tasks. We
introduce PAFFA (Premeditated Actions For Fast Agents), a method that makes
LLMs faster and more accurate in completing tasks on the internet using a novel
inference-time technique that requires no task-specific training. PAFFA
constructs an 'Action Library', leveraging the parametric knowledge of the base
LLM to pre-compute browser interaction patterns that generalize across tasks.
By strategically re-using LLM inference across tasks - either via 'Dist-Map'
for task-agnostic identification of key interactive web elements, or 'Unravel'
for first-encounter, stateful exploration of novel tasks/sites) - PAFFA
drastically reduces inference time tokens by 87% while maintaining robust
performance (achieving 0.57 vs. 0.50 step accuracy compared to baseline).
Further, Unravel's ability to update its action library based on explorations
allows generalization and adaptation to unseen websites. In sum, this work
exhibits that LLM reasoning sequences can generalize across prompts, offering a
way to scale inference-time techniques for internet-scale data with sublinear
token count.


## AI 摘要

研究人员提出PAFFA方法，通过构建"动作库"预计算浏览器交互模式，显著提升大语言模型(LLM)的网页任务处理效率。该方法采用两种策略：通用元素识别的"Dist-Map"和首次探索的"Unravel"，减少87%推理时间的同时保持准确性(步骤准确率0.57 vs 基线0.50)。Unravel还能根据探索更新动作库，适应新网站。研究表明LLM推理序列可跨提示泛化，为互联网规模数据处理提供了次线性增长的推理方案。该方法无需任务特定训练，有效解决了传统HTML解析方法计算成本高、易出错的问题。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-07T07:02:24Z
- **目录日期**: 2025-04-07
