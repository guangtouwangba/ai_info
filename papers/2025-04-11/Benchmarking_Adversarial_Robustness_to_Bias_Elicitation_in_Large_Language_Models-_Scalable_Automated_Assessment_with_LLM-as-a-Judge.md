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

大型语言模型(LLMs)在AI领域取得突破性进展，但也存在内嵌偏见问题，可能强化刻板印象并影响公平性。这些偏见源于训练数据中的历史不平等、语言不平衡和对抗性操纵。研究提出可扩展的评估框架CLEAR-Bias，通过多任务方法系统探测模型在不同社会文化维度的偏见，采用LLM-as-a-Judge自动评分量化鲁棒性，并利用越狱技术检测安全机制漏洞。分析显示模型大小与安全性存在权衡关系，领域专用模型(如医疗)也面临偏见挑战。该研究为开发更公平、更鲁棒的语言模型提供了评估工具和数据集支持。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-11T23:01:20Z
- **目录日期**: 2025-04-11
