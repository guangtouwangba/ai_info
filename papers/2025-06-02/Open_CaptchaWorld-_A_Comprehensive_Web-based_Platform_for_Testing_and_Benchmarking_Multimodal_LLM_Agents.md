# Open CaptchaWorld: A Comprehensive Web-based Platform for Testing and Benchmarking Multimodal LLM Agents

**URL**: http://arxiv.org/abs/2505.24878v1

## 原始摘要

CAPTCHAs have been a critical bottleneck for deploying web agents in
real-world applications, often blocking them from completing end-to-end
automation tasks. While modern multimodal LLM agents have demonstrated
impressive performance in static perception tasks, their ability to handle
interactive, multi-step reasoning challenges like CAPTCHAs is largely untested.
To address this gap, we introduce Open CaptchaWorld, the first web-based
benchmark and platform specifically designed to evaluate the visual reasoning
and interaction capabilities of MLLM-powered agents through diverse and dynamic
CAPTCHA puzzles. Our benchmark spans 20 modern CAPTCHA types, totaling 225
CAPTCHAs, annotated with a new metric we propose: CAPTCHA Reasoning Depth,
which quantifies the number of cognitive and motor steps required to solve each
puzzle. Experimental results show that humans consistently achieve near-perfect
scores, state-of-the-art MLLM agents struggle significantly, with success rates
at most 40.0% by Browser-Use Openai-o3, far below human-level performance,
93.3%. This highlights Open CaptchaWorld as a vital benchmark for diagnosing
the limits of current multimodal agents and guiding the development of more
robust multimodal reasoning systems. Code and Data are available at this https
URL.


## AI 摘要

Open CaptchaWorld是首个基于网页的基准测试平台，专门评估多模态大语言模型(MLLM)代理在动态验证码(CAPTCHA)任务中的视觉推理和交互能力。该平台涵盖20种现代验证码类型共225个样本，并引入"验证码推理深度"新指标量化解决每个谜题所需的认知和操作步骤。实验显示人类准确率达93.3%，而最先进的MLLM代理最高仅达40%，远低于人类水平。该研究揭示了当前多模态代理的局限性，为开发更强大的多模态推理系统提供了重要基准。相关代码和数据已开源。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-06-02T23:00:53Z
- **目录日期**: 2025-06-02
