# Assessing Consistency and Reproducibility in the Outputs of Large Language Models: Evidence Across Diverse Finance and Accounting Tasks

**URL**: http://arxiv.org/abs/2503.16974v2

## 原始摘要

This study provides the first comprehensive assessment of consistency and
reproducibility in Large Language Model (LLM) outputs in finance and accounting
research. We evaluate how consistently LLMs produce outputs given identical
inputs through extensive experimentation with 50 independent runs across five
common tasks: classification, sentiment analysis, summarization, text
generation, and prediction. Using three OpenAI models (GPT-3.5-turbo,
GPT-4o-mini, and GPT-4o), we generate over 3.4 million outputs from diverse
financial source texts and data, covering MD&amp;As, FOMC statements, finance news
articles, earnings call transcripts, and financial statements. Our findings
reveal substantial but task-dependent consistency, with binary classification
and sentiment analysis achieving near-perfect reproducibility, while complex
tasks show greater variability. More advanced models do not consistently
demonstrate better consistency and reproducibility, with task-specific patterns
emerging. LLMs significantly outperform expert human annotators in consistency
and maintain high agreement even where human experts significantly disagree. We
further find that simple aggregation strategies across 3-5 runs dramatically
improve consistency. We also find that aggregation may come with an additional
benefit of improved accuracy for sentiment analysis when using newer models.
Simulation analysis reveals that despite measurable inconsistency in LLM
outputs, downstream statistical inferences remain remarkably robust. These
findings address concerns about what we term "G-hacking," the selective
reporting of favorable outcomes from multiple Generative AI runs, by
demonstrating that such risks are relatively low for finance and accounting
tasks.


## AI 摘要

这项研究首次全面评估了大型语言模型(LLM)在金融会计研究中的输出一致性和可复现性。通过对5类常见任务(分类、情感分析、摘要、文本生成和预测)进行50次独立实验，使用3种OpenAI模型生成340万条输出。结果显示：二元分类和情感分析几乎完美复现，复杂任务变异性更大；高级模型不一定更稳定；LLM一致性显著优于人类专家；3-5次运行的简单聚合可大幅提升稳定性，新模型的情感分析准确性也能提高。尽管存在输出不一致，下游统计推断仍保持稳健，缓解了"G-hacking"的担忧。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-03-27T23:02:06Z
- **目录日期**: 2025-03-27
