# How Hungry is AI? Benchmarking Energy, Water, and Carbon Footprint of LLM Inference

**URL**: http://arxiv.org/abs/2505.09598v1

## 原始摘要

As large language models (LLMs) spread across industries, understanding their
environmental footprint at the inference level is no longer optional; it is
essential. However, most existing studies exclude proprietary models, overlook
infrastructural variability and overhead, or focus solely on training, even as
inference increasingly dominates AI's environmental impact. To bridge this gap,
this paper introduces a novel infrastructure-aware benchmarking framework for
quantifying the environmental footprint of LLM inference across 30
state-of-the-art models as deployed in commercial data centers. Our framework
combines public API performance data with region-specific environmental
multipliers and statistical inference of hardware configurations. We
additionally utilize cross-efficiency Data Envelopment Analysis (DEA) to rank
models by performance relative to environmental cost. Our results show that o3
and DeepSeek-R1 emerge as the most energy-intensive models, consuming over 33
Wh per long prompt, more than 70 times the consumption of GPT-4.1 nano, and
that Claude-3.7 Sonnet ranks highest in eco-efficiency. While a single short
GPT-4o query consumes 0.43 Wh, scaling this to 700 million queries/day results
in substantial annual environmental impacts. These include electricity use
comparable to 35,000 U.S. homes, freshwater evaporation matching the annual
drinking needs of 1.2 million people, and carbon emissions requiring a
Chicago-sized forest to offset. These findings illustrate a growing paradox:
although individual queries are efficient, their global scale drives
disproportionate resource consumption. Our study provides a standardized,
empirically grounded methodology for benchmarking the sustainability of LLM
deployments, laying a foundation for future environmental accountability in AI
development and sustainability standards.


## AI 摘要

随着大语言模型(LLMs)的广泛应用，量化其推理阶段的环境影响变得至关重要。本研究提出了首个基础设施感知的基准框架，评估了30个商业数据中心部署的先进LLM模型。结果显示：o3和DeepSeek-R1能耗最高(单次长提示超33Wh)，是GPT-4.1 nano的70倍；Claude-3.7 Sonnet生态效率最佳。单个GPT-4o短查询耗电0.43Wh，但日处理7亿次查询时，年耗电相当于3.5万美国家庭，淡水蒸发量满足120万人年饮水需求，碳排放需芝加哥大小的森林抵消。研究表明，尽管单次查询高效，但全球规模仍导致巨大资源消耗。该研究为AI可持续发展提供了标准化评估方法。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-15T22:01:10Z
- **目录日期**: 2025-05-15
