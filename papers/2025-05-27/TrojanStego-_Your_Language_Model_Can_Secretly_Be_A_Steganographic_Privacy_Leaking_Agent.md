# TrojanStego: Your Language Model Can Secretly Be A Steganographic Privacy Leaking Agent

**URL**: http://arxiv.org/abs/2505.20118v1

## 原始摘要

As large language models (LLMs) become integrated into sensitive workflows,
concerns grow over their potential to leak confidential information. We propose
TrojanStego, a novel threat model in which an adversary fine-tunes an LLM to
embed sensitive context information into natural-looking outputs via linguistic
steganography, without requiring explicit control over inference inputs. We
introduce a taxonomy outlining risk factors for compromised LLMs, and use it to
evaluate the risk profile of the threat. To implement TrojanStego, we propose a
practical encoding scheme based on vocabulary partitioning learnable by LLMs
via fine-tuning. Experimental results show that compromised models reliably
transmit 32-bit secrets with 87% accuracy on held-out prompts, reaching over
97% accuracy using majority voting across three generations. Further, they
maintain high utility, can evade human detection, and preserve coherence. These
results highlight a new class of LLM data exfiltration attacks that are
passive, covert, practical, and dangerous.


## AI 摘要

研究人员提出了一种名为TrojanStego的新型威胁模型，通过微调大语言模型（LLM）将敏感信息隐藏在看似自然的输出中，实现隐蔽数据泄露。该方法基于词汇划分的编码方案，实验显示被攻击模型能以87%的准确率传输32位密钥（经多数表决可达97%），同时保持高可用性和隐蔽性。这种被动、隐蔽的攻击方式突显了LLM数据泄露的新风险，可能在不控制输入的情况下实现信息窃取，具有实际威胁性且难以被人工检测。研究还提出了评估LLM安全风险的分类框架。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-27T22:02:05Z
- **目录日期**: 2025-05-27
