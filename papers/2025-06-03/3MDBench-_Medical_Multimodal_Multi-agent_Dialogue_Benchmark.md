# 3MDBench: Medical Multimodal Multi-agent Dialogue Benchmark

**URL**: http://arxiv.org/abs/2504.13861v2

## 原始摘要

Though Large Vision-Language Models (LVLMs) are being actively explored in
medicine, their ability to conduct telemedicine consultations combining
accurate diagnosis with professional dialogue remains underexplored. In this
paper, we present 3MDBench (Medical Multimodal Multi-agent Dialogue Benchmark),
an open-source framework for simulating and evaluating LVLM-driven telemedical
consultations. 3MDBench simulates patient variability through four
temperament-based Patient Agents and an Assessor Agent that jointly evaluate
diagnostic accuracy and dialogue quality. It includes 3013 cases across 34
diagnoses drawn from real-world telemedicine interactions, combining textual
and image-based data. The experimental study compares diagnostic strategies for
popular LVLMs, including GPT-4o-mini, LLaVA-3.2-11B-Vision-Instruct, and
Qwen2-VL-7B-Instruct. We demonstrate that multimodal dialogue with internal
reasoning improves F1 score by 6.5% over non-dialogue settings, highlighting
the importance of context-aware, information-seeking questioning. Moreover,
injecting predictions from a diagnostic convolutional network into the LVLM's
context boosts F1 by up to 20%. Source code is available at
https://anonymous.4open.science/r/3mdbench_acl-0511.


## AI 摘要

3MDBench是一个开源框架，用于评估大型视觉语言模型(LVLMs)在远程医疗咨询中的表现。它通过4种性格模拟的患者代理和评估代理，综合诊断准确性和对话质量，涵盖34种诊断的3013个真实案例。实验比较了GPT-4o-mini等主流LVLMs，发现多模态对话结合内部推理使F1分数提升6.5%，上下文感知的提问策略效果显著。此外，将诊断卷积网络的预测结果融入LVLM上下文可进一步提升F1分数达20%。该研究为LVLMs在专业医疗对话中的应用提供了新见解。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-03T15:01:23Z
- **目录日期**: 2025-06-03
