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

大型语言模型(LLMs)在推动AI进步的同时，其社会应用也引发了关于偏见的担忧。这些偏见源于训练数据的不平等、语言失衡和对抗性操纵。本研究提出可扩展的评估框架，通过多任务测试模型在不同社会文化维度的偏见表现，采用LLM-as-a-Judge方法量化安全性，并利用越狱技术检测防御漏洞。研究分析了主流大/小模型的偏见问题及其对安全性的影响，评估了医疗等关键领域专用模型的安全性，并发布CLEAR-Bias数据集用于系统评估。结果显示模型规模与安全性之间存在关键权衡，为开发更公平、鲁棒的模型提供参考。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-13T00:25:25Z
- **目录日期**: 2025-04-13
