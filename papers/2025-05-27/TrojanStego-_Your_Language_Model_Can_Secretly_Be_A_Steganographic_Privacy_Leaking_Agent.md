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

研究人员提出了一种名为TrojanStego的新型威胁模型，通过微调大语言模型（LLM）将敏感信息以隐写术方式嵌入自然文本输出中，无需控制推理输入。该模型采用基于词汇分区的编码方案，实验显示被入侵模型能以87%的准确率传输32位机密信息，三次生成多数投票后准确率可达97%以上。这些模型在保持高实用性、连贯性的同时能逃避人工检测，凸显了一种被动、隐蔽、实用且危险的LLM数据泄露攻击方式。研究还提出了评估LLM风险因素的分类框架。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-05-27T12:02:20Z
- **目录日期**: 2025-05-27
