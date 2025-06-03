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

本文提出了3MDBench开源框架，用于评估大视觉语言模型(LVLMs)在远程医疗咨询中的表现。该框架通过4种性格的患者代理和评估代理，模拟真实场景下的3013个病例（覆盖34种诊断）。实验比较了GPT-4o-mini等主流模型，发现：1）多模态对话结合内部推理使F1值提升6.5%；2）在LVLM上下文中注入卷积网络的诊断预测可提升F1达20%。研究突出了情境感知和信息寻求式提问的重要性，为医疗AI发展提供了新思路。代码已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-03T05:01:22Z
- **目录日期**: 2025-06-03
