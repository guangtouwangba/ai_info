# Benchmarking Adversarial Robustness to Bias Elicitation in Large Language Models: Scalable Automated Assessment with LLM-as-a-Judge

**URL**: http://arxiv.org/abs/2504.07887v1

## 原始摘要

Large Language Models (LLMs) have revolutionized artificial intelligence,
driving advancements in machine translation, summarization, and conversational
agents. However, their increasing integration into critical societal domains
has raised concerns about embedded biases, which can perpetuate stereotypes and
compromise fairness. These biases stem from various sources, including
historical inequalities in training data, linguistic imbalances, and
adversarial manipulation. Despite mitigation efforts, recent studies indicate
that LLMs remain vulnerable to adversarial attacks designed to elicit biased
responses. This work proposes a scalable benchmarking framework to evaluate LLM
robustness against adversarial bias elicitation. Our methodology involves (i)
systematically probing models with a multi-task approach targeting biases
across various sociocultural dimensions, (ii) quantifying robustness through
safety scores using an LLM-as-a-Judge approach for automated assessment of
model responses, and (iii) employing jailbreak techniques to investigate
vulnerabilities in safety mechanisms. Our analysis examines prevalent biases in
both small and large state-of-the-art models and their impact on model safety.
Additionally, we assess the safety of domain-specific models fine-tuned for
critical fields, such as medicine. Finally, we release a curated dataset of
bias-related prompts, CLEAR-Bias, to facilitate systematic vulnerability
benchmarking. Our findings reveal critical trade-offs between model size and
safety, aiding the development of fairer and more robust future language
models.


## AI 摘要

大型语言模型(LLMs)在推动AI进步的同时，其潜在偏见问题引发关注。这些偏见源于训练数据中的历史不平等、语言不平衡等因素，可能加剧社会刻板印象。研究提出可扩展的评估框架CLEAR-Bias，通过多任务探测、安全评分量化及越狱技术测试，系统评估模型对抗偏见诱导的鲁棒性。分析发现：模型大小与安全性存在权衡关系；特定领域(如医疗)微调模型也存在安全隐患。研究揭示了当前先进模型普遍存在的偏见漏洞，并发布标准化评估数据集，为开发更公平、鲁棒的下一代语言模型提供重要参考。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-13T21:01:32Z
- **目录日期**: 2025-04-13
