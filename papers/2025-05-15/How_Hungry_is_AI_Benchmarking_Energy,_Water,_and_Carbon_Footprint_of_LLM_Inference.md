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

本研究提出了一种新型基础设施感知基准框架，用于评估30种商用大语言模型(LLM)推理阶段的环境足迹。研究发现：o3和DeepSeek-R1能耗最高，单次长提示耗电超33Wh，是GPT-4.1 nano的70倍；Claude-3.7 Sonnet生态效率最佳。虽然单次GPT-4o短查询仅耗0.43Wh，但日处理7亿次查询的年环境影响惊人：耗电量相当于3.5万美国家庭、淡水蒸发量满足120万人年饮水需求、碳排放需芝加哥大小的森林才能抵消。研究表明，尽管单次查询高效，但全球规模仍导致资源消耗失衡，亟需建立AI可持续发展的标准化评估方法。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-15T12:01:22Z
- **目录日期**: 2025-05-15
